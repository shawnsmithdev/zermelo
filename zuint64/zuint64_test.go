package zuint64

import (
	"log"
	"math"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	test := [7]uint64{3, 1000, 1, 100, 0, 999, math.MaxUint64}
	Sort(test[:])

	if !uint64sAreSorted(test[:]) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", test)
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(data []uint64) bool {
		buffer := make([]uint64, len(data))
		SortBYOB(data, buffer)
		return uint64sAreSorted(data)
	}
	config := quick.Config{MaxCountScale: 100}

	if err := quick.Check(test, &config); err != nil {
		t.Error(err)
	}
}

func TestSortEmpty(t *testing.T) {
	test := []uint64{}
	Sort(test)
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
			log.Printf("Value at index %v (%v) was less than at index %v (%v)\n",
				idx, x, idx-1, data[idx-1])
			return false
		}
	}
	return true
}
