package zermelo

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"testing"
)

// Benchmarks
const testTinySize = 64       //  64 * 64bit = 512 B (worse that sort.Sort)
const testSmallSize = 256     // 256 * 64bit = 2 KB (break even is around here)
const testMediumSize = 1024   // ~1k * 64bit = 8 KB
const testLargeSize = 1 << 20 // ~1M * 64bit = 8 MB

// []uint64
func BenchmarkZSortUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), randInteger[uint64](), Sort)
}
func BenchmarkZSorterUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), randInteger[uint64](), New().Sort)
}
func BenchmarkGoSortUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), randInteger[uint64](), goSorter[uint64])
}
func BenchmarkZSortUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), randInteger[uint64](), Sort)
}
func BenchmarkZSorterUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), randInteger[uint64](), New().Sort)
}
func BenchmarkGoSortUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), randInteger[uint64](), goSorter[uint64])
}
func BenchmarkZSortUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), randInteger[uint64](), Sort)
}
func BenchmarkZSorterUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), randInteger[uint64](), New().Sort)
}
func BenchmarkGoSortUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), randInteger[uint64](), goSorter[uint64])
}
func BenchmarkZSortUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), Sort)
}
func BenchmarkZSorterUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), New().Sort)
}
func BenchmarkGoSortUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), goSorter[uint64])
}

// []float64
func BenchmarkZSortFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), randFloat64(), Sort)
}
func BenchmarkZSorterFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), randFloat64(), New().Sort)
}
func BenchmarkGoSortFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), randFloat64(), goSorter[float64])
}
func BenchmarkZSortFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), randFloat64(), Sort)
}
func BenchmarkZSorterFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), randFloat64(), New().Sort)
}
func BenchmarkGoSortFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), randFloat64(), goSorter[float64])
}
func BenchmarkZSortFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), randFloat64(), Sort)
}
func BenchmarkZSorterFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), randFloat64(), New().Sort)
}
func BenchmarkGoSortFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), randFloat64(), goSorter[float64])
}
func BenchmarkZSortFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), randFloat64(), Sort)
}
func BenchmarkZSorterFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), randFloat64(), New().Sort)
}
func BenchmarkGoSortFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), randFloat64(), goSorter[float64])
}

func BenchmarkZSortSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), Sort)
}
func BenchmarkZSorterSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), New().Sort)
}
func BenchmarkGoSortSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), randInteger[uint64](), goSorter[uint64])
}

// Benchmarking Utility Functions

type sorter func(any) error

func goSorter[T constraints.Ordered](x any) error {
	slices.Sort(x.([]T))
	return nil
}

// for bench b, tests s by copying rnd to x and sorting x repeatedly
func testSortBencher[T constraints.Ordered](b *testing.B, rnd, x []T, rng func() T, s sorter) {
	fillSlice[T](rnd, rng)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(x, rnd)
		if err := s(x); err != nil {
			b.Fatal(err)
		}
	}
}

func testSortedBencher[T constraints.Ordered](b *testing.B, rnd, x []T, rng func() T, s sorter) {
	fillSlice[T](rnd, rng)
	slices.Sort(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(x, rnd)
		if err := s(x); err != nil {
			b.Fatal(err)
		}
	}
}
