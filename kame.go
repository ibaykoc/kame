// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var hasInitialized = false
var window *Window

// TurnOn needs update & draw function that will be called every frame sequentially
// and returns Window and also error if something goes wrong.
// And finally, don't forget to Turn Off
// Have Fun <3
func TurnOn2D(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	return TurnOn(WindowConfig{
		Title:           "KAME 2D",
		Width:           960,
		Height:          800,
		Resizable:       true,
		TargetFPS:       60,
		BackgroundColor: Color{0.5, 0.5, 0.5, 1},
		CameraType:      Orthographic,
	}, updateFunc, drawFunc)
}

func TurnOn3D(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	return TurnOn(WindowConfig{
		Title:           "KAME 3D",
		Width:           960,
		Height:          800,
		Resizable:       true,
		TargetFPS:       60,
		BackgroundColor: Color{0.5, 0.5, 0.5, 1},
		CameraType:      Perspective,
	}, updateFunc, drawFunc)
}

func TurnOn(config WindowConfig, updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	runtime.LockOSThread()
	if hasInitialized {
		return nil, errors.New("kame Has Turned On")
	}
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	hasInitialized = true
	var err error
	window, err = createWindow(config, updateFunc, drawFunc)
	if err != nil {
		return nil, err
	}
	var d *Drawer

	if config.CameraType == Perspective {
		d, err = newDrawer3D(config.BackgroundColor)
	} else {
		d, err = newDrawer2D(config.BackgroundColor)
	}
	if err != nil {
		panic(err)
	}
	window.drawer = d
	window.glfwWindow.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		window.width = width
		window.height = height
		window.drawer.changeSize(int32(width), int32(height))
		if window.OnSizeChange != nil {
			window.OnSizeChange(width, height)
		}
	})

	i := newInput(window.glfwWindow)
	window.input = i
	window.glfwWindow.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		window.input.glfwInputHandler(key, action)
	})

	window.glfwWindow.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		window.input.glfwMousePosHandler(xpos, ypos)
	})

	window.glfwWindow.SetDropCallback(func(w *glfw.Window, names []string) {
		for _, filePath := range names {
			if window.onDropFile != nil {
				window.onDropFile(filePath)
			}
		}
	})
	fmt.Println("kame Turned On")
	return window, nil
}

// TurnOff turn off kame
func TurnOff() {
	window.Close()
	glfw.Terminate()
	fmt.Println("kame Turned Off")
}
