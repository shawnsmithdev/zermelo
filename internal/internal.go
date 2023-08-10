package internal

import (
	"crypto/rand"
)

const maxSize uint = 64

// Integer is a constraint that permits any integer type.
type Integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Detect returns bit size and min value of T
func Detect[T Integer]() (uint, T) {
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

// FillSlice will fill a slice with values returned by gen.
func FillSlice[T any](x []T, gen func() T) {
	for i := range x {
		x[i] = gen()
	}
}

// RandInteger returns a function that generates random integers, including negative values.
func RandInteger[I Integer]() func() I {
	tSize, _ := Detect[I]()
	buf := make([]byte, tSize/8)
	return func() I {
		var result I
		_, _ = rand.Read(buf)
		for i, val := range buf {
			result |= I(val) << (i * 8)
		}
		return result
	}
}
