package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	testGiveUpRace         = 2 * compSortCutoff64
	testRaceAttemptsAtSize = 32
	testOldSorterSize      = 1024
)

func TestUnsignedSorters(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testIntSorters(t, randInteger[uint]())
	testIntSorters(t, randInteger[uintptr]())
	testIntSorters(t, randInteger[uint64]())
	testIntSorters(t, randInteger[uint32]())
	testIntSorters(t, randInteger[uint16]())
	testIntSorters(t, randInteger[uint8]())
}

func TestSignedSorters(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testIntSorters(t, randInteger[int]())
	testIntSorters(t, randInteger[int64]())
	testIntSorters(t, randInteger[int32]())
	testIntSorters(t, randInteger[int16]())
	testIntSorters(t, randInteger[int8]())
}

func TestFloatSorters(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testFloatSorters[float64](t, randFloat64(false), false)
	testFloatSorters[float32](t, randFloat32(false), false)
	testFloatSorters[float64](t, randFloat64(true), true)
	testFloatSorters[float32](t, randFloat32(true), true)
}

func TestOldSorters(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testOldSorter[int](t, randInteger[int]())
	testOldSorter[int64](t, randInteger[int64]())
	testOldSorter[int32](t, randInteger[int32]())
	testOldSorter[int16](t, randInteger[int16]())
	testOldSorter[int8](t, randInteger[int8]())
	testOldSorter[uint](t, randInteger[uint]())
	testOldSorter[uintptr](t, randInteger[uintptr]())
	testOldSorter[uint64](t, randInteger[uint64]())
	testOldSorter[uint32](t, randInteger[uint32]())
	testOldSorter[uint16](t, randInteger[uint16]())
	testOldSorter[uint8](t, randInteger[uint8]())
	testOldSorter[float32](t, randFloat32(false))
	testOldSorter[float64](t, randFloat64(false))
}

func testOldSorter[T constraints.Ordered](t *testing.T, rgen func() T) {
	toTest := make([]T, testOldSorterSize)
	fillSlice(toTest, rgen)
	control := make([]T, testOldSorterSize)
	copy(control, toTest)
	testSorter := New()
	goStart := time.Now()
	slices.Sort(control)
	goDelta := time.Now().Sub(goStart)
	zStart := time.Now()
	_ = testSorter.Sort(toTest)
	zDelta := time.Now().Sub(zStart)
	for i := 0; i < len(control); i++ {
		if control[i] != toTest[i] {
			t.Fatalf("(%T) different result at i=%d, control=%v, test=%v", toTest, i, control[i], toTest[i])
		}
	}
	t.Logf("(%T) goDelta: %v, zDelta: %v", toTest, goDelta, zDelta)
	if !slices.IsSorted(toTest) {
		t.Fatalf("(%T) should have been sorted", toTest)
	}
	rand.Shuffle(len(toTest), func(i, j int) {
		toTest[i], toTest[j] = toTest[j], toTest[i]
	})
	if slices.IsSorted(toTest) {
		t.Fatalf("(%T) should have not been sorted", toTest)
	}
	copied, _ := testSorter.CopySort(toTest)
	if slices.IsSorted(toTest) {
		t.Fatalf("(%T) should have not been sorted", toTest)
	}
	if asType, ok := copied.([]T); ok {
		if !slices.IsSorted(asType) {
			t.Fatalf("(%T) should have been sorted", toTest)
		}
	} else {
		t.Fatalf("(%T) should have been able to cast", toTest)
	}
}

func testIntSorters[I constraints.Integer](t *testing.T, rgen func() I) {
	testSorter := newIntSorter[I]()
	testSorter.setCutoff(0) // prevent comparison sort cutoff for testing
	testSorters[I](t, rgen, testSorter, slices.Sort[I], "slices.Sort")
	testSorters[I](t, rgen, testSorter, sortSort[I], "sort.Sort")
}

func testFloatSorters[F constraints.Float](t *testing.T, rgen func() F, nans bool) {
	testSorter := newFloatSorter[F]()
	testSorter.setCutoff(0) // prevent comparison sort cutoff for testing
	if nans {
		testSorters[F](t, rgen, testSorter, nanSortSort[F], "sort.Sort")
	} else {
		testSorters[F](t, rgen, testSorter, slices.Sort[F], "slices.Sort")
	}
}

type nSorter[N constraints.Ordered] interface {
	Sort(x []N)
}

func testSorters[N constraints.Ordered](t *testing.T, rgen func() N, sorter nSorter[N], gsort func([]N), name string) {
	var toTest []N
	var attempts int
	for size := 3; size < testGiveUpRace; size++ {
		attempts = testRaceAttemptsAtSize
		toTest = make([]N, size)
		control := make([]N, size)
		for attempts > 0 {
			fillSlice(toTest, rgen)
			copy(control, toTest)
			gstart := time.Now()
			gsort(control)
			gdelta := time.Now().Sub(gstart)
			zstart := time.Now()
			sorter.Sort(toTest)
			zdelta := time.Now().Sub(zstart)
			for idx, val := range control {
				if val != toTest[idx] && !(isNaN(val) && isNaN(toTest[idx])) {
					t.Fatal(control, toTest)
				}
			}
			if !slices.IsSorted(toTest) {
				t.Fatal(control, toTest)
			}
			if zdelta < gdelta {
				attempts--
			} else {
				break
			}
		}
		if attempts == 0 {
			t.Logf("%T: Won %v in a row vs. %s at size %d", toTest, testRaceAttemptsAtSize, name, size)
			return
		}
	}
	if attempts != 0 {
		t.Logf("%T: Gave up racing vs. %s at size %d", toTest, name, testGiveUpRace)
	}
}

type sortable[N constraints.Ordered] []N

func (x sortable[N]) Len() int           { return len(x) }
func (x sortable[N]) Less(i, j int) bool { return x[i] < x[j] }
func (x sortable[N]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type nanSortable[N constraints.Ordered] []N

func (x nanSortable[N]) Len() int           { return len(x) }
func (x nanSortable[N]) Less(i, j int) bool { return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j])) }
func (x nanSortable[N]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func sortSort[T constraints.Ordered](x []T)    { sort.Sort(sortable[T](x)) }
func nanSortSort[T constraints.Ordered](x []T) { sort.Sort(nanSortable[T](x)) }

func fillSlice[T any](x []T, gen func() T) {
	for i := range x {
		x[i] = gen()
	}
}
