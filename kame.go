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
func TurnOn(updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	runtime.LockOSThread()
	if hasInitialized {
		return nil, errors.New("kame Has Turned On")
	}
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	hasInitialized = true
	var err error
	window, err = createDefaultWindow(updateFunc, drawFunc)
	if err != nil {
		return nil, err
	}
	fmt.Println("kame Turned On")
	return window, nil
}

func TurnOnConfigured(config WindowConfig, updateFunc updateFunc, drawFunc drawFunc) (*Window, error) {
	runtime.LockOSThread()
	if hasInitialized {
		return nil, errors.New("kame Has Turned On")
	}
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	hasInitialized = true
	var err error
	window, err = createWindowWithConfig(config, updateFunc, drawFunc)
	if err != nil {
		return nil, err
	}
	fmt.Println("kame Turned On")
	return window, nil
}

// TurnOff turn off kame
func TurnOff() {
	window.Close()
	glfw.Terminate()
	fmt.Println("kame Turned Off")
}
