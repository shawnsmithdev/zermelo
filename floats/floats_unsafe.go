package floats

import (
	"github.com/shawnsmithdev/zermelo/v2"
	"golang.org/x/exp/constraints"
	"reflect"
	"unsafe"
)

// unsafeFlipSortFlip converts float slices to unsigned, flips some bits to allow sorting, sorts and unflips.
// F and U must be the same bit size, and len(buf) must be >= len(x)
// This will not work if NaNs are present in x. Remove them first.
func unsafeFlipSortFlip[F constraints.Float, U constraints.Unsigned](x, b []F, size uint) {
	xu := unsafeSliceConvert[F, U](x)
	bu := unsafeSliceConvert[F, U](b)
	floatFlip[U](xu, U(1)<<(size-1))
	zermelo.SortBYOB(xu, bu)
	floatUnflip[U](xu, U(1)<<(size-1))
}

func floatFlip[U constraints.Unsigned](y []U, topBit U) {
	for idx, val := range y {
		if val&topBit == topBit {
			y[idx] = val ^ (^U(0))
		} else {
			y[idx] = val ^ topBit
		}
	}
}

func floatUnflip[U constraints.Unsigned](y []U, topBit U) {
	for idx, val := range y {
		if val&topBit == topBit {
			y[idx] = val ^ topBit
		} else {
			y[idx] = val ^ (^U(0))
		}
	}
}

// unsafeSliceConvert takes a slice of one type and returns a slice of another type using the same memory
// for the backing array. F and U obviously must be exactly the same size for this to work.
//
// This must only be used to temporarily treat elements in a slice as though they were of a different type.
// One must not modify the length or capacity of either the given or returned slice
// while the returned slice is still in scope.
//
// If x goes out of scope, the returned slice becomes invalid, as they share memory but the garbage collector is
// unaware of the returned slice and may invalidate that memory. Working around this may require
// use of `runtime.KeepAlive(x)`.
//
func unsafeSliceConvert[F any, U any](x []F) []U {
	var result []U
	xHeader := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	resultHeader := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	resultHeader.Data = xHeader.Data
	resultHeader.Len = xHeader.Len
	resultHeader.Cap = xHeader.Cap
	return result
}
