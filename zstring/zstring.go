// Radix sort for []string.

// This code is alpha quality, and the API can change without notice.
// This will likely run faster when strings have similar sizes.
package zstring

import (
	"github.com/shawnsmithdev/zermelo/zuint32"
)

func Sort(x []string) {
	SortBYOB(x, make([]string, len(x)))
}

func SortBYOB(x, buffer []string) {
	if len(x) > len(buffer) {
		panic("Buffer too small")
	}
	if len(x) < 2 {
		return
	}

	firstPass(x, buffer[:len(x)])
	to := x
	from := buffer[:len(x)]
	maxLen := len(from[len(from)-1])
	shortest := len(from) - 1 // Shortest word being considered
	for byteOffset := maxLen - 1; byteOffset >= 0; byteOffset-- {
		for shortest > 0 && len(from[shortest-1]) > byteOffset {
			shortest--
		}
		var counts [256]int // Keep track of the number of elements for each kind of byte
		var offset [256]int // Keep track of where room is made for byte groups in the buffer

		for _, val := range from[shortest:] {
			counts[val[byteOffset]]++
		}
		// Find target bucket offsets
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}

		// Rebucket while copying to other buffer
		for _, val := range from[shortest:] {
			key := val[byteOffset]         // Get the digit
			to[shortest+offset[key]] = val // Copy the element to the digit's bucket
			offset[key]++                  // One less space, move the offset
		}
		// On next pass copy data the other way
		to, from = from, to
	}
	if maxLen&1 == 0 {
		copy(x, from)
	}
}

func firstPass(from, to []string) {
	sizeCounts := make(map[uint32]uint32)
	for _, s := range from {
		sizeCounts[uint32(len(s))]++
	}
	var sizes []uint32
	for size := range sizeCounts {
		sizes = append(sizes, size)
	}
	zuint32.Sort(sizes)
	offset := uint32(0)
	sizeOffsets := make(map[uint32]uint32)
	for _, size := range sizes {
		sizeOffsets[size] = offset
		offset += sizeCounts[size]
	}

	for _, s := range from {
		to[sizeOffsets[uint32(len(s))]] = s
		sizeOffsets[uint32(len(s))]++
	}
	copy(from, to)
}
