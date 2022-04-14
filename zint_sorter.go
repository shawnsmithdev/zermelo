package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// IntSorter describes types that can sort integer slices
type IntSorter[I constraints.Integer] interface {
	// Sort sorts integer slices.
	Sort(x []I)
}

type zIntSorter[I constraints.Integer] struct {
	buf            []I
	compSortCutoff int
}

func (z *zIntSorter[I]) Sort(x []I) {
	if len(x) < z.compSortCutoff {
		slices.Sort(x)
		return
	}
	if len(z.buf) < len(x) {
		z.buf = make([]I, allocSize(len(z.buf), len(x)))
	}
	SortIntegersBYOB(x, z.buf)
}

// NewIntSorter creates a new IntSorter that will use radix sort on large slices and reuses buffers.
// The first sort creates a buffer the same size as the slice being sorted and keeps it for future use.
// Later sorts may grow this buffer as needed. The IntSorter returned is not thread safe.
// Using this sorter can be much faster than repeat calls to SortIntegers.
func NewIntSorter[I constraints.Integer]() IntSorter[I] {
	result := &zIntSorter[I]{}
	if size, _ := detect[I](); size == 64 {
		result.compSortCutoff = compSortCutoff64
	} else {
		result.compSortCutoff = compSortCutoff
	}
	return result
}
