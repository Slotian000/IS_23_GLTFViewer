package Wrappers

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"os"
)

type Shader struct {
	ID uint32
}

func NewShaderFromFile(path string, shaderType uint32) (Shader, error) {
	shader := Shader{ID: gl.CreateShader(shaderType)}

	file, err := os.ReadFile(path)
	if err != nil {
		return Shader{}, err
	}

	source, free := gl.Strs(string(file) + " \x00")
	defer free()

	gl.ShaderSource(shader.ID, 1, source, nil)
	gl.CompileShader(shader.ID)

	var status, length int32
	if gl.GetShaderiv(shader.ID, gl.COMPILE_STATUS, &status); status == gl.FALSE {
		gl.GetShaderiv(shader.ID, gl.INFO_LOG_LENGTH, &length)
		infoLog := make([]byte, length)
		gl.GetShaderInfoLog(shader.ID, length, nil, &infoLog[0])
		return shader, errors.New(string(infoLog))
	}

	return shader, nil
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.ID)
}
