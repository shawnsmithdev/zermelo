package zermelo

import (
	"github.com/shawnsmithdev/zermelo/zfloat32"
	"github.com/shawnsmithdev/zermelo/zfloat64"
	"github.com/shawnsmithdev/zermelo/zint32"
	"github.com/shawnsmithdev/zermelo/zint64"
	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
	"log"
	"testing"
	"time"
)

// These test makes sure that the MinSize before using sort.Sort() is roughly large enough.

const breakEvenStartSize = 4

func TestBreakEvenFloat32(t *testing.T) {
	size := breakEvenTest(
		make([]float32, zfloat32.MinSize),
		make([]float32, zfloat32.MinSize),
		zfloat32.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []float32 race at length %v\n", size)
}

func TestBreakEvenFloat64(t *testing.T) {
	size := breakEvenTest(
		make([]float64, zfloat64.MinSize),
		make([]float64, zfloat64.MinSize),
		zfloat64.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []float64 race at length %v\n", size)
}

func TestBreakEvenInt32(t *testing.T) {
	size := breakEvenTest(
		make([]int32, zint32.MinSize),
		make([]int32, zint32.MinSize),
		zint32.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []int32 race at length %v\n", size)
}

func TestBreakEvenInt64(t *testing.T) {
	size := breakEvenTest(
		make([]int64, zint64.MinSize),
		make([]int64, zint64.MinSize),
		zint64.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []int64 race at length %v\n", size)
}

func TestBreakEvenUint32(t *testing.T) {
	size := breakEvenTest(
		make([]uint32, zuint32.MinSize),
		make([]uint32, zuint32.MinSize),
		zuint32.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []uint32 race at length %v\n", size)
}

func TestBreakEvenUint64(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize),
		make([]uint64, zuint64.MinSize),
		zuint64.MinSize,
		t, genTestData)
	log.Printf("Zermelo won []uint64 race at length %v\n", size)
}

func TestBreakEvenUint64Sorted(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize),
		make([]uint64, zuint64.MinSize),
		zuint64.MinSize,
		t, genSortedTestData)
	log.Printf("Zermelo won []uint64 sorted race at length %v\n", size)
}

func TestBreakEvenUint64Reversed(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize),
		make([]uint64, zuint64.MinSize),
		zuint64.MinSize,
		t, genReversedTestData)
	log.Printf("Zermelo won []uint64 reversed race at length %v\n", size)
}

func breakEvenTest(g, r interface{}, minSize uint, t *testing.T, genFunc func(interface{})) uint {
	gsort := goSorter
	zsort := newRawSorter().Sort
	for size := uint(breakEvenStartSize); size < (2 * minSize); size++ {
		var retry uint
		max := uint(65)
		for retry = 0; retry < max; retry++ {
			genFunc(g)
			x := sliceSlice(r, size)
			sliceCopy(x, g)
			start := time.Now().UnixNano()
			zsort(x)
			ztime := time.Now().UnixNano() - start
			sliceCopy(x, g)
			start = time.Now().UnixNano()
			gsort(x)
			gtime := time.Now().UnixNano() - start
			if ztime > gtime && retry > 0 { // Always throw away first run
				break
			}
		}
		if retry == max {
			return size
		}
	}
	t.FailNow()
	return 0
}

// Slices x to x[:newlen]
func sliceSlice(x interface{}, newlen uint) interface{} {
	switch xAsCase := x.(type) {
	case []float32:
		return xAsCase[:newlen]
	case []float64:
		return xAsCase[:newlen]
	case []int:
		return xAsCase[:newlen]
	case []int32:
		return xAsCase[:newlen]
	case []int64:
		return xAsCase[:newlen]
	case []uint:
		return xAsCase[:newlen]
	case []uint32:
		return xAsCase[:newlen]
	case []uint64:
		return xAsCase[:newlen]
	default:
		panic("not supported")
	}
}
