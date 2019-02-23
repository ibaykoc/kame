// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type KwindowID int

type KwindowController struct {
	window *kwindow
}

func (kwCon KwindowController) ID() KwindowID {
	return kwCon.window.id
}

func (kwCon KwindowController) Close() {
	kwCon.window.Close()
}

func (kwCon KwindowController) EnableCameraMovementControl(enable bool) {
	kwCon.window.cameraControlEnabled = enable
}

func (kwCon KwindowController) LockCursor(lock bool) {
	kwCon.window.LockCursor(lock)
}

type updateFunc func(timeSinceLastFrame float32)
type drawFunc func(kwindowDrawer *KwindowDrawer)
type processInputFunc func(windowInput KwindowInput)
type onDropFileFunc func(mouseX, mouseY float32, filePath string)

// Kwindow window for kame
type kwindow struct {
	id                            KwindowID
	title                         string
	width, height                 int
	targetFps                     int
	input                         *KwindowInput
	hasClose                      bool
	processInputFunc              processInputFunc // Called every frame before update, receive current input state
	updateFunc                    updateFunc       // Called every frame before draw, received delta time (1 = meets targetFps)
	drawFunc                      drawFunc         // Called every frame after update, received drawer
	onDropFile                    onDropFileFunc   // Called when mouse dropped file onto window
	OnSizeChange                  func(newWidth int, newHeight int)
	lastFrameStartTime            time.Time
	glfwWindow                    *glfw.Window
	kwindowDrawer                 KwindowDrawer
	cameraFPSControlEnabled       bool
	isFullScreen                  bool
	windowedHeight, windowedWidth int
	windowedX, windowedY          int
	resizable                     bool
	cameraControlEnabled          bool
}

// CreateKwindow create Kwindow
func newKwindow(config kwindowBuilder) (*kwindow, error) {
	if config.targetFPS <= 0 {
		return nil, fmt.Errorf("Target FPS should not be less than 1")
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	if !config.resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	var err error
	var glfwWindow *glfw.Window
	var wX, wY, wW, wH int
	if config.fullscreen {
		monitor := glfw.GetPrimaryMonitor()
		mode := monitor.GetVideoMode()
		if config.windowed {
			glfwWindow, err = glfw.CreateWindow(mode.Width, mode.Height, config.title, nil, nil)
			config.fullscreen = false
		} else {
			glfwWindow, err = glfw.CreateWindow(mode.Width, mode.Height, config.title, monitor, nil)
			wW = mode.Width / 2
			wH = mode.Height / 2
			wX = wW / 2
			wY = wH / 2
		}
	} else {
		if config.width <= 0 || config.height <= 0 {
			return nil, fmt.Errorf("\n***\tWidth or Height configuration should not be less than 1")
		}
		glfwWindow, err = glfw.CreateWindow(config.width, config.height, config.title, nil, nil)
		if config.center {
			mode := glfw.GetPrimaryMonitor().GetVideoMode()
			wX = mode.Width/2 - config.width/2
			wY = mode.Height/2 - config.height/2
		} else {
			wX = config.x
			wY = config.y
		}
		glfwWindow.SetPos(wX, wY)
	}
	if err != nil {
		return nil, err
	}
	glfwWindow.MakeContextCurrent()
	var updateFunc updateFunc
	updateFunc = config.updateFunc
	var processInputFunc processInputFunc
	processInputFunc = config.processInputFunc
	var drawFunc drawFunc
	drawFunc = config.drawFunc
	var onDropFileFunc onDropFileFunc
	onDropFileFunc = config.onDropFileFunc

	glfwWindow.MakeContextCurrent()
	w := kwindow{
		title:                   config.title,
		width:                   config.width,
		height:                  config.height,
		targetFps:               config.targetFPS,
		lastFrameStartTime:      time.Now(),
		glfwWindow:              glfwWindow,
		isFullScreen:            config.fullscreen,
		windowedX:               wX,
		windowedY:               wY,
		windowedWidth:           wW,
		windowedHeight:          wH,
		cameraFPSControlEnabled: false,
		processInputFunc:        processInputFunc,
		updateFunc:              updateFunc,
		drawFunc:                drawFunc,
		onDropFile:              onDropFileFunc,
		resizable:               config.resizable,
	}
	kinput := newKinput(&w)
	w.input = &kinput
	w.glfwWindow.SetKeyCallback(func(glfwWindow *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		w.input.glfwKeyInputHandler(key, action)
	})

	w.glfwWindow.SetCursorPosCallback(func(glfwWindow *glfw.Window, xpos float64, ypos float64) {
		w.input.glfwMousePosHandler(xpos, ypos)
	})

	w.glfwWindow.SetDropCallback(func(glfwWindow *glfw.Window, names []string) {
		for _, filePath := range names {
			if w.onDropFile != nil {
				w.onDropFile(w.input.mouseX, w.input.mouseY, filePath)
			}
		}
	})

	w.glfwWindow.SetScrollCallback(func(glfwWin *glfw.Window, xoff, yoff float64) {
		w.input.glfwMouseScrollHandler(xoff, yoff)
	})

	w.glfwWindow.SetMouseButtonCallback(func(glfwWin *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		w.input.glfwMouseButtonHandler(button, action)
	})

	w.glfwWindow.SetSizeCallback(func(glfwWin *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
		if w.kwindowDrawer != nil {
			w.kwindowDrawer.onWindowSizeChange(float32(width), float32(height))
		}
	})

	debug.pf("Kwindow %s successfuly created\n", config.title)
	return &w, nil
}

func (w *kwindow) SetOnDropFileFunc(onDropFileFunc onDropFileFunc) {
	w.onDropFile = onDropFileFunc
}

func (w *kwindow) EnableCameraMovementControl() {
	w.cameraFPSControlEnabled = true
}

func (w *kwindow) DisableCameraFPSControl() {
	w.cameraFPSControlEnabled = false
}

func (w *kwindow) LockCursor(lock bool) {
	if lock {
		w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	} else {
		w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (w *kwindow) UnlockCursor() {
	w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}

func (w *kwindow) Start() {
	// w.kdrawer.start()
}

func (w *kwindow) run() {
	if w.hasClose {
		return
	}
	// Process fps and calculate deltatime
	dt := float32(time.Since(w.lastFrameStartTime).Seconds())
	if desireDiff := 1.0/float32(w.targetFps) - dt; desireDiff > 0 {
		time.Sleep(time.Duration(desireDiff * 1000000000))
		dt = float32(time.Since(w.lastFrameStartTime).Seconds())
	}
	dt *= float32(w.targetFps)
	w.lastFrameStartTime = time.Now()

	if !w.hasClose && w.processInputFunc != nil {
		w.processInputFunc((*w.input))
	}

	// Do update
	if !w.hasClose && w.updateFunc != nil {
		w.updateFunc(dt)
	}
	// Do draw
	if w.kwindowDrawer != nil {
		if w.cameraControlEnabled {
			(*w.kwindowDrawer.GetCamera()).updateFPSControl((*w.input), dt) // ..updateFPSControl(dt)
		}
		w.kwindowDrawer.clear()
		if !w.hasClose && w.drawFunc != nil {
			w.drawFunc(&w.kwindowDrawer)
		}
		w.kwindowDrawer.draw()
	}
	if w.hasClose {
		return
	}
	// Show drawn
	w.glfwWindow.SwapBuffers()
	if w.input.GetKeyStat(KeyLeftAlt) == Press && w.input.GetKeyStat(KeyF4) == Press || w.glfwWindow.ShouldClose() {
		w.Close()
		return
	}
	if w.input.GetKeyStat(KeyLeftAlt) == Press && w.input.GetKeyStat(KeyEnter) == JustPress && w.resizable {
		w.ToggleFullscreen()
	}
	// Get input should be in this order
	// Otherwise just release and just press input won't work
	w.input.update()
}

func (w *kwindow) ToggleFullscreen() {
	w.SetFullscreen(!w.isFullScreen)
}

func (w *kwindow) SetFullscreen(enable bool) {
	if enable == w.isFullScreen {
		return
	}
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	if w.isFullScreen {
		w.glfwWindow.SetMonitor(nil, w.windowedX, w.windowedY, w.windowedWidth, w.windowedHeight, 0)
	} else {
		wX, wY := w.glfwWindow.GetPos()
		w.windowedX = wX
		w.windowedY = wY
		wW, wH := w.glfwWindow.GetSize()
		w.windowedWidth = wW
		w.windowedHeight = wH
		w.glfwWindow.SetMonitor(monitor, 0, 0, mode.Width, mode.Height, mode.RefreshRate)
	}
	w.isFullScreen = !w.isFullScreen
}

func (w *kwindow) SetOnSizeChangeCallback(callback glfw.SizeCallback) {
	w.glfwWindow.SetSizeCallback(callback)
}

func (w *kwindow) GetPosition() (x int, y int) {
	return w.glfwWindow.GetPos()
}

func (w *kwindow) SetPosition(x int, y int) {
	w.glfwWindow.SetPos(x, y)
}

func (w *kwindow) GetSize() (width, height int) {
	return w.width, w.height
}

func (w *kwindow) Move(x int, y int) {
	wX, wY := w.glfwWindow.GetPos()
	w.glfwWindow.SetPos(wX+x, wY+y)
}

func (w *kwindow) HasClosed() bool {
	return w.hasClose
}

func (w *kwindow) Close() {
	if w.hasClose {
		return
	}
	debug.pf("kwindow %s close\n", w.title)
	delete(windows, w.id)
	w.hasClose = true
	// w.kdrawer.dispose()
	w.glfwWindow.Destroy()
}

func (w *kwindow) GetInput() *KwindowInput {
	return w.input
}

// func (w *Kwindow) GetCameraFrustum() Frustum {
// 	return w.kdrawer.camera.getFrustum()
// }
