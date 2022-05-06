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

// returns a function that returns random float32s, without NaNs
func randFloat32() func() float32 {
	return noNans(randFloat32WithNans())
}

// returns a function that returns random float64s, without NaNs
func randFloat64() func() float64 {
	return noNans(randFloat64WithNans())
}

// returns a function that returns random float32s, including NaNs
func randFloat32WithNans() func() float32 {
	rng := randInteger[uint32]()
	return func() float32 {
		return math.Float32frombits(rng())
	}
}

// returns a function that returns random float64s, including NaNs
func randFloat64WithNans() func() float64 {
	rng := randInteger[uint64]()
	return func() float64 {
		return math.Float64frombits(rng())
	}
}

func noNans[F constraints.Float](rng func() F) func() F {
	return func() F {
		for {
			if result := rng(); !math.IsNaN(float64(result)) {
				return result
			}
		}
	}
}
