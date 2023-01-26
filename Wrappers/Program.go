package Wrappers

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Program struct {
	ID uint32
}

func (p *Program) Use() {
	gl.UseProgram(p.ID)
}

func NewProgram(shaders ...Shader) (Program, error) {
	program := Program{ID: gl.CreateProgram()}

	for _, shader := range shaders {
		gl.AttachShader(program.ID, shader.ID)
	}
	gl.LinkProgram(program.ID)

	var status, length int32
	if gl.GetProgramiv(program.ID, gl.LINK_STATUS, &status); status == gl.FALSE {
		gl.GetProgramiv(program.ID, gl.INFO_LOG_LENGTH, &length)
		infoLog := make([]byte, length)
		gl.GetProgramInfoLog(program.ID, length, nil, &infoLog[0])
		return Program{}, errors.New(string(infoLog))
	}
	return Program{ID: program.ID}, nil
}
