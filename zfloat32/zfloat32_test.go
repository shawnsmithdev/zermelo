package zfloat32

import (
	"log"
	"math"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	r := testData()
	b1 := make([]uint32, len(r))
	b2 := make([]uint32, len(r))

	SortBYOB(r, b1, b2)
	if !isSorted(r) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", r)
		t.FailNow()
	}
}

func TestSortCopy(t *testing.T) {
	x := testData()
	y := SortCopy(x)

	if !isSorted(y) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", y)
		t.FailNow()
	}
	if x[0] != 3.1415 || x[10] != math.SmallestNonzeroFloat32 {
		log.Printf("Slice should have not have been modified.\n")
		log.Printf("Data was %v", y)
		t.FailNow()
	}
}

func TestSortEmpty(t *testing.T) {
	r := []float32{}
	Sort(r)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
	b1 := make([]uint32, len(r))
	b2 := make([]uint32, len(r))
	SortBYOB(r, b1, b2)
	if len(r) != 0 {
		log.Printf("Should have been empty\n")
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(r []float32) bool {
		b1 := make([]uint32, len(r))
		b2 := make([]uint32, len(r))
		SortBYOB(r, b1, b2)
		return isSorted(r)
	}
	config := quick.Config{MaxCountScale: 100}

	if err := quick.Check(test, &config); err != nil {
		t.Error(err)
	}
}

func testData() []float32 {
	return []float32{
		3.1415, -1000.0, -1.0, 100.5, 0, 999,
		float32(math.NaN()),
		float32(math.Inf(0)),
		float32(math.Inf(-1)),
		math.MaxFloat32,
		math.SmallestNonzeroFloat32,
		-math.MaxFloat32,
		-math.SmallestNonzeroFloat32,
	}
}

func isSorted(x []float32) bool {
	for idx, val := range x {
		if idx > 0 && val < x[idx-1] {
			return false
		}
	}
	return true
}
