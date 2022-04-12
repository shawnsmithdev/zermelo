// Package zfloat32 implements radix sort for []float32.
// This package is deprecated
package zfloat32

import (
	"github.com/shawnsmithdev/zermelo"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	// This is deprecated and no longer used
	MinSize = 256
)

// Sort sorts x
func Sort(x []float32) {
	zermelo.SortFloats(x)
}

// SortCopy is similar to Sort(), but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []float32) []float32 {
	y := make([]float32, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []float32) {
	zermelo.SortFloatsBYOB(x, buffer)
}
