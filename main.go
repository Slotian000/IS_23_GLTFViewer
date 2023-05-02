package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	_ "github.com/go-gl/mathgl/mgl32"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"openGL/Utils"
	"openGL/Wrappers"
	"runtime"
)

const (
	WindowWidth  = 800
	WindowHeight = 600
)

var (
	lastFrame = 0.0

	shaders = map[string]Wrappers.Program{
		"PNT":  LoadProgram("Shaders/PNT"),
		"PNTT": LoadProgram("Shaders/PNTT"),
	}
)

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
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Hello!", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetKeyCallback(KeyCallback)
	window.SetCursorPosCallback(CursorPosCallback)
	window.SetScrollCallback(ScrollCallBack)
	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	if err = gl.Init(); err != nil {
		panic(err)
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.5, 0.5, 1.0)
	meshes := test()
	loop(window, meshes)
}

func DeltaTime() float32 {
	currentFrame := glfw.GetTime()
	deltaTime := currentFrame - lastFrame
	lastFrame = currentFrame
	return float32(deltaTime)
}

func loop(window *glfw.Window, meshes []Mesh) {
	vertexShader, err := Wrappers.NewShaderFromFile("Sources/VertexShader.vert", gl.VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
	}

	fragmentShader, err := Wrappers.NewShaderFromFile("Sources/FragmentShader.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
	}

	program, err := Wrappers.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		fmt.Println(err)
	}

	program.Use()
	camera := Utils.NewCamera(WindowWidth, WindowHeight)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		program.SetMat4("model", mgl32.Translate3D(0, 0, 0))

		for _, mesh := range meshes {
			mesh.VAO.Bind()
			gl.DrawElementsWithOffset(gl.TRIANGLES, int32(mesh.VAO.Count), gl.UNSIGNED_INT, 0)
		}

		window.SwapBuffers()
		glfw.PollEvents()
		camera.Update(DeltaTime(), float32(CursorX), float32(CursorY), float32(ScrollY), Active, Bindings)
		program.SetMat4("view", camera.View)
		program.SetMat4("projection", camera.Projection)
	}
}

func framebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func LoadProgram(path string) Wrappers.Program {
	vertexShader, err := Wrappers.NewShaderFromFile(path+"/VertexShader.vert", gl.VERTEX_SHADER)
	if err != nil {
		fmt.Println(err)
	}
	fragmentShader, err := Wrappers.NewShaderFromFile(path+"/FragmentShader.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println(err)
	}
	program, err := Wrappers.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		fmt.Println(err)
	}
	return program
}
