// Package zuint64 implements radix sort for []uint64.
// This package is deprecated.
package zuint64

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
func Sort(x []uint64) {
	zermelo.SortIntegers(x)
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
// This is deprecated. Use slices.Clone and zermelo.SortIntegers.
func SortCopy(x []uint64) []uint64 {
	y := slices.Clone(x)
	zermelo.SortIntegers(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
// This is deprecated, use zermelo.SortIntegersBYOB instead.
func SortBYOB(x, buffer []uint64) {
	zermelo.SortIntegersBYOB(x, buffer)
}
