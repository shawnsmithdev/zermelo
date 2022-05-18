package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"runtime"
)

const (
	compSortCutoffFloat32 = 128
	compSortCutoffFloat64 = 384
)

// SortFloats sorts float slices. If the slice is large enough, radix sort is used by allocating a new buffer.
func SortFloats[F constraints.Float](x []F) {
	x = sortNaNs(x)
	if len(x) < 2 {
		return
	}
	is32 := isFloat32[F]()
	if len(x) < compSortCutoffFloat32 || (!is32 && len(x) < compSortCutoffFloat64) {
		slices.Sort(x)
		return
	}
	sortFloatsBYOB(x, make([]F, len(x)), is32)
}

// SortFloatsBYOB sorts float slices with radix sort using the provided buffer.
// len(buffer) must be greater or equal to len(x).
func SortFloatsBYOB[F constraints.Float](x, buffer []F) {
	x = sortNaNs(x)
	if len(x) >= 2 {
		sortFloatsBYOB(x, buffer, isFloat32[F]())
	}
}

func sortFloatsBYOB[F constraints.Float](x, buf []F, is32 bool) {
	if is32 {
		unsafeFlipSortFlip[F, uint32](
			x, unsafeSliceConvert[F, uint32](buf), 32)
	} else {
		unsafeFlipSortFlip[F, uint64](
			x, unsafeSliceConvert[F, uint64](buf), 64)
	}
	runtime.KeepAlive(buf) // avoid gc as buf is never used directly
}

// isFloat32 returns true if F is float32, false if float64
func isFloat32[F constraints.Float]() bool {
	return F(math.SmallestNonzeroFloat32)/2 == 0
}

// isNaN returns true only if x is a float32 or float64 representing a NaN value, as only NaN is not equal itself.
func isNaN[C comparable](x C) bool { return x != x }

// sortNaNs put nans up front, similar to sort.Float64s, returning a slice of x excluding those nans
func sortNaNs[F constraints.Float](x []F) []F {
	nans := 0
	for idx, val := range x {
		if isNaN(val) {
			x[idx] = x[nans]
			x[nans] = val
			nans++
		}
	}
	return x[nans:]
}
