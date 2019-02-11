// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Window for kame
type Window struct {
	Title                string
	width                int
	height               int
	targetFps            int
	ShouldClose          bool
	hasClose             bool
	OnUpdate             func(deltaTime float64) // Called every frame before draw, received delta time (1 = meets targetFps)
	OnDraw               func()
	OnSizeChangeCallback func(newWidth int, newHeight int)
	lastFrameStartTime   time.Time
	glfwWindow           *glfw.Window
}

func newWindow(title string, width int, height int, targetFps int, glfwWindow *glfw.Window) *Window {
	window := Window{
		Title:              title,
		width:              width,
		height:             height,
		targetFps:          targetFps,
		lastFrameStartTime: time.Now(),
		glfwWindow:         glfwWindow,
	}
	glfwWindow.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		window.width = width
		window.height = height
		if window.OnSizeChangeCallback != nil {
			window.OnSizeChangeCallback(width, height)
		}
	})
	return &window
}

func (w *Window) Run() {
	dt := time.Since(w.lastFrameStartTime).Seconds()
	if desireDiff := 1.0/float64(w.targetFps) - dt; desireDiff > 0 {
		time.Sleep(time.Duration(desireDiff * 1000000000))
		dt = time.Since(w.lastFrameStartTime).Seconds()
	}
	dt *= float64(w.targetFps)
	w.lastFrameStartTime = time.Now()
	glfw.PollEvents()
	if w.OnUpdate != nil {
		w.OnUpdate(dt)
	}

	gl.Clear(gl.COLOR_BUFFER_BIT)
	if w.OnDraw != nil {
		w.OnDraw()
	}
	w.glfwWindow.SwapBuffers()
	w.ShouldClose = w.glfwWindow.ShouldClose()
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

func (w *Window) Move(x int, y int) {
	wX, wY := w.glfwWindow.GetPos()
	w.glfwWindow.SetPos(wX+x, wY+y)
}

func (w *Window) HasClose() bool {
	return w.hasClose
}

func (w *Window) Close() {
	w.hasClose = true
	w.glfwWindow.Destroy()
}
