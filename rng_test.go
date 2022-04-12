package zermelo

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
)

// most of the rand functions are like Int63 returning unsigned values
// this is generic and fills all bits randomly
func randInteger[I constraints.Integer]() func() I {
	tSize, _ := detect[I]()
	buf := make([]byte, tSize/8)
	return func() I {
		var result I
		rand.Read(buf)
		for i, val := range buf {
			result |= I(val) << (i * 8)
		}
		return result
	}
}

// TODO: as of go 1.18, slices.Sort doesn't handle NaNs at all. So neither do we. So we don't test them anymore.
// TODO: In practice NaNs will still go in front of other values.

// returns a function that returns random non-NaN float64s
func randFloat64() func() float64 {
	uintRng := randInteger[uint64]()
	return func() float64 {
		for {
			if x := math.Float64frombits(uintRng()); !math.IsNaN(x) {
				return x
			}
		}
	}
}

// returns a function that returns random non-NaN float32s
func randFloat32() func() float32 {
	uintRng := randInteger[uint32]()
	return func() float32 {
		for {
			if x := math.Float32frombits(uintRng()); !math.IsNaN(float64(x)) {
				return x
			}
		}
	}
}
