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

type cutoffIntSorter[I constraints.Integer] interface {
	IntSorter[I]
	setCutoff(int)
}

type zIntSorter[I constraints.Integer] struct {
	buf            []I
	compSortCutoff int
	minval         I
	size           uint
}

func (z *zIntSorter[I]) Sort(x []I) {
	if len(x) < z.compSortCutoff {
		slices.Sort(x)
		return
	}
	if len(z.buf) < len(x) {
		z.buf = make([]I, allocSize(len(z.buf), len(x)))
	}
	sortIntegersBYOB(x, z.buf, z.size, z.minval)
}

func (z *zIntSorter[I]) setCutoff(cutoff int) {
	z.compSortCutoff = cutoff
}

// NewIntSorter creates a new IntSorter that will use radix sort on large slices and reuses buffers.
// The first sort creates a buffer the same size as the slice being sorted and keeps it for future use.
// Later sorts may grow this buffer as needed. The IntSorter returned is not thread safe.
// Using this sorter can be much faster than repeat calls to SortIntegers.
func NewIntSorter[I constraints.Integer]() IntSorter[I] {
	return newIntSorter[I]()
}

func newIntSorter[I constraints.Integer]() cutoffIntSorter[I] {
	size, minval := detect[I]()
	cutoff := compSortCutoff
	if size == 64 {
		cutoff = compSortCutoff64
	}
	result := &zIntSorter[I]{
		compSortCutoff: cutoff,
		minval:         minval,
		size:           size,
	}
	return result
}

// Given an existing buffer capacity and a requested one, finds a new buffer size.
// For the first alloc this will equal requested size, then after at it leaves
// a 25% buffer for future growth.
func allocSize(bufCap, reqLen int) int {
	if bufCap == 0 {
		return reqLen
	}
	return 5 * reqLen / 4
}
