// Zermelo is a library for sorting slices in Go.
package zermelo

import (
	"sort"
)

// The radix size using during radix sorts - a byte.
const rSortRadix = 8

// Slices smaller than this use sort.Sort() instead of radix sort.
const rSortMinSize = 256

// Sorts x with the library if it is a large enough, supported slice type.
// Otherwise it tries to sort x using sort.Sort().
// If it is a supported type and doesn't implement sort.Interface, this does nothing.
func Sort(x interface{}) {
	switch xAsCase := x.(type) {
	case []uint32:
		SortUint32(xAsCase)
	case []uint64:
		SortUint64(xAsCase)
	case []int32:
		SortInt32(xAsCase)
	case []int64:
		SortInt64(xAsCase)
	case sort.Interface:
		sort.Sort(xAsCase)
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
