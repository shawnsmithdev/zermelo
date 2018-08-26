// Package zermelo is a library for sorting slices in Go.
package zermelo // import "github.com/shawnsmithdev/zermelo"

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

// Sort attempts to sort x.
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
	case []string:
		sort.Strings(xAsCase)
	case sort.Interface:
		sort.Sort(xAsCase)
	default:
		return errors.New("type not supported")
	}
	return nil
}
