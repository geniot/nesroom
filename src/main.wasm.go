//go:build wasm

package main

import rl "github.com/BrownNPC/Raylib-Go-Wasm/raylib"

func main() {
	application := NewApplication(true)
	var update = func() {
		application.Update()
		application.Render()
	}
	rl.SetMainLoop(update)
	for !rl.WindowShouldClose() {
		update()
	}
	application.Exit()
}
