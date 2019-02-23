package kame

type kwindowBuilder struct {
	title                                   string
	width, height, x, y, targetFPS          int
	fullscreen, windowed, center, resizable bool
	processInputFunc                        processInputFunc
	updateFunc                              updateFunc
	drawFunc                                drawFunc
	onDropFileFunc                          onDropFileFunc
}

func KwindowBuilder() *kwindowBuilder {
	return &kwindowBuilder{}
}

func (kwb *kwindowBuilder) SetTitle(title string) *kwindowBuilder {
	kwb.title = title
	return kwb
}

func (kwb *kwindowBuilder) SetSize(width, height int) *kwindowBuilder {
	kwb.width = width
	kwb.height = height
	return kwb
}

func (kwb *kwindowBuilder) SetPosition(x, y int) *kwindowBuilder {
	kwb.x = x
	kwb.y = y
	return kwb
}

func (kwb *kwindowBuilder) SetTargetFPS(targetFPS int) *kwindowBuilder {
	kwb.targetFPS = targetFPS
	return kwb
}

func (kwb *kwindowBuilder) SetProcessInputFunc(processInputFunc processInputFunc) *kwindowBuilder {
	kwb.processInputFunc = processInputFunc
	return kwb
}
func (kwb *kwindowBuilder) SetUpdateFunc(updateFunc updateFunc) *kwindowBuilder {
	kwb.updateFunc = updateFunc
	return kwb
}

func (kwb *kwindowBuilder) SetDrawFunc(drawFunc drawFunc) *kwindowBuilder {
	kwb.drawFunc = drawFunc
	return kwb
}

func (kwb *kwindowBuilder) SetOnDropFileFunc(onDropFileFunc onDropFileFunc) *kwindowBuilder {
	kwb.onDropFileFunc = onDropFileFunc
	return kwb
}

func (kwb *kwindowBuilder) IsWindowed() *kwindowBuilder {
	kwb.windowed = true
	return kwb
}

func (kwb *kwindowBuilder) IsResizable() *kwindowBuilder {
	kwb.resizable = true
	return kwb
}

func (kwb *kwindowBuilder) IsFullscreen() *kwindowBuilder {
	kwb.fullscreen = true
	return kwb
}

func (kwb *kwindowBuilder) Build() (KwindowController, error) {
	w, err := newKwindow(*kwb)
	if err != nil {
		return KwindowController{}, err
	}
	w.id = KwindowID(len(windows))
	windows[w.id] = w
	return KwindowController{
		window: windows[w.id],
	}, nil
}
