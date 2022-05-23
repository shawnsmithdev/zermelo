package internal

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

const maxSize uint = 64

// Detect returns bit size and min value of T
func Detect[T constraints.Integer]() (uint, T) {
	// 'fffe' has all but least bit set
	fffe := (^T(0)) ^ T(1)

	// find size of T in bits
	size := maxSize
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

func FillSlice[T any](x []T, gen func() T) {
	for i := range x {
		x[i] = gen()
	}
}

func RandInteger[I constraints.Integer]() func() I {
	tSize, _ := Detect[I]()
	buf := make([]byte, tSize/8)
	return func() I {
		var result I
		rand.Read(buf)
		for i, val := range buf {
			result |= I(val) << (i * 8)
		}
		return result
	}
}
