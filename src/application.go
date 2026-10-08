package main

import (
	"nesroom/src/nes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TICK         float64 = 1.0 / (60.0 * 4.0)
	SAMPLE_RATE          = 44100
	MenuSceneKey         = iota
	GameSceneKey
	ControlsSceneKey
)

var (
	ZERO_VECTOR2 = rl.Vector2{}
)

type Application struct {
	isWeb             bool
	scenes            map[int]Scene
	drawTarget        rl.RenderTexture2D
	currentSceneIndex int
	sourceRect        rl.Rectangle
	destRect          rl.Rectangle
	timestamp         float64
	audioChannel      chan float32
	audioStream       rl.AudioStream
}

func NewApplication(iw bool) *Application {

	app := Application{}
	app.isWeb = iw

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogTrace)
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable) //should be set before window initialization!
	rl.InitWindow(TspWinWidth, TspWinHeight, "NESRoom")
	if !app.isWeb {
		scaleFactor := int32(3)
		rl.SetWindowSize(int(nes.ScreenLogicalWidth*scaleFactor), int(nes.ScreenLogicalHeight*scaleFactor))
		rl.SetWindowMonitor(0) //used for testing on multiple monitors
	}

	rl.SetAudioStreamBufferSizeDefault(4096)
	rl.InitAudioDevice()
	app.audioStream = rl.LoadAudioStream(SAMPLE_RATE, 32, 1)
	app.audioChannel = make(chan float32, SAMPLE_RATE)
	rl.PlayAudioStream(app.audioStream)
	rl.SetAudioStreamCallback(app.audioStream, app.Callback)

	setTextStyle(defaultTextStyle)

	// scenes
	app.scenes = make(map[int]Scene)
	app.scenes[GameSceneKey] = NewGameScene(&app)
	app.currentSceneIndex = GameSceneKey
	app.timestamp = rl.GetTime()

	//debug
	//app.currentSceneIndex = controlsSceneKey

	app.drawTarget = rl.LoadRenderTexture(nes.ScreenLogicalWidth, nes.ScreenLogicalHeight)
	app.onResize()

	return &app
}

func (a *Application) Callback(out []float32, frames int) {
	for i := range frames {
		out[i] = <-a.audioChannel
	}
}

func (a *Application) onResize() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	a.sourceRect = rl.NewRectangle(0, 0, float32(nes.ScreenLogicalWidth), float32(-nes.ScreenLogicalHeight))
	ratioX := screenWidth / float32(nes.ScreenLogicalWidth)
	ratioY := screenHeight / float32(nes.ScreenLogicalHeight)
	resizeRatio := If(ratioX < ratioY, ratioX, ratioY)
	a.destRect = rl.NewRectangle(
		(screenWidth-(float32(nes.ScreenLogicalWidth)*resizeRatio))*0.5,
		(screenHeight-(float32(nes.ScreenLogicalHeight)*resizeRatio))*0.5,
		float32(nes.ScreenLogicalWidth)*resizeRatio,
		float32(nes.ScreenLogicalHeight)*resizeRatio,
	)
}

func (a *Application) ProcessInput() {
	if rl.IsWindowResized() {
		a.onResize()
	}
	a.scenes[a.currentSceneIndex].ProcessInput()
}

func (a *Application) Update() {
	timestamp := rl.GetTime()
	delta := timestamp - a.timestamp
	a.timestamp = timestamp
	a.scenes[a.currentSceneIndex].Update(delta)
}

func (a *Application) Render() {
	rl.BeginDrawing()
	a.scenes[a.currentSceneIndex].Render(a.drawTarget)
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(a.drawTarget.Texture,
		a.sourceRect,
		a.destRect,
		ZERO_VECTOR2, 0, rl.White)
	rl.EndDrawing()
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() ||
		a.scenes[a.currentSceneIndex].ShouldExit() ||
		rl.IsGamepadButtonDown(GamePadId, TspMenuCode) && rl.IsGamepadButtonDown(GamePadId, TspStartCode)
}

func (a *Application) Exit() {
	rl.UnloadAudioStream(a.audioStream)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
