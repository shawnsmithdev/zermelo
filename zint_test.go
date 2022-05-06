package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	// Const bit size thanks to kostya-sh@github
	bitSize uint = 1 << (5 + (^uint(0))>>32&1)
)

func TestSortIntegersBYOBEmpty(t *testing.T) {
	SortIntegersBYOB([]int{}, []int{})
}

func TestSortIntegers(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSortIntegers[int8](t, randInteger[int8]())
	testSortIntegers[int16](t, randInteger[int16]())
	testSortIntegers[int32](t, randInteger[int32]())
	testSortIntegers[int64](t, randInteger[int64]())
	testSortIntegers[int](t, randInteger[int]())
	testSortIntegers[uint8](t, randInteger[uint8]())
	testSortIntegers[uint16](t, randInteger[uint16]())
	testSortIntegers[uint32](t, randInteger[uint32]())
	testSortIntegers[uint64](t, randInteger[uint64]())
	testSortIntegers[uintptr](t, randInteger[uintptr]())
	testSortIntegers[uint](t, randInteger[uint]())
}

func testSortIntegers[N constraints.Integer](t *testing.T, rng func() N) {
	toTest := make([]N, 2*compSortCutoff64)
	fillSlice(toTest, rng)
	control := make([]N, len(toTest))
	copy(control, toTest)
	gstart := time.Now()
	slices.Sort(control)
	gdelta := time.Now().Sub(gstart)
	zstart := time.Now()
	SortIntegers(toTest)
	zdelta := time.Now().Sub(zstart)
	if !slices.Equal(control, toTest) {
		t.Fatal(control, toTest)
	}
	t.Logf("[%T] gsort:%v, zsort:%v", toTest[0], gdelta, zdelta)
}

func TestDetect(t *testing.T) {
	testDetect[uint](t, bitSize, 0)
	testDetect[uint8](t, 8, 0)
	testDetect[uint16](t, 16, 0)
	testDetect[uint32](t, 32, 0)
	testDetect[uint64](t, 64, 0)
	testDetect[int](t, bitSize, math.MinInt)
	testDetect[int8](t, 8, math.MinInt8)
	testDetect[int16](t, 16, math.MinInt16)
	testDetect[int32](t, 32, math.MinInt32)
	testDetect[int64](t, 64, math.MinInt64)
}

func testDetect[I constraints.Integer](t *testing.T, size uint, min I) {
	start := time.Now()
	detectedSize, detectedMin := detect[I]()
	delta := time.Now().Sub(start)
	if size != detectedSize {
		t.Fatalf("%T: Wrong size, expected %v, got %v", I(0), size, detectedSize)
	}
	if detectedMin != min {
		t.Fatalf("%T: Wrong min, expected %v, got %v", I(0), min, detectedMin)
	}
	t.Logf("%T: detect in %v", I(0), delta)
}
