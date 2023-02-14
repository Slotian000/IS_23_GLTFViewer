package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"openGL/Wrappers"
)

type Camera struct {
	cameraPos   mgl32.Vec3
	cameraFront mgl32.Vec3
	cameraUp    mgl32.Vec3

	speed float32

	program *Wrappers.Program
}
