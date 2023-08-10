package floats

import (
	"github.com/shawnsmithdev/zermelo/v2"
	"slices"
)

// cutoffSorter is a Sorter with adjustable comparison sort cutoff, for testing.
type cutoffSorter[F Float] interface {
	zermelo.Sorter[F]
	withCutoff(int) cutoffSorter[F]
}

type floatSorter[F Float, U zermelo.Unsigned] struct {
	uintSorter     zermelo.Sorter[U]
	compSortCutoff int
	topBit         U
}

func (s *floatSorter[F, U]) Sort(x []F) {
	x = sortNaNs(x)
	if len(x) < 2 {
		return
	}
	if len(x) < s.compSortCutoff {
		slices.Sort(x)
		return
	}

	y := unsafeSliceConvert[F, U](x)
	floatFlip[U](y, s.topBit)
	s.uintSorter.Sort(y)
	floatUnflip[U](y, s.topBit)
}

func (s *floatSorter[F, U]) withCutoff(cutoff int) cutoffSorter[F] {
	s.compSortCutoff = cutoff
	return s
}

// NewFloatSorter creates a new Sorter for float slices that will use radix sort on large slices and reuses buffers.
// The first sort creates a buffer the same size as the slice being sorted and keeps it for future use.
// Later sorts may grow this buffer as needed. The FloatSorter returned is not thread safe.
// Using this sorter can be much faster than repeat calls to SortFloats.
func NewFloatSorter[F Float]() zermelo.Sorter[F] {
	return newFloatSorter[F]()
}

func newFloatSorter[F Float]() cutoffSorter[F] {
	if isFloat32[F]() {
		return &floatSorter[F, uint32]{
			uintSorter:     zermelo.NewSorter[uint32](),
			compSortCutoff: compSortCutoffFloat32,
			topBit:         uint32(1) << 31,
		}
	}
	return &floatSorter[F, uint64]{
		uintSorter:     zermelo.NewSorter[uint64](),
		compSortCutoff: compSortCutoffFloat64,
		topBit:         uint64(1) << 63,
	}
}
