package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"math/rand"
	"testing"
	"time"
)

//TODO: This needs some dry'ing

const (
	testRequireWinRace     = false // racing can fail due to instrumenting, low precision os clocks, etc
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
	testFloatSorters[float64](t, randFloat64())
	testFloatSorters[float32](t, randFloat32())
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
	testOldSorter[float32](t, randFloat32())
	testOldSorter[float64](t, randFloat64())
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
	rand.Seed(time.Now().UnixNano())
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
	var toTest []I
	testSorter := newIntSorter[I]()

	// prevent comparison sort cutoff for testing
	testSorter.setCutoff(0)

	var attempts int
	for size := 3; size < testGiveUpRace; size++ {
		attempts = testRaceAttemptsAtSize
		toTest = make([]I, size)
		for attempts > 0 {
			fillSlice(toTest, rgen)
			if testIntSort[I](t, toTest, testSorter) {
				attempts--
			} else {
				break
			}
		}
		if attempts == 0 {
			t.Logf("%T: Won %v in a row at size %d", toTest, testRaceAttemptsAtSize, size)
			return
		}
	}
	if attempts != 0 {
		if testRequireWinRace {
			t.Fatalf("[%T] attempts: %d, size: %d", toTest, attempts, len(toTest))
		} else {
			t.Logf("%T: Gave up racing at size %d", toTest, testGiveUpRace)
		}
	}
}

func testIntSort[T constraints.Integer](t *testing.T, toTest []T, zsort IntSorter[T]) bool {
	control := make([]T, len(toTest))
	copy(control, toTest)
	gstart := time.Now()
	slices.Sort(control)
	gdelta := time.Now().Sub(gstart)
	zstart := time.Now()
	zsort.Sort(toTest)
	zdelta := time.Now().Sub(zstart)
	for idx, val := range control {
		if val != toTest[idx] {
			t.Fatal(control, toTest)
		}
	}
	if !slices.IsSorted(toTest) {
		t.Fatal(control, toTest)
	}
	return zdelta < gdelta
}

func testFloatSorters[F constraints.Float](t *testing.T, rgen func() F) {
	var toTest []F
	testSorter := newFloatSorter[F]()
	// prevent comparison sort cutoff for testing
	testSorter.setCutoff(0)

	for size := 0; size < testGiveUpRace; size++ {
		toTest = make([]F, size)
		attempts := testRaceAttemptsAtSize
		for attempts > 0 {
			fillSlice(toTest, rgen)
			if testFloatSorter[F](t, toTest, testSorter) && size > 2 {
				attempts--
			} else {
				attempts = -1
			}
			if !slices.IsSorted(toTest) {
				t.Fatal(toTest)
			}
		}
		if attempts == 0 {
			t.Logf("%T: Won %v in a row at size %d", toTest, testRaceAttemptsAtSize, size)
			return
		}
	}
	if testRequireWinRace {
		t.Fail()
	} else {
		t.Logf("%T: Gave up racing at size %d", toTest, testGiveUpRace)
	}
}

func testFloatSorter[T constraints.Float](t *testing.T, toTest []T, zsort FloatSorter[T]) bool {
	control := make([]T, len(toTest))
	copy(control, toTest)
	gstart := time.Now()
	slices.Sort(control)
	gdelta := time.Now().Sub(gstart)
	zstart := time.Now()
	zsort.Sort(toTest)
	zdelta := time.Now().Sub(zstart)
	for idx, val := range control {
		if math.IsNaN(float64(val)) {
			if !math.IsNaN(float64(toTest[idx])) {
				t.Fatal(control, toTest)
			}
		} else if val != toTest[idx] {
			t.Fatal(control, toTest)
		}
	}
	if !slices.IsSorted(toTest) {
		t.Fatal(control, toTest)
	}
	return zdelta < gdelta
}
