package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene interface {
	ProcessInput()
	Update(delta float64)
	Render(drawTarget rl.RenderTexture2D)
	ShouldExit() bool
}
