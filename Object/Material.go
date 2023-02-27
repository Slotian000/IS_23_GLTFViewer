package Object

import (
	"os"
	"strings"
)

type Material struct {
	Ka    []float32
	Kd    []float32
	Ks    []float32
	D     float32
	Tr    float32
	Ns    float32
	Illum int
	MapKa []byte
}

func ReadMaterialsFromFile(path string) (map[string]Material, error) {
	materials := make(map[string]Material)
	file, err := os.ReadFile(path)
	if err != nil {
		return materials, err
	}

	name := ""
	current := Material{}

	for _, line := range strings.Split(string(file), "\n") {

		switch {

		case line == "":
		case strings.HasPrefix(line, "#"):
		case strings.HasPrefix(line, "newmtl "):
			materials[name] = current
			name = strings.TrimPrefix(line, "newmtl ")
			current = Material{}
		case strings.HasPrefix(line, "Ka"):
			current.Ka = readThreeFloats(line, "Ka ")
		case strings.HasPrefix(line, "Kd"):
			current.Kd = readThreeFloats(line, "Kd ")
		case strings.HasPrefix(line, "Ks"):
			current.Ks = readThreeFloats(line, "Ks ")
		case strings.HasPrefix(line, "D"):
			current.D = readFloat(line, "D ")
		case strings.HasPrefix(line, "Tr"):
			current.Tr = readFloat(line, "Tr ")
		case strings.HasPrefix(line, "Ns"):
			current.Ns = readFloat(line, "Ns ")
		case strings.HasPrefix(line, "Illum"):
			current.Illum = readInt(line, "Illum ")
			//case strings.HasPrefix(line, "MapKa"):
			//current.Illum = readInt(line, "Illum ")
		}
	}
	materials[name] = current
	return materials, nil
}
