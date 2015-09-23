package zint

import (
	"log"
	"math"
	"reflect"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	test := [8]int{}
	if reflect.TypeOf(int(0)).Bits() == 32 {
		log.Printf("Testing int as 32 bit\n")
		test = [8]int{3, -1000, -1, 100, 0, 999, int(math.MaxInt32), int(math.MinInt32)}
	} else {
		log.Printf("Testing int as 64 bit\n")
		test = [8]int{3, -1000, -1, 100, 0, 999, int(math.MaxInt64), int(math.MinInt64)}
	}
	SortBYOB(test[:], make([]int, 8))

	if !intsAreSorted(test[:]) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", test)
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(data []int) bool {
		Sort(data)
		return intsAreSorted(data)
	}

	if err := quick.Check(test, nil); err != nil {
		t.Error(err)
	}
}

func TestSortEmpty(t *testing.T) {
	test := []int{}
	Sort(test)
	if len(test) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}
}

func intsAreSorted(data []int) bool {
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
