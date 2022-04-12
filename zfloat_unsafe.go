package zermelo

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"runtime"
	"unsafe"
)

// unsafeFlipSortFlip converts float slices to unsigned, flips some bits to allow sorting, sorts and unflips.
// F and U must be the same bit size, and len(buf) must be >= len(x)
// This will not work if NaNs are present in x. Remove them first.
func unsafeFlipSortFlip[F constraints.Float, S ~[]F, U constraints.Unsigned](x S, b []U, size uint) {
	y := unsafeSliceConvert[F, []F, U](x)
	allBits := ^U(0)
	topBit := U(1) << (size - 1)

	// flip
	for idx, val := range y {
		if val&topBit == topBit {
			y[idx] = val ^ allBits
		} else {
			y[idx] = val ^ topBit
		}
	}

	// sort
	sortIntegersBYOB(y, b, size, 0)

	// unflip
	for idx, val := range y {
		if val&topBit == topBit {
			y[idx] = val ^ topBit
		} else {
			y[idx] = val ^ allBits
		}
	}
	runtime.KeepAlive(x)
}

// unsafeSliceConvert takes a slice of one type and returns a slice
// of another type using the same memory for the backing array.
// If x goes out of scope, the returned slice becomes invalid.
// F and U must have the same bit size (64 or 32).
func unsafeSliceConvert[F constraints.Float, S ~[]F, U constraints.Unsigned](x S) []U {
	var result []U
	xHeader := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	resultHeader := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	resultHeader.Data = xHeader.Data
	resultHeader.Len = xHeader.Len
	resultHeader.Cap = xHeader.Cap
	return result
}
