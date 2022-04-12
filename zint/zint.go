// Package zint implements radix sort for []int.
// This package is deprecated
package zint

import (
	"github.com/shawnsmithdev/zermelo"
)

const (
	// MinSize is the minimum size of a slice that will be radix sorted by Sort.
	// This is deprecated and no longer used
	MinSize = 256
)

// Sort sorts x using a Radix sort (Small slices are sorted with slices.Sort() instead).
func Sort(x []int) {
	zermelo.SortIntegers(x)
}

// SortCopy is similar to Sort, but returns a sorted copy of x, leaving x unmodified.
func SortCopy(x []int) []int {
	y := make([]int, len(x))
	copy(y, x)
	Sort(y)
	return y
}

// SortBYOB sorts a []int using a Radix sort, using supplied buffer space. Panics if
// len(x) does not equal len(buffer). Uses radix sort even on small slices.
func SortBYOB(x, buffer []int) {
	zermelo.SortIntegersBYOB(x, buffer)
}
