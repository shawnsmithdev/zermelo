package zermelo

import (
	"errors"
	"github.com/shawnsmithdev/zermelo/zfloat32"
	"github.com/shawnsmithdev/zermelo/zfloat64"
	"github.com/shawnsmithdev/zermelo/zint"
	"github.com/shawnsmithdev/zermelo/zint32"
	"github.com/shawnsmithdev/zermelo/zint64"
	"sort"
	"github.com/shawnsmithdev/zermelo/zuint"
	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
)

// A Sorter can sort things like slices. Returns an error on unsupported types.
type Sorter interface {
	Sort(x interface{}) error
	CopySort(x interface{}) (interface{}, error)
}

// Reuseable buffers
type zSorter struct {
	bufFloat32 []float32
	bufFloat64 []float64
	bufInt     []int
	bufInt32   []int32
	bufInt64   []int64
	bufUint    []uint
	bufUint32  []uint32
	bufUint64  []uint64
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

func (z *zSorter) prepFloat32(size int) {
	if cap(z.bufFloat32) < size {
		z.bufFloat32 = make([]float32, allocSize(cap(z.bufFloat32), size))
	}
}

func (z *zSorter) prepFloat64(size int) {
	if cap(z.bufFloat64) < size {
		z.bufFloat64 = make([]float64, allocSize(cap(z.bufFloat64), size))
	}
}

func (z *zSorter) prepInt(size int) {
	if cap(z.bufInt) < size {
		z.bufInt = make([]int, allocSize(cap(z.bufInt), size))
	}
}

func (z *zSorter) prepInt32(size int) {
	if cap(z.bufInt32) < size {
		z.bufInt32 = make([]int32, allocSize(cap(z.bufInt32), size))
	}
}

func (z *zSorter) prepInt64(size int) {
	if cap(z.bufInt64) < size {
		z.bufInt64 = make([]int64, allocSize(cap(z.bufInt64), size))
	}
}

func (z *zSorter) prepUint(size int) {
	if cap(z.bufUint) < size {
		z.bufUint = make([]uint, allocSize(cap(z.bufUint), size))
	}
}

func (z *zSorter) prepUint32(size int) {
	if cap(z.bufUint32) < size {
		z.bufUint32 = make([]uint32, allocSize(cap(z.bufUint32), size))
	}
}

func (z *zSorter) prepUint64(size int) {
	if cap(z.bufUint64) < size {
		z.bufUint64 = make([]uint64, allocSize(cap(z.bufUint64), size))
	}
}

// Checks that buffers are large enough.
func (z *zSorter) prepBuffers(x interface{}) {
	switch xAsCase := x.(type) {
	case []float32:
		z.prepFloat32(len(xAsCase))
	case []float64:
		z.prepFloat64(len(xAsCase))
	case []int:
		z.prepInt(len(xAsCase))
	case []int32:
		z.prepInt32(len(xAsCase))
	case []int64:
		z.prepInt64(len(xAsCase))
	case []uint:
		z.prepUint(len(xAsCase))
	case []uint32:
		z.prepUint32(len(xAsCase))
	case []uint64:
		z.prepUint64(len(xAsCase))
	}
}

func (z *zSorter) Sort(x interface{}) error {
	z.prepBuffers(x)
	switch xAsCase := x.(type) {
	case []float32:
		zfloat32.SortBYOB(xAsCase, z.bufFloat32[:len(xAsCase)])
	case []float64:
		zfloat64.SortBYOB(xAsCase, z.bufFloat64[:len(xAsCase)])
	case []int:
		zint.SortBYOB(xAsCase, z.bufInt[:len(xAsCase)])
	case []int32:
		zint32.SortBYOB(xAsCase, z.bufInt32[:len(xAsCase)])
	case []int64:
		zint64.SortBYOB(xAsCase, z.bufInt64[:len(xAsCase)])
	case []string:
		sort.Strings(xAsCase)
	case []uint:
		zuint.SortBYOB(xAsCase, z.bufUint[:len(xAsCase)])
	case []uint32:
		zuint32.SortBYOB(xAsCase, z.bufUint32[:len(xAsCase)])
	case []uint64:
		zuint64.SortBYOB(xAsCase, z.bufUint64[:len(xAsCase)])
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
	case []string:
		y := make([]string, len(xAsCase))
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

// New creates a Sorter that reuses buffers on repeated Sort() or CopySort() calls on the same type.
// This is not thread safe. CopySort() does not support passthrough of sort.Interface values.
func New() Sorter {
	return new(zSorter)
}
