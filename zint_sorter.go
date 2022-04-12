package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type IntSorter[I constraints.Integer] interface {
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

func NewIntSorter[I constraints.Integer]() IntSorter[I] {
	result := &zIntSorter[I]{}
	if size, _ := detect[I](); size == 64 {
		result.compSortCutoff = compSortCutoff64
	} else {
		result.compSortCutoff = compSortCutoff
	}
	return result
}
