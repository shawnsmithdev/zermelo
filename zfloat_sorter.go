package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
)

type FloatSorter[F constraints.Float] interface {
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
