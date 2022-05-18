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

// returns a function that returns random floats
func randFloat[F constraints.Float, U constraints.Unsigned](fromBits func(U) F, nans bool) func() F {
	rng := randInteger[U]()
	return func() F {
		for {
			if result := fromBits(rng()); nans || !isNaN(result) {
				return result
			}
		}
	}
}

// returns a function that returns random float32s
func randFloat32(nans bool) func() float32 {
	return randFloat[float32, uint32](math.Float32frombits, nans)
}

// returns a function that returns random float64s
func randFloat64(nans bool) func() float64 {
	return randFloat[float64, uint64](math.Float64frombits, nans)
}
