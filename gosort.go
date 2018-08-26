package zermelo

import (
	"errors"
	"sort"
)

func goSortFloat32(x []float32) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

func goSortInt32(x []int32) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

func goSortInt64(x []int64) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

func goSortUint(x []uint) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

func goSortUint32(x []uint32) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}
func goSortUint64(x []uint64) {
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
}

// Attempts to run the best the sort package has to offer for the given type
func goSorter(x interface{}) error {
	switch xAsCase := x.(type) {
	case []float64:
		sort.Float64s(xAsCase)
	case []int:
		sort.Ints(xAsCase)
	case []string:
		sort.Strings(xAsCase)

	case []float32:
		goSortFloat32(xAsCase)
	case []int32:
		goSortInt32(xAsCase)
	case []int64:
		goSortInt64(xAsCase)
	case []uint:
		goSortUint(xAsCase)
	case []uint32:
		goSortUint32(xAsCase)
	case []uint64:
		goSortUint64(xAsCase)
	default:
		return errors.New("type not supported")
	}
	return nil
}
