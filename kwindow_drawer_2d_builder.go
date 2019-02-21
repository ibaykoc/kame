package kame

type kwindowDrawer2DBuilder struct {
	backgroundColor                Kcolor
	windowWidth, windowHeight, ppu float32
}

func KwindowDrawer2DBuilder() *kwindowDrawer2DBuilder {
	return &kwindowDrawer2DBuilder{
		ppu: 50,
	}
}

func (kwd2db *kwindowDrawer2DBuilder) SetBackgroundColor(color Kcolor) *kwindowDrawer2DBuilder {
	kwd2db.backgroundColor = color
	return kwd2db
}

func (kwd2db *kwindowDrawer2DBuilder) SetPixelPerUnit(ppu float32) *kwindowDrawer2DBuilder {
	kwd2db.ppu = ppu
	return kwd2db
}

func (kwd2db *kwindowDrawer2DBuilder) BuildTo(kwindowID KwindowID) (KwindowDrawer2DController, error) {
	kwd2db.windowWidth = float32(windows[kwindowID].width)
	kwd2db.windowHeight = float32(windows[kwindowID].height)
	drawer, err := newKwindowDrawer2D(*kwd2db)
	if err != nil {
		return KwindowDrawer2DController{}, err
	}
	windows[kwindowID].kwindowDrawer = &drawer
	return KwindowDrawer2DController{
		KwindowDrawerController: KwindowDrawerController{
			kwindowDrawer: &drawer.kwindowDrawer,
		},
		kwindowDrawer2D: &drawer,
	}, nil
}
