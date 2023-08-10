package floats

import (
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"golang.org/x/exp/constraints"
	"math"
	"slices"
	"sort"
	"testing"
)

const (
	testSize = 2 * compSortCutoffFloat64
)

func TestSort(t *testing.T) {
	testSort[float32](t, randFloat32(false), false, false)
	testSort[float64](t, randFloat64(false), false, false)
}

func TestSortNaNs(t *testing.T) {
	testSort[float32](t, randFloat32(true), false, true)
	testSort[float64](t, randFloat64(true), false, true)
}

func TestSortBYOB(t *testing.T) {
	testSort[float32](t, randFloat32(false), true, false)
	testSort[float32](t, randFloat32(false), true, false)
}

func TestSortNaNsBYOB(t *testing.T) {
	testSort[float32](t, randFloat32(true), true, true)
	testSort[float32](t, randFloat32(true), true, true)
}

func testSort[N constraints.Float](t *testing.T, rng func() N, byob, nans bool) {
	var buf []N
	if byob {
		buf = make([]N, testSize)
	}
	for i := 0; i < testSize; i++ {
		toTest := make([]N, i)
		internal.FillSlice(toTest, rng)
		control := slices.Clone(toTest)
		if nans {
			sortSort[N](control)
		} else {
			slices.Sort(control)
		}
		if byob {
			SortFloatsBYOB(toTest, buf[:i])
		} else {
			SortFloats(toTest)
		}
		if !floatSlicesEqual(control, toTest) {
			t.Fatal(control, toTest)
		}
	}
}

// returns a function that returns random float32s
func randFloat32(nans bool) func() float32 {
	return randFloat[float32, uint32](math.Float32frombits, nans)
}

// returns a function that returns random float64s
func randFloat64(nans bool) func() float64 {
	return randFloat[float64, uint64](math.Float64frombits, nans)
}

// randFloat returns a function that returns random floats
func randFloat[F constraints.Float, U constraints.Unsigned](fromBits func(U) F, nans bool) func() F {
	rng := internal.RandInteger[U]()
	return func() F {
		for {
			if result := fromBits(rng()); nans || !isNaN(result) {
				return result
			}
		}
	}
}

type sortable[F constraints.Float] []F

func (s sortable[F]) Len() int           { return len(s) }
func (s sortable[F]) Less(i, j int) bool { return s[i] < s[j] || (isNaN(s[i]) && !isNaN(s[j])) }
func (s sortable[F]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortSort[F constraints.Float](x []F) {
	sort.Sort(sortable[F](x))
}

func floatSlicesEqual[F constraints.Float](x, y []F) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] && !(isNaN(x[i]) && isNaN(y[i])) {
			return false
		}
	}
	return true
}
