// Package zuint64 implements radix sort for []uint64.
// This package is deprecated
package zuint64

import (
	"github.com/shawnsmithdev/zermelo"
	"golang.org/x/exp/slices"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize = 256
)

// Sort Sort x using a Radix sort (Small slices are sorted with slices.Sort() instead).
func Sort(x []uint64) {
	if len(x) < MinSize {
		slices.Sort(x)
	} else {
		zermelo.SortIntegers(x)
	}
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []uint64) []uint64 {
	y := make([]uint64, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts x using a Radix sort, using supplied buffer space. Panics if
// len(x) is greater than len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []uint64) {
	zermelo.SortIntegersBYOB(x, buffer)
}
