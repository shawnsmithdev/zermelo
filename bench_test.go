package zermelo

import (
	"testing"
)

// Benchmarks

// []string
func BenchmarkZSortStringT(b *testing.B) {
	zermeloSortBencher(b, make([]string, TEST_TINY_SIZE), make([]string, TEST_TINY_SIZE))
}
func BenchmarkGoSortStringT(b *testing.B) {
	goSortBencher(b, make([]string, TEST_TINY_SIZE), make([]string, TEST_TINY_SIZE))
}
func BenchmarkZSortStringS(b *testing.B) {
	zermeloSortBencher(b, make([]string, TEST_SMALL_SIZE), make([]string, TEST_SMALL_SIZE))
}
func BenchmarkGoSortStringS(b *testing.B) {
	goSortBencher(b, make([]string, TEST_SMALL_SIZE), make([]string, TEST_SMALL_SIZE))
}
func BenchmarkZSortString(b *testing.B) {
	zermeloSortBencher(b, make([]string, TEST_SIZE), make([]string, TEST_SIZE))
}
func BenchmarkGoSortString(b *testing.B) {
	goSortBencher(b, make([]string, TEST_SIZE), make([]string, TEST_SIZE))
}
func BenchmarkZSortStringB(b *testing.B) {
	zermeloSortBencher(b, make([]string, TEST_BIG_SIZE), make([]string, TEST_BIG_SIZE))
}
func BenchmarkGoSortStringB(b *testing.B) {
	goSortBencher(b, make([]string, TEST_BIG_SIZE), make([]string, TEST_BIG_SIZE))
}

// []uint64
func BenchmarkZSortUint64T(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, TEST_TINY_SIZE), make([]uint64, TEST_TINY_SIZE))
}
func BenchmarkGoSortUint64T(b *testing.B) {
	goSortBencher(b, make([]uint64, TEST_TINY_SIZE), make([]uint64, TEST_TINY_SIZE))
}
func BenchmarkZSortUint64S(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, TEST_SMALL_SIZE), make([]uint64, TEST_SMALL_SIZE))
}
func BenchmarkGoSortUint64S(b *testing.B) {
	goSortBencher(b, make([]uint64, TEST_SMALL_SIZE), make([]uint64, TEST_SMALL_SIZE))
}
func BenchmarkZSortUint64(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, TEST_SIZE), make([]uint64, TEST_SIZE))
}
func BenchmarkGoSortUint64(b *testing.B) {
	goSortBencher(b, make([]uint64, TEST_SIZE), make([]uint64, TEST_SIZE))
}
func BenchmarkZSortUint64B(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, TEST_BIG_SIZE), make([]uint64, TEST_BIG_SIZE))
}
func BenchmarkGoSortUint64B(b *testing.B) {
	goSortBencher(b, make([]uint64, TEST_BIG_SIZE), make([]uint64, TEST_BIG_SIZE))
}

// []float64
func BenchmarkZSortFloat64T(b *testing.B) {
	zermeloSortBencher(b, make([]float64, TEST_TINY_SIZE), make([]float64, TEST_TINY_SIZE))
}
func BenchmarkGoSortFloat64T(b *testing.B) {
	goSortBencher(b, make([]float64, TEST_TINY_SIZE), make([]float64, TEST_TINY_SIZE))
}
func BenchmarkZSortFloat64S(b *testing.B) {
	zermeloSortBencher(b, make([]float64, TEST_SMALL_SIZE), make([]float64, TEST_SMALL_SIZE))
}
func BenchmarkGoSortFloat64S(b *testing.B) {
	goSortBencher(b, make([]float64, TEST_SMALL_SIZE), make([]float64, TEST_SMALL_SIZE))
}
func BenchmarkZSortFloat64(b *testing.B) {
	zermeloSortBencher(b, make([]float64, TEST_SIZE), make([]float64, TEST_SIZE))
}
func BenchmarkGoSortFloat64(b *testing.B) {
	goSortBencher(b, make([]float64, TEST_SIZE), make([]float64, TEST_SIZE))
}
func BenchmarkZSortFloat64B(b *testing.B) {
	zermeloSortBencher(b, make([]float64, TEST_BIG_SIZE), make([]float64, TEST_BIG_SIZE))
}
func BenchmarkGoSortFloat64B(b *testing.B) {
	goSortBencher(b, make([]float64, TEST_BIG_SIZE), make([]float64, TEST_BIG_SIZE))
}

func BenchmarkZSortSorted(b *testing.B) {
	zermeloSortSortedBencher(b, make([]uint64, TEST_BIG_SIZE), make([]uint64, TEST_BIG_SIZE))
}
func BenchmarkGoSortSorted(b *testing.B) {
	goSortSortedBencher(b, make([]uint64, TEST_BIG_SIZE), make([]uint64, TEST_BIG_SIZE))
}

// Benchmarking Utility Functions

// these benchmark a type, storing the random values in rnd, copying them to x, and sorting x
func zermeloSortBencher(b *testing.B, rnd interface{}, x interface{}) {
	genTestData(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		Sort(x)
	}
}

func goSortBencher(b *testing.B, rnd interface{}, x interface{}) {
	genTestData(rnd)
	b.ResetTimer()
	gsort := newGoSorter(rnd)
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		gsort(x)
	}
}

func zermeloSortSortedBencher(b *testing.B, rnd []uint64, x []uint64) {
	genSortedTestData(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		Sort(x)
	}
}

func goSortSortedBencher(b *testing.B, rnd []uint64, x []uint64) {
	genSortedTestData(rnd)
	b.ResetTimer()
	gsort := newGoSorter(rnd)
	for i := 0; i < b.N; i++ {
		sliceCopy(x, rnd)
		gsort(x)
	}
}
