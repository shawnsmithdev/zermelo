package zermelo

import (
	"testing"
	"math/rand"
	"sort"
)

const TEST_TINY_SIZE  = 1 << 6  // * 64bit = 512 B (worse that sort.Sort)
const TEST_SMALL_SIZE = 1 << 8  // * 64bit = 2 KB (break even is around here)
const TEST_SIZE       = 1 << 16 // * 64bit = 512 KB
const TEST_BIG_SIZE   = 1 << 20 // * 64bit = 8 MB

type uint32Sortable []uint32
type uint64Sortable []uint64

func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func TestSortUint32Empty(t *testing.T) {
	empty := make(uint32Sortable, 0)
	SortUint32(empty)
	if len(empty) != 0 {
		t.FailNow()
	}
}

func TestSortUint64Empty(t *testing.T) {
	empty := make(uint64Sortable, 0)
	SortUint64(empty)
	if len(empty) != 0 {
		t.FailNow()
	}
}

func TestSortUint32(t *testing.T) {
	var godata [TEST_SIZE]uint32
	g := godata[:]
	genTestDataUint32(g)
	var rdata [TEST_SIZE]uint32
	r := rdata[:]
	copy(r, g)
	sort.Sort(uint32Sortable(g))
	SortUint32(r)
	for i, val := range g {
		if r[i] != val {
			t.FailNow()
		}
	}
}

func TestSortUint64(t *testing.T) {
	var godata [TEST_SIZE]uint64
	g := godata[:]
	genTestDataUint64(g)
	var rdata [TEST_SIZE]uint64
	r := rdata[:]
	copy(r, g)
	sort.Sort(uint64Sortable(g))
	SortUint64(r)
	for i, val := range g {
		if r[i] != val {
			t.FailNow()
		}
	}
}

// Benchmarks

func BenchmarkZermeloSortUint64Tiny(b *testing.B) {
	var rnd, a [TEST_TINY_SIZE]uint64
	zermeloSortUint64Bencher(b, rnd[:], a[:])
}
func BenchmarkGoSortUint64Tiny(b *testing.B) {
	var rnd, a [TEST_TINY_SIZE]uint64
	goSortUint64Bencher(b, rnd[:], a[:])
}

func BenchmarkZermeloSortUint64Small(b *testing.B) {
	var rnd, a [TEST_SMALL_SIZE]uint64
	zermeloSortUint64Bencher(b, rnd[:], a[:])
}
func BenchmarkGoSortUint64Small(b *testing.B) {
	var rnd, a [TEST_SMALL_SIZE]uint64
	goSortUint64Bencher(b, rnd[:], a[:])
}

func BenchmarkZermeloSortUint64(b *testing.B) {
	var rnd, a [TEST_SIZE]uint64
	zermeloSortUint64Bencher(b, rnd[:], a[:])
}
func BenchmarkGoSortUint64(b *testing.B) {
	var rnd, a [TEST_SIZE]uint64
	goSortUint64Bencher(b, rnd[:], a[:])
}

func BenchmarkZermeloSortUint64Big(b *testing.B) {
	var rnd, a [TEST_BIG_SIZE]uint64
	zermeloSortUint64Bencher(b, rnd[:], a[:])
}
func BenchmarkGoSortUint64Big(b *testing.B) {
	var rnd, a [TEST_BIG_SIZE]uint64
	goSortUint64Bencher(b, rnd[:], a[:])
}

// Utility Functions

func zermeloSortUint64Bencher(b *testing.B, rnd []uint64, a []uint64) {
	genTestDataUint64(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(a, rnd)
		SortUint64(a)
	}
}

func goSortUint64Bencher(b *testing.B, rnd []uint64, a []uint64) {
	genTestDataUint64(rnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(a, rnd)
		sort.Sort(uint64Sortable(a))
	}
}

func genTestDataUint32(data []uint32) {
	for i, _ := range data {
		data[i] = rand.Uint32()
	}
}

func genTestDataUint64(data []uint64) {
	for i, _ := range data {
		data[i] = uint64(rand.Int63())
	}
}
