// Package zermelo is a library for sorting slices in Go.
package zermelo // import "github.com/shawnsmithdev/zermelo/v2"

import (
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"slices"
)

const (
	radix            uint = 8
	compSortCutoff64      = 256
	compSortCutoff        = 128
)

// Sort sorts integer slices. If the slice is large enough, radix sort is used by allocating a new buffer.
func Sort[T Integer](x []T) {
	if len(x) < 2 {
		return
	}
	size, minval := internal.Detect[T]()
	if len(x) < compSortCutoff || (size == 64 && len(x) < compSortCutoff64) {
		slices.Sort(x)
	} else {
		sortBYOB(x, make([]T, len(x)), size, minval)
	}
}

// SortBYOB sorts integer slices with radix sort using the provided buffer.
// len(buffer) must be greater or equal to len(x).
func SortBYOB[T Integer](x, buffer []T) {
	if len(x) >= 2 {
		size, minval := internal.Detect[T]()
		sortBYOB(x, buffer, size, minval)
	}
}

func sortBYOB[T Integer](x, buffer []T, size uint, minval T) {
	from := x
	to := buffer[:len(x)]

	var keyOffset uint
	for keyOffset = 0; keyOffset < size; keyOffset += radix {
		var (
			offset [256]int // Keep track of where room is made for byte groups in the buffer
			prev   = minval
			key    uint8
			sorted = true
		)

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
			break
		}

		// Find target bucket offsets
		var watermark int
		if minval != 0 && keyOffset == size-radix {
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

	// copy from buffer if done during odd turn
	if radix&keyOffset == radix {
		copy(to, from)
	}
}
