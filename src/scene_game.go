package main

import (
	"bytes"
	"embed"
	"nesroom/src/nes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
)

type GameScene struct {
	a              *Application
	console        *nes.Console
	gameDrawTarget rl.RenderTexture2D
	gameSourceRect rl.Rectangle
	gameDestRect   rl.Rectangle
	debugGrid      *DebugGrid
}

func NewGameScene(a *Application) *GameScene {
	gs := GameScene{}
	gs.a = a
	bytesData, _ := resList.ReadFile("res/tetris.nes")
	reader := bytes.NewReader(bytesData)
	gs.console, _ = nes.NewConsole(reader)
	gs.gameDrawTarget = rl.LoadRenderTexture(nes.ScreenLogicalWidth, nes.ScreenLogicalHeight)
	gs.gameSourceRect = rl.NewRectangle(0, 0, float32(nes.ScreenLogicalWidth), float32(nes.ScreenLogicalHeight))
	gs.gameDestRect = rl.NewRectangle(0, 0, float32(nes.ScreenLogicalWidth), float32(nes.ScreenLogicalHeight))
	gs.debugGrid = NewDebugGrid(&gs)
	return &gs
}

func (gs *GameScene) ProcessInput() {
	turbo := gs.console.PPU.Frame%6 < 3
	buttons1 := [8]bool{false, false, false, false, false, false, false, false}
	buttons1[nes.ButtonA] = rl.IsKeyDown(rl.KeyZ) || turbo && rl.IsKeyDown(rl.KeyA) || rl.IsGamepadButtonDown(GamePadId, TspACode)
	buttons1[nes.ButtonB] = rl.IsKeyDown(rl.KeyX) || turbo && rl.IsKeyDown(rl.KeyS) || rl.IsGamepadButtonDown(GamePadId, TspBCode)
	buttons1[nes.ButtonSelect] = rl.IsKeyDown(rl.KeyRightShift) || rl.IsGamepadButtonDown(GamePadId, TspSelectCode)
	buttons1[nes.ButtonStart] = rl.IsKeyDown(rl.KeyEnter) || rl.IsGamepadButtonDown(GamePadId, TspStartCode)
	buttons1[nes.ButtonUp] = rl.IsKeyDown(rl.KeyUp) || rl.IsGamepadButtonDown(GamePadId, TspUpCode)
	buttons1[nes.ButtonDown] = rl.IsKeyDown(rl.KeyDown) || rl.IsGamepadButtonDown(GamePadId, TspDownCode)
	buttons1[nes.ButtonLeft] = rl.IsKeyDown(rl.KeyLeft) || rl.IsGamepadButtonDown(GamePadId, TspLeftCode)
	buttons1[nes.ButtonRight] = rl.IsKeyDown(rl.KeyRight) || rl.IsGamepadButtonDown(GamePadId, TspRightCode)
	gs.console.SetButtons1(buttons1)
}

func (gs *GameScene) Update(delta float64) {
	if delta > 1 {
		delta = 0
	}
	gs.console.StepSeconds(delta)
}

func (gs *GameScene) Render(drawTarget rl.RenderTexture2D) {
	rl.UpdateTexture(gs.gameDrawTarget.Texture, gs.console.Buffer())
	rl.BeginTextureMode(drawTarget)
	{
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(gs.gameDrawTarget.Texture, gs.gameSourceRect, gs.gameDestRect, ZERO_VECTOR2, 0, rl.White)
		//gs.debugGrid.Render()
	}
	rl.EndTextureMode()
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape)
}
