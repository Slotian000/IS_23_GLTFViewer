package main

import (
	"encoding/binary"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/qmuntal/gltf"
	"math"
	"openGL/Wrappers"
	"strings"
)

type Mesh struct {
	Positions     []float32 //Buffer
	NormalCoords  []float32 //Buffer
	TextureCoords []float32 //Buffer
	Tangents      []float32 //Buffer
	Translation   []float32
	Rotation      []float32
	Scale         []float32
	Indices       []uint32
	VAO           Wrappers.VAO
	Program       Wrappers.Program
	Material      Material
	Model         mgl32.Mat4
}

type Material struct {
	Texture Wrappers.Texture
}

type Buffer struct {
	Buffer []interface{}
	Stride int
}

var VertexAttributes = []Wrappers.VertexAttribute{
	Wrappers.NewVertexAttribute(3, gl.FLOAT, false), //Position
	Wrappers.NewVertexAttribute(3, gl.FLOAT, false), //Normal
	Wrappers.NewVertexAttribute(2, gl.FLOAT, false), //Texture
	Wrappers.NewVertexAttribute(3, gl.FLOAT, false), //Tangent
}

func test(path string) []Mesh {
	doc, err := gltf.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	meshes := make([]Mesh, 0, 0)
	textures := make([]Wrappers.Texture, 0, 0)
	materials := make([]Material, 0, 0)

	for _, image := range doc.Images {
		texture, _ := Wrappers.NewTexture("Sources/" + image.URI)
		textures = append(textures, texture)
	}

	for _, material := range doc.Materials {
		if material.PBRMetallicRoughness.BaseColorTexture != nil {
			materials = append(materials, Material{Texture: textures[material.PBRMetallicRoughness.BaseColorTexture.Index]})
		} else {
			x := material.PBRMetallicRoughness.BaseColorFactor
			materials = append(materials, Material{Texture: Wrappers.NewSolidTexture(x[0], x[1], x[2], x[3])})
		}
	}

	for _, node := range doc.Nodes {
		if node.Mesh != nil {
			attributes := doc.Meshes[*node.Mesh].Primitives[0].Attributes
			key := ""
			vertexAttributes := make([]Wrappers.VertexAttribute, 0, 0)

			mesh := Mesh{
				Indices:     Uint16BufferAsUint32Buffer(doc, *doc.Meshes[*node.Mesh].Primitives[0].Indices),
				Translation: node.Translation[:],
				Rotation:    node.Rotation[:],
				Scale:       node.Scale[:],
				Model:       translate(node.Translation, node.Rotation, node.Scale),
			}
			if accessor, ok := attributes["POSITION"]; ok {
				mesh.Positions = Float32Buffer(doc, accessor)
				key += "P"
				vertexAttributes = append(vertexAttributes, VertexAttributes[0])
			}
			if accessor, ok := attributes["NORMAL"]; ok {
				mesh.NormalCoords = Float32Buffer(doc, accessor)
				key += "N"
				vertexAttributes = append(vertexAttributes, VertexAttributes[1])
			}
			if accessor, ok := attributes["TEXCOORD_0"]; ok {
				mesh.TextureCoords = Float32Buffer(doc, accessor)
				key += "T"
				vertexAttributes = append(vertexAttributes, VertexAttributes[2])
			}
			if accessor, ok := attributes["TANGENT"]; ok {
				mesh.Tangents = Float32Buffer(doc, accessor)
				key += "X"
				vertexAttributes = append(vertexAttributes, VertexAttributes[3])
			}

			raw := make([]float32, 0, 0)
			for i, j := 0, 0; i < len(mesh.Positions); i, j = i+3, j+2 {
				raw = append(raw, mesh.Positions[i:i+3]...)
				if strings.Contains(key, "N") {
					raw = append(raw, mesh.NormalCoords[i:i+3]...)
				}
				if strings.Contains(key, "T") {
					raw = append(raw, mesh.TextureCoords[j:j+2]...)
				}
				if strings.Contains(key, "X") {
					raw = append(raw, mesh.Tangents[i:i+3]...)
				}
			}
			mesh.VAO = Wrappers.NewVAOWithEBO(raw, mesh.Indices, gl.STATIC_DRAW, vertexAttributes...)
			mesh.Program = Programs[key]
			mesh.Material = materials[*doc.Meshes[*node.Mesh].Primitives[0].Material]
			meshes = append(meshes, mesh)
		}
	}
	return meshes
}

func Float32Buffer(doc *gltf.Document, a uint32) []float32 {
	accessor := doc.Accessors[a]
	bv := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[bv.Buffer].Data[bv.ByteOffset+accessor.ByteOffset:]
	result := make([]float32, 0, accessor.Count)

	for i := 0; i < int(accessor.Count*accessor.Type.Components()*4); i += 4 {
		result = append(result, math.Float32frombits(binary.LittleEndian.Uint32(buffer[i:i+4])))
	}
	return result
}

func Uint16BufferAsUint32Buffer(doc *gltf.Document, a uint32) []uint32 {
	accessor := doc.Accessors[a]
	bv := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[bv.Buffer].Data[bv.ByteOffset+accessor.ByteOffset:]
	result := make([]uint32, 0, accessor.Count)

	for i := 0; i < int(accessor.Count*accessor.Type.Components()*2); i += 2 {
		result = append(result, uint32(binary.LittleEndian.Uint16([]byte{buffer[i], buffer[i+1]})))
	}
	return result
}

func translate(translation [3]float32, rotation [4]float32, scale [3]float32) mgl32.Mat4 {
	model := mgl32.Scale3D(scale[0], scale[1], scale[2])
	//model := mgl32.Scale3D(scale[0]*.01, scale[1]*.01, scale[2]*.01)
	//model := mgl32.Scale3D(scale[0]*10, scale[1]*10, scale[2]*10)
	//model := mgl32.Scale3D(1, 1, 1)
	//model := mgl32.Scale3D(10, 10, 10)
	model.Mul4(mgl32.HomogRotate3DX(rotation[0]))
	model.Mul4(mgl32.HomogRotate3DY(rotation[1]))
	model.Mul4(mgl32.HomogRotate3DZ(rotation[2]))
	model.Mul4(mgl32.Translate3D(translation[0], translation[0], translation[0]))
	return model
}
