package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type kwindowDrawer2DBuilder struct {
	backgroundColor                mgl32.Vec4
	windowWidth, windowHeight, ppu float32
}

func KwindowDrawer2DBuilder() *kwindowDrawer2DBuilder {
	return &kwindowDrawer2DBuilder{
		ppu: 50,
	}
}

func (kwd2db *kwindowDrawer2DBuilder) SetBackgroundColor(color mgl32.Vec4) *kwindowDrawer2DBuilder {
	kwd2db.backgroundColor = color
	return kwd2db
}

func (kwd2db *kwindowDrawer2DBuilder) SetPixelPerUnit(ppu float32) *kwindowDrawer2DBuilder {
	kwd2db.ppu = ppu
	return kwd2db
}

func (kwd2db *kwindowDrawer2DBuilder) BuildTo(kwindowID KwindowID) (KwindowDrawer2DID, error) {
	kwd2db.windowWidth = float32(windows[kwindowID].width)
	kwd2db.windowHeight = float32(windows[kwindowID].height)
	drawer, err := newKwindowDrawer2D(*kwd2db)
	if err != nil {
		return KwindowDrawer2DID(-1), err
	}
	windows[kwindowID].kwindowDrawer = &drawer
	return KwindowDrawer2DID(kwindowID), nil
}
