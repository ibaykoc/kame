package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type kwindowDrawer3DBuilder struct {
	backgroundColor                Kcolor
	windowWidth, windowHeight, fov float32
}

func KwindowDrawer3DBuilder() *kwindowDrawer3DBuilder {
	return &kwindowDrawer3DBuilder{
		fov: mgl32.DegToRad(90),
	}
}

func (kwd3Db *kwindowDrawer3DBuilder) SetBackgroundColor(color Kcolor) *kwindowDrawer3DBuilder {
	kwd3Db.backgroundColor = color
	return kwd3Db
}

func (kwd3Db *kwindowDrawer3DBuilder) SetFieldOfView(fov float32) *kwindowDrawer3DBuilder {
	kwd3Db.fov = fov
	return kwd3Db
}

func (kwd3Db *kwindowDrawer3DBuilder) BuildTo(kwindowID KwindowID) (KwindowDrawer3DController, error) {
	kwd3Db.windowWidth = float32(windows[kwindowID].width)
	kwd3Db.windowHeight = float32(windows[kwindowID].height)
	drawer, err := newKwindowDrawer3D(*kwd3Db)
	if err != nil {
		return KwindowDrawer3DController{}, err
	}
	windows[kwindowID].kwindowDrawer = &drawer
	return KwindowDrawer3DController{
		KwindowDrawerController: KwindowDrawerController{
			kwindowDrawer: &drawer.kwindowDrawer,
		},
		kwindowDrawer3D: &drawer,
	}, nil
}
