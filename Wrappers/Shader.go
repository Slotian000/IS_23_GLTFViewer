package Wrappers

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"os"
	"strings"
)

type Shader struct {
	handle uint32
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.handle)
}

func NewFromFile(path string, shaderType uint32) (Shader, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Shader{}, err
	}

	handle := gl.CreateShader(shaderType)
	source, freeFn := gl.Strs(string(file))
	defer freeFn()

	gl.ShaderSource(handle, 1, source, nil)
	gl.CompileShader(handle)
	shader := Shader{handle: handle}

	return Shader{}, shader.compileError()
}

func (s *Shader) compileError() error {
	var success int32 = 0
	if gl.GetShaderiv(s.handle, gl.COMPILE_STATUS, &success); success != 0 {
		infoLog := strings.Repeat("\x00", int(512))
		gl.GetShaderInfoLog(s.handle, int32(len(infoLog)), nil, gl.Str(infoLog))
		return errors.New("SHADER::COMPILE_FAILURE")
	}
	return nil
}
