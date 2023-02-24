package Object

import (
	"os"
	"strings"
)

type Material struct {
	Ka    [3]float32
	Kd    [3]float32
	Ks    [3]float32
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

	current := ""

	for _, line := range strings.Split(string(file), "\n") {

		switch {

		case line == "":
		case strings.HasPrefix(line, "#"):
		case strings.HasPrefix(line, "newmtl "):
			current := strings.TrimPrefix(line, "newmtl ")
			materials[current] = Material{}

		case strings.HasPrefix(line, "Ka ") || strings.HasPrefix(line, "Kd ") || strings.HasPrefix(line, "Ks "):

		}
	}
}
