package main

import (
	"nesroom/src/nes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugGrid struct {
	scene     *GameScene
	texture   rl.RenderTexture2D
	sourceRec rl.Rectangle
	destRec   rl.Rectangle
}

func NewDebugGrid(scene *GameScene) *DebugGrid {
	debugGrid := &DebugGrid{}
	debugGrid.scene = scene
	debugGrid.sourceRec = rl.NewRectangle(0, 0, float32(nes.ScreenLogicalWidth), -float32(nes.ScreenLogicalHeight)) //see https://github.com/raysan5/raylib/issues/3803
	debugGrid.destRec = rl.NewRectangle(0, 0, float32(nes.ScreenLogicalWidth), float32(nes.ScreenLogicalHeight))

	debugGrid.texture = rl.LoadRenderTexture(nes.ScreenLogicalWidth, nes.ScreenLogicalHeight)
	rl.BeginTextureMode(debugGrid.texture)
	rl.DrawRectangleLinesEx(rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(nes.ScreenLogicalWidth),
		Height: float32(nes.ScreenLogicalHeight),
	}, 1, rl.Yellow)
	rl.EndTextureMode()

	return debugGrid
}

func (debugGrid *DebugGrid) Update(_ int64) {
}

func (debugGrid *DebugGrid) Render() {
	rl.DrawFPS(5, 5)
	rl.DrawTexturePro(debugGrid.texture.Texture, debugGrid.sourceRec, debugGrid.destRec, ZERO_VECTOR2, 0, rl.White)
}
