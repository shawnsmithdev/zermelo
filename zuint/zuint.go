// Package zuint implements radix sort for []uint.
// This package is deprecated.
package zuint

import (
	"github.com/shawnsmithdev/zermelo"
	"golang.org/x/exp/slices"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	// This is deprecated and no longer used
	MinSize = 256
)

// Sort sorts x.
// This is deprecated, use zermelo.SortIntegers instead.
func Sort(x []uint) {
	zermelo.SortIntegers(x)
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
// This is deprecated. Use slices.Clone and zermelo.SortIntegers.
func SortCopy(x []uint) []uint {
	y := slices.Clone(x)
	zermelo.SortIntegers(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
// This is deprecated, use zermelo.SortIntegersBYOB instead.
func SortBYOB(x, buffer []uint) {
	zermelo.SortIntegersBYOB(x, buffer)
}
