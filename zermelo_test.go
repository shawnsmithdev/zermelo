package zermelo

import (
	"github.com/shawnsmithdev/zermelo/zfloat32"
	"github.com/shawnsmithdev/zermelo/zfloat64"
	"github.com/shawnsmithdev/zermelo/zint"
	"github.com/shawnsmithdev/zermelo/zint32"
	"github.com/shawnsmithdev/zermelo/zint64"
	"github.com/shawnsmithdev/zermelo/zuint"
	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const TEST_TINY_SIZE = 1 << 6  //   64 * 64bit = 512 B (worse that sort.Sort)
const TEST_SMALL_SIZE = 1 << 8 //  256 * 64bit = 2 KB (break even is around here)
const TEST_SIZE = 1 << 16      // ~64k * 64bit = 512 KB
const TEST_BIG_SIZE = 1 << 20  //  ~1M * 64bit = 8 MB 

// Compare results of using reflection api instead of directly calling sort func

func TestReflectSortFloat32(t *testing.T) {
	g := make([]float32, TEST_SIZE)
	r := make([]float32, TEST_SIZE)
	genTestDataFloat32(g)
	copy(r, g)
	zfloat32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortFloat64(t *testing.T) {
	g := make([]float64, TEST_SIZE)
	r := make([]float64, TEST_SIZE)
	genTestDataFloat64(g)
	copy(r, g)
	zfloat64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt(t *testing.T) {
	g := make([]int, TEST_SIZE)
	r := make([]int, TEST_SIZE)
	genTestDataInt(g)
	copy(r, g)
	zint.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt32(t *testing.T) {
	g := make([]int32, TEST_SIZE)
	r := make([]int32, TEST_SIZE)
	genTestDataInt32(g)
	copy(r, g)
	zint32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt64(t *testing.T) {
	g := make([]int64, TEST_SIZE)
	r := make([]int64, TEST_SIZE)
	genTestDataInt64(g)
	copy(r, g)
	zint64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint(t *testing.T) {
	g := make([]uint, TEST_SIZE)
	r := make([]uint, TEST_SIZE)
	genTestDataUint(g)
	copy(r, g)
	zuint.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint32(t *testing.T) {
	g := make([]uint32, TEST_SIZE)
	r := make([]uint32, TEST_SIZE)
	genTestDataUint32(g)
	copy(r, g)
	zuint32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint64(t *testing.T) {
	g := make([]uint64, TEST_SIZE)
	r := make([]uint64, TEST_SIZE)
	genTestDataUint64(g)
	copy(r, g)
	zuint64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%d]\tact: [%d]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
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
		Sort(a)
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

func genTestDataUint(data []uint) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = uint(rand.Uint32())
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
// We generate random int64 between 0 and 2^32 - 1
// Then we subtract 2^31
func genTestDataInt(data []int) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = int(rand.Int63n(1<<32) - (1 << 31))
	}
}

// Same process here as with genTestDataInt
func genTestDataInt32(data []int32) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = int32(rand.Int63n(1<<32) - (1 << 31))
	}
}

// For int64 instead we'll just roll the sign bit seperately
// This does results in 0 being twice as likely as any other number
// but the chances are so very remote that it doesn't matter for tests
func genTestDataInt64(data []int64) {
	rand.Seed(time.Now().UnixNano())
	var tmp int64
	var isNeg bool
	for i, _ := range data {
		tmp = rand.Int63()
		isNeg = 1 == (1 & rand.Uint32())
		if isNeg {
			data[i] = 0 - tmp
		} else {
			data[i] = tmp
		}
	}
}

func genTestDataFloat32(data []float32) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = rand.Float32()
	}
}

func genTestDataFloat64(data []float64) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range data {
		data[i] = rand.Float64()
	}
}



// Implements sort.Interface for float32[]
type float32Sortable []float32
func (r float32Sortable) Len() int           { return len(r) }
func (r float32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r float32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// float64[] is provided by sort.Float64Slice
// int is provided by sort.IntSlice

// Implements sort.Interface for int32[]
type int32Sortable []int32
func (r int32Sortable) Len() int           { return len(r) }
func (r int32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Implements sort.Interface for int64[]
type int64Sortable []int64
func (r int64Sortable) Len() int           { return len(r) }
func (r int64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Implements sort.Interface for uint[]
type uintSortable []uint
func (r uintSortable) Len() int           { return len(r) }
func (r uintSortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uintSortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Implements sort.Interface for uint32[]
type uint32Sortable []uint32
func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Implements sort.Interface for uint64[]
type uint64Sortable []uint64
func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
