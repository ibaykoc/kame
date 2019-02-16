package kame

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position mgl32.Vec3
	front    mgl32.Vec3
	right    mgl32.Vec3
	up       mgl32.Vec3

	worldUp mgl32.Vec3

	fov   float32
	pitch float32
	yaw   float32
	roll  float32
}

func createCamera(fov float32) Camera {
	c := Camera{
		position: mgl32.Vec3{0, 0, 2},
		front:    mgl32.Vec3{0, 0, -1},
		worldUp:  mgl32.Vec3{0, 1, 0},
		fov:      fov,
		yaw:      270,
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

func (c *Camera) updateFPSControl(timeSinceLastFrame float32) {
	moveSpeed := 0.05 * timeSinceLastFrame
	rotateSensitivity := 0.5 * timeSinceLastFrame
	input := window.input
	moveXInput := float32(0)
	moveZInput := float32(0)
	if input.GetKeyStat(KeyLeftShift) == Press {
		moveSpeed *= 5
	}
	if input.GetKeyStat(KeyW) == Press {
		moveZInput++
	}
	if input.GetKeyStat(KeyS) == Press {
		moveZInput--
	}
	if input.GetKeyStat(KeyD) == Press {
		moveXInput++
	}
	if input.GetKeyStat(KeyA) == Press {
		moveXInput--
	}
	mDX := input.MouseDeltaX * rotateSensitivity
	mDY := input.MouseDeltaY * rotateSensitivity
	c.Move(moveXInput*moveSpeed, 0, moveZInput*moveSpeed)
	c.Rotate(-mDY, mDX, 0)
}
