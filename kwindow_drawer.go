package kame

import "github.com/go-gl/mathgl/mgl32"

type KwindowDrawerController struct {
	kwindowDrawer *kwindowDrawer
}

type KwindowDrawer interface {
	clear()
	draw()
	AppendDrawable(drawable Kdrawable, translation mgl32.Mat4)
	GetCamera() *kdrawerCamera
	onWindowSizeChange(newWidth, newHeight float32)
}

type kwindowDrawer struct {
	backgroundColor  Kcolor
	kdrawerCamera    kdrawerCamera
	defaultkshaderID kshaderID
	kshaders         map[kshaderID]*kshader
	kmeshes          map[kmeshID]kmesh
	ktextures        map[KtextureID]ktexture
}

func (wdCon KwindowDrawerController) DefaultShaderID() kshaderID {
	return wdCon.kwindowDrawer.defaultkshaderID
}

func (wdCon KwindowDrawerController) StoreTexturePNG(path string) (KtextureID, error) {
	ktex, err := newktextureFromPNG(path)
	if err != nil {
		return KtextureID(0), err
	}
	wdCon.kwindowDrawer.ktextures[ktex.id] = ktex
	return ktex.id, nil
}
