// Package zfloat64 implements radix sort for []float64.
// This package is deprecated.
package zfloat64

import (
	"github.com/shawnsmithdev/zermelo"
	"golang.org/x/exp/slices"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	// This is deprecated and no longer used
	MinSize = 384
)

// Sort sorts x.
// This is deprecated, use zermelo.SortFloats instead.
func Sort(x []float64) {
	zermelo.SortFloats(x)
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
// This is deprecated. Use slices.Clone and zermelo.SortFloats
func SortCopy(x []float64) []float64 {
	y := slices.Clone(x)
	zermelo.SortFloats(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
// This is deprecated, use zermelo.SortFloatsBYOB instead.
func SortBYOB(x, buffer []float64) {
	zermelo.SortFloatsBYOB(x, buffer)
}
