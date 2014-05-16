package zermelo

import (
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const TEST_TINY_SIZE = 1 << 6  // * 64bit = 512 B (worse that sort.Sort)
const TEST_SMALL_SIZE = 1 << 8 // * 64bit = 2 KB (break even is around here)
const TEST_SIZE = 1 << 16      // * 64bit = 512 KB
const TEST_BIG_SIZE = 1 << 20  // * 64bit = 8 MB

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
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
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
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
	}
}

func TestSortInt32(t *testing.T) {
	var godata [TEST_SMALL_SIZE]int32
	g := godata[:]
	genTestDataInt32(g)
	var rdata [TEST_SMALL_SIZE]int32
	r := rdata[:]
	copy(r, g)
	sort.Sort(int32Sortable(g))
	SortInt32(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
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
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = rand.Uint32()
	}
}

func genTestDataUint64(data []uint64) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = uint64(rand.Int63())
	}
}

// rand doesn't make generating random singed values easy
// We generate a random uin64 between 0 and 2^32 - 1
// Then we subtract 2^31
func genTestDataInt32(data []int32) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = int32(rand.Int63n(1<<32) - (1 << 31))
	}
}
