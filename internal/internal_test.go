package internal

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
	"testing"
	"time"
)

const intSize uint = 1 << (5 + (^uint(0))>>32&1)

func TestDetect(t *testing.T) {
	testDetect[uint](t, intSize, 0)
	testDetect[uint8](t, 8, 0)
	testDetect[uint16](t, 16, 0)
	testDetect[uint32](t, 32, 0)
	testDetect[uint64](t, 64, 0)
	testDetect[int](t, intSize, math.MinInt)
	testDetect[int8](t, 8, math.MinInt8)
	testDetect[int16](t, 16, math.MinInt16)
	testDetect[int32](t, 32, math.MinInt32)
	testDetect[int64](t, 64, math.MinInt64)
}

func testDetect[I constraints.Integer](t *testing.T, size uint, min I) {
	start := time.Now()
	detectedSize, detectedMin := Detect[I]()
	delta := time.Now().Sub(start)
	if size != detectedSize {
		t.Fatalf("%T: Wrong size, expected %v, got %v", I(0), size, detectedSize)
	}
	if detectedMin != min {
		t.Fatalf("%T: Wrong min, expected %v, got %v", I(0), min, detectedMin)
	}
	t.Logf("%T: detect in %v", I(0), delta)
}

func TestFillSlice(t *testing.T) {
	testFillSlice(t, "test")
	testFillSlice(t, "foo")
	testFillSlice(t, "bar")
}

func testFillSlice(t *testing.T, toFill string) {
	test := make([]string, 128)
	n := 0
	FillSlice(test, func() string {
		result := fmt.Sprintf("%s%d", toFill, n)
		n++
		return result
	})
	for i, val := range test {
		if val != fmt.Sprintf("%s%d", toFill, i) {
			t.Fatal("wrong value filled")
		}
	}
}

func TestRandInteger(t *testing.T) {
	t.Log(time.Now()) // tests are cached and that can be confusing
	rand.Seed(time.Now().UnixNano())
	// just print results so it can be checked by a human
	randIntPrint(t, RandInteger[uint]()())
	randIntPrint(t, RandInteger[uint8]()())
	randIntPrint(t, RandInteger[uint16]()())
	randIntPrint(t, RandInteger[uint32]()())
	randIntPrint(t, RandInteger[uintptr]()())
	randIntPrint(t, RandInteger[int]()())
	randIntPrint(t, RandInteger[int8]()())
	randIntPrint(t, RandInteger[int16]()())
	randIntPrint(t, RandInteger[int32]()())
	randIntPrint(t, RandInteger[int64]()())
}

func randIntPrint(t *testing.T, x any) {
	t.Logf("RandInteger[%T] %v", x, x)
}
