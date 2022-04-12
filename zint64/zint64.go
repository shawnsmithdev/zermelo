// Package zint64 implements radix sort for []int64.
// This package is deprecated
package zint64

import (
	"github.com/shawnsmithdev/zermelo"
	"golang.org/x/exp/slices"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	MinSize = 256
)

// Sort sorts x using a Radix sort (Small slices are sorted with slices.Sort() instead).
func Sort(x []int64) {
	if len(x) < MinSize {
		slices.Sort(x)
	} else {
		zermelo.SortIntegers(x)
	}
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []int64) []int64 {
	y := make([]int64, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts a []int64 using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []int64) {
	zermelo.SortIntegersBYOB(x, buffer)
}
