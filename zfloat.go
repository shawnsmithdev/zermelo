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
	if len(x) >= 2 {
		is32 := isFloat32[F]()
		if len(x) < compSortCutoffFloat32 || (!is32 && len(x) < compSortCutoffFloat64) {
			slices.Sort(x)
		} else {
			sortFloatsBYOB(x, make([]F, len(x)), is32)
		}
	}
}

// SortFloatsBYOB sorts float slices with radix sort using the provided buffer.
// len(buf) must be greater or equal to len(x).
func SortFloatsBYOB[F constraints.Float](x, buf []F) {
	if len(x) >= 2 {
		sortFloatsBYOB(x, buf, isFloat32[F]())
	}
}

func sortFloatsBYOB[F constraints.Float](x, buf []F, is32 bool) {
	// Put nans up front and skip them, similar to sort.Float64s
	nans := 0
	for idx, val := range x {
		if math.IsNaN(float64(val)) {
			x[idx] = x[nans]
			x[nans] = val
			nans++
		}
	}
	if len(x)-nans >= 2 {
		if is32 {
			unsafeFlipSortFlip[F, uint32](
				x[nans:], unsafeSliceConvert[F, uint32](buf), 32)
		} else {
			unsafeFlipSortFlip[F, uint64](
				x[nans:], unsafeSliceConvert[F, uint64](buf), 64)
		}
		runtime.KeepAlive(buf) // avoid gc as buf is never used directly
	}
}
