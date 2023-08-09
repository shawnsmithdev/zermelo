package zermelo

import (
	"github.com/shawnsmithdev/zermelo/v2/internal"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices" // TODO: replace with stdlib when go1.21+
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"
)

// Benchmarks
const testTinySize = compSortCutoff
const testSmallSize = compSortCutoff64
const testMediumSize = 1024   // ~1k * 64bit = 8 KB
const testLargeSize = 1 << 20 // ~1M * 64bit = 8 MB

// tiny32
func BenchmarkSortSortInt32T(b *testing.B) {
	testSortBencher[int32](b, testTinySize, sortSort[int32])
}
func BenchmarkZSortInt32T(b *testing.B) {
	testSortBencher[int32](b, testTinySize, Sort[int32])
}
func BenchmarkZSorterInt32T(b *testing.B) {
	testSortBencher[int32](b, testTinySize, newSorter[int32]().withCutoff(0).Sort)
}

// tiny
func BenchmarkSortSortUint64T(b *testing.B) {
	testSortBencher[uint64](b, testTinySize, sortSort[uint64])
}
func BenchmarkZSortUint64T(b *testing.B) {
	testSortBencher[uint64](b, testTinySize, Sort[uint64])
}
func BenchmarkZSorterUint64T(b *testing.B) {
	testSortBencher[uint64](b, testTinySize, newSorter[uint64]().withCutoff(0).Sort)
}

// small
func BenchmarkSortSortUint64S(b *testing.B) {
	testSortBencher[uint64](b, testSmallSize, sortSort[uint64])
}
func BenchmarkZSortUint64S(b *testing.B) {
	testSortBencher[uint64](b, testSmallSize, Sort[uint64])
}
func BenchmarkZSorterUint64S(b *testing.B) {
	testSortBencher[uint64](b, testSmallSize, newSorter[uint64]().withCutoff(0).Sort)
}

// medium
func BenchmarkSortSortUint64M(b *testing.B) {
	testSortBencher[uint64](b, testMediumSize, sortSort[uint64])
}
func BenchmarkZSortUint64M(b *testing.B) {
	testSortBencher[uint64](b, testMediumSize, Sort[uint64])
}
func BenchmarkZSorterUint64M(b *testing.B) {
	testSortBencher[uint64](b, testMediumSize, newSorter[uint64]().withCutoff(0).Sort)
}

// large
func BenchmarkSortSortUint64L(b *testing.B) {
	testSortBencher[uint64](b, testLargeSize, sortSort[uint64])
}

func BenchmarkZSortUint64L(b *testing.B) {
	testSortBencher[uint64](b, testLargeSize, Sort[uint64])
}
func BenchmarkZSorterUint64L(b *testing.B) {
	testSortBencher[uint64](b, testLargeSize, newSorter[uint64]().withCutoff(0).Sort)
}

// presorted
func BenchmarkSortSortSorted(b *testing.B) {
	testBencher[uint64](b, sortSort[uint64],
		sortedTestData[uint64](internal.RandInteger[uint64](), testSmallSize))
}
func BenchmarkZSortSorted(b *testing.B) {
	testBencher[uint64](b, Sort[uint64],
		sortedTestData[uint64](internal.RandInteger[uint64](), testSmallSize))
}
func BenchmarkZSorterSorted(b *testing.B) {
	testBencher[uint64](b, NewSorter[uint64]().Sort,
		sortedTestData[uint64](internal.RandInteger[uint64](), testSmallSize))
}

func testSortBencher[T constraints.Integer](b *testing.B, size int, sortFunc func([]T)) {
	rand.Seed(time.Now().UnixNano()) // TODO: remove when go1.20+
	testBencher(b, sortFunc, testDataFromRng[T](internal.RandInteger[T](), size))
}

// for bench b, tests s by copying rnd to x and sorting x repeatedly
func testBencher[T constraints.Ordered](b *testing.B, sortFunc func([]T), getTestData func(n int) [][]T) {
	b.StopTimer()
	rnd := getTestData(b.N)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sortFunc(rnd[i])
	}
}

type sortable[I constraints.Integer] []I

func (s sortable[I]) Len() int           { return len(s) }
func (s sortable[I]) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable[I]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortSort[I constraints.Integer](x []I) {
	sort.Sort(sortable[I](x))
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
func sortedTestData[T constraints.Ordered](rng func() T, size int) func(int) [][]T {
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
