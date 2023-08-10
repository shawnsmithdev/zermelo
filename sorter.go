package zermelo

import (
	"cmp"
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"slices"
)

// Sorter describes types that can sort slices.
type Sorter[T cmp.Ordered] interface {
	// Sort sorts slices in ascending order.
	Sort(x []T)
}

// cutoffSorter is a Sorter with adjustable comparison sort cutoff, for testing.
type cutoffSorter[T Integer] interface {
	Sorter[T]
	withCutoff(int) cutoffSorter[T]
}

type sorter[I Integer] struct {
	buf            []I
	compSortCutoff int
	minval         I
	size           uint
}

func (s *sorter[I]) Sort(x []I) {
	if len(x) < s.compSortCutoff {
		slices.Sort(x)
		return
	}
	if len(s.buf) < len(x) {
		s.buf = make([]I, allocSize(len(s.buf), len(x)))
	}
	sortBYOB(x, s.buf, s.size, s.minval)
}

func (s *sorter[I]) withCutoff(cutoff int) cutoffSorter[I] {
	s.compSortCutoff = cutoff
	return s
}

// NewSorter creates a new Sorter that will use radix sort on large slices and reuses buffers.
// The first sort creates a buffer the same size as the slice being sorted and keeps it for future use.
// Later sorts may grow this buffer as needed. The Sorter returned is not thread safe.
// Using this sorter can be much faster than repeat calls to Sort.
func NewSorter[I Integer]() Sorter[I] {
	return newSorter[I]()
}

func newSorter[I Integer]() cutoffSorter[I] {
	size, minval := internal.Detect[I]()
	cutoff := compSortCutoff
	if size == 64 {
		cutoff = compSortCutoff64
	}
	return &sorter[I]{
		compSortCutoff: cutoff,
		minval:         minval,
		size:           size,
	}
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
