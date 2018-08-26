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
	// if true, use go stdlib sort on small slices
	useGoSort bool
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

func (z *zSorter) prepFloat32(size int) []float32 {
	if cap(z.bufFloat32) < size {
		z.bufFloat32 = make([]float32, allocSize(cap(z.bufFloat32), size))
	}
	return z.bufFloat32
}

func (z *zSorter) prepFloat64(size int) []float64 {
	if cap(z.bufFloat64) < size {
		z.bufFloat64 = make([]float64, allocSize(cap(z.bufFloat64), size))
	}
	return z.bufFloat64
}

func (z *zSorter) prepInt(size int) []int {
	if cap(z.bufInt) < size {
		z.bufInt = make([]int, allocSize(cap(z.bufInt), size))
	}
	return z.bufInt
}

func (z *zSorter) prepInt32(size int) []int32 {
	if cap(z.bufInt32) < size {
		z.bufInt32 = make([]int32, allocSize(cap(z.bufInt32), size))
	}
	return z.bufInt32
}

func (z *zSorter) prepInt64(size int) []int64 {
	if cap(z.bufInt64) < size {
		z.bufInt64 = make([]int64, allocSize(cap(z.bufInt64), size))
	}
	return z.bufInt64
}

func (z *zSorter) prepUint(size int) []uint {
	if cap(z.bufUint) < size {
		z.bufUint = make([]uint, allocSize(cap(z.bufUint), size))
	}
	return z.bufUint
}

func (z *zSorter) prepUint32(size int) []uint32 {
	if cap(z.bufUint32) < size {
		z.bufUint32 = make([]uint32, allocSize(cap(z.bufUint32), size))
	}
	return z.bufUint32
}

func (z *zSorter) prepUint64(size int) []uint64 {
	if cap(z.bufUint64) < size {
		z.bufUint64 = make([]uint64, allocSize(cap(z.bufUint64), size))
	}
	return z.bufUint64
}

func (z *zSorter) Sort(x interface{}) error {
	switch xAsCase := x.(type) {
	case []float32:
		if z.useGoSort && len(xAsCase) < zfloat32.MinSize {
			goSortFloat32(xAsCase)
		} else {
			zfloat32.SortBYOB(xAsCase, z.prepFloat32(len(xAsCase)))
		}
	case []float64:
		if z.useGoSort && len(xAsCase) < zfloat64.MinSize {
			sort.Float64s(xAsCase)
		} else {
			zfloat64.SortBYOB(xAsCase, z.prepFloat64(len(xAsCase)))
		}
	case []int:
		if z.useGoSort && len(xAsCase) < zint.MinSize {
			sort.Ints(xAsCase)
		} else {
			zint.SortBYOB(xAsCase, z.prepInt(len(xAsCase)))
		}
	case []int32:
		if z.useGoSort && len(xAsCase) < zint32.MinSize {
			goSortInt32(xAsCase)
		} else {
			zint32.SortBYOB(xAsCase, z.prepInt32(len(xAsCase)))
		}
	case []int64:
		if z.useGoSort && len(xAsCase) < zint64.MinSize {
			goSortInt64(xAsCase)
		} else {
			zint64.SortBYOB(xAsCase, z.prepInt64(len(xAsCase)))
		}
	case []uint:
		if z.useGoSort && len(xAsCase) < zuint.MinSize {
			goSortUint(xAsCase)
		} else {
			zuint.SortBYOB(xAsCase, z.prepUint(len(xAsCase)))
		}
	case []uint32:
		if z.useGoSort && len(xAsCase) < zuint32.MinSize {
			goSortUint32(xAsCase)
		} else {
			zuint32.SortBYOB(xAsCase, z.prepUint32(len(xAsCase)))
		}
	case []uint64:
		if z.useGoSort && len(xAsCase) < zuint64.MinSize {
			goSortUint64(xAsCase)
		} else {
			zuint64.SortBYOB(xAsCase, z.prepUint64(len(xAsCase)))
		}
	case []string:
		sort.Strings(xAsCase)
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
	return &zSorter{useGoSort: true}
}

// Same as New(), but will not uses go sort on small slices
func newRawSorter() Sorter {
	return &zSorter{useGoSort: false}
}
