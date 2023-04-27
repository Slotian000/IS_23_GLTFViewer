package Object

import (
	"fmt"
	"github.com/qmuntal/gltf"
)

/*
type Group struct {
	Name     string
	Material Material
	Indices  []int
	Smooth   int
}

type Object struct {
	Materials map[string]Material
	Groups    []Group
}

func NewObject(path string) (Object, error) {
	verticies := make([]float32, 0)
	textureCords := make([]float32, 0)
	normals := make([]float32, 0)
	stride := 0

	fmt.Print(verticies, textureCords, normals, stride)

	object := Object{}
	current := Group{Name: ""}

	file, err := os.ReadFile(path)
	if err != nil {
		return object, err
	}

	for _, line := range strings.Split(string(file), "\n") {
		switch {
		case line == "":
		case strings.HasPrefix(line, "#"):
		case strings.HasPrefix(line, "mtllib"):
			if materials, err := ReadMaterialsFromFile(strings.TrimPrefix(line, "mtllib ")); err != nil {
				return Object{}, err
			} else {
				object.Materials = materials
			}
		case strings.HasPrefix(line, "v"):
			verticies = append(verticies, readFloats(line, "v ")...)
		case strings.HasPrefix(line, "vt"):
			verticies = append(verticies, readFloats(line, "vt ")...)
		case strings.HasPrefix(line, "vn"):
			verticies = append(verticies, readFloats(line, "vn ")...)
		case strings.HasPrefix(line, "g"):
			object.Groups = append(object.Groups, current)
			current = Group{Name: strings.TrimPrefix(line, "g ")}
		case strings.HasPrefix(line, "usemtl"):
			if material, ok := object.Materials[strings.TrimPrefix(line, "usemtl ")]; ok {
				current.Material = material
			}
		case strings.HasPrefix(line, "s"):
			current.Smooth = readInts(line, "s ")[0]
		case strings.HasPrefix(line, "f"):
			current.Indices, stride = readFace(line)
		}
	}
	return Object{}, err
}

func readFloats(line string, prefix string) []float32 {
	floats := make([]float32, 0)
	for _, f := range strings.Split(strings.TrimPrefix(line, prefix), " ") {
		if f != "" {
			if float, err := strconv.ParseFloat(f, 32); err != nil {
				floats = append(floats, float32(float))
			}
		}
	}
	return floats
}

func readInts(line string, prefix string) []int {
	ints := make([]int, 0)
	for _, i := range strings.Split(strings.TrimPrefix(line, prefix), " ") {
		if i != "" {
			if number, err := strconv.Atoi(i); err != nil {
				ints = append(ints, number)
			}
		}
	}
	return ints
}

func readFace(line string) ([]int, int) {
	ints := make([]int, 0)
	stride := 0

	for _, face := range strings.Split(strings.TrimPrefix(line, "f "), " ") {
		for i, cord := range strings.Split(strings.TrimPrefix(face, "/"), " ") {
			if cord == "" {
				number, err := strconv.Atoi(cord)
				if err != nil {
					ints = append(ints, number)
				}
			}
			stride = i
		}
	}
	return ints, stride
}
*/

func Test() {
	/*doc, _ := gltf.Open(path)

	for _, camera := range doc.Cameras {
		camera
	}
	*/

	doc, _ := gltf.Open("Sources/AnimatedCube.gltf")

	fmt.Println(doc.Materials[0])

}
