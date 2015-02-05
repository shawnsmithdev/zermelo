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

// Random test data generators
func genTestDataFloat32(x []float32) {
	for i, _ := range x {
		x[i] = rand.Float32()
	}
}

func genTestDataFloat64(x []float64) {
	for i, _ := range x {
		x[i] = rand.Float64()
	}
}

func genTestDataUint(x []uint) {
	for i, _ := range x {
		x[i] = uint(rand.Uint32())
	}
}

func genTestDataUint32(x []uint32) {
	for i, _ := range x {
		x[i] = rand.Uint32()
	}
}

func genTestDataUint64(x []uint64) {
	for i, _ := range x {
		x[i] = uint64(rand.Int63())
	}
}

// rand doesn't make generating random singed values easy
// We generate random int64 between 0 and 2^32 - 1
// Then we subtract 2^31
func genTestDataInt(x []int) {
	for i, _ := range x {
		x[i] = int(rand.Int63n(1<<32) - (1 << 31))
	}
}

// Same process here as with genTestDataInt
func genTestDataInt32(x []int32) {
	for i, _ := range x {
		x[i] = int32(rand.Int63n(1<<32) - (1 << 31))
	}
}

// For int64 instead we'll just roll the sign bit seperately
// This does results in 0 being twice as likely as any other number
// but the chances are so very remote that it doesn't matter for tests
func genTestDataInt64(x []int64) {
	var tmp int64
	var isNeg bool
	for i, _ := range x {
		tmp = rand.Int63()
		isNeg = 1 == (1 & rand.Uint32())
		if isNeg {
			x[i] = 0 - tmp
		} else {
			x[i] = tmp
		}
	}
}

// Reflection based utility functions.

func genTestData(x interface{}) {
	rand.Seed(time.Now().UnixNano())
	switch xAsCase := x.(type) {
	case []float32:
		genTestDataFloat32(xAsCase)
	case []float64:
		genTestDataFloat64(xAsCase)
	case []int:
		genTestDataInt(xAsCase)
	case []int32:
		genTestDataInt32(xAsCase)
	case []int64:
		genTestDataInt64(xAsCase)
	case []uint:
		genTestDataUint(xAsCase)
	case []uint32:
		genTestDataUint32(xAsCase)
	case []uint64:
		genTestDataUint64(xAsCase)
	default:
		panic("not supported")
	}
}

// Copies content of y to x. Will fail if x and y aren't the same kind and size of slice
func sliceCopy(x interface{}, y interface{}) {
	switch yAsCase := y.(type) {
	case []float32:
		copy(x.([]float32), yAsCase)
	case []float64:
		copy(x.([]float64), yAsCase)
	case []int:
		copy(x.([]int), yAsCase)
	case []int32:
		copy(x.([]int32), yAsCase)
	case []int64:
		copy(x.([]int64), yAsCase)
	case []uint:
		copy(x.([]uint), yAsCase)
	case []uint32:
		copy(x.([]uint32), yAsCase)
	case []uint64:
		copy(x.([]uint64), yAsCase)
	default:
		panic("not supported")
	}
}

// Go's not embracing immutability is one of its weaknesses.
type sorter func(interface{})

// Attempts to return the best the sort package has to offer for the given type
func newGoSorter(x interface{}) sorter {
	switch x.(type) {
	case []float32:
		return func(y interface{}) {
			sort.Sort(float32Sortable(y.([]float32)))
		}
	case []float64:
		return func(y interface{}) {
			sort.Float64s(y.([]float64))
		}
	case []int:
		return func(y interface{}) {
			sort.Ints(y.([]int))
		}
	case []int32:
		return func(y interface{}) {
			sort.Sort(int32Sortable(y.([]int32)))
		}
	case []int64:
		return func(y interface{}) {
			sort.Sort(int64Sortable(y.([]int64)))
		}
	case []uint:
		return func(y interface{}) {
			sort.Sort(uintSortable(y.([]uint)))
		}
	case []uint32:
		return func(y interface{}) {
			sort.Sort(uint32Sortable(y.([]uint32)))
		}
	case []uint64:
		return func(y interface{}) {
			sort.Sort(uint64Sortable(y.([]uint64)))
		}
	default:
		panic("not supported")
	}
}

// Only Float64 and int are directly sortable by the sort package.
// So these implement sort.Interface for everything else
type float32Sortable []float32

func (r float32Sortable) Len() int           { return len(r) }
func (r float32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r float32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type int32Sortable []int32

func (r int32Sortable) Len() int           { return len(r) }
func (r int32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type int64Sortable []int64

func (r int64Sortable) Len() int           { return len(r) }
func (r int64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r int64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type uintSortable []uint

func (r uintSortable) Len() int           { return len(r) }
func (r uintSortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uintSortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type uint32Sortable []uint32

func (r uint32Sortable) Len() int           { return len(r) }
func (r uint32Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint32Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type uint64Sortable []uint64

func (r uint64Sortable) Len() int           { return len(r) }
func (r uint64Sortable) Less(i, j int) bool { return r[i] < r[j] }
func (r uint64Sortable) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
