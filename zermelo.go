// Zermelo is a library for sorting slices in Go.
package zermelo

import (
	"errors"
	"github.com/shawnsmithdev/zermelo/zfloat32"
	"github.com/shawnsmithdev/zermelo/zfloat64"
	"github.com/shawnsmithdev/zermelo/zint"
	"github.com/shawnsmithdev/zermelo/zint32"
	"github.com/shawnsmithdev/zermelo/zint64"
	"github.com/shawnsmithdev/zermelo/zuint"
	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
	"sort"
)

// Attempts to sort x.
//
// If x is a supported slice type, this library will be used to sort it. Otherwise,
// if x implements sort.Interface it will passthrough to the sort.Sort() algorithm.
// Returns an error on unsupported types.
func Sort(x interface{}) error {
	switch xAsCase := x.(type) {
	case []float32:
		zfloat32.Sort(xAsCase)
	case []float64:
		zfloat64.Sort(xAsCase)
	case []int:
		zint.Sort(xAsCase)
	case []int32:
		zint32.Sort(xAsCase)
	case []int64:
		zint64.Sort(xAsCase)
	case []uint:
		zuint.Sort(xAsCase)
	case []uint32:
		zuint32.Sort(xAsCase)
	case []uint64:
		zuint64.Sort(xAsCase)
	case sort.Interface:
		sort.Sort(xAsCase)
	default:
		return errors.New("type not supported")
	}
	return nil
}

// A Sorter can sort things like slices. Returns an error on unsupported types.
type Sorter interface {
	Sort(x interface{}) error
	CopySort(x interface{}) (interface{}, error)
}

// Reuseable buffers
type zSorter struct {
	bufInt         []int
	bufInt32       []int32
	bufInt64       []int64
	bufUint        []uint
	bufUint32Alpha []uint32
	bufUint32Beta  []uint32
	bufUint64Alpha []uint64
	bufUint64Beta  []uint64
}

func (z *zSorter) prepBuffers(x interface{}) {
	switch xAsCase := x.(type) {
	case []float32:
		if cap(z.bufUint32Alpha) < len(xAsCase) {
			z.bufUint32Alpha = make([]uint32, len(xAsCase))
		}
		if cap(z.bufUint32Beta) < len(xAsCase) {
			z.bufUint32Beta = make([]uint32, len(xAsCase))
		}
	case []float64:
		if cap(z.bufUint64Alpha) < len(xAsCase) {
			z.bufUint64Alpha = make([]uint64, len(xAsCase))
		}
		if cap(z.bufUint64Beta) < len(xAsCase) {
			z.bufUint64Beta = make([]uint64, len(xAsCase))
		}
	case []int:
		if cap(z.bufInt) < len(xAsCase) {
			z.bufInt = make([]int, len(xAsCase))
		}
	case []int32:
		if cap(z.bufInt32) < len(xAsCase) {
			z.bufInt32 = make([]int32, len(xAsCase))
		}
	case []int64:
		if cap(z.bufInt64) < len(xAsCase) {
			z.bufInt64 = make([]int64, len(xAsCase))
		}
	case []uint:
		if cap(z.bufUint) < len(xAsCase) {
			z.bufUint = make([]uint, len(xAsCase))
		}
	case []uint32:
		if cap(z.bufUint32Alpha) < len(xAsCase) {
			z.bufUint32Alpha = make([]uint32, len(xAsCase))
		}
	case []uint64:
		if cap(z.bufUint64Alpha) < len(xAsCase) {
			z.bufUint64Alpha = make([]uint64, len(xAsCase))
		}
	}
}

func (z *zSorter) Sort(x interface{}) error {
	z.prepBuffers(x)
	switch xAsCase := x.(type) {
	case []float32:
		zfloat32.SortBYOB(xAsCase, z.bufUint32Alpha, z.bufUint32Beta)
	case []float64:
		zfloat64.SortBYOB(xAsCase, z.bufUint64Alpha, z.bufUint64Beta)
	case []int:
		zint.SortBYOB(xAsCase, z.bufInt)
	case []int32:
		zint32.SortBYOB(xAsCase, z.bufInt32)
	case []int64:
		zint64.SortBYOB(xAsCase, z.bufInt64)
	case []uint:
		zuint.SortBYOB(xAsCase, z.bufUint)
	case []uint32:
		zuint32.SortBYOB(xAsCase, z.bufUint32Alpha)
	case []uint64:
		zuint64.SortBYOB(xAsCase, z.bufUint64Alpha)
	case sort.Interface:
		sort.Sort(xAsCase)
	default:
		return errors.New("type not supported")
	}
	return nil
}

func (z *zSorter) CopySort(x interface{}) (interface{}, error) {
	y := makeCopy(x)
	if y == nil {
		return x, errors.New("type not supported")
	}
	err := z.Sort(y)
	return y, err
}

func makeCopy(x interface{}) interface{} {
	switch xAsCase := x.(type) {
	case []float32:
		y := make([]float32, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []float64:
		y := make([]float64, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []int:
		y := make([]int, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []int32:
		y := make([]int32, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []int64:
		y := make([]int64, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []uint:
		y := make([]uint, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []uint32:
		y := make([]uint32, len(xAsCase))
		copy(y, xAsCase)
		return y
	case []uint64:
		y := make([]uint64, len(xAsCase))
		copy(y, xAsCase)
		return y
	default:
		return nil
	}
}

// Creates a Sorter that reuses buffers on repeated Sort() or CopySort() calls on the same type.
// This is not thread safe. CopySort() does not support passthrough of sort.Interface values.
func New() Sorter {
	return new(zSorter)
}
