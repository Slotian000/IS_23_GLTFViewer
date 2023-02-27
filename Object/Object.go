package Object

import (
	"fmt"
	"os"
	"strings"
)

type Group struct {
	Name     string
	Material string
	Indices  []uint32
}

type Object struct {
	Materials map[string]Material
	Groups    []Group
}

func NewObject(path string) (Object, error) {
	object := Object{}
	file, err := os.ReadFile(path)
	if err != nil {
		return object, err
	}
	group := Group{Name: ""}

	for _, line := range strings.Split(string(file), "\n") {
		switch {
		case line == "":
		case strings.HasPrefix(line, "#"):
		case strings.HasPrefix(line, "mtllib"):
			if materials, err := ReadMaterialsFromFile(strings.SplitAfter(line, " ")[1]); err != nil {
				return Object{}, err
			} else {
				object.Materials = materials
			}
		case strings.HasPrefix(line, "g"):

		}

	}
}

func readThreeFloats(line string, prefix string) []float32 {
	var x, y, z float32
	if _, err := fmt.Sscanf(strings.TrimPrefix(line, prefix), "%f %f %f", &x, &y, &z); err != nil {
		return []float32{0, 0, 0}
	}
	return []float32{x, y, z}
}

func readFloat(line string, prefix string) float32 {
	var x float32
	if _, err := fmt.Sscanf(strings.TrimPrefix(line, prefix), "%f", &x); err != nil {
		return 0
	}
	return x
}

func readInt(line string, prefix string) int {
	var x int
	if _, err := fmt.Sscanf(strings.TrimPrefix(line, prefix), "%d", &x); err != nil {
		return 0
	}
	return x
}

/*
type Object struct {
	path string
	VAO  Wrappers.VAO
}

var (
	stride     = 8
	attributes = []Wrappers.VertexAttribute{
		Wrappers.NewVertexAttribute(3, gl.FLOAT, false), //vertices
		Wrappers.NewVertexAttribute(2, gl.FLOAT, false), // texture cords
		Wrappers.NewVertexAttribute(3, gl.FLOAT, false), // normals
	}
)
func NewObjectFromFile(path string) (Object, error) {
	object := Object{path: path}
	file, err := os.ReadFile(path)
	if err != nil {
		return Object{}, err
	}
	cords := make([]float32, 0, 50)
	textureCords := make([]float32, 0, 50)
	normals := make([]float32, 0, 50)

	faces := make([]int, 0, 50)

	for _, line := range strings.Split(string(file), "\n") {
		vals := strings.Split(line[:len(line)-1], " ")

		switch vals[0] {
		case "v", "vn":
			x, _ := strconv.ParseFloat(vals[1], 32)
			y, _ := strconv.ParseFloat(vals[2], 32)
			z, _ := strconv.ParseFloat(vals[3], 32)
			if vals[0] == "v" {
				cords = append(cords, float32(x), float32(y), float32(z))
			} else {
				normals = append(normals, float32(x), float32(y), float32(z))
			}
		case "vt":
			x, _ := strconv.ParseFloat(vals[1], 32)
			y, _ := strconv.ParseFloat(vals[2], 32)
			textureCords = append(textureCords, float32(x), float32(y))
		case "f":
			for i := 1; i < 4; i++ {
				face := strings.Split(vals[i], "/")
				x, _ := strconv.Atoi(face[0])
				y, _ := strconv.Atoi(face[0])
				z, _ := strconv.Atoi(face[0])
				faces = append(faces, x, y, z)
			}
		}
	}

	vertices := make([]float32, 0)

	for i := 0; i < len(faces); i += 3 {
		vertices = append(vertices, cords[(faces[i]-1)*3:((faces[i]-1)*3)+3]...)
		vertices = append(vertices, textureCords[(faces[i+1]-1)*2:((faces[i+1]-1)*2)+2]...)
		vertices = append(vertices, normals[(faces[i+2]-1)*3:((faces[i+2]-1)*3)+3]...)
	}

	//vertices, indices := genEBO(vertices, stride)
	//object.VAO = Wrappers.NewVAOWithEBO(vertices, indices, gl.STATIC_DRAW, attributes...)
	object.VAO = Wrappers.NewVAO(vertices, gl.STATIC_DRAW, attributes...)

	for _, v := range vertices {
		fmt.Print(v)
		fmt.Print(",")
	}

	return object, err
}

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
*/
