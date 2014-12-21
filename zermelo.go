// Zermelo is a library for sorting slices in Go.
package zermelo

import (
	"sort"
)

// The radix size using during radix sorts - a byte.
const rSortRadix = 8

// Slices smaller than this use sort.Sort() instead of radix sort.
const rSortMinSize = 256

// Attempts to sort x. If x is a supported slice type, this library will be
// be used to sort it. Otherwise, this attempts to sort x using sort.Sort().
// If x is not a supported type, and doesn't implement sort.Interface, this does nothing.
func Sort(x interface{}) {
	switch xAsCase := x.(type) {
	case []uint:
		SortUint(xAsCase)
	case []uint32:
		SortUint32(xAsCase)
	case []uint64:
		SortUint64(xAsCase)
	case []int:
		SortInt(xAsCase)
	case []int32:
		SortInt32(xAsCase)
	case []int64:
		SortInt64(xAsCase)
	case []float32:
		SortFloat32(xAsCase)
	case []float64:
		SortFloat64(xAsCase)
	case sort.Interface:
		sort.Sort(xAsCase)
	}
}

// Sorts a []uint using a Radix sort.
func SortUint(r []uint) {
	sort.Sort(uintSortable(r)) // TODO
}

// Sorts a []uint32 using a Radix sort.
func SortUint32(r []uint32) {
	if len(r) < rSortMinSize {
		sort.Sort(uint32Sortable(r))
	} else {
		buffer := make([]uint32, len(r))
		rsortUint32BYOB(r, buffer)
	}
}

// Sorts a []uint64 using a Radix sort.
func SortUint64(r []uint64) {
	if len(r) < rSortMinSize {
		sort.Sort(uint64Sortable(r))
	} else {
		buffer := make([]uint64, len(r))
		rsortUint64BYOB(r, buffer)
	}
}

// Sorts a []int using a Radix sort.
func SortInt(r []int) {
	sort.Sort(sort.IntSlice(r)) // TODO
}

// Sorts a []int32 using a Radix sort.
func SortInt32(r []int32) {
	if len(r) < rSortMinSize {
		sort.Sort(int32Sortable(r))
	} else {
		buffer := make([]int32, len(r))
		rsortInt32BYOB(r, buffer)
	}
}

// Sorts a []int64 using a Radix sort.
func SortInt64(r []int64) {
	if len(r) < rSortMinSize {
		sort.Sort(int64Sortable(r))
	} else {
		buffer := make([]int64, len(r))
		rsortInt64BYOB(r, buffer)
	}
}

// Sorts a []float32 using a Radix sort.
func SortFloat32(r []float32) {
	sort.Sort(float32Sortable(r)) // TODO
}

// Sorts a []float64 using a Radix sort.
func SortFloat64(r []float64) {
	sort.Sort(sort.Float64Slice(r)) // TODO
}
