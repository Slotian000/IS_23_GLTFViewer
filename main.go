package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"openGL/Wrappers"
	"runtime"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var vertices = []float32{
	0.5, 0.5, 0.0, 1.0, 0.0, 0.0, //1.0, 1.0,
	0.5, -0.5, 0.0, 0.0, 1.0, 0.0, //1.0, 0.0,
	-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, //0.0, 1.0,
	0.5, -0.5, 0.0, 0.0, 1.0, 0.0, //1.0, 0.0,
	-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, //0.0, 0.0,
	-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, //0.0, 1.0,
}

var indices = []uint32{
	0, 1, 3,
	1, 2, 3,
}

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Hello!", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(keyCallback)
	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	if err = gl.Init(); err != nil {
		panic(err)
	}
	loop(window)
}

func loop(window *glfw.Window) {

	vertexShader, err := Wrappers.NewShaderFromFile("Sources/default.vert", gl.VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
	}

	fragmentShader, err := Wrappers.NewShaderFromFile("Sources/default.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
	}

	program, err := Wrappers.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		fmt.Println(err)
	}

	attributes := []Wrappers.VertexAttribute{
		Wrappers.NewVertexAttribute(3, gl.FLOAT, false),
		Wrappers.NewVertexAttribute(3, gl.FLOAT, false),
	}

	VAO := Wrappers.NewVAO(vertices, gl.STATIC_DRAW, attributes...)

	for !window.ShouldClose() {
		gl.ClearColor(0.2, 0.5, 0.5, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		VAO.Bind()

		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func framebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
