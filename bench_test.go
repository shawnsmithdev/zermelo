package zermelo

import (
	"testing"
)

// Benchmarks

// []uint64
func BenchmarkZSortUint64T(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize))
}
func BenchmarkGoSortUint64T(b *testing.B) {
	goSortBencher(b, make([]uint64, testTinySize), make([]uint64, testTinySize))
}
func BenchmarkZSortUint64S(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize))
}
func BenchmarkGoSortUint64S(b *testing.B) {
	goSortBencher(b, make([]uint64, testSmallSize), make([]uint64, testSmallSize))
}
func BenchmarkZSortUint64(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, testSize), make([]uint64, testSize))
}
func BenchmarkGoSortUint64(b *testing.B) {
	goSortBencher(b, make([]uint64, testSize), make([]uint64, testSize))
}
func BenchmarkZSortUint64B(b *testing.B) {
	zermeloSortBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize))
}
func BenchmarkGoSortUint64B(b *testing.B) {
	goSortBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize))
}

// []float64
func BenchmarkZSortFloat64T(b *testing.B) {
	zermeloSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize))
}
func BenchmarkGoSortFloat64T(b *testing.B) {
	goSortBencher(b, make([]float64, testTinySize), make([]float64, testTinySize))
}
func BenchmarkZSortFloat64S(b *testing.B) {
	zermeloSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize))
}
func BenchmarkGoSortFloat64S(b *testing.B) {
	goSortBencher(b, make([]float64, testSmallSize), make([]float64, testSmallSize))
}
func BenchmarkZSortFloat64(b *testing.B) {
	zermeloSortBencher(b, make([]float64, testSize), make([]float64, testSize))
}
func BenchmarkGoSortFloat64(b *testing.B) {
	goSortBencher(b, make([]float64, testSize), make([]float64, testSize))
}
func BenchmarkZSortFloat64B(b *testing.B) {
	zermeloSortBencher(b, make([]float64, testBigSize), make([]float64, testBigSize))
}
func BenchmarkGoSortFloat64B(b *testing.B) {
	goSortBencher(b, make([]float64, testBigSize), make([]float64, testBigSize))
}

func BenchmarkZSortSorted(b *testing.B) {
	zermeloSortSortedBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize))
}
func BenchmarkGoSortSorted(b *testing.B) {
	goSortSortedBencher(b, make([]uint64, testBigSize), make([]uint64, testBigSize))
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
