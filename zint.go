package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

const (
	radix            uint = 8
	maxSize               = 64
	compSortCutoff64      = 256
	compSortCutoff        = 128
)

// SortIntegers sorts integer slices. If the slice is large enough, radix sort is used by allocating a new buffer.
func SortIntegers[T constraints.Integer](x []T) {
	if len(x) < 2 {
		return
	}
	size, minval := detect[T]()
	if (size == 64 && len(x) < compSortCutoff64) || len(x) < compSortCutoff {
		slices.Sort(x)
		return
	}
	sortIntegersBYOB(x, make([]T, len(x)), size, minval)
}

// SortIntegersBYOB sorts integer slices. If the slice is large enough, radix sort is used with the provided buffer.
// len(buf) must be greater or equal to len(x).
func SortIntegersBYOB[T constraints.Integer](x, buffer []T) {
	size, minval := detect[T]()
	sortIntegersBYOB(x, buffer, size, minval)
}

func sortIntegersBYOB[T constraints.Integer](x, buffer []T, size uint, minval T) {
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

// returns bit size and min value of T
func detect[T constraints.Integer]() (uint, T) {
	// 'fffe' has all but least bit set
	fffe := (^T(0)) ^ T(1)

	// find size of T in bits
	size := uint(maxSize)
	for fffe<<(size>>1) == 0 {
		size = size >> 1
	}

	// if 'ffff' is positive, T is unsigned
	if ^T(0) > 0 {
		return size, 0
	}
	// T is signed, min val is '8000'
	return size, fffe << (size - 2)
}
