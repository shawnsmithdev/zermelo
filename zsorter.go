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

// A Sorter can sort things like slices.
type Sorter interface {
	// Sort attempts to sort x, returning an error if unable to sort.
	Sort(x interface{}) error
	// CopySort returns a sorted copy of x, or an error if unable to copy or sort.
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

func (z *zSorter) sortFloat32(x []float32) {
	size := len(x)
	if z.useGoSort && size < zfloat32.MinSize {
		goSortFloat32(x)
		return
	}
	if len(z.bufFloat32) < size {
		z.bufFloat32 = make([]float32, allocSize(len(z.bufFloat32), size))
	}
	zfloat32.SortBYOB(x, z.bufFloat32)
}

func (z *zSorter) sortFloat64(x []float64)  {
	size := len(x)
	if z.useGoSort && size < zfloat64.MinSize {
		sort.Float64s(x)
		return
	}
	if len(z.bufFloat64) < size {
		z.bufFloat64 = make([]float64, allocSize(len(z.bufFloat64), size))
	}
	zfloat64.SortBYOB(x, z.bufFloat64)
}

func (z *zSorter) sortInt(x []int)  {
	size := len(x)
	if z.useGoSort && size < zint.MinSize {
		sort.Ints(x)
		return
	}
	if len(z.bufInt) < size {
		z.bufInt = make([]int, allocSize(len(z.bufInt), size))
	}
	zint.SortBYOB(x, z.bufInt)
}

func (z *zSorter) sortInt32(x []int32) {
	size := len(x)
	if z.useGoSort && size < zint32.MinSize {
		goSortInt32(x)
		return
	}
	if len(z.bufInt32) < size {
		z.bufInt32 = make([]int32, allocSize(len(z.bufInt32), size))
	}
	zint32.SortBYOB(x, z.bufInt32)
}

func (z *zSorter) sortInt64(x []int64) {
	size := len(x)
	if z.useGoSort && size < zint64.MinSize {
		goSortInt64(x)
		return
	}
	if len(z.bufInt64) < size {
		z.bufInt64 = make([]int64, allocSize(len(z.bufInt64), size))
	}
	zint64.SortBYOB(x, z.bufInt64)
}

func (z *zSorter) sortUint(x []uint) {
	size := len(x)
	if z.useGoSort && size < zuint.MinSize {
		goSortUint(x)
		return
	}
	if len(z.bufUint) < size {
		z.bufUint = make([]uint, allocSize(len(z.bufUint), size))
	}
	zuint.SortBYOB(x, z.bufUint)
}

func (z *zSorter) sortUint32(x []uint32) {
	size := len(x)
	if z.useGoSort && size < zuint32.MinSize {
		goSortUint32(x)
		return
	}
	if len(z.bufUint32) < size {
		z.bufUint32 = make([]uint32, allocSize(len(z.bufUint32), size))
	}
	zuint32.SortBYOB(x, z.bufUint32)
}

func (z *zSorter) sortUint64(x []uint64) {
	size := len(x)
	if z.useGoSort && size < zuint64.MinSize {
		goSortUint64(x)
		return
	}
	if len(z.bufUint64) < size {
		z.bufUint64 = make([]uint64, allocSize(len(z.bufUint64), size))
	}
	zuint64.SortBYOB(x, z.bufUint64)
}

func (z *zSorter) Sort(x interface{}) error {
	switch xAsCase := x.(type) {
	case []float32:
		z.sortFloat32(xAsCase)
	case []float64:
		z.sortFloat64(xAsCase)
	case []int:
		z.sortInt(xAsCase)
	case []int32:
		z.sortInt32(xAsCase)
	case []int64:
		z.sortInt64(xAsCase)
	case []uint:
		z.sortUint(xAsCase)
	case []uint32:
		z.sortUint32(xAsCase)
	case []uint64:
		z.sortUint64(xAsCase)
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
