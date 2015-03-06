// Radix sort for []string (alpha, sort.Strings() is faster right now).
package zstring

import (
	"sort"
)

func Sort(x []string) {
	SortBYOB(x, make([]string, len(x)))
}

func trim (x string) string {
	end := sort.Search(len(x), func(i int) bool { return x[i] == 0 })
	return x[:end]
}

func SortBYOB(x, buffer []string) {
	if len(x) > len(buffer) {
		panic("Buffer too small")
	}
	if len(x) < 2 {
		return
	}

	from := x
	to := buffer[:len(x)]

	// Find longest string and add padding
	maxLen := 0
	for _, val := range from {
		if len(val) > maxLen {
			maxLen = len(val)
		}
	}
	for i, val := range from {
		newString := make([]byte, maxLen)
		copy(newString, val)
		to[i] = string(newString)
	}
	to, from = from, to

	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	for byteOffset := maxLen - 1; byteOffset >= 0; byteOffset-- {
		var counts [256]int // Keep track of the number of elements for each kind of byte
		for _, val := range from {
			counts[val[byteOffset]]++
		}
		// Find target bucket offsets
		offset[0] = 0
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}

		// Rebucket while copying to other buffer
		for _, val := range from {
			key := val[byteOffset]         // Get the digit
			to[offset[key]] = val // Copy the element to the digit's bucket
			offset[key]++                  // One less space, move the offset
		}
		// On next pass copy data the other way
		to, from = from, to
	}
	// Trim, and make sure final result is in correct buffer
	if maxLen&1 == 0 {
		for i, val := range buffer[:len(x)] {
			x[i] = trim(val)
		}
	} else {
		for i, val := range x {
			x[i] = trim(val)
		}
	}
}
