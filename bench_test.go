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
func BenchmarkZSortUint64(b *testing.B) {
	testSortBencher(b, make([]uint64, testSize), make([]uint64, testSize), Sort)
}
func BenchmarkZSorterUint64(b *testing.B) {
	testSortBencher(b, make([]uint64, testSize), make([]uint64, testSize), New().Sort)
}
func BenchmarkGoSortUint64(b *testing.B) {
	testSortBencher(b, make([]uint64, testSize), make([]uint64, testSize), goSorter)
}
func BenchmarkZSortUint64B(b *testing.B) {
	testSortBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), Sort)
}
func BenchmarkZSorterUint64B(b *testing.B) {
	testSortBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), New().Sort)
}
func BenchmarkGoSortUint64B(b *testing.B) {
	testSortBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), goSorter)
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
func BenchmarkZSortFloat64(b *testing.B) {
	testSortBencher(b, make([]float64, testSize), make([]float64, testSize), Sort)
}
func BenchmarkZSorterFloat64(b *testing.B) {
	testSortBencher(b, make([]float64, testSize), make([]float64, testSize), New().Sort)
}
func BenchmarkGoSortFloat64(b *testing.B) {
	testSortBencher(b, make([]float64, testSize), make([]float64, testSize), goSorter)
}
func BenchmarkZSortFloat64B(b *testing.B) {
	testSortBencher(b, make([]float64, testBigSize), make([]float64, testBigSize), Sort)
}
func BenchmarkZSorterFloat64B(b *testing.B) {
	testSortBencher(b, make([]float64, testBigSize), make([]float64, testBigSize), New().Sort)
}
func BenchmarkGoSortFloat64B(b *testing.B) {
	testSortBencher(b, make([]float64, testBigSize), make([]float64, testBigSize), goSorter)
}

func BenchmarkZSortSorted(b *testing.B) {
	testSortedBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), Sort)
}
func BenchmarkZSorterSorted(b *testing.B) {
	testSortedBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), New().Sort)
}
func BenchmarkGoSortSorted(b *testing.B) {
	testSortedBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize), goSorter)
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
