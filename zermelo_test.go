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

const testTinySize = 1 << 6  //   64 * 64bit = 512 B (worse that sort.Sort)
const testSmallSize = 1 << 8 //  256 * 64bit = 2 KB (break even is around here)
const testSize = 1 << 16     // ~64k * 64bit = 512 KB
const testBigSize = 1 << 20  //  ~1M * 64bit = 8 MB

// Compare results of using reflection api instead of directly calling sort func

func TestReflectSortFloat32(t *testing.T) {
	g := make([]float32, testSize)
	r := make([]float32, testSize)
	genTestDataFloat32(g)
	copy(r, g)
	zfloat32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortFloat64(t *testing.T) {
	g := make([]float64, testSize)
	r := make([]float64, testSize)
	genTestDataFloat64(g)
	copy(r, g)
	zfloat64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt(t *testing.T) {
	g := make([]int, testSize)
	r := make([]int, testSize)
	genTestDataInt(g)
	copy(r, g)
	zint.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt32(t *testing.T) {
	g := make([]int32, testSize)
	r := make([]int32, testSize)
	genTestDataInt32(g)
	copy(r, g)
	zint32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortInt64(t *testing.T) {
	g := make([]int64, testSize)
	r := make([]int64, testSize)
	genTestDataInt64(g)
	copy(r, g)
	zint64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint(t *testing.T) {
	g := make([]uint, testSize)
	r := make([]uint, testSize)
	genTestDataUint(g)
	copy(r, g)
	zuint.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint32(t *testing.T) {
	g := make([]uint32, testSize)
	r := make([]uint32, testSize)
	genTestDataUint32(g)
	copy(r, g)
	zuint32.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestReflectSortUint64(t *testing.T) {
	g := make([]uint64, testSize)
	r := make([]uint64, testSize)
	genTestDataUint64(g)
	copy(r, g)
	zuint64.Sort(g)
	Sort(r)
	for i, val := range g {
		if r[i] != val {
			log.Printf("exp: [%v]\tact: [%v]\n", val, r[i])
			t.FailNow()
		}
		if i > 0 && val < r[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestSorterFloat64(t *testing.T) {
	g := make([]float64, testSize)
	genTestData(g)

	sorter := New()
	sorter.Sort(g)
	for i, val := range g {
		if i > 0 && val < g[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

func TestSorterString(t *testing.T) {
	g := make([]string, testSize)
	genTestData(g)

	sorter := New()
	sorter.Sort(g)
	for i, val := range g {
		if i > 0 && val < g[i-1] {
			log.Printf("Not Sorted!")
			t.FailNow()
		}
	}
}

// Test data generators

func genSortedTestData(x interface{}) {
	switch xAsCase := x.(type) {
	case []uint64:
		for idx := range xAsCase {
			xAsCase[idx] = uint64(idx)
		}
	}
}

func genReversedTestData(x interface{}) {
	switch xAsCase := x.(type) {
	case []uint64:
		for idx := range xAsCase {
			xAsCase[idx] = uint64(len(xAsCase)) - uint64(idx)
		}
	}
}

// Random test data generators
func genTestDataFloat32(x []float32) {
	for i := range x {
		x[i] = rand.Float32()
	}
}

func genTestDataFloat64(x []float64) {
	for i := range x {
		x[i] = rand.Float64()
	}
}

func genTestDataUint(x []uint) {
	for i := range x {
		x[i] = uint(rand.Uint32())
	}
}

func genTestDataUint32(x []uint32) {
	for i := range x {
		x[i] = rand.Uint32()
	}
}

func genTestDataUint64(x []uint64) {
	for i := range x {
		x[i] = uint64(rand.Int63())
	}
}

func genTestDataString(x []string) {
	for i := range x {
		wordLen := (rand.Int() & 31) + 1 // 1-32 long
		word := make([]byte, wordLen)
		r := uint64(rand.Int63())
		ri := 3
		for wi := range word {
			shift := uint8(ri * 8)
			word[wi] = 'a' + (uint8((r&(uint64(0x7F)<<shift))>>shift) % 26)
			ri--
			if ri < 0 {
				ri = 3
				r = uint64(rand.Int63())
			}
		}
		x[i] = string(word)
	}
}

// rand doesn't make generating random singed values easy
// We generate random int64 between 0 and 2^32 - 1
// Then we subtract 2^31
func genTestDataInt(x []int) {
	for i := range x {
		x[i] = int(rand.Int63n(1<<32) - (1 << 31))
	}
}

// Same process here as with genTestDataInt
func genTestDataInt32(x []int32) {
	for i := range x {
		x[i] = int32(rand.Int63n(1<<32) - (1 << 31))
	}
}

// For int64 instead we'll just roll the sign bit separately
// This does results in 0 being twice as likely as any other number
// but the chances are so very remote that it doesn't matter for tests
func genTestDataInt64(x []int64) {
	var tmp int64
	var isNeg bool
	for i := range x {
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
	case []string:
		genTestDataString(xAsCase)
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
	case []string:
		copy(x.([]string), yAsCase)
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
	case []float64:
		return func(y interface{}) {
			sort.Float64s(y.([]float64))
		}
	case []int:
		return func(y interface{}) {
			sort.Ints(y.([]int))
		}
	case []float32:
		return func(y interface{}) {
			z := y.([]float32)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	case []int32:
		return func(y interface{}) {
			z := y.([]int32)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	case []int64:
		return func(y interface{}) {
			z := y.([]int64)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	case []uint:
		return func(y interface{}) {
			z := y.([]uint)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	case []uint32:
		return func(y interface{}) {
			z := y.([]uint32)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	case []uint64:
		return func(y interface{}) {
			z := y.([]uint64)
			sort.Slice(z, func(i, j int) bool { return z[i] < z[j] })
		}
	default:
		panic("not supported")
	}
}
