package main

import (
	"encoding/binary"
	"fmt"
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

func load() {
	doc, err := gltf.Open("AnimatedCube.gltf")
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

			mesh.VAO = Wrappers.NewVAOWithEBO()
			meshes = append(meshes, mesh)
		}
	}
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
	for i := 0; i < len(buffer)-2; i += 2 {
		result = append(result, uint32(binary.BigEndian.Uint16(buffer[i:i+2])))
	}
	return result
}
