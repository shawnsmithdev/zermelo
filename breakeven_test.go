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
		make([]float32, zfloat32.MinSize * 2),
		make([]float32, zfloat32.MinSize * 2),
		zfloat32.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []float32 race at length %v\n", size)
}

func TestBreakEvenFloat64(t *testing.T) {
	size := breakEvenTest(
		make([]float64, zfloat64.MinSize * 2),
		make([]float64, zfloat64.MinSize * 2),
		zfloat64.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []float64 race at length %v\n", size)
}

func TestBreakEvenInt32(t *testing.T) {
	size := breakEvenTest(
		make([]int32, zint32.MinSize * 2),
		make([]int32, zint32.MinSize * 2),
		zint32.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []int32 race at length %v\n", size)
}

func TestBreakEvenInt64(t *testing.T) {
	size := breakEvenTest(
		make([]int64, zint64.MinSize * 2),
		make([]int64, zint64.MinSize * 2),
		zint64.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []int64 race at length %v\n", size)
}

func TestBreakEvenUint32(t *testing.T) {
	size := breakEvenTest(
		make([]uint32, zuint32.MinSize * 2),
		make([]uint32, zuint32.MinSize * 2),
		zuint32.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []uint32 race at length %v\n", size)
}

func TestBreakEvenUint64(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize * 2),
		make([]uint64, zuint64.MinSize * 2),
		zuint64.MinSize * 2,
		t, genTestData)
	log.Printf("Zermelo won []uint64 race at length %v\n", size)
}

func TestBreakEvenUint64Sorted(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize * 2),
		make([]uint64, zuint64.MinSize * 2),
		zuint64.MinSize * 2,
		t, genSortedTestData)
	log.Printf("Zermelo won []uint64 sorted race at length %v\n", size)
}

func TestBreakEvenUint64Reversed(t *testing.T) {
	size := breakEvenTest(
		make([]uint64, zuint64.MinSize * 2),
		make([]uint64, zuint64.MinSize * 2),
		zuint64.MinSize * 2,
		t, genReversedTestData)
	log.Printf("Zermelo won []uint64 reversed race at length %v\n", size)
}

func breakEvenTest(g, r interface{}, failSize uint, t *testing.T, genFunc func(interface{})) uint {
	gsort := goSorter
	zsort := newRawSorter().Sort
	for size := uint(breakEvenStartSize); size < failSize; size++ {
		var retry uint
		max := uint(128)
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
