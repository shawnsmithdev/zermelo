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
	b1 := make([]uint64, len(r))
	b2 := make([]uint64, len(r))

	SortBYOB(r, b1, b2)
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
	r := []float64{}
	Sort(r)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
	b1 := make([]uint64, len(r))
	b2 := make([]uint64, len(r))
	SortBYOB(r, b1, b2)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(r []float64) bool {
		b1 := make([]uint64, len(r))
		b2 := make([]uint64, len(r))
		SortBYOB(r, b1, b2)
		return sort.Float64sAreSorted(r)
	}
	config := quick.Config{MaxCountScale: 100}

	if err := quick.Check(test, &config); err != nil {
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
