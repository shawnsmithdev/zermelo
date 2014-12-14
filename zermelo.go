// Zermelo is a library for sorting slices in Go.
package zermelo

import (
	"sort"
	"reflect"
)

// The radix size using during radix sorts - a byte.
const rSortRadix = 8

// Slices smaller than this use sort.Sort() instead of radix sort.
const rSortMinSize = 256

func Sort(x interface{}) {
	xVal := reflect.ValueOf(x)
	xKind := xVal.Kind()
	if (xKind != reflect.Slice) {
		return
	}
	xElemKind := reflect.TypeOf(x).Elem().Kind()
	switch xElemKind {
	case reflect.Uint32:
		SortUint32(xVal.Interface().([]uint32))
	case reflect.Uint64:
		SortUint64(xVal.Interface().([]uint64))
	case reflect.Int32:
		SortInt32(xVal.Interface().([]int32))
	case reflect.Int64:
		SortInt64(xVal.Interface().([]int64))
	}
}

// Sorts a []uint64 using a Radix sort.  This uses O(n) extra memory
func SortUint64(r []uint64) {
	if len(r) < rSortMinSize {
		sort.Sort(uint64Sortable(r))
	} else {
		buffer := make([]uint64, len(r))
		rsortUint64BYOB(r, buffer)
	}
}

// Sorts a []uint32 using a Radix sort.  This uses O(n) extra memory
func SortUint32(r []uint32) {
	if len(r) < rSortMinSize {
		sort.Sort(uint32Sortable(r))
	} else {
		buffer := make([]uint32, len(r))
		rsortUint32BYOB(r, buffer)
	}
}

// Sorts a []int32 using a Radix sort.  This uses O(n) extra memory
func SortInt32(r []int32) {
	if len(r) < rSortMinSize {
		sort.Sort(int32Sortable(r))
	} else {
		buffer := make([]int32, len(r))
		rsortInt32BYOB(r, buffer)
	}
}

// Sorts a []int64 using a Radix sort.  This uses O(n) extra memory
func SortInt64(r []int64) {
	if len(r) < rSortMinSize {
		sort.Sort(int64Sortable(r))
	} else {
		buffer := make([]int64, len(r))
		rsortInt64BYOB(r, buffer)
	}
}
