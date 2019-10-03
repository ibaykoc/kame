package kame

type kgameWindowBuilder struct {
	wb *kwindowBuilder
}

func KgameWindowBuilder() *kgameWindowBuilder {
	return &kgameWindowBuilder{
		wb: &kwindowBuilder{},
	}
}

func (wb *kgameWindowBuilder) SetTitle(title string) *kgameWindowBuilder {
	wb.wb.title = title
	return wb
}

func (wb *kgameWindowBuilder) SetSize(width, height int) *kgameWindowBuilder {
	wb.wb.width = width
	wb.wb.height = height
	return wb
}

func (wb *kgameWindowBuilder) SetPosition(x, y int) *kgameWindowBuilder {
	wb.wb.x = x
	wb.wb.y = y
	return wb
}

func (wb *kgameWindowBuilder) IsWindowed() *kgameWindowBuilder {
	wb.wb.windowed = true
	return wb
}

func (wb *kgameWindowBuilder) IsResizable() *kgameWindowBuilder {
	wb.wb.resizable = true
	return wb
}

func (wb *kgameWindowBuilder) IsFullscreen() *kgameWindowBuilder {
	wb.wb.fullscreen = true
	return wb
}

func (wb *kgameWindowBuilder) BuildWith(scenes []Scene) (*KGameWindow, error) {
	wb.wb.SetTargetFPS(60)
	kgw := KGameWindow{
		scenes: scenes,
	}
	wb.wb.SetDrawFunc(kgw.draw)
	wb.wb.SetUpdateFunc(kgw.update)
	wb.wb.SetProcessInputFunc(kgw.processInput)
	w, err := newKwindow(*wb.wb)
	if err != nil {
		return nil, err
	}
	w.id = KwindowID(len(windows))
	windows[w.id] = w
	kgw.KwindowController = &KwindowController{
		window: windows[w.id],
	}
	return &kgw, nil
}
