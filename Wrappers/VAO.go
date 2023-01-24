package Wrappers

import "github.com/go-gl/gl/v4.1-core/gl"

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
	ID       uint32
	bufferID uint32
	count    int
}

func (v *VAO) Bind() {
	gl.BindVertexArray(v.ID)
}

func (v *VAO) UnBind() {
	gl.BindVertexArray(0)
}

func (v *VAO) Delete() {
	gl.DeleteBuffers(1, &v.bufferID)
	gl.DeleteBuffers(1, &v.ID)
}

func NewVAO(vertices []float32, usage uint32, attributes ...VertexAttribute) VAO {
	byteStride, elementsStride := 0, 0
	for _, attribute := range attributes {
		byteStride += attribute.ByteSize
		elementsStride += int(attribute.VecSize)
	}

	if len(vertices)%elementsStride != 0 {
		panic("Vertices does not match stride")
	}

	result := VAO{count: len(vertices) / elementsStride}
	gl.GenVertexArrays(1, &result.ID)
	gl.GenBuffers(1, &result.bufferID)

	result.Bind()
	gl.BindBuffer(gl.ARRAY_BUFFER, result.bufferID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), usage)

	offSet := 0
	for i, attribute := range attributes {
		gl.VertexAttribPointerWithOffset(uint32(i), attribute.VecSize, attribute.GlType, attribute.Normalized, int32(byteStride), uintptr(offSet))
		gl.EnableVertexAttribArray(uint32(i))
		offSet += attribute.ByteSize
	}

	result.UnBind()
	return result
}
