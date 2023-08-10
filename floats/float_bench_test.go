package floats

import (
	"cmp"
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"runtime"
	"slices"
	"sync"
	"testing"
)

// Benchmarks
const testTinySize = compSortCutoffFloat32
const testSmallSize = compSortCutoffFloat64
const testMediumSize = 1024   // ~1k * 64bit = 8 KB
const testLargeSize = 1 << 20 // ~1M * 64bit = 8 MB

// tiny32
func BenchmarkSortSortF32T(b *testing.B) {
	testBencher[float32](b, sortSort[float32],
		testDataFromRng[float32](randFloat32(false), testTinySize))
}
func BenchmarkZSortF32T(b *testing.B) {
	testBencher[float32](b, SortFloats[float32],
		testDataFromRng[float32](randFloat32(false), testTinySize))
}
func BenchmarkZSorterF32T(b *testing.B) {
	testBencher[float32](b, NewFloatSorter[float32]().Sort,
		testDataFromRng[float32](randFloat32(false), testTinySize))
}

// tiny
func BenchmarkSortSortF64T(b *testing.B) {
	testBencher[float64](b, sortSort[float64],
		testDataFromRng[float64](randFloat64(false), testTinySize))
}
func BenchmarkZSortF64T(b *testing.B) {
	testBencher[float64](b, SortFloats[float64],
		testDataFromRng[float64](randFloat64(false), testTinySize))
}
func BenchmarkZSorterF64T(b *testing.B) {
	testBencher[float64](b, NewFloatSorter[float64]().Sort,
		testDataFromRng[float64](randFloat64(false), testTinySize))
}

// small
func BenchmarkSortSortF64S(b *testing.B) {
	testBencher[float64](b, sortSort[float64],
		testDataFromRng[float64](randFloat64(false), testSmallSize))
}
func BenchmarkZSortF64S(b *testing.B) {
	testBencher[float64](b, SortFloats[float64],
		testDataFromRng[float64](randFloat64(false), testSmallSize))
}
func BenchmarkZSorterF64S(b *testing.B) {
	testBencher[float64](b, NewFloatSorter[float64]().Sort,
		testDataFromRng[float64](randFloat64(false), testSmallSize))
}

// medium
func BenchmarkSortSortF64M(b *testing.B) {
	testBencher[float64](b, sortSort[float64],
		testDataFromRng[float64](randFloat64(false), testMediumSize))
}
func BenchmarkZSortF64M(b *testing.B) {
	testBencher[float64](b, SortFloats[float64],
		testDataFromRng[float64](randFloat64(false), testMediumSize))
}
func BenchmarkZSorterF64M(b *testing.B) {
	testBencher[float64](b, NewFloatSorter[float64]().Sort,
		testDataFromRng[float64](randFloat64(false), testMediumSize))
}

// large
func BenchmarkSortSortF64L(b *testing.B) {
	testBencher[float64](b, sortSort[float64],
		testDataFromRng[float64](randFloat64(false), testLargeSize))
}
func BenchmarkZSortF64L(b *testing.B) {
	testBencher[float64](b, SortFloats[float64],
		testDataFromRng[float64](randFloat64(false), testLargeSize))
}
func BenchmarkZSorterF64L(b *testing.B) {
	testBencher[float64](b, NewFloatSorter[float64]().Sort,
		testDataFromRng[float64](randFloat64(false), testLargeSize))
}

// presorted
func BenchmarkSortSortSorted(b *testing.B) {
	testBencher[float64](b, sortSort[float64],
		sortedTestData[float64](randFloat64(false), testSmallSize))
}
func BenchmarkZSortSorted(b *testing.B) {
	testBencher[float64](b, SortFloats[float64],
		sortedTestData[float64](randFloat64(false), testSmallSize))
}
func BenchmarkZSorterSorted(b *testing.B) {
	testBencher[float64](b, NewFloatSorter[float64]().Sort,
		sortedTestData[float64](randFloat64(false), testSmallSize))
}

// for bench b, tests s by copying rnd to x and sorting x repeatedly
func testBencher[T any](b *testing.B, sortFunc func([]T), getTestData func(n int) [][]T) {
	b.StopTimer()
	rnd := getTestData(b.N)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sortFunc(rnd[i])
	}
}

// testDataFromRng returns a function that generates tables of test data
// using the given random value generator and slice size.
func testDataFromRng[T any](rng func() T, size int) func(int) [][]T {
	return func(n int) [][]T {
		result := make([][]T, n)
		for i := 0; i < n; i++ {
			result[i] = make([]T, size)
			internal.FillSlice(result[i], rng)
		}
		return result
	}
}

// sortedTestData creates a function that generates tables of presorted test data
// using the given random value generator and slice size.
func sortedTestData[T cmp.Ordered](rng func() T, size int) func(int) [][]T {
	return func(n int) [][]T {
		result := testDataFromRng[T](rng, size)(n)
		var wg sync.WaitGroup
		cpus := runtime.NumCPU()
		for cpu := 0; cpu < cpus; cpu++ {
			wg.Add(1)
			go func(c int) {
				defer wg.Done()
				for i := c; i < len(result); i += cpus {
					slices.Sort(result[i])
				}
			}(cpu)
		}
		wg.Wait()
		return result
	}
}
