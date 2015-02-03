package zfloat32

import (
	"sort"
)

// Sorts a []float32 using a Radix sort. (eventually)
func Sort(r []float32) {
	sort.Sort(float32Sortable(r)) // TODO
}

type float32Sortable []float32

func (r float32Sortable) Len() int           { return len(r) }
func (r float32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r float32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
