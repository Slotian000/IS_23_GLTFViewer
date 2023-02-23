package Wrappers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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
	Count           int
}

func (v *VAO) Bind() {
	gl.BindVertexArray(v.ID)
}

func (v *VAO) UnBind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

}

func (v *VAO) Delete() {
	gl.DeleteBuffers(1, &v.vertexBufferID)
	gl.DeleteBuffers(1, &v.ID)
}

func getStrides(attributes []VertexAttribute) (int, int) {
	byteStride, elementsStride := 0, 0
	for _, attribute := range attributes {
		byteStride += attribute.ByteSize
		elementsStride += int(attribute.VecSize)
	}

	return byteStride, elementsStride

}

func addAttributes(attributes []VertexAttribute, stride int32) {
	offSet := 0
	for i, attribute := range attributes {
		gl.VertexAttribPointerWithOffset(uint32(i), attribute.VecSize, attribute.GlType, attribute.Normalized, stride, uintptr(offSet))
		gl.EnableVertexAttribArray(uint32(i))
		offSet += attribute.ByteSize
	}
}

func NewVAO(vertices []float32, usage uint32, attributes ...VertexAttribute) VAO {
	byteStride, elementsStride := getStrides(attributes)

	if len(vertices)%elementsStride != 0 {
		panic("Vertices does not match stride")
	}

	result := VAO{Count: len(vertices) / elementsStride}
	gl.GenVertexArrays(1, &result.ID)
	gl.GenBuffers(1, &result.vertexBufferID)
	result.Bind()

	gl.BindBuffer(gl.ARRAY_BUFFER, result.vertexBufferID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), usage)

	addAttributes(attributes, int32(byteStride))
	result.UnBind()
	return result
}

func NewVAOWithEBO(vertices []float32, indices []uint32, usage uint32, attributes ...VertexAttribute) VAO {
	byteStride, elementsStride := getStrides(attributes)

	if len(vertices)%elementsStride != 0 {
		panic("Vertices does not match stride")
	}

	result := VAO{Count: len(vertices) / elementsStride}
	gl.GenVertexArrays(1, &result.ID)
	gl.GenBuffers(1, &result.vertexBufferID)
	gl.GenBuffers(1, &result.elementBufferID)
	result.Bind()

	gl.BindBuffer(gl.ARRAY_BUFFER, result.vertexBufferID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), usage)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, result.elementBufferID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), usage)

	addAttributes(attributes, int32(byteStride))
	result.UnBind()

	return result
}
