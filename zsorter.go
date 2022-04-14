package zermelo

import (
	"errors"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Sorter can sort slices
type Sorter interface {
	// Sort attempts to sort x, returning an error if unable to sort.
	Sort(x any) error
	// CopySort returns a sorted copy of x, or an error if unable to copy or sort.
	CopySort(x any) (interface{}, error)
}

// New creates a Sorter that reuses buffers on repeated Sort() or CopySort() calls on the same type.
// This is not thread safe. CopySort() does not support passthrough of sort.Interface values.
// This style of sorter is deprecated, use NewIntSorter/NewFloatSorter
func New() Sorter {
	return &zSorter{}
}

type zSorter struct {
	ints     cutoffIntSorter[int]
	int64s   cutoffIntSorter[int64]
	int32s   cutoffIntSorter[int32]
	int16s   cutoffIntSorter[int16]
	int8s    cutoffIntSorter[int8]
	uints    cutoffIntSorter[uint]
	uintptrs cutoffIntSorter[uintptr]
	uint64s  cutoffIntSorter[uint64]
	uint32s  cutoffIntSorter[uint32]
	uint16s  cutoffIntSorter[uint16]
	uint8s   cutoffIntSorter[uint8]
	float64s cutoffFloatSorter[float64]
	float32s cutoffFloatSorter[float32]
}

func maybeIntSort[I constraints.Integer](x []I, sorter cutoffIntSorter[I]) (cutoffIntSorter[I], bool) {
	if sorter == nil {
		result := newIntSorter[I]()
		result.Sort(x)
		return result, true
	}
	sorter.Sort(x)
	return sorter, false
}

func maybeFloatSort[F constraints.Float](x []F, sorter cutoffFloatSorter[F]) (cutoffFloatSorter[F], bool) {
	if sorter == nil {
		result := newFloatSorter[F]()
		result.Sort(x)
		return result, true
	}
	sorter.Sort(x)
	return sorter, false
}

func (z *zSorter) Sort(x any) error {
	switch xAsCase := x.(type) {
	case []int:
		if newSorter, isNew := maybeIntSort[int](xAsCase, z.ints); isNew {
			z.ints = newSorter
		}
	case []int64:
		if newSorter, isNew := maybeIntSort[int64](xAsCase, z.int64s); isNew {
			z.int64s = newSorter
		}
	case []int32:
		if newSorter, isNew := maybeIntSort[int32](xAsCase, z.int32s); isNew {
			z.int32s = newSorter
		}
	case []int16:
		if newSorter, isNew := maybeIntSort[int16](xAsCase, z.int16s); isNew {
			z.int16s = newSorter
		}
	case []int8:
		if newSorter, isNew := maybeIntSort[int8](xAsCase, z.int8s); isNew {
			z.int8s = newSorter
		}
	case []uint:
		if newSorter, isNew := maybeIntSort[uint](xAsCase, z.uints); isNew {
			z.uints = newSorter
		}
	case []uintptr:
		if newSorter, isNew := maybeIntSort[uintptr](xAsCase, z.uintptrs); isNew {
			z.uintptrs = newSorter
		}
	case []uint64:
		if newSorter, isNew := maybeIntSort[uint64](xAsCase, z.uint64s); isNew {
			z.uint64s = newSorter
		}
	case []uint32:
		if newSorter, isNew := maybeIntSort[uint32](xAsCase, z.uint32s); isNew {
			z.uint32s = newSorter
		}
	case []uint16:
		if newSorter, isNew := maybeIntSort[uint16](xAsCase, z.uint16s); isNew {
			z.uint16s = newSorter
		}
	case []uint8:
		if newSorter, isNew := maybeIntSort[uint8](xAsCase, z.uint8s); isNew {
			z.uint8s = newSorter
		}
	case []float32:
		if newSorter, isNew := maybeFloatSort[float32](xAsCase, z.float32s); isNew {
			z.float32s = newSorter
		}
	case []float64:
		if newSorter, isNew := maybeFloatSort[float64](xAsCase, z.float64s); isNew {
			z.float64s = newSorter
		}
	default:
		return errors.New("unsupported type")
	}
	return nil
}

func maybeCopySort[T any](z *zSorter, x []T) ([]T, error) {
	sorted := slices.Clone(x)
	return sorted, z.Sort(sorted)
}

func (z *zSorter) CopySort(x any) (any, error) {
	switch xAsCase := x.(type) {
	case []int:
		return maybeCopySort[int](z, xAsCase)
	case []int64:
		return maybeCopySort[int64](z, xAsCase)
	case []int32:
		return maybeCopySort[int32](z, xAsCase)
	case []int16:
		return maybeCopySort[int16](z, xAsCase)
	case []int8:
		return maybeCopySort[int8](z, xAsCase)
	case []uint:
		return maybeCopySort[uint](z, xAsCase)
	case []uintptr:
		return maybeCopySort[uintptr](z, xAsCase)
	case []uint64:
		return maybeCopySort[uint64](z, xAsCase)
	case []uint32:
		return maybeCopySort[uint32](z, xAsCase)
	case []uint16:
		return maybeCopySort[uint16](z, xAsCase)
	case []uint8:
		return maybeCopySort[uint8](z, xAsCase)
	case []float32:
		return maybeCopySort[float32](z, xAsCase)
	case []float64:
		return maybeCopySort[float64](z, xAsCase)
	default:
		return nil, errors.New("unsupported type")
	}
}
