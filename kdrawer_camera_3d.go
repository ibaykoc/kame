package kame

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Kfrustum struct {
	NearPlane Krect
	FarPlane  Krect
}

type KdrawerCamera3DController struct {
	camera *kdrawerCamera3D
}

// func (cam3dCon KdrawerCamera3DController) ScreenToWorldPos(screenPos mgl32.Vec2, zPercent float32) mgl32.Vec3 {
// 	// MousePos on  19.00, 16.0000006.2f
// 	// MouseWorldPos on [9.770483 -9.820934 -0.090001106]
// 	// Near Width : 0.2 Far Width: 200
// 	// MousePos on   8.00, 589.0000006.2f
// 	// MouseWorldPos on [9.955467 -0.1849836 -0.090001106]
// 	f := cam3dCon.camera.frustum()
// 	xNormalize := screenPos.X() / cam3dCon.camera.windowWidth
// 	yNormalize := (cam3dCon.camera.windowHeight - screenPos.Y()) / cam3dCon.camera.windowHeight

// 	nearXPos := lerpf(f.NearPlane.BottomLeft.X(), f.NearPlane.BottomRight.X(), xNormalize)
// 	nearYPos := lerpf(f.NearPlane.BottomLeft.Y(), f.NearPlane.TopLeft.Y(), yNormalize)
// 	nearZPos := lerpf(f.NearPlane.BottomLeft.Z(), f.NearPlane.TopRight.Z(), zPercent)
// 	nearPos := mgl32.Vec3{nearXPos, nearYPos, nearZPos}

// 	farXPos := lerpf(f.FarPlane.BottomLeft.X(), f.FarPlane.BottomRight.X(), xNormalize)
// 	farYPos := lerpf(f.FarPlane.BottomLeft.Y(), f.FarPlane.TopLeft.Y(), yNormalize)
// 	farZPos := lerpf(f.FarPlane.BottomLeft.Z(), f.FarPlane.TopRight.Z(), zPercent)
// 	farPos := mgl32.Vec3{farXPos, farYPos, farZPos}

// 	worldPos := lerpV3(nearPos, farPos, zPercent)
// 	// debug.pf("xT: %1.2f, yT: %1.2f, nearPos: %v, farPos: %v, worldPos: %v\n", xNormalize, yNormalize, nearPos, farPos, worldPos)
// 	debug.pf("xT: %1.2f, yT: %1.2f, camFarBottomLeft: %v, camFarTopRight: %v, farPos: %v, worldPos: %v\n", xNormalize, yNormalize, f.FarPlane.BottomLeft, farPos, worldPos)
// 	return worldPos
// }

func lerpV3(start, end mgl32.Vec3, percent float32) mgl32.Vec3 {
	return start.Mul(1 - percent).Add(end.Mul(percent))
}

func lerpf(start, end, percent float32) float32 {
	return (1-percent)*start + percent*end
}

func (cam3dCon KdrawerCamera3DController) Frustum() Kfrustum {
	return cam3dCon.camera.frustum()
}

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

func (c *kdrawerCamera3D) frustum() Kfrustum {
	var f Kfrustum

	nearCenter := c.position.Add(c.front.Mul(c.near))
	farCenter := c.position.Add(c.front.Mul(c.far))

	nearHeight := float32(2) * float32(math.Tan(float64((c.fov / 2)))) * c.near
	farHeight := float32(2) * float32(math.Tan(float64((c.fov)/2))) * c.far
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

func (cam3dCon KdrawerCamera3DController) ScreenToWorldPos(screenPos mgl32.Vec3) mgl32.Vec3 {

	nearPos, err := mgl32.UnProject(
		mgl32.Vec3{screenPos.X(), cam3dCon.camera.windowHeight - screenPos.Y(), 0},
		cam3dCon.camera.viewMatrix(), cam3dCon.camera.projectionMatrix(), 0, 0, int(cam3dCon.camera.windowWidth), int(cam3dCon.camera.windowHeight))
	if err != nil {
		panic(err)
	}
	dir := nearPos.Add(cam3dCon.camera.position.Mul(-1)).Normalize()
	pos := nearPos.Add(dir.Mul(screenPos.Z()))
	return pos
}
