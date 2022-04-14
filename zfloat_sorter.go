package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
)

// FloatSorter describes types that can sort float slices
type FloatSorter[F constraints.Float] interface {
	// Sort sorts float slices
	Sort(x []F)
}

type zFloatSorter[F constraints.Float, U constraints.Unsigned] struct {
	size           uint
	buf            []U
	compSortCutoff int
}

func (z *zFloatSorter[F, U]) Sort(x []F) {
	if len(x) < z.compSortCutoff {
		slices.Sort(x)
		return
	}
	if len(z.buf) < len(x) {
		z.buf = make([]U, allocSize(len(z.buf), len(x)))
	}
	unsafeFlipSortFlip[F, []F, U](x, z.buf, z.size)
}

// NewFloatSorter creates a new FloatSorter that will use radix sort on large slices and reuses buffers.
// The first sort creates a buffer the same size as the slice being sorted and keeps it for future use.
// Later sorts may grow this buffer as needed. The FloatSorter returned is not thread safe.
// Using this sorter can be much faster than repeat calls to SortFloats.
func NewFloatSorter[F constraints.Float]() FloatSorter[F] {
	if isFloat32[F]() {
		return &zFloatSorter[F, uint32]{
			size:           32,
			compSortCutoff: compSortCutoffFloat32,
		}
	}
	return &zFloatSorter[F, uint64]{
		size:           64,
		compSortCutoff: compSortCutoffFloat64,
	}
}

func isFloat32[F constraints.Float]() bool {
	return F(math.SmallestNonzeroFloat32)/2 == 0
}
