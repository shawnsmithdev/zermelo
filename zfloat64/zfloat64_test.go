package zfloat64

import (
	"log"
	"sort"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	r := []float64{3.1415, -1000.0, -1.0, 100.5, 0, 999}
	Sort(r)

	if !sort.Float64sAreSorted(r) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", r)
		t.FailNow()
	}
}

func TestSortEmpty(t *testing.T) {
	r := []float64{}
	Sort(r)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(r []float64) bool {
		Sort(r)
		return sort.Float64sAreSorted(r)
	}
	config := quick.Config{MaxCountScale: 100}

	if err := quick.Check(test, &config); err != nil {
		t.Error(err)
	}
}
