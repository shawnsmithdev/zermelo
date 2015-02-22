// Radix sort for []uint32.
package zuint32

import (
	"sort"
)

// Calling zuint32.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 128

const radix = 8

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []uint32) {
	if len(x) < MinSize {
		sort.Sort(uint32Sortable(x))
	} else {
		buffer := make([]uint32, len(x))
		SortBYOB(x, buffer)
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []uint32) []uint32 {
	y := make([]uint32, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// Sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices..
func SortBYOB(x, buffer []uint32) {
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

func sortPass(from, to []uint32, pass uint) {
	byteOffset := pass * radix
	byteMask := uint32(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value

	for _, elem := range from {
		// For each elem to sort, fetch the byte at current radix
		passByte = uint8((elem & byteMask) >> byteOffset)
		// inc count of bytes of this type
		counts[passByte]++
	}

	// Make room for each group of bytes by calculating offset of each
	offset[0] = 0
	for i := 1; i < len(offset); i++ {
		offset[i] = offset[i-1] + counts[i-1]
	}

	// Swap values between the buffers by radix
	for _, elem := range from {
		passByte = uint8((elem & byteMask) >> byteOffset) // Get the byte of each element at the radix
		to[offset[passByte]] = elem                       // Copy the element depending on byte offsets
		offset[passByte]++                                // One less space, move the offset
	}
}

func checkSlices(a, b []uint32) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}

type uint32Sortable []uint32

func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
