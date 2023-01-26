package Utils

import "reflect"

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
