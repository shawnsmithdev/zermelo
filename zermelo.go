// Package zermelo is a library for sorting slices in Go.
package zermelo // import "github.com/shawnsmithdev/zermelo"

import (
	"errors"
	"sort"
)

// Sort attempts to sort x.
//
// If x is a supported slice type, this library will be used to sort it. Otherwise,
// if x implements sort.Interface it will passthrough to the sort.Sort() algorithm.
// Returns an error on unsupported types.
func Sort(x any) error {
	switch xAsCase := x.(type) {
	case []float32:
		SortFloats(xAsCase)
	case []float64:
		SortFloats(xAsCase)
	case []int:
		SortIntegers(xAsCase)
	case []int8:
		SortIntegers(xAsCase)
	case []int16:
		SortIntegers(xAsCase)
	case []int32:
		SortIntegers(xAsCase)
	case []int64:
		SortIntegers(xAsCase)
	case []uint:
		SortIntegers(xAsCase)
	case []uint8:
		SortIntegers(xAsCase)
	case []uint16:
		SortIntegers(xAsCase)
	case []uint32:
		SortIntegers(xAsCase)
	case []uint64:
		SortIntegers(xAsCase)
	case []uintptr:
		SortIntegers(xAsCase)
	case []string:
		sort.Strings(xAsCase)
	case sort.Interface:
		sort.Sort(xAsCase)
	default:
		return errors.New("type not supported")
	}
	return nil
}
