// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"errors"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var hasInitialized = false
var debug kdebug = true
var windows = make(map[KwindowID]*Kwindow)

// Initialize kame
func Initialize() error {
	runtime.LockOSThread()
	if hasInitialized {
		return errors.New("kame Should not initialize more than one")
	}
	if err := glfw.Init(); err != nil {
		return err
	}
	hasInitialized = true
	debug.pf("kame successfuly initialized\n")
	return nil
}

func GetMonitorSize() (int, int) {
	vMode := glfw.GetPrimaryMonitor().GetVideoMode()
	h := vMode.Height
	w := vMode.Width
	return w, h
}

func ShouldClose() bool {
	return len(windows) <= 0
}

func DoMagic() {
	for _, w := range windows {
		w.run()
	}
}

// // TurnOn2D Create 2D World PPU (Pixel Per Unit: 50), (0,0) at center
// func TurnOn2D(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
// 	return TurnOn(WindowConfig{
// 		Title:           "KAME 2D",
// 		Width:           960,
// 		Height:          800,
// 		Resizable:       true,
// 		TargetFPS:       60,
// 		BackgroundColor: Color{0.25, 0.25, 0.25, 1},
// 		CameraType:      Orthographic,
// 	}, updateFunc, drawFunc)
// }

// func GameOn2D(scenes []Scene) (*GameWindow, error) {
// 	gw := &GameWindow{}
// 	w, err := TurnOn2D(gw.update, gw.draw)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gw.Window = w
// 	gw.initialize(scenes)
// 	return gw, nil
// }

// func GameOn3D(scenes []Scene) (*GameWindow, error) {
// 	gw := &GameWindow{}
// 	w, err := TurnOn3D(gw.update, gw.draw)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gw.Window = w
// 	gw.initialize(scenes)
// 	return gw, nil
// }

// func TurnOn3D(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
// 	return TurnOn(WindowConfig{
// 		Title:           "KAME 3D",
// 		Width:           960,
// 		Height:          800,
// 		Resizable:       true,
// 		TargetFPS:       60,
// 		BackgroundColor: Color{0.5, 0.5, 0.5, 1},
// 		CameraType:      Perspective,
// 	}, updateFunc, drawFunc)
// }

// func TurnOn(config WindowConfig, updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {

// 	var err error
// 	window, err = createWindow(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var d *KDrawer

// 	if config.CameraType == Perspective {
// 		d, err = newDrawer3D(config.BackgroundColor)
// 	} else {
// 		d, err = newDrawer2D(config.BackgroundColor)
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	window.kdrawer = d
// 	window.glfwWindow.SetSizeCallback(func(w *glfw.Window, width int, height int) {
// 		window.width = width
// 		window.height = height
// 		window.kdrawer.changeSize(int32(width), int32(height))
// 		if window.OnSizeChange != nil {
// 			window.OnSizeChange(width, height)
// 		}
// 	})

// 	i := newInput(window.glfwWindow)
// 	window.input = i
// 	window.glfwWindow.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
// 		window.input.glfwInputHandler(key, action)
// 	})

// 	window.glfwWindow.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
// 		window.input.glfwMousePosHandler(xpos, ypos)
// 	})

// 	window.glfwWindow.SetDropCallback(func(w *glfw.Window, names []string) {
// 		for _, filePath := range names {
// 			if window.onDropFile != nil {
// 				window.onDropFile(filePath)
// 			}
// 		}
// 	})
// 	window.updateFunc = updateFunc
// 	window.drawFunc = drawFunc
// 	fmt.Println("kame Turned On")
// 	return window, nil
// }

// // TurnOff turn off kame
// func TurnOff() {
// 	window.Close()
// 	glfw.Terminate()
// 	fmt.Println("kame Turned Off")
// }
