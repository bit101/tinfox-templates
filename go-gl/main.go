// Package main is the starting point of the app.
package main

//revive:disable:unused-parameter

import (
	"${MODULE}/app"
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	window *glfw.Window
)

func init() {
	runtime.LockOSThread()
}

// =================================================================
// This file is mainly boilerplate. Customize the app in app/app.go
// =================================================================

func main() {
	initGLFW()
	initWindow()
	initGL()
	appLoop()
	glfw.Terminate()
}

func appLoop() {
	for !window.ShouldClose() {
		app.ProcessInput(window)
		app.RenderFrame()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initGLFW() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func initGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func initWindow() {
	win, err := glfw.CreateWindow(app.Width, app.Height, app.Name, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window = win
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(framebufferSizeCallback)
}

func framebufferSizeCallback(window *glfw.Window, w, h int) {
	gl.Viewport(0, 0, int32(w), int32(h))
}
