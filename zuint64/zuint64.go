// Package zuint64 implements radix sort for []uint64.
package zuint64

import (
	"sort"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize      = 256
	radix   uint = 8
	bitSize uint = 64
)

// Sort Sort x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []uint64) {
	if len(x) < MinSize {
		sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	} else {
		buffer := make([]uint64, len(x))
		SortBYOB(x, buffer)
	}
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []uint64) []uint64 {
	y := make([]uint64, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []uint64) {
	if len(x) > len(buffer) {
		panic("Buffer too small")
	}
	if len(x) < 2 {
		return
	}

	// Each pass processes a byte offset, copying back and forth between slices
	from := x
	to := buffer[:len(x)]
	var key uint8

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += radix {
		keyMask := uint64(0xFF << keyOffset) // Current 'digit' to look at
		var offset [256]int                  // Keep track of where groups start
		sorted := true                       // Check for already sorted
		prev := uint64(0)                    // if elem is always >= prev it is already sorted
		for _, elem := range from {
			key = uint8((elem & keyMask) >> keyOffset) // fetch the byte at current 'digit'
			offset[key]++                              // count of elems to put in this digit's bucket

			if sorted { // Detect sorted
				sorted = elem >= prev
				prev = elem
			}
		}

		if sorted { // Short-circuit sorted
			if (keyOffset/radix)%2 == 1 {
				copy(to, from)
			}
			return
		}

		// Find target bucket offsets
		watermark := offset[0] - offset[0] // Like := 0, but inherits the type.
		for i, count := range offset {
			offset[i] = watermark
			watermark += count
		}

		// Rebucket while copying to other buffer
		for _, elem := range from {
			key = uint8((elem & keyMask) >> keyOffset) // Get the digit
			to[offset[key]] = elem                     // Copy the element to the digit's bucket
			offset[key]++                              // One less space, move the offset
		}
		// On next pass copy data the other way
		to, from = from, to
	}
}
