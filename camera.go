package kame

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type ProjectionType int

const (
	Orthographic ProjectionType = iota
	Perspective
)

type Camera struct {
	position       mgl32.Vec3
	front          mgl32.Vec3
	right          mgl32.Vec3
	up             mgl32.Vec3
	projectionType ProjectionType
	worldUp        mgl32.Vec3
	pixelPerUnit   float32
	fov            float32
	pitch          float32
	yaw            float32
	roll           float32
}

func createCamera3D(fov float32) Camera {
	c := Camera{
		position:       mgl32.Vec3{0, 0, 10},
		front:          mgl32.Vec3{0, 0, -1},
		worldUp:        mgl32.Vec3{0, 1, 0},
		fov:            fov,
		yaw:            270,
		projectionType: Perspective,
	}
	c.updateVectors()
	return c
}

func createCamera2D(pixelPerUnit float32) Camera {
	c := Camera{
		position:       mgl32.Vec3{0, 0, 10},
		front:          mgl32.Vec3{0, 0, -1},
		worldUp:        mgl32.Vec3{0, 1, 0},
		yaw:            270,
		pixelPerUnit:   pixelPerUnit,
		projectionType: Orthographic,
	}
	c.updateVectors()
	return c
}

func (c *Camera) Move(x, y, z float32) {
	c.position = c.position.Add(c.front.Mul(z)).
		Add(c.right.Mul(x)).
		Add(c.up.Mul(y))
}

func (c *Camera) Rotate(pitch, yaw, roll float32) {
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

func (c *Camera) updateVectors() {
	fX := math.Cos(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	fY := math.Sin(float64(mgl32.DegToRad(c.pitch)))
	fZ := math.Sin(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	c.front = mgl32.Vec3{float32(fX), float32(fY), float32(fZ)}.Normalize()
	c.right = c.front.Cross(c.worldUp).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
}

func (c *Camera) viewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		c.position, c.position.Add(c.front), c.up,
	)
}

func (c *Camera) projectionMatrix() mgl32.Mat4 {
	if c.projectionType == Orthographic {
		return mgl32.Ortho(0, float32(window.width)/c.pixelPerUnit, -float32(window.height)/c.pixelPerUnit, 0, -100, 100)
	}
	return mgl32.Perspective(c.fov, float32(window.width)/float32(window.height), 0.1, 100)
}

func (c *Camera) updateFPSControl(timeSinceLastFrame float32) {
	moveSpeed := 0.05 * timeSinceLastFrame
	rotateSensitivity := 0.5 * timeSinceLastFrame
	input := window.input
	moveHInput := float32(0)
	moveVInput := float32(0)
	if input.GetKeyStat(KeyLeftShift) == Press {
		moveSpeed *= 5
	}
	if input.GetKeyStat(KeyW) == Press {
		moveVInput++
	}
	if input.GetKeyStat(KeyS) == Press {
		moveVInput--
	}
	if input.GetKeyStat(KeyD) == Press {
		moveHInput++
	}
	if input.GetKeyStat(KeyA) == Press {
		moveHInput--
	}
	if c.projectionType == Perspective {
		c.Move(moveHInput*moveSpeed, 0, moveVInput*moveSpeed)
		mDX := input.MouseDeltaX * rotateSensitivity
		mDY := input.MouseDeltaY * rotateSensitivity
		c.Rotate(-mDY, mDX, 0)
	} else {
		c.Move(moveHInput*moveSpeed, moveVInput*moveSpeed, 0)
	}
}
