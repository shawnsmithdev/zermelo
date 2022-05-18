package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math/rand"
	"testing"
	"time"
)

const (
	testSize = 512
)

func TestSortEmpty(t *testing.T) {
	testSortEmpty[int8](t)
	testSortEmpty[int16](t)
	testSortEmpty[int32](t)
	testSortEmpty[int64](t)
	testSortEmpty[int](t)
	testSortEmpty[uint8](t)
	testSortEmpty[uint16](t)
	testSortEmpty[uint32](t)
	testSortEmpty[uint64](t)
	testSortEmpty[uintptr](t)
	testSortEmpty[uint](t)
	testSortEmpty[float32](t)
	testSortEmpty[float64](t)
}

func testSortEmpty[N any](t *testing.T) {
	if err := Sort([]N{}); err != nil {
		t.Fail()
	}
}

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSort[int8](t, randInteger[int8]())
	testSort[int16](t, randInteger[int16]())
	testSort[int32](t, randInteger[int32]())
	testSort[int64](t, randInteger[int64]())
	testSort[int](t, randInteger[int]())
	testSort[uint8](t, randInteger[uint8]())
	testSort[uint16](t, randInteger[uint16]())
	testSort[uint32](t, randInteger[uint32]())
	testSort[uint64](t, randInteger[uint64]())
	testSort[uintptr](t, randInteger[uintptr]())
	testSort[uint](t, randInteger[uint]())
	testSort[float32](t, randFloat32(false))
	testSort[float64](t, randFloat64(false))
}

func testSort[N constraints.Ordered](t *testing.T, rng func() N) {
	toTest := make([]N, testSize)
	fillSlice(toTest, rng)
	control := make([]N, len(toTest))
	copy(control, toTest)
	gstart := time.Now()
	slices.Sort(control)
	gdelta := time.Now().Sub(gstart)
	zstart := time.Now()
	if err := Sort(toTest); err != nil {
		t.Fatal(err)
	}
	zdelta := time.Now().Sub(zstart)
	if !slices.Equal(control, toTest) {
		t.Fatal(control, toTest)
	}
	t.Logf("[%T] gsort:%v, zsort:%v", toTest[0], gdelta, zdelta)
}
