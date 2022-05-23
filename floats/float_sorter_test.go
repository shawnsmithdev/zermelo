package floats

import (
	"github.com/shawnsmithdev/zermelo/v2"
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSorter(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	if !testSorter[float32](randFloat32(false), false, false) {
		t.Fatal("failed float32")
	}
	if !testSorter[float64](randFloat64(false), false, false) {
		t.Fatal("failed float64")
	}
	if !testSorter[float32](randFloat32(false), false, true) {
		t.Fatal("failed float32 cutoff")
	}
	if !testSorter[float64](randFloat64(false), false, true) {
		t.Fatal("failed float64 cutoff")
	}
}

func TestSorterNaNs(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	if !testSorter[float32](randFloat32(true), true, false) {
		t.Fatal("failed float32 nans")
	}
	if !testSorter[float64](randFloat64(true), true, false) {
		t.Fatal("failed float64 nans")
	}
	if !testSorter[float32](randFloat32(true), true, true) {
		t.Fatal("failed float32 nans cutoff")
	}
	if !testSorter[float64](randFloat64(true), true, true) {
		t.Fatal("failed float64 nans cutoff")
	}
}
func TestSorterOnlyNaNs(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	if !testSorter[float32](func() float32 { return float32(math.NaN()) }, true, false) {
		t.Fatal("failed float32 onlynans")
	}
	if !testSorter[float64](math.NaN, true, false) {
		t.Fatal("failed float64 onlynans")
	}
	if !testSorter[float32](func() float32 { return float32(math.NaN()) }, true, true) {
		t.Fatal("failed float32 onlynans cutoff")
	}
	if !testSorter[float64](math.NaN, true, true) {
		t.Fatal("failed float64 onlynans cutoff")
	}
}

func testSorter[F constraints.Float](gen func() F, nans bool, cutoff bool) bool {
	var test zermelo.Sorter[F]
	if cutoff {
		test = NewFloatSorter[F]()
	} else {
		test = newFloatSorter[F]().withCutoff(0)
	}
	toTest := make([]F, testSize)
	for i := 0; i < testSize; i++ {
		internal.FillSlice(toTest[:i], gen)
		control := slices.Clone(toTest[:i])
		if nans {
			sortSort[F](control)
		} else {
			slices.Sort(control)
		}
		copied := slices.Clone(toTest[:i])
		test.Sort(copied)
		if !floatSlicesEqual[F](copied, control) {
			return false
		}
		test.Sort(toTest[:i])
		if !floatSlicesEqual[F](toTest[:i], control) {
			return false
		}
	}
	return true
}
