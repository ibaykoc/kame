// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type updateFunc func(timeSinceLastFrame float32)
type drawFunc func(drawer *KDrawer)
type onDropFileFunc func(filePath string)

// Window for kame
type Window struct {
	title                         string
	width, height                 int
	targetFps                     int
	input                         Input
	WannaClose                    bool
	hasClose                      bool
	updateFunc                    updateFunc     // Called every frame before draw, received delta time (1 = meets targetFps)
	drawFunc                      drawFunc       // Called every frame after draw, received drawer to draw something
	onDropFile                    onDropFileFunc // Called when mouse dropped file onto window
	OnSizeChange                  func(newWidth int, newHeight int)
	lastFrameStartTime            time.Time
	glfwWindow                    *glfw.Window
	kdrawer                       *KDrawer
	cameraFPSControlEnabled       bool
	isFullScreen                  bool
	windowedHeight, windowedWidth int
	windowedX, windowedY          int
	resizable                     bool
	CameraType                    ProjectionType
}

type WindowConfig struct {
	Title           string
	Fullscreen      bool
	Windowed        bool
	TargetFPS       int
	Center          bool
	Width           int
	Height          int
	Resizable       bool
	CameraType      ProjectionType
	BackgroundColor Color
}

func createWindow(config WindowConfig) (*Window, error) {
	if !hasInitialized {
		return nil, errors.New("Kame should be initialized first")
	}
	if config.TargetFPS <= 0 {
		return nil, fmt.Errorf("\n***\tTarget FPS configuration should not be less than 1")
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	if !config.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	var err error
	var glfwWindow *glfw.Window

	var wX, wY, wW, wH int
	if config.Fullscreen {
		monitor := glfw.GetPrimaryMonitor()
		mode := monitor.GetVideoMode()
		if config.Windowed {
			glfwWindow, err = glfw.CreateWindow(mode.Width, mode.Height, config.Title, nil, nil)
			config.Fullscreen = false
		} else {
			glfwWindow, err = glfw.CreateWindow(mode.Width, mode.Height, config.Title, monitor, nil)
			wW = mode.Width / 2
			wH = mode.Height / 2
			wX = wW / 2
			wY = wH / 2
		}
	} else {
		if config.Width <= 0 || config.Height <= 0 {
			return nil, fmt.Errorf("\n***\tWidth or Height configuration should not be less than 1")
		}
		glfwWindow, err = glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
		if config.Center {
			mode := glfw.GetPrimaryMonitor().GetVideoMode()
			wX = mode.Width/2 - config.Width/2
			wY = mode.Height/2 - config.Height/2
			glfwWindow.SetPos(wX, wY)
		} else {
			wX, wY = glfwWindow.GetPos()
		}
	}
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()
	w := Window{
		title:                   config.Title,
		width:                   config.Width,
		height:                  config.Height,
		targetFps:               config.TargetFPS,
		lastFrameStartTime:      time.Now(),
		glfwWindow:              glfwWindow,
		isFullScreen:            config.Fullscreen,
		windowedX:               wX,
		windowedY:               wY,
		windowedWidth:           wW,
		windowedHeight:          wH,
		cameraFPSControlEnabled: false,
		resizable:               config.Resizable,
	}
	return &w, nil
}

func (w *Window) SetOnDropFileFunc(onDropFileFunc onDropFileFunc) {
	w.onDropFile = onDropFileFunc
}

func (w *Window) EnableCameraMovementControl() {
	w.cameraFPSControlEnabled = true
}

func (w *Window) DisableCameraFPSControl() {
	w.cameraFPSControlEnabled = false
}

func (w *Window) LockCursor() {
	w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
}

func (w *Window) UnlockCursor() {
	w.glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}

func (w *Window) DoMagic() {
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

	// Update window should close stat
	w.WannaClose = w.glfwWindow.ShouldClose()

	if w.cameraFPSControlEnabled {
		w.kdrawer.camera.updateFPSControl(dt)
	}

	// Do update
	if !w.hasClose {
		w.updateFunc(dt)
	}
	// Do draw
	w.kdrawer.clear()
	if !w.hasClose {
		w.drawFunc(w.kdrawer)
	}
	w.kdrawer.applyDraw()

	if w.hasClose {
		return
	}
	// Show drawn
	w.glfwWindow.SwapBuffers()
	if w.input.GetKeyStat(KeyLeftAlt) == Press && w.input.GetKeyStat(KeyF4) == Press {
		w.Close()
		return
	}
	if w.input.GetKeyStat(KeyLeftAlt) == Press && w.input.GetKeyStat(KeyEnter) == JustPress && w.resizable {
		w.ToggleFullscreen()
	}
	// Get input should be in this order
	// Otherwise just release and just press input won't work
	w.input.update()
	glfw.PollEvents()
}

func (w *Window) ToggleFullscreen() {
	w.SetFullscreen(!w.isFullScreen)
}

func (w *Window) SetFullscreen(enable bool) {
	if enable == w.isFullScreen {
		return
	}
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	if window.isFullScreen {
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
	window.isFullScreen = !window.isFullScreen
}

func (w *Window) SetOnSizeChangeCallback(callback glfw.SizeCallback) {
	w.glfwWindow.SetSizeCallback(callback)
}

func (w *Window) GetPosition() (x int, y int) {
	return w.glfwWindow.GetPos()
}

func (w *Window) SetPosition(x int, y int) {
	w.glfwWindow.SetPos(x, y)
}

func (w *Window) GetSize() (width, height int) {
	return w.width, w.height
}

func (w *Window) Move(x int, y int) {
	wX, wY := w.glfwWindow.GetPos()
	w.glfwWindow.SetPos(wX+x, wY+y)
}

func (w *Window) HasClosed() bool {
	return w.hasClose
}

func (w *Window) Close() {
	if w.hasClose {
		return
	}
	w.WannaClose = true
	w.hasClose = true
	w.kdrawer.dispose()
	w.glfwWindow.Destroy()
}

func (w *Window) GetInput() Input {
	return w.input
}

func (w Window) GetCameraFrustum() Frustum {
	return w.kdrawer.camera.getFrustum()
}
