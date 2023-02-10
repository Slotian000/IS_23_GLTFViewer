package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	_ "github.com/go-gl/mathgl/mgl32"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"openGL/Wrappers"
	"runtime"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var (
	vertices = []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	cubePositions = []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{1.5, -2.2, -2.5},
		{3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{1.3, 1.0, -1.5},
	}

	cameraPos   = mgl32.Vec3{0, 0, 3}
	cameraFront = mgl32.Vec3{0, 0, -1}
	cameraUp    = mgl32.Vec3{0, 1, 0}

	deltaTime = 0.0
	lastFrame = 0.0
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

	texture, err := Wrappers.NewTexture("Sources/container.jpg")
	if err != nil {
		fmt.Println(err)
	}
	texture.Bind()

	attributes := []Wrappers.VertexAttribute{
		Wrappers.NewVertexAttribute(3, gl.FLOAT, false),
		//Wrappers.NewVertexAttribute(3, gl.FLOAT, false),
		Wrappers.NewVertexAttribute(2, gl.FLOAT, false),
	}

	VAO := Wrappers.NewVAO(vertices, gl.STATIC_DRAW, attributes...)

	program.Use()
	VAO.Bind()

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/float32(windowHeight), 0.1, 100.0)
	program.SetMat4("projection", projection)

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.5, 0.5, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		currentFrame := glfw.GetTime()
		deltaTime = currentFrame - lastFrame
		lastFrame = currentFrame

		for _, vec := range cubePositions {
			program.SetMat4("model", mgl32.Translate3D(vec.X(), vec.Y(), vec.Z()))
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		view := mgl32.LookAtV(cameraPos, cameraPos.Add(cameraFront), cameraUp)
		program.SetMat4("view", view)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	speed := float32(2.5 * deltaTime)

	if key == glfw.KeyW {
		cameraPos = cameraPos.Add(cameraFront.Mul(speed))
	}
	if key == glfw.KeyS {
		cameraPos = cameraPos.Sub(cameraFront.Mul(speed))
	}
	if key == glfw.KeyA {
		cameraPos = cameraPos.Sub(cameraFront.Cross(cameraUp).Normalize().Mul(speed))
	}
	if key == glfw.KeyD {
		cameraPos = cameraPos.Add(cameraFront.Cross(cameraUp).Normalize().Mul(speed))
	}
	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}

}

func framebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

/*
	cameraPosition := mgl32.Vec3{0.0, 0.0, 3.0}
	cameraTarget := mgl32.Vec3{0.0, 0.0, 0.0}
	cameraDirection := cameraPosition.Sub(cameraTarget).Normalize()

	up := mgl32.Vec3{0.0, 1.0, 0.0}
	cameraRight := up.Cross(cameraDirection).Normalize()

	_ = cameraDirection.Cross(cameraRight).Normalize()
*/
