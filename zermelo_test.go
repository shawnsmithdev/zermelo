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
	if !slices.Equal(control, toTest) {
		t.Fatal(control, toTest)
	}
	t.Logf("[%T] gsort:%v, zsort:%v", toTest[0], gdelta, zdelta)
}
