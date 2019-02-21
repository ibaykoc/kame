package kame

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type kdrawerCamera2D struct {
	position                  mgl32.Vec3
	front                     mgl32.Vec3
	right                     mgl32.Vec3
	near                      float32
	far                       float32
	up                        mgl32.Vec3
	worldUp                   mgl32.Vec3
	pixelPerUnit              float32
	fov                       float32
	pitch                     float32
	yaw                       float32
	roll                      float32
	windowWidth, windowHeight float32
}

func newkdrawerCamera2D(windowWidth, windowHeight, pixelPerUnit float32) kdrawerCamera2D {
	c := kdrawerCamera2D{
		position:     mgl32.Vec3{0, 0, 10},
		front:        mgl32.Vec3{0, 0, -1},
		worldUp:      mgl32.Vec3{0, 1, 0},
		yaw:          270,
		near:         0.1,
		far:          100,
		pixelPerUnit: pixelPerUnit,
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
	}
	c.updateVectors()
	return c
}

func (c *kdrawerCamera2D) updateVectors() {
	fX := math.Cos(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	fY := math.Sin(float64(mgl32.DegToRad(c.pitch)))
	fZ := math.Sin(float64(mgl32.DegToRad(c.yaw))) * math.Cos(float64(mgl32.DegToRad(c.pitch)))
	c.front = mgl32.Vec3{float32(fX), float32(fY), float32(fZ)}.Normalize()
	c.right = c.front.Cross(c.worldUp).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
}

func (c *kdrawerCamera2D) viewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		c.position, c.position.Add(c.front), c.up,
	)
}

func (c *kdrawerCamera2D) projectionMatrix() mgl32.Mat4 {
	return mgl32.Ortho(
		-c.windowWidth/c.pixelPerUnit/2,
		c.windowHeight/c.pixelPerUnit/2,
		-c.windowHeight/c.pixelPerUnit/2,
		c.windowHeight/c.pixelPerUnit/2,
		c.near, c.far)
}

func (c *kdrawerCamera2D) Move(x, y, z float32) {
	c.position = c.position.Add(c.front.Mul(z)).
		Add(c.right.Mul(x)).
		Add(c.up.Mul(y))
}

func (c *kdrawerCamera2D) updateFPSControl(timeSinceLastFrame float32) {
	moveSpeed := 0.05 * timeSinceLastFrame
	// rotateSensitivity := 0.5 * timeSinceLastFrame
	if windows[1] == nil {
		return
	}
	// fmt.Printf("Update cam control\n")
	input := windows[1].input
	// println(input.yScroll)
	c.pixelPerUnit *= (1 + (input.yScroll * 0.1))
	moveSpeed /= c.pixelPerUnit / 70
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
	c.Move(moveHInput*moveSpeed, moveVInput*moveSpeed, 0)

}
