package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type kdrawerCamera interface {
	viewMatrix() mgl32.Mat4
	projectionMatrix() mgl32.Mat4
	updateFPSControl(windowInput KwindowInput, timeSinceLastFrame float32)
	onWindowSizeChange(newWidth, newHeight float32)
	frustum() Kfrustum
}
