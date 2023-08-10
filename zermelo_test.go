package zermelo

import (
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"golang.org/x/exp/constraints"
	"slices"
	"testing"
)

const (
	// Const int size thanks to kostya-sh@github
	intSize  uint = 1 << (5 + (^uint(0))>>32&1)
	testSize      = 2 * compSortCutoff64
)

func TestSort(t *testing.T) {
	testSort[int8](t, internal.RandInteger[int8](), false)
	testSort[int16](t, internal.RandInteger[int16](), false)
	testSort[int32](t, internal.RandInteger[int32](), false)
	testSort[int64](t, internal.RandInteger[int64](), false)
	testSort[int](t, internal.RandInteger[int](), false)
	testSort[uint8](t, internal.RandInteger[uint8](), false)
	testSort[uint16](t, internal.RandInteger[uint16](), false)
	testSort[uint32](t, internal.RandInteger[uint32](), false)
	testSort[uint64](t, internal.RandInteger[uint64](), false)
	testSort[uintptr](t, internal.RandInteger[uintptr](), false)
	testSort[uint](t, internal.RandInteger[uint](), false)
}

func TestSortBYOB(t *testing.T) {
	testSort[int8](t, internal.RandInteger[int8](), true)
	testSort[int16](t, internal.RandInteger[int16](), true)
	testSort[int32](t, internal.RandInteger[int32](), true)
	testSort[int64](t, internal.RandInteger[int64](), true)
	testSort[int](t, internal.RandInteger[int](), true)
	testSort[uint8](t, internal.RandInteger[uint8](), true)
	testSort[uint16](t, internal.RandInteger[uint16](), true)
	testSort[uint32](t, internal.RandInteger[uint32](), true)
	testSort[uint64](t, internal.RandInteger[uint64](), true)
	testSort[uintptr](t, internal.RandInteger[uintptr](), true)
	testSort[uint](t, internal.RandInteger[uint](), true)
}

func testSort[N constraints.Integer](t *testing.T, rng func() N, byob bool) {
	for i := 0; i <= testSize; i++ {
		toTest := make([]N, i)
		internal.FillSlice(toTest, rng)
		control := slices.Clone(toTest)
		slices.Sort(control)
		if byob {
			SortBYOB(toTest, make([]N, i))
		} else {
			Sort(toTest)
		}
		if !slices.Equal(control, toTest) {
			t.Fatal(control, toTest)
		}
	}
}
