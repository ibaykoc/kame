package kame

import "github.com/go-gl/mathgl/mgl32"

type KwindowDrawerID int

type KwindowDrawer interface {
	DefaultShaderID() kshaderID
	clear()
	AppendDrawable(drawable Kdrawable, translation mgl32.Mat4)
	draw()
	GetCamera() *kdrawerCamera
}

func (kwdID KwindowDrawerID) DefaultShaderID() kshaderID {
	return windows[KwindowID(kwdID)].kwindowDrawer.DefaultShaderID()
}

func (kwdID KwindowDrawerID) CreateDrawable() kshaderID {
	return windows[KwindowID(kwdID)].kwindowDrawer.DefaultShaderID()
}
