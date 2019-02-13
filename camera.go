package kame

import (
	"math"

	mgl "github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position mgl.Vec3
	front    mgl.Vec3
	up       mgl.Vec3

	pitch float32
	yaw   float32
	roll  float32
}

func CreateCamera(x, y, z float32) Camera {
	c := Camera{
		position: mgl.Vec3{0, 0, -10},
		front:    mgl.Vec3{0, 0, 1},
		up:       mgl.Vec3{0, 1, 0},
		yaw:      90,
	}
	return c
}

func (c *Camera) Move(x, y, z float32) {
	c.position = c.position.Add(c.front.Mul(z)).
		Add(c.front.Cross(c.up).Normalize().Mul(x))
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
	fX := math.Cos(float64(mgl.DegToRad(c.yaw))) * math.Cos(float64(mgl.DegToRad(c.pitch)))
	fY := math.Sin(float64(mgl.DegToRad(c.pitch)))
	fZ := math.Sin(float64(mgl.DegToRad(c.yaw))) * math.Cos(float64(mgl.DegToRad(c.pitch)))
	c.front = mgl.Vec3{float32(fX), float32(fY), float32(fZ)}.Normalize()
}

func (c *Camera) viewMatrix() mgl.Mat4 {
	return mgl.LookAtV(
		c.position, c.position.Add(c.front), c.up,
	)
}
