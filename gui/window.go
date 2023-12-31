package gui

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)


func init() {
	runtime.LockOSThread()
}

func Window(width int, height int) {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
