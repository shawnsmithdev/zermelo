package zint

import (
	"reflect"
	"sort"
)

// Calling zint.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []int) {
	if len(x) < MinSize {
		sort.Sort(intSortable(x))
	} else {
		buffer := make([]int, len(x))
		SortBYOB(x, buffer)
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []int) []int {
	y := make([]int, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// Sorts a []int using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []int) {
	checkSlices(x, buffer)
	copy(buffer, x)

	// Reflection because we don't know if int is 32 or 64 bits.
	passCount := uint(reflect.TypeOf(int(0)).Bits() / radix)
	for pass := uint(0); pass < passCount; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			sortPass(x[:], buffer[:], pass, pass == passCount-1)
		} else {
			sortPass(buffer[:], x[:], pass, pass == passCount-1)
		}
	}
}

func sortPass(from, to []int, pass uint, last bool) {
	byteOffset := pass * radix
	byteMask := int(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value
	for _, elem := range from {
		// For each elem to sort, fetch the byte at current radix
		passByte = uint8((elem & byteMask) >> byteOffset)
		// inc count of bytes of this type
		counts[passByte]++
	}

	if last {
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

func checkSlices(a, b []int) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}

type intSortable []int

func (r intSortable) Len() int           { return len(r) }
func (r intSortable) Less(i, j int) bool { return r[i] < r[j] }
func (r intSortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
