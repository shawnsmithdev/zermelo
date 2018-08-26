package zermelo

import (
	"testing"
)

// Benchmarks

// []uint64
func BenchmarkZSortUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), Sort)
}
func BenchmarkZSorterUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), New().Sort)
}
func BenchmarkGoSortUint64T(b *testing.B) {
	testSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize), goSorter)
}
func BenchmarkZSortUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), Sort)
}
func BenchmarkZSorterUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), New().Sort)
}
func BenchmarkGoSortUint64S(b *testing.B) {
	testSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize), goSorter)
}
func BenchmarkZSortUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), Sort)
}
func BenchmarkZSorterUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), New().Sort)
}
func BenchmarkGoSortUint64M(b *testing.B) {
	testSortBencher(b, make([]uint64, testMediumSize), make([]uint64, testMediumSize), goSorter)
}
func BenchmarkZSortUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), Sort)
}
func BenchmarkZSorterUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), New().Sort)
}
func BenchmarkGoSortUint64L(b *testing.B) {
	testSortBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), goSorter)
}

// []float64
func BenchmarkZSortFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), Sort)
}
func BenchmarkZSorterFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), New().Sort)
}
func BenchmarkGoSortFloat64T(b *testing.B) {
	testSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize), goSorter)
}
func BenchmarkZSortFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), Sort)
}
func BenchmarkZSorterFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), New().Sort)
}
func BenchmarkGoSortFloat64S(b *testing.B) {
	testSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize), goSorter)
}
func BenchmarkZSortFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), Sort)
}
func BenchmarkZSorterFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), New().Sort)
}
func BenchmarkGoSortFloat64M(b *testing.B) {
	testSortBencher(b, make([]float64, testMediumSize), make([]float64, testMediumSize), goSorter)
}
func BenchmarkZSortFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), Sort)
}
func BenchmarkZSorterFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), New().Sort)
}
func BenchmarkGoSortFloat64L(b *testing.B) {
	testSortBencher(b, make([]float64, testLargeSize), make([]float64, testLargeSize), goSorter)
}

func BenchmarkZSortSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), Sort)
}
func BenchmarkZSorterSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), New().Sort)
}
func BenchmarkGoSortSortedL(b *testing.B) {
	testSortedBencher(b, make([]uint64, testLargeSize), make([]uint64, testLargeSize), goSorter)
}

// Benchmarking Utility Functions

type sorter func(interface{}) error

// for bench b, tests s by copying rnd to x and sorting x repeatedly
func testSortBencher(b *testing.B, rnd, x interface{}, s sorter) {
	genTestData(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		if err := s(x); err != nil {
			b.Fatal(err)
		}
	}
}

func testSortedBencher(b *testing.B, rnd, x []uint64, s sorter) {
	genSortedTestData(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		if err := s(x); err != nil {
			b.Fatal(err)
		}
	}
}
