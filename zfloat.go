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

func SortFloats[F constraints.Float](x []F) {
	if len(x) < 2 {
		return
	}
	is32 := isFloat32[F]()
	if (is32 && len(x) < compSortCutoffFloat32) || (!is32 && len(x) < compSortCutoffFloat64) {
		slices.Sort(x)
		return
	}
	sortFloatsBYOB(x, make([]F, len(x)), is32)
}

func SortFloatsBYOB[F constraints.Float](x, buf []F) {
	sortFloatsBYOB(x, buf, isFloat32[F]())
}

func sortFloatsBYOB[F constraints.Float](x, buf []F, is32 bool) {
	if len(x) < 2 {
		return
	}
	// TODO NaN handling is undocumented, untested behavior
	// Don't sort NaNs, just put them up front and skip them
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
			unsafeFlipSortFlip[F, []F, uint32](
				x[nans:], unsafeSliceConvert[F, []F, uint32](buf), 32)
		} else {
			unsafeFlipSortFlip[F, []F, uint64](
				x[nans:], unsafeSliceConvert[F, []F, uint64](buf), 64)
		}
	}
	runtime.KeepAlive(buf)
}
