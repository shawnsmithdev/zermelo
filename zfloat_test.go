package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSortFloatsEmpty(t *testing.T) {
	SortFloats([]float64{})
}

func TestSortFloatsBYOBEmpty(t *testing.T) {
	SortFloatsBYOB([]float64{}, []float64{})
}

func TestSortFloats(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testSortFloats[float32](t, randFloat32(false))
	testSortFloats[float64](t, randFloat64(false))
}

func testSortFloats[F constraints.Float](t *testing.T, rng func() F) {
	toTest := make([]F, 2*compSortCutoffFloat64)
	fillSlice(toTest, rng)
	control := make([]F, len(toTest))
	copy(control, toTest)
	gstart := time.Now()
	slices.Sort(control)
	gdelta := time.Now().Sub(gstart)
	zstart := time.Now()
	SortFloats(toTest)
	zdelta := time.Now().Sub(zstart)
	if !slices.Equal(control, toTest) {
		t.Fatal(control, toTest)
	}
	checkSorted(t, toTest)
	t.Logf("[%T] gsort:%v, zsort:%v", toTest[0], gdelta, zdelta)
}

func TestSortFloatsWithNans(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	rng := randFloat64(true)
	toTest := make([]float64, 2*compSortCutoffFloat64)
	fillSlice(toTest, rng)
	SortFloats(toTest)
	checkSorted(t, toTest)
}

func checkSorted[F constraints.Float](t *testing.T, x []F) {
	nans := 0
	for i := 0; i < len(x); i++ {
		if math.IsNaN(float64(x[i])) {
			if i == nans {
				nans++
			} else {
				t.Fatal("unexpected NaN out of order", x)
			}
		}
	}
	if !slices.IsSorted(x[nans:]) {
		t.Fatal("should have sorted non-NaNs", x)
	}
}
