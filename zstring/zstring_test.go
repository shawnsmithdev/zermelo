package zstring

import (
	"sort"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	for _, test := range testDatas {
		x := test
		if sort.StringsAreSorted(x) {
			t.FailNow()
		}
		Sort(x)
		if !sort.StringsAreSorted(x) {
			t.FailNow()
		}
	}
}

func TestSortRand(t *testing.T) {
	test := func(data []string) bool {
		buffer := make([]string, len(data))
		SortBYOB(data, buffer)
		return sort.StringsAreSorted(data)
	}

	if err := quick.Check(test, nil); err != nil {
		t.Error(err)
	}
}

var testDatas = [][]string{
	{"AAAA", "BBBB", "ABCD", "DDDD", "ZAAA", "世界"},
	{"coffee", "sandwich", "banana", "pizza", "apple", "soda"},
	{"coffee", "sandwich", "banana", "apple", "soda"},
}
