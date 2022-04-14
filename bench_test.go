package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Benchmarks
const testTinySize = compSortCutoff
const testSmallSize = compSortCutoff64
const testMediumSize = 1024   // ~1k * 64bit = 8 KB
const testLargeSize = 1 << 20 // ~1M * 64bit = 8 MB

// []uint64
func BenchmarkZSortUint64T(b *testing.B) {
	sortFunc := func(x []uint64) {
		SortIntegersBYOB[uint64](x, make([]uint64, testTinySize))
	}
	testIntSortBencher[uint64](b, testTinySize, sortFunc)
}
func BenchmarkZSorterUint64T(b *testing.B) {
	s := newIntSorter[uint64]()
	s.setCutoff(0)
	testIntSortBencher[uint64](b, testTinySize, s.Sort)
}
func BenchmarkGoSortUint64T(b *testing.B) {
	testIntSortBencher[uint64](b, testTinySize, slices.Sort[uint64])
}
func BenchmarkZSortUint64S(b *testing.B) {
	testIntSortBencher[uint64](b, testSmallSize, SortIntegers[uint64])
}
func BenchmarkZSorterUint64S(b *testing.B) {
	testIntSortBencher[uint64](b, testSmallSize, NewIntSorter[uint64]().Sort)
}
func BenchmarkGoSortUint64S(b *testing.B) {
	testIntSortBencher[uint64](b, testSmallSize, slices.Sort[uint64])
}
func BenchmarkZSortUint64M(b *testing.B) {
	testIntSortBencher[uint64](b, testMediumSize, SortIntegers[uint64])
}
func BenchmarkZSorterUint64M(b *testing.B) {
	testIntSortBencher[uint64](b, testMediumSize, NewIntSorter[uint64]().Sort)
}
func BenchmarkGoSortUint64M(b *testing.B) {
	testIntSortBencher[uint64](b, testMediumSize, slices.Sort[uint64])
}
func BenchmarkZSortUint64L(b *testing.B) {
	testIntSortBencher[uint64](b, testLargeSize, SortIntegers[uint64])
}
func BenchmarkZSorterUint64L(b *testing.B) {
	testIntSortBencher[uint64](b, testLargeSize, NewIntSorter[uint64]().Sort)
}
func BenchmarkGoSortUint64L(b *testing.B) {
	testIntSortBencher[uint64](b, testLargeSize, slices.Sort[uint64])
}

// []float64
func BenchmarkZSortFloat64T(b *testing.B) {
	sortFunc := func(x []float64) {
		SortFloatsBYOB(x, make([]float64, testTinySize))
	}
	testFloatSortBencher(b, testTinySize, sortFunc)
}
func BenchmarkZSorterFloat64T(b *testing.B) {
	s := newFloatSorter[float64]()
	s.setCutoff(0)
	testFloatSortBencher(b, testTinySize, s.Sort)
}
func BenchmarkGoSortFloat64T(b *testing.B) {
	testFloatSortBencher(b, testTinySize, slices.Sort[float64])
}
func BenchmarkZSortFloat64S(b *testing.B) {
	testFloatSortBencher(b, testSmallSize, SortFloats[float64])
}
func BenchmarkZSorterFloat64S(b *testing.B) {
	testFloatSortBencher(b, testSmallSize, NewFloatSorter[float64]().Sort)
}
func BenchmarkGoSortFloat64S(b *testing.B) {
	testFloatSortBencher(b, testSmallSize, slices.Sort[float64])
}
func BenchmarkZSortFloat64M(b *testing.B) {
	testFloatSortBencher(b, testMediumSize, SortFloats[float64])
}
func BenchmarkZSorterFloat64M(b *testing.B) {
	testFloatSortBencher(b, testMediumSize, NewFloatSorter[float64]().Sort)
}
func BenchmarkGoSortFloat64M(b *testing.B) {
	testFloatSortBencher(b, testMediumSize, slices.Sort[float64])
}
func BenchmarkZSortFloat64L(b *testing.B) {
	testFloatSortBencher(b, testLargeSize, SortFloats[float64])
}
func BenchmarkZSorterFloat64L(b *testing.B) {
	testFloatSortBencher(b, testLargeSize, NewFloatSorter[float64]().Sort)
}
func BenchmarkGoSortFloat64L(b *testing.B) {
	testFloatSortBencher(b, testLargeSize, slices.Sort[float64])
}

func sortedTestData[T constraints.Integer](size int) func(int) [][]T {
	return func(n int) [][]T {
		result := testDataFromRng[T](randInteger[T](), size)(n)
		var wg sync.WaitGroup
		cpus := runtime.NumCPU()
		for cpu := 0; cpu < cpus; cpu++ {
			wg.Add(1)
			go func(c int) {
				defer wg.Done()
				presorter := NewIntSorter[T]()
				for i := c; i < len(result); i += cpus {
					presorter.Sort(result[i])
				}
			}(cpu)
		}
		wg.Wait()
		return result
	}
}

// presorted
func BenchmarkZSortSorted(b *testing.B) {
	testBencher[uint64](b, SortIntegers[uint64], sortedTestData[uint64](testSmallSize))
}
func BenchmarkZSorterSorted(b *testing.B) {
	testBencher[uint64](b, NewIntSorter[uint64]().Sort, sortedTestData[uint64](testSmallSize))
}
func BenchmarkGoSortSorted(b *testing.B) {
	testBencher[uint64](b, slices.Sort[uint64], sortedTestData[uint64](testSmallSize))
}

type sorter[T any] func([]T)

func testDataFromRng[T any](rng func() T, size int) func(int) [][]T {
	return func(n int) [][]T {
		result := make([][]T, n)
		for i := 0; i < n; i++ {
			result[i] = make([]T, size)
			fillSlice(result[i], rng)
		}
		return result
	}
}

func testIntSortBencher[T constraints.Integer](b *testing.B, size int, s sorter[T]) {
	rand.Seed(time.Now().UnixNano())
	rng := randInteger[T]()
	testBencher(b, s, testDataFromRng[T](rng, size))
}

func testFloatSortBencher(b *testing.B, size int, s sorter[float64]) {
	rand.Seed(time.Now().UnixNano())
	testBencher(b, s, testDataFromRng[float64](randFloat64(), size))
}

// for bench b, tests s by copying rnd to x and sorting x repeatedly
func testBencher[T constraints.Ordered](b *testing.B, s sorter[T], getTestData func(n int) [][]T) {
	b.StopTimer()
	rnd := getTestData(b.N)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s(rnd[i])
	}
}
