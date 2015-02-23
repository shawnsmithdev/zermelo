package zuint64

import (
	"log"
	"math"
	"testing"
	"testing/quick"
)

func TestSortBYOB(t *testing.T) {
	test := []uint64{3, 1000, 1, 100, 0, 999, math.MaxUint64}
	b := make([]uint64, len(test))
	SortBYOB(test, b)

	if !uint64sAreSorted(test) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", test)
		t.FailNow()
	}
}

func genTestData(size uint) []uint64 {
	r := make([]uint64, size)
	for i := range r {
		r[i] = uint64(len(r) - i)
	}
	return r
}

func runSortTest(t *testing.T, size uint) {
	x := genTestData(size)
	Sort(x)
	if !uint64sAreSorted(x) {
		log.Printf("Should have sorted slice with len=%v\n", len(x))
		t.FailNow()
	}
}

func TestSortSmall(t *testing.T) {
	runSortTest(t, MinSize-1)
}

func TestSortBig(t *testing.T) {
	runSortTest(t, MinSize)
}

func runSortCopyTest(t *testing.T, size uint) {
	x := genTestData(size)
	if uint64sAreSorted(x) {
		log.Printf("Should NOT have sorted data in generated slice.\n")
		log.Printf("Data was %v", x)
		t.FailNow()
	}
	c := SortCopy(x)
	if !uint64sAreSorted(c) {
		log.Printf("Should have sorted copied slice.\n")
		log.Printf("Data was %v", c)
		t.FailNow()
	}
	if uint64sAreSorted(x) {
		log.Printf("Should NOT have sorted original slice.\n")
		log.Printf("Data was %v", x)
		t.FailNow()
	}
}

func TestSortCopySmall(t *testing.T) {
	runSortCopyTest(t, MinSize-1)
}

func TestSortCopyBig(t *testing.T) {
	runSortCopyTest(t, MinSize)
}

func TestSortRand(t *testing.T) {
	test := func(data []uint64) bool {
		buffer := make([]uint64, len(data))
		SortBYOB(data, buffer)
		return uint64sAreSorted(data)
	}

	if err := quick.Check(test, nil); err != nil {
		t.Error(err)
	}
}

func TestEmpty(t *testing.T) {
	test := []uint64{}
	Sort(test)
	if len(test) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}
	x := SortCopy(test)
	if len(test) != 0 || len(x) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}
	SortBYOB(test, test)
	if len(test) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}

}

func uint64sAreSorted(data []uint64) bool {
	for idx, x := range data {
		if idx == 0 {
			continue
		}
		if x < data[idx-1] {
			return false
		}
	}
	return true
}
