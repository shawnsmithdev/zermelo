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
	testSize = 512
	// Const bit size thanks to kostya-sh@github
	bitSize uint = 1 << (5 + (^uint(0))>>32&1)
)

// numerical is a constraint that permits any integer or floating point type
type numerical interface {
	constraints.Integer | constraints.Float
}

func TestEmpty(t *testing.T) {
	testEmpty[int](t)
	testEmpty[float64](t)
}

func testEmpty[N numerical](t *testing.T) {
	if err := Sort([]N{}); err != nil {
		t.Fail()
	}
}

func TestSortSigned(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSort[int8](t, randInteger[int8]())
	testSort[int16](t, randInteger[int16]())
	testSort[int32](t, randInteger[int32]())
	testSort[int64](t, randInteger[int64]())
	testSort[int](t, randInteger[int]())
}

func TestSortUnsigned(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSort[uint8](t, randInteger[uint8]())
	testSort[uint16](t, randInteger[uint16]())
	testSort[uint32](t, randInteger[uint32]())
	testSort[uint64](t, randInteger[uint64]())
	testSort[uintptr](t, randInteger[uintptr]())
	testSort[uint](t, randInteger[uint]())
}

func TestSortFloats(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSort[float32](t, randFloat32())
	testSort[float64](t, randFloat64())
}

func testSort[N numerical](t *testing.T, rng func() N) {
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
	sliceCompare(toTest, control, t)
	t.Logf("[%T] gsort:%v, zsort:%v", toTest[0], gdelta, zdelta)
}

func sliceCompare[I comparable](toTest, control []I, t *testing.T) {
	if !slices.Equal(control, toTest) {
		t.Fatal(control, toTest)
	}
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

func fillSlice[T any](x []T, gen func() T) {
	for i := range x {
		x[i] = gen()
	}
}
