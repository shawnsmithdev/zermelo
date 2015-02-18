// Radix sort for []uint64.
package zuint64

import (
	"sort"
)

// Calling zuint64.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8
const bitSize = 64

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
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []uint64) {
	if x == nil || buffer == nil {
		panic("Slices must not be nil")
	} else if len(x) > len(buffer) {
		panic("Buffer too small")
	}

	// Each pass processes a byte offset, copying back and forth between slices
	from := x
	to := buffer[:len(x)]
	for byteOffset := uint(0); byteOffset < bitSize; byteOffset += radix {
		from, to = sortPass(from, to, byteOffset)
	}
}

// Does a pass of radix sort, copying data between the slices
func sortPass(from, to []uint64, byteOffset uint) (newfrom, newto []uint64) {
	byteMask := uint64(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value

	for _, elem := range from {
		// For each value to sort, fetch the byte at current radix
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
	// Next pass copy data the other way
	return to, from
}

// Implements sort.Interface for small slices
type uint64Sortable []uint64

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
