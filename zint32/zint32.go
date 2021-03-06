// Package zint32 implements radix sort for []int32.
package zint32

import (
	"math"
	"sort"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize      = 128
	radix   uint = 8
	bitSize uint = 32
)

// Sort sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []int32) {
	if len(x) < MinSize {
		sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	} else {
		buffer := make([]int32, len(x))
		SortBYOB(x, buffer)
	}
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []int32) []int32 {
	y := make([]int32, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts a []int32 using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []int32) {
	if len(x) > len(buffer) {
		panic("Buffer too small")
	}
	if len(x) < 2 {
		return
	}

	from := x
	to := buffer[:len(x)]
	var key uint8

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += radix {
		sorted := true
		var prev int32 = math.MinInt32
		var offset [256]int // Keep track of where room is made for byte groups in the buffer
		for _, elem := range from {
			// For each elem to sort, fetch the byte at current radix
			key = uint8(elem >> keyOffset)
			// inc count of bytes of this type
			offset[key]++
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
		watermark := 0
		if keyOffset == bitSize-radix {
			// Handle signed values
			// Negatives
			for i := 128; i < len(offset); i++ {
				count := offset[i]
				offset[i] = watermark
				watermark += count
			}
			// Positives
			for i := 0; i < 128; i++ {
				count := offset[i]
				offset[i] = watermark
				watermark += count
			}
		} else {
			for i, count := range offset {
				offset[i] = watermark
				watermark += count
			}
		}

		// Swap values between the buffers by radix
		for _, elem := range from {
			key = uint8(elem >> keyOffset) // Get the byte of each element at the radix
			to[offset[key]] = elem         // Copy the element depending on byte offsets
			offset[key]++
		}

		// Reverse buffers on each pass
		from, to = to, from
	}
}
