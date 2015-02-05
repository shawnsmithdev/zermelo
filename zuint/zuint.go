package zuint

import (
	"reflect"
	"sort"
)

// Calling zuint.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8

// This won't fit in a 32 bits uint, but will fit a 64 bit uint
const tooBig = 0x1000000000

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []uint) {
	if len(x) < MinSize {
		sort.Sort(uintSortable(x))
	} else {
		buffer := make([]uint, len(x))
		SortBYOB(x, buffer)
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []uint) []uint {
	y := make([]uint, len(x))
	if len(x) < MinSize {
		copy(x, y)
		sort.Sort(uintSortable(y))
	} else {
		buffer := make([]uint, len(x))
		SortBYOB(x, buffer)
	}
	return y
}

// Sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices..
func SortBYOB(x, buffer []uint) {
	checkSlices(x, buffer)
	copy(buffer, x)

	// Reflection because we don't know if uint is 32 or 64 bits.
	passCount := uint(reflect.TypeOf(uint(0)).Bits() / radix)
	for pass := uint(0); pass < passCount; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			sortPass(x[:], buffer[:], pass)
		} else {
			sortPass(buffer[:], x[:], pass)
		}
	}
}

func sortPass(from, to []uint, pass uint) {
	byteOffset := pass * radix
	byteMask := uint(0xFF << byteOffset)
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

func checkSlices(a, b []uint) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}

type uintSortable []uint

func (r uintSortable) Len() int           { return len(r) }
func (r uintSortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uintSortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
