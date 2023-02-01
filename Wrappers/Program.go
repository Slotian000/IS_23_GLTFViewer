package Wrappers

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Program struct {
	ID uint32
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

func (p *Program) Use() {
	gl.UseProgram(p.ID)
}

func (p *Program) GetAttributeLocation(attribute string) int32 {
	return gl.GetAttribLocation(p.ID, gl.Str(attribute+"\x00"))
}

func (p *Program) SetMat4(name string, mat4 mgl32.Mat4) {
	pos := p.getAttributeLocation(name)
	gl.Uniform1fv(pos, 1, &mat4[0])
}
