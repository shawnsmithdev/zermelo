// Radix sort for []int32.
package zint32

import (
	"sort"
)

// Calling zint32.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []int32) {
	if len(x) < MinSize {
		sort.Sort(int32Sortable(x))
	} else {
		buffer := make([]int32, len(x))
		SortBYOB(x, buffer)
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []int32) []int32 {
	y := make([]int32, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// Sorts a []int32 using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []int32) {
	checkSlices(x, buffer)

	// Radix is a byte, 4 bytes to a uint32
	for pass := uint(0); pass < 4; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			sortPass(x[:], buffer[:], pass)
		} else {
			sortPass(buffer[:], x[:], pass)
		}
	}
}

func sortPass(from, to []int32, pass uint) {
	byteOffset := pass * radix
	byteMask := int32(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value
	for _, elem := range from {
		// For each elem to sort, fetch the byte at current radix
		passByte = uint8((elem & byteMask) >> byteOffset)
		// inc count of bytes of this type
		counts[passByte]++
	}

	if pass == 3 {
		// Handle signed values
		// Count negative elements (last 128 counts)
		negCnt := 0
		for i := 128; i < 256; i++ {
			negCnt += counts[i]
		}

		offset[0] = negCnt // Start of positives
		offset[128] = 0    // Start of negatives
		for i := 1; i < 128; i++ {
			// Positive values
			offset[i] = offset[i-1] + counts[i-1]
			// Negative values
			offset[i+128] = offset[i+127] + counts[i+127]
		}
	} else {
		offset[0] = 0
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}
	}

	// Swap values between the buffers by radix
	for _, elem := range from {
		passByte = uint8((elem & byteMask) >> byteOffset) // Get the byte of each element at the radix
		to[offset[passByte]] = elem                       // Copy the element depending on byte offsets
		offset[passByte]++
	}
}

func checkSlices(a, b []int32) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}

type int32Sortable []int32

func (r int32Sortable) Len() int           { return len(r) }
func (r int32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
