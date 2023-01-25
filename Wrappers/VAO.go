package Wrappers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"reflect"
)

type VertexAttribute struct {
	VecSize    int32
	GlType     uint32
	Normalized bool
	ByteSize   int
}

func typeSize(GlType uint32) int {
	switch GlType {
	case gl.FLOAT:
		return 4
	case gl.INT:
		return 4
	default:
		panic("Unknown Type")
	}
}

func NewVertexAttribute(VecSize int, GlType uint32, Normalized bool) VertexAttribute {
	vertexAttribute := VertexAttribute{
		VecSize:    int32(VecSize),
		GlType:     GlType,
		Normalized: Normalized,
		ByteSize:   typeSize(GlType) * VecSize,
	}
	return vertexAttribute
}

type VAO struct {
	ID              uint32
	vertexBufferID  uint32
	elementBufferID uint32
	UsingEBO        bool
	count           int
}

func (v *VAO) Bind() {
	gl.BindVertexArray(v.ID)
}

func (v *VAO) UnBind() {
	gl.BindVertexArray(0)
}

func (v *VAO) Delete() {
	gl.DeleteBuffers(1, &v.vertexBufferID)
	gl.DeleteBuffers(1, &v.ID)
}

func genEBOData(vertices []float32, stride int) ([]float32, []uint32) {
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

func NewVAO(vertices []float32, useEBO bool, usage uint32, attributes ...VertexAttribute) VAO {
	byteStride, elementsStride := 0, 0
	for _, attribute := range attributes {
		byteStride += attribute.ByteSize
		elementsStride += int(attribute.VecSize)
	}

	if len(vertices)%elementsStride != 0 {
		panic("Vertices does not match stride")
	}

	var indices []uint32
	if useEBO {
		vertices, indices = genEBOData(vertices, elementsStride)
	}

	result := VAO{count: len(vertices) / elementsStride}
	gl.GenVertexArrays(1, &result.ID)
	gl.GenBuffers(1, &result.vertexBufferID)

	if useEBO {
		gl.GenBuffers(1, &result.elementBufferID)
	}

	result.Bind()
	gl.BindBuffer(gl.ARRAY_BUFFER, result.vertexBufferID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), usage)

	if useEBO {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, result.elementBufferID)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), usage)
	}

	offSet := 0
	for i, attribute := range attributes {
		gl.VertexAttribPointerWithOffset(uint32(i), attribute.VecSize, attribute.GlType, attribute.Normalized, int32(byteStride), uintptr(offSet))
		gl.EnableVertexAttribArray(uint32(i))
		offSet += attribute.ByteSize
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	result.UnBind()

	return result
}

/*
var VAO, VBO, EBO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.GenBuffers(1, &EBO)

	gl.BindVertexArray(VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
	gl.EnableVertexAttribArray(1)
*/
