package Utils

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var Bindings = make(map[int]glfw.Key)

func init() {

}

type Camera struct {
	View        mgl32.Mat4 `json:"view"`
	Projection  mgl32.Mat4 `json:"projection"`
	CameraPos   mgl32.Vec3 `json:"cameraPos"`
	cameraFront mgl32.Vec3 `json:"cameraFront"`
	cameraUp    mgl32.Vec3 `json:"cameraUp"`
	fov         float32    `json:"fov"`
	aspectRatio float32    `json:"aspectRatio"`
	lastX       float32    `json:"lastX"`
	lastY       float32    `json:"lastY"`
	scrollY     float32    `json:"scrollY"`
	yaw         float32    `json:"yaw"`
	pitch       float32    `json:"pitch"`
	speed       float32    `json:"speed"`
	sensitivity float32    `json:"sensitivity"`
}

func NewCamera(width float32, height float32) Camera {
	camera := Camera{
		CameraPos:   mgl32.Vec3{0, 0, 3},
		cameraFront: mgl32.Vec3{0, 0, -1},
		cameraUp:    mgl32.Vec3{0, 1, 0},
		aspectRatio: width / height,
		lastX:       width / 2,
		lastY:       height / 2,
		fov:         45,
		yaw:         -90.0,
		pitch:       0,
		speed:       5,
		sensitivity: .05,
	}
	return camera
}

func (c *Camera) Update(deltaTime float32, cursorX float32, cursorY float32, scrollY float32, Active map[glfw.Key]bool, Bindings map[string]glfw.Key) {
	xOffset := (cursorX - c.lastX) * c.sensitivity
	yOffset := (c.lastY - cursorY) * c.sensitivity
	c.lastX, c.lastY = cursorX, cursorY

	c.yaw += xOffset
	c.pitch = MinMax(-89, c.pitch+yOffset, 89)
	c.fov = MinMax(1, c.fov-scrollY, 90)

	if Active[Bindings["WireFrame"]] {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}
	if Active[Bindings["CameraForward"]] {
		c.CameraPos = c.CameraPos.Add(c.cameraFront.Mul(deltaTime * c.speed))
	}
	if Active[Bindings["CameraBack"]] {
		c.CameraPos = c.CameraPos.Sub(c.cameraFront.Mul(deltaTime * c.speed))
	}
	if Active[Bindings["CameraLeft"]] {
		c.CameraPos = c.CameraPos.Sub(c.cameraFront.Cross(c.cameraUp).Normalize().Mul(deltaTime * c.speed))
	}
	if Active[Bindings["CameraRight"]] {
		c.CameraPos = c.CameraPos.Add(c.cameraFront.Cross(c.cameraUp).Normalize().Mul(deltaTime * c.speed))
	}
	if Active[Bindings["CameraSprint"]] {
		c.fov = 120
		c.speed = 10
	} else {
		c.speed = 5
	}

	c.cameraFront = mgl32.Vec3{
		Cos(mgl32.DegToRad(c.yaw)) * Cos(mgl32.DegToRad(c.pitch)),
		Sin(mgl32.DegToRad(c.pitch)),
		Sin(mgl32.DegToRad(c.yaw)) * Cos(mgl32.DegToRad(c.pitch)),
	}.Normalize()

	c.View = mgl32.LookAtV(c.CameraPos, c.CameraPos.Add(c.cameraFront), c.cameraUp)
	c.Projection = mgl32.Perspective(mgl32.DegToRad(c.fov), c.aspectRatio, 0.1, 1000.0)
}
