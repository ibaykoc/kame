// Copyright 2019 ibaykoc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kame

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var hasInitialized = false

func Init() error {
	runtime.LockOSThread()
	if hasInitialized {
		return errors.New("Can't initialize kame more than once")
	}
	if err := glfw.Init(); err != nil {
		return err
	}
	hasInitialized = true
	fmt.Println("kame succesfully initialize")
	return nil
}

func CreateWindow(title string, windowWidth int, windowHeight int, targetFps int, backgroundColor Color) (*Window, error) {
	if !hasInitialized {
		return nil, errors.New("Kame should be initialized first")
	}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwWindow, err := glfw.CreateWindow(windowWidth, windowHeight, title, nil, nil)
	if err != nil {
		return nil, err
	}

	// Center Window
	vMode := glfw.GetPrimaryMonitor().GetVideoMode()
	mW := vMode.Width
	mH := vMode.Height
	glfwWindow.SetPos(mW/2-windowWidth/2, mH/2-windowHeight/2)
	glfwWindow.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}
	// version := gl.GoStr(gl.GetString(gl.VERSION))
	// fmt.Println("OpenGL initialized: version", version)
	gl.ClearColor(
		backgroundColor.R,
		backgroundColor.G,
		backgroundColor.B,
		backgroundColor.A)

	window := newWindow(
		title,
		windowWidth,
		windowHeight,
		targetFps,
		glfwWindow,
	)
	return window, nil
}

func Quit() {
	glfw.Terminate()
	fmt.Println("kame succesfully quit")
}
