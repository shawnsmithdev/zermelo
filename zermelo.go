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

// Attempts to sort x. If x is a supported slice type, this library will be
// be used to sort it. Otherwise, this attempts to sort x using sort.Sort().
// If x is not a supported type, and doesn't implement sort.Interface, an error is returned
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

type Sorter interface {
	Sort(x interface{})
}

type CopySorter interface {
	Sorter
	CopySort(x interface{}) interface{}
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

func (z *zSorter) Sort(x interface{}) {
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
	}
}

// A Sorter that reuses buffers on repeated Sort() calls on the same type. Not thread safe.
func New() Sorter {
	return new(zSorter)
}
