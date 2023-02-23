package Utils

import (
	"math"
)

func Sin(f float32) float32 {
	return float32(math.Sin(float64(f)))
}
func Cos(f float32) float32 {
	return float32(math.Cos(float64(f)))
}
func MinMax(min float32, val float32, max float32) float32 {
	if val > max {
		return max
	}
	if val < min {
		return min
	}
	return val

}
