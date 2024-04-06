package app

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	// Width is the width of the window
	Width = 400
	// Height is the height of the window
	Height = 400
	// Name is the name of the app
	Name = "OpenGL App"
)

// ProcessInput handles any mouse or keyboard input.
func ProcessInput(window *glfw.Window) {
	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

// RenderFrame determines what gets rendered on each frame.
func RenderFrame() {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
