package Utils

import (
	"math"
	"reflect"
)

func genEBO(vertices []float32, stride int) ([]float32, []uint32) {
	newVertices := make([]float32, 0)
	indices := make([]uint32, 0)

	for i := 0; i < len(vertices); i += stride {
		current := vertices[i : i+stride]
		index := len(newVertices)

		for j := 0; j < len(vertices); j += stride {
			if reflect.DeepEqual(vertices[j:j+stride], current) {
				index = j
				break
			}
		}
		if index == len(newVertices) {
			newVertices = append(newVertices, current...)
		}
		indices = append(indices, uint32(index/stride))
	}

	return newVertices, indices
}

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
