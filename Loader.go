package main

import (
	"encoding/binary"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/qmuntal/gltf"
	"math"
	"openGL/Wrappers"
)

type Mesh struct {
	Positions     []float32
	NormalCoords  []float32
	TextureCoords []float32
	Tangents      []float32
	Indices       []uint32
	VAO           Wrappers.VAO
}

type Image struct {
	Data []byte
}

var VertexAttributes = []Wrappers.VertexAttribute{
	Wrappers.NewVertexAttribute(3, gl.FLOAT, false), // Position
	Wrappers.NewVertexAttribute(3, gl.FLOAT, false), // Normal
	Wrappers.NewVertexAttribute(2, gl.FLOAT, false), // Texture
	Wrappers.NewVertexAttribute(2, gl.FLOAT, false), // Tangent
}

func test() []Mesh {
	doc, err := gltf.Open("Sources/AnimatedCube.gltf")
	if err != nil {
		fmt.Println(err)
	}
	meshes := make([]Mesh, 0, 0)

	for _, node := range doc.Nodes {
		if node.Mesh != nil {
			attributes := doc.Meshes[*node.Mesh].Primitives[0].Attributes
			mesh := Mesh{Indices: Uint16BufferAsUint32Buffer(doc, *doc.Meshes[*node.Mesh].Primitives[0].Indices)}

			if accessor, ok := attributes["POSITION"]; ok {
				mesh.Positions = Float32Buffer(doc, accessor)
			}
			if accessor, ok := attributes["NORMAL"]; ok {
				mesh.NormalCoords = Float32Buffer(doc, accessor)
			}
			if accessor, ok := attributes["TEXCOORD_0"]; ok {
				mesh.TextureCoords = Float32Buffer(doc, accessor)
			}
			if accessor, ok := attributes["TANGENT"]; ok {
				mesh.Tangents = Float32Buffer(doc, accessor)
			}

			raw := make([]float32, 0, len(mesh.Positions)*4)
			for i := 0; i*3 < len(mesh.Positions)/3; i++ {
				raw = append(raw, mesh.Positions[i*3:i*3+3]...)
				raw = append(raw, mesh.NormalCoords[i*3:i*3+3]...)
				raw = append(raw, mesh.TextureCoords[i*2:i*2+2]...)
				raw = append(raw, mesh.Tangents[i*2:i*2+2]...)
			}

			mesh.VAO = Wrappers.NewVAOWithEBO(raw, mesh.Indices, gl.STATIC_DRAW, VertexAttributes...)
			meshes = append(meshes, mesh)
		}
	}
	return meshes
}

func Float32Buffer(doc *gltf.Document, accessor uint32) []float32 {
	bv := doc.BufferViews[*doc.Accessors[accessor].BufferView]
	buffer := doc.Buffers[bv.Buffer].Data[bv.ByteOffset : bv.ByteOffset+bv.ByteLength]
	result := make([]float32, 0, len(buffer)/4)
	for i := 0; i < len(buffer)-4; i += 4 {
		result = append(result, math.Float32frombits(binary.LittleEndian.Uint32(buffer[i:i+4])))
	}
	return result
}

func Uint16BufferAsUint32Buffer(doc *gltf.Document, accessor uint32) []uint32 {
	bv := doc.BufferViews[*doc.Accessors[accessor].BufferView]
	buffer := doc.Buffers[bv.Buffer].Data[bv.ByteOffset : bv.ByteOffset+bv.ByteLength]
	result := make([]uint32, 0, len(buffer)/2)
	for i := len(buffer) - 1; i >= 1; i -= 2 {
		result = append([]uint32{uint32(binary.BigEndian.Uint16([]byte{buffer[i], buffer[i-1]}))}, result...)
	}
	return result
}

/*
	images := make([]Image, 0, 0)
	for _, image := range doc.Images {
		bv := doc.BufferViews[*doc.Accessors[*image.BufferView].BufferView]
		images = append(images, Image{Data: doc.Buffers[bv.Buffer].Data[bv.ByteOffset : bv.ByteOffset+bv.ByteLength]})
	}

	for _, material := range doc.Materials{
		material.


	}

*/
