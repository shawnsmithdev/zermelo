package zuint

import (
	"log"
	"math"
	"reflect"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	test := [7]uint{}
	if reflect.TypeOf(int(0)).Bits() == 32 {
		log.Printf("Testing int as 32 bit\n")
		test = [7]uint{3, 1000, 1, 100, 0, 999, math.MaxInt32}
	} else {
		log.Printf("Testing int as 64 bit\n")
		test = [7]uint{3, 1000, 1, 100, 0, 999, math.MaxInt64}
	}
	Sort(test[:])

	if !uintsAreSorted(test[:]) {
		log.Printf("Should have sorted slice.\n")
		log.Printf("Data was %v", test)
		t.FailNow()
	}
}

func TestSortRand(t *testing.T) {
	test := func(data []uint) bool {
		buffer := make([]uint, len(data))
		SortBYOB(data, buffer)
		return uintsAreSorted(data)
	}
	config := quick.Config{MaxCountScale: 100}

	if err := quick.Check(test, &config); err != nil {
		t.Error(err)
	}
}

func TestSortEmpty(t *testing.T) {
	test := []uint{}
	Sort(test)
	if len(test) != 0 {
		log.Printf("Should still be empty\n")
		t.FailNow()
	}
}

func uintsAreSorted(data []uint) bool {
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
