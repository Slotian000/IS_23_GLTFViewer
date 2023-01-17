package Wrappers

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type Program struct {
	handle uint32
}

func (p *Program) Use() {
	gl.UseProgram(p.handle)
}

func New(shaders ...Shader) (program Program, err error) {
	program.handle = gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program.handle, shader.handle)
		shader.Delete()
	}
	gl.LinkProgram(program.handle)

	var success int32 = 0
	if gl.GetShaderiv(program.handle, gl.LINK_STATUS, &success); success != 0 {
		infoLog := strings.Repeat("\x00", 512)
		gl.GetShaderInfoLog(program.handle, 512, nil, gl.Str(infoLog))
		err = errors.New("PROGRAM::LINKING_FAILURE")
	}
	return
}
