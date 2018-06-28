// Package zint32 implements radix sort for []int32.
package zint32

import (
	"sort"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize        = 128
	radix    uint  = 8
	bitSize  uint  = 32
	minInt32 int32 = -1 << 31
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
	var offset [256]int // Keep track of where groups start

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += radix {
		keyMask := int32(0xFF << keyOffset)
		sorted := true
		prev := minInt32
		var counts [256]int // Keep track of the number of elements for each kind of byte
		for _, elem := range from {
			// For each elem to sort, fetch the byte at current radix
			key = uint8((elem & keyMask) >> keyOffset)
			// inc count of bytes of this type
			counts[key]++
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

		if keyOffset == bitSize-radix {
			// Last pass. Handle signed values
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
			for i := 1; i < 256; i++ {
				offset[i] = offset[i-1] + counts[i-1]
			}
		}

		// Swap values between the buffers by radix
		for _, elem := range from {
			key = uint8((elem & keyMask) >> keyOffset) // Get the byte of each element at the radix
			to[offset[key]] = elem                     // Copy the element depending on byte offsets
			offset[key]++
		}
		// Reverse buffers on each pass
		to, from = from, to
	}
}
