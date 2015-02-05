// zfloat64 sorts []float64 by copying the data to new uint64 backed buffers, before sorting them
// with zuint64, and copying the sorted floats back. This means this allocates twice the additional
// memory that integer based sorts in zermelo like zuint64 usually do.
//
// However, if memory is available, this is much faster than sort.Float64s() for large slices.
package zfloat64

import (
	"github.com/shawnsmithdev/zermelo/zuint64"
	"math"
	"sort"
)

// Calling zfloat64.Sort() on slices smaller than this will result is sorting with sort.Sort() instead.
const MinSize = 256

const radix = 8

// Sorts x using a Radix sort (Small slices are sorted with sort.Sort() instead).
func Sort(x []float64) {
	if len(x) < MinSize {
		sort.Float64s(x)
	} else {
		SortBYOB(x, make([]uint64, len(x)), make([]uint64, len(x)))
	}
}

// Similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []float64) []float64 {
	y := make([]uint64, len(x))
	for idx, val := range x {
		y[idx] = floatFlip(math.Float64bits(val))
	}
	zuint64.Sort(y)

	z := make([]float64, len(x))
	for idx, val := range y {
		z[idx] = math.Float64frombits(floatFlop(val))
	}
	return z
}

// Sorts x using a Radix sort, using supplied buffer space y and z. Panics if
// len(x) does not equal len(y) or len(z). Uses radix sort even on small slices..
func SortBYOB(x []float64, y, z []uint64) {
	for idx, val := range x {
		y[idx] = floatFlip(math.Float64bits(val))
	}
	zuint64.SortBYOB(y, z)
	for idx, val := range y {
		x[idx] = math.Float64frombits(floatFlop(val))
	}
}

// Converts a uint64 that represents a true float to one sorts properly
func floatFlip(x uint64) uint64 {
	if (x & 0x8000000000000000) == 0x8000000000000000 {
		return x ^ 0xFFFFFFFFFFFFFFFF
	}
	return x ^ 0x8000000000000000
}

// Inverse of floatFlip()
func floatFlop(x uint64) uint64 {
	if (x & 0x8000000000000000) == 0 {
		return x ^ 0xFFFFFFFFFFFFFFFF
	}
	return x ^ 0x8000000000000000
}
