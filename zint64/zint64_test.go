package zint64

import (
	"log"
	"math"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	test := [8]int64{3, -1000, -1, 100, 0, 999, math.MaxInt64, math.MinInt64}
	Sort(test[:])

	if !int64sAreSorted(test[:]) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", test)
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(data []int64) bool {
		Sort(data)
		return int64sAreSorted(data)
	}

	if err := quick.Check(test, nil); err != nil {
		t.Error(err)
	}
}

func TestSortEmpty(t *testing.T) {
	test := []int64{}
	Sort(test)
	if len(test) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}
}

func int64sAreSorted(data []int64) bool {
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
