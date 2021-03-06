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
	hW := c.windowWidth / c.pixelPerUnit / 2
	hH := c.windowHeight / c.pixelPerUnit / 2
	return mgl32.Ortho(
		-hW,
		hW,
		-hH,
		hH,
		c.near, c.far)
}

func (c *kdrawerCamera2D) Move(x, y, z float32) {
	c.position = c.position.Add(c.front.Mul(z)).
		Add(c.right.Mul(x)).
		Add(c.up.Mul(y))
}

func (c *kdrawerCamera2D) onWindowSizeChange(newWidth, newHeight float32) {
	c.windowWidth = newWidth
	c.windowHeight = newHeight
}

func (c *kdrawerCamera2D) frustum() Kfrustum {
	var f Kfrustum

	nearCenter := c.position.Add(c.front.Mul(c.near))
	farCenter := c.position.Add(c.front.Mul(c.far))

	nearHeight := c.windowHeight / c.pixelPerUnit
	farHeight := nearHeight
	nearWidth := nearHeight * c.windowWidth / c.windowHeight
	farWidth := farHeight * c.windowWidth / c.windowHeight

	nearTopLeft := nearCenter.Add(c.right.Mul(-nearWidth / 2)).Add(c.up.Mul(nearHeight / 2))
	nearTopRight := nearCenter.Add(c.right.Mul(nearWidth / 2)).Add(c.up.Mul(nearHeight / 2))
	nearBottomLeft := nearCenter.Add(c.right.Mul(-nearWidth / 2)).Add(c.up.Mul(-nearHeight / 2))
	nearBottomRight := nearCenter.Add(c.right.Mul(nearWidth / 2)).Add(c.up.Mul(-nearHeight / 2))

	farTopLeft := farCenter.Add(c.right.Mul(-farWidth / 2)).Add(c.up.Mul(farHeight / 2))
	farTopRight := farCenter.Add(c.right.Mul(farWidth / 2)).Add(c.up.Mul(farHeight / 2))
	farBottomLeft := farCenter.Add(c.right.Mul(-farWidth / 2)).Add(c.up.Mul(-farHeight / 2))
	farBottomRight := farCenter.Add(c.right.Mul(farWidth / 2)).Add(c.up.Mul(-farHeight / 2))

	nearRect := Krect{
		TopLeft:     nearTopLeft,
		TopRight:    nearTopRight,
		BottomLeft:  nearBottomLeft,
		BottomRight: nearBottomRight,
		Center:      nearCenter,
	}
	farRect := Krect{
		TopLeft:     farTopLeft,
		TopRight:    farTopRight,
		BottomLeft:  farBottomLeft,
		BottomRight: farBottomRight,
		Center:      farCenter,
	}
	f.NearPlane = nearRect
	f.FarPlane = farRect

	return f
}

func (c *kdrawerCamera2D) updateFPSControl(windowInput KwindowInput, timeSinceLastFrame float32) {
	moveSpeed := 0.05 * timeSinceLastFrame
	c.pixelPerUnit *= (1 + (windowInput.yScroll * 0.1))
	moveSpeed /= c.pixelPerUnit / 70
	if windowInput.mouseButtonStats[MouseButtonLeft] == Press {
		c.Move(-windowInput.mouseDeltaX/c.pixelPerUnit, windowInput.mouseDeltaY/c.pixelPerUnit, 0)
		return
	}
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
	c.Move(moveHInput*moveSpeed, moveVInput*moveSpeed, 0)
}
