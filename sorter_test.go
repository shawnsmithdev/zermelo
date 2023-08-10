package zermelo

import (
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"golang.org/x/exp/constraints"
	"slices"
	"testing"
)

func TestSorter(t *testing.T) {
	testSorter[int8](t, internal.RandInteger[int8](), false)
	testSorter[int8](t, internal.RandInteger[int8](), true)
	testSorter[int16](t, internal.RandInteger[int16](), false)
	testSorter[int16](t, internal.RandInteger[int16](), true)
	testSorter[int32](t, internal.RandInteger[int32](), false)
	testSorter[int32](t, internal.RandInteger[int32](), true)
	testSorter[int64](t, internal.RandInteger[int64](), false)
	testSorter[int64](t, internal.RandInteger[int64](), true)
	testSorter[int](t, internal.RandInteger[int](), false)
	testSorter[int](t, internal.RandInteger[int](), true)
	testSorter[uint8](t, internal.RandInteger[uint8](), false)
	testSorter[uint8](t, internal.RandInteger[uint8](), true)
	testSorter[byte](t, internal.RandInteger[byte](), false)
	testSorter[byte](t, internal.RandInteger[byte](), true)
	testSorter[uint16](t, internal.RandInteger[uint16](), false)
	testSorter[uint16](t, internal.RandInteger[uint16](), true)
	testSorter[uint32](t, internal.RandInteger[uint32](), false)
	testSorter[uint32](t, internal.RandInteger[uint32](), true)
	testSorter[uint64](t, internal.RandInteger[uint64](), false)
	testSorter[uint64](t, internal.RandInteger[uint64](), true)
	testSorter[uintptr](t, internal.RandInteger[uintptr](), false)
	testSorter[uintptr](t, internal.RandInteger[uintptr](), true)
	testSorter[uint](t, internal.RandInteger[uint](), false)
	testSorter[uint](t, internal.RandInteger[uint](), true)
}

func testSorter[I constraints.Integer](t *testing.T, gen func() I, cutoff bool) {
	var test Sorter[I]
	if cutoff {
		test = NewSorter[I]()
	} else {
		test = newSorter[I]().withCutoff(0)
	}
	toTest := make([]I, testSize)
	for i := 0; i < testSize; i++ {
		internal.FillSlice(toTest[:i], gen)
		control := slices.Clone(toTest[:i])
		slices.Sort(control)
		copied := slices.Clone(toTest[:i])
		test.Sort(copied)
		if !slices.Equal(copied, control) {
			t.Fatal("cutoff=", cutoff, ", copied != control", copied, control)
		}
		test.Sort(toTest[:i])
		if !slices.Equal(toTest[:i], control) {
			t.Fatal("cutoff=", cutoff, "toTest ! control", toTest, control)
		}
	}
}
