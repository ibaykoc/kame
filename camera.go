package kame

import (
	"math"

	mgl "github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position mgl.Vec3
	front    mgl.Vec3
	right    mgl.Vec3
	up       mgl.Vec3

	worldUp mgl.Vec3

	pitch float32
	yaw   float32
	roll  float32
}

func CreateCamera() Camera {
	c := Camera{
		position: mgl.Vec3{0, 0, 10},
		front:    mgl.Vec3{0, 0, -1},
		worldUp:  mgl.Vec3{0, 1, 0},
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
	fX := math.Cos(float64(mgl.DegToRad(c.yaw))) * math.Cos(float64(mgl.DegToRad(c.pitch)))
	fY := math.Sin(float64(mgl.DegToRad(c.pitch)))
	fZ := math.Sin(float64(mgl.DegToRad(c.yaw))) * math.Cos(float64(mgl.DegToRad(c.pitch)))
	c.front = mgl.Vec3{float32(fX), float32(fY), float32(fZ)}.Normalize()
	c.right = c.front.Cross(c.worldUp).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
	// fmt.Printf("pos:(%6.3f, %6.3f, %6.3f), pyr:(%6.3f, %6.3f, %6.3f), front:(%6.3f, %6.3f, %6.3f), up:(%6.3f, %6.3f, %6.3f), right:(%6.3f, %6.3f, %6.3f)\n", c.position.X(), c.position.Y(), c.position.Z(), c.pitch, c.yaw, c.roll, c.front.X(), c.front.Y(), c.front.Z(), c.up.X(), c.up.Y(), c.up.Z(), c.right.X(), c.right.Y(), c.right.Z())
}

func (c *Camera) viewMatrix() mgl.Mat4 {
	return mgl.LookAtV(
		c.position, c.position.Add(c.front), c.up,
	)
}
