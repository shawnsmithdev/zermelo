package zermelo

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"unsafe"
)

// unsafeFlipSortFlip converts float slices to unsigned, flips some bits to allow sorting, sorts and unflips.
// F and U must be the same bit size, and len(buf) must be >= len(x)
// This will not work if NaNs are present in x. Remove them first.
func unsafeFlipSortFlip[F constraints.Float, U constraints.Unsigned](x []F, b []U, size uint) {
	y := unsafeSliceConvert[F, U](x)
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
}

// unsafeSliceConvert takes a slice of one type and returns a slice
// of another type using the same memory for the backing array.
//
// This must only be used to temporarily treat elements in a slice as though they were of a different type.
// One must not modify the length or capacity of either the given or returned slice
// while the returned slice is still in scope.
//
// If x goes out of scope, the returned slice becomes invalid, as they share memory but the garbage collector is
// unaware of the returned slice and may invalidate that memory. Working around this may require
// use of `runtime.KeepAlive(x)`.
//
// F and U must have the same bit size (64 or 32).
func unsafeSliceConvert[F any, U any](x []F) []U {
	var result []U
	xHeader := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	resultHeader := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	resultHeader.Data = xHeader.Data
	resultHeader.Len = xHeader.Len
	resultHeader.Cap = xHeader.Cap
	return result
}
