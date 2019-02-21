package kame

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type kdrawerCamera3D struct {
	position                  mgl32.Vec3
	front                     mgl32.Vec3
	right                     mgl32.Vec3
	near                      float32
	far                       float32
	up                        mgl32.Vec3
	worldUp                   mgl32.Vec3
	fov                       float32
	pitch                     float32
	yaw                       float32
	roll                      float32
	windowWidth, windowHeight float32
}

func newkdrawerCamera3D(windowWidth, windowHeight, fov float32) kdrawerCamera3D {
	c := kdrawerCamera3D{
		position:     mgl32.Vec3{0, 0, 10},
		front:        mgl32.Vec3{0, 0, -1},
		worldUp:      mgl32.Vec3{0, 1, 0},
		yaw:          270,
		near:         0.1,
		far:          100,
		fov:          fov,
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
	}
	c.updateVectors()
	return c
}

func (c *kdrawerCamera3D) Rotate(pitch, yaw, roll float32) {
	c.pitch += pitch
	c.yaw += yaw
	c.roll += roll
	if c.pitch > 89.0 {
		c.pitch = 89.0
	} else if c.pitch < -89.0 {
		c.pitch = -89.0
	}
	c.updateVectors()
}

func (c *kdrawerCamera3D) updateVectors() {
	fX := math.Cos(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	fY := math.Sin(float64(mgl32.DegToRad(c.pitch)))
	fZ := math.Sin(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	c.front = mgl32.Vec3{float32(fX), float32(fY), float32(fZ)}.Normalize()
	c.right = c.front.Cross(c.worldUp).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
}

func (c *kdrawerCamera3D) viewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		c.position, c.position.Add(c.front), c.up,
	)
}

func (c *kdrawerCamera3D) projectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(c.fov, c.windowWidth/c.windowHeight, c.near, c.far)
}

func (c *kdrawerCamera3D) Move(x, y, z float32) {
	c.position = c.position.Add(c.front.Mul(z)).
		Add(c.right.Mul(x)).
		Add(c.up.Mul(y))
}

func (c *kdrawerCamera3D) onWindowSizeChange(newWidth, newHeight float32) {
	c.windowWidth = newWidth
	c.windowHeight = newHeight
}

func (c *kdrawerCamera3D) updateFPSControl(windowInput KwindowInput, timeSinceLastFrame float32) {
	moveSpeed := 0.05 * timeSinceLastFrame
	rotateSensitivity := 0.5 * timeSinceLastFrame
	moveHInput := float32(0)
	moveVInput := float32(0)
	if windowInput.GetKeyStat(KeyLeftShift) == Press {
		moveSpeed *= 5
	}
	if windowInput.GetKeyStat(KeyW) == Press {
		moveVInput++
	}
	if windowInput.GetKeyStat(KeyS) == Press {
		moveVInput--
	}
	if windowInput.GetKeyStat(KeyD) == Press {
		moveHInput++
	}
	if windowInput.GetKeyStat(KeyA) == Press {
		moveHInput--
	}
	c.Move(moveHInput*moveSpeed, 0, moveVInput*moveSpeed)
	mDX := windowInput.mouseDeltaX * rotateSensitivity
	mDY := windowInput.mouseDeltaY * rotateSensitivity
	c.Rotate(-mDY, mDX, 0)
}
