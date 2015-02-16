package zuint64

import (
	"sort"
)

// Calling zuint64.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []uint64) {
	if len(x) < MinSize {
		sort.Sort(uint64Sortable(x))
	} else {
		buffer := make([]uint64, len(x))
		SortBYOB(x, buffer)
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []uint64) []uint64 {
	y := make([]uint64, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// Sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices..
func SortBYOB(x, buffer []uint64) {
	checkSlices(x, buffer)
	copy(buffer, x)

	// Radix is a byte, 8 bytes to a uint64
	for pass := uint(0); pass < 8; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			sortPass(x[:], buffer[:], pass)
		} else {
			sortPass(buffer[:], x[:], pass)
		}
	}
}

func sortPass(from, to []uint64, pass uint) {
	byteOffset := pass * radix
	byteMask := uint64(0xFF << byteOffset)
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

func checkSlices(a, b []uint64) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}

type uint64Sortable []uint64

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
