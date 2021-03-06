// Package zfloat64 implements radix sort for []float64.
package zfloat64

import (
	"math"
	"sort"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize      = 256
	radix   uint = 8
	bitSize uint = 64
)

// Sort sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []float64) {
	if len(x) < MinSize {
		sort.Float64s(x)
	} else {
		SortBYOB(x, make([]float64, len(x)))
	}
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []float64) []float64 {
	y := make([]float64, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []float64) {
	if len(x) > len(buffer) {
		panic("Buffer too small")
	}
	if len(x) < 2 {
		return
	}

	// Don't sort NaNs, just put them up front and skip them
	nans := 0
	for idx, val := range x {
		if math.IsNaN(val) {
			x[idx] = x[nans]
			x[nans] = val
			nans++
		}
	}

	// Each pass processes a byte offset, copying back and forth between slices
	from := x[nans:]
	to := buffer[:len(from)]
	var key uint8
	var uintVal uint64

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += radix {
		var offset [256]int // Keep track of where room is made for byte groups in the buffer
		sorted := true      // Check for already sorted
		prev := float64(0)  // if elem is always >= prev it is already sorted

		for _, val := range from {
			uintVal = floatFlip(math.Float64bits(val))
			key = uint8(uintVal >> keyOffset) // fetch the byte at current 'digit'
			offset[key]++                     // count of values to put in this digit's bucket

			if sorted { // Detect sorted
				sorted = val >= prev
				prev = val
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
		for i, count := range offset {
			offset[i] = watermark
			watermark += count
		}

		// Rebucket while copying to other buffer
		for _, val := range from {
			uintVal = floatFlip(math.Float64bits(val))
			key = uint8(uintVal >> keyOffset) // Get the digit
			to[offset[key]] = val             // Copy the element to the digit's bucket
			offset[key]++                     // One less space, move the offset
		}

		// Reverse buffers on each pass
		from, to = to, from
	}
}

// Converts a uint64 that represents a true float to one that sorts properly
func floatFlip(x uint64) uint64 {
	if (x & 0x8000000000000000) == 0x8000000000000000 {
		return x ^ 0xFFFFFFFFFFFFFFFF
	}
	return x ^ 0x8000000000000000
}
