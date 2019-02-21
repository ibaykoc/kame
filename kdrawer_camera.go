package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type kdrawerCamera interface {
	viewMatrix() mgl32.Mat4
	updateFPSControl(timeSinceLastFrame float32)
}
