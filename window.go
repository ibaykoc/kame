// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"errors"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type updateFunc func(timeSinceLastFrame float64)
type drawFunc func(drawer *Drawer)
type onDropFileFunc func(filePath string)

// Window for kame
type Window struct {
	title              string
	width              int
	height             int
	targetFps          int
	input              Input
	WannaClose         bool
	hasClose           bool
	updateFunc         updateFunc     // Called every frame before draw, received delta time (1 = meets targetFps)
	drawFunc           drawFunc       // Called every frame after draw, received drawer to draw something
	onDropFile         onDropFileFunc // Called when mouse dropped file onto window
	OnSizeChange       func(newWidth int, newHeight int)
	lastFrameStartTime time.Time
	glfwWindow         *glfw.Window
	drawer             *Drawer
}

// CreateWindow with default value
func createWindow(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	if !hasInitialized {
		return nil, errors.New("Kame should be initialized first")
	}

	windowWidth := 960
	windowHeight := 800
	windowTitle := "KAME"
	windowTargetFPS := 60
	windowBackgroundColor := Color{0.5, 0.5, 0.5, 1}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwWindow, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()
	window := Window{
		title:              windowTitle,
		width:              windowWidth,
		height:             windowHeight,
		targetFps:          windowTargetFPS,
		updateFunc:         updateFunc,
		drawFunc:           drawFunc,
		lastFrameStartTime: time.Now(),
		glfwWindow:         glfwWindow,
	}
	d, err := newDrawer(&window, windowBackgroundColor)
	if err != nil {
		panic(err)
	}
	window.drawer = d
	glfwWindow.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		window.width = width
		window.height = height
		window.drawer.changeSize(int32(width), int32(height))
		if window.OnSizeChange != nil {
			window.OnSizeChange(width, height)
		}
	})

	i := newInput(glfwWindow)
	window.input = i
	glfwWindow.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		window.input.glfwInputHandler(key, action)
	})

	glfwWindow.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		window.input.glfwMousePosHandler(xpos, ypos)
	})

	glfwWindow.SetDropCallback(func(w *glfw.Window, names []string) {
		for _, filePath := range names {
			if window.onDropFile != nil {
				window.onDropFile(filePath)
			}
		}
	})
	return &window, nil
}
func (w *Window) SetOnDropFileFunc(onDropFileFunc onDropFileFunc) {
	w.onDropFile = onDropFileFunc
}
func (w *Window) DoMagic() {
	if w.hasClose {
		return
	}
	// Process fps and calculate deltatime
	dt := time.Since(w.lastFrameStartTime).Seconds()
	if desireDiff := 1.0/float64(w.targetFps) - dt; desireDiff > 0 {
		time.Sleep(time.Duration(desireDiff * 1000000000))
		dt = time.Since(w.lastFrameStartTime).Seconds()
	}
	dt *= float64(w.targetFps)
	w.lastFrameStartTime = time.Now()

	// Update window should close stat
	w.WannaClose = w.glfwWindow.ShouldClose()

	// Do update
	if !w.hasClose {
		w.updateFunc(dt)
	}
	// Do draw
	w.drawer.clear()
	if !w.hasClose {
		w.drawFunc(w.drawer)
	}
	if w.hasClose {
		return
	}
	// Show drawn
	w.glfwWindow.SwapBuffers()

	// Get input
	w.input.update()
	glfw.PollEvents()
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
	w.drawer.dispose()
	w.glfwWindow.Destroy()
}

func (w *Window) GetInput() Input {
	return w.input
}
