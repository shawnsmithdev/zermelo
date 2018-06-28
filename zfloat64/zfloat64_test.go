package zfloat64

import (
	"log"
	"math"
	"sort"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	r := testData()
	SortBYOB(r, make([]float64, len(r)))
	if !sort.Float64sAreSorted(r) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", r)
		t.FailNow()
	}
}

func TestSortCopy(t *testing.T) {
	x := testData()
	y := SortCopy(x)

	if !sort.Float64sAreSorted(y) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", y)
		t.FailNow()
	}
	if x[0] != 3.1415 || x[10] != math.SmallestNonzeroFloat64 {
		log.Printf("Slice should have not have been modified.\n")
		log.Printf("Data was %v", y)
		t.FailNow()
	}
}

func TestSortEmpty(t *testing.T) {
	var r []float64
	Sort(r)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
	SortBYOB(r, make([]float64, len(r)))
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
	Sort(nil)
	SortBYOB(nil, nil)
}

func TestSortRand(t *testing.T) {
	test := func(r []float64) bool {
		SortBYOB(r, make([]float64, len(r)))
		return sort.Float64sAreSorted(r)
	}

	if err := quick.Check(test, nil); err != nil {
		t.Error(err)
	}
}

func testData() []float64 {
	return []float64{
		3.1415, -1000.0, -1.0, 100.5, 0, 999,
		math.NaN(), math.Inf(0), math.Inf(-1),
		math.MaxFloat64,
		math.SmallestNonzeroFloat64,
		-math.MaxFloat64,
		-math.SmallestNonzeroFloat64,
	}
}
