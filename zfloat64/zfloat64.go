package zfloat64

import (
	"sort"
)

// Sorts a []float64 using a Radix sort. (eventually)
func Sort(r []float64) {
	sort.Sort(sort.Float64Slice(r)) // TODO
}
