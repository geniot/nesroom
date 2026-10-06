module nesroom

go 1.26.1

require (
	github.com/BrownNPC/Raylib-Go-Wasm/raylib v0.0.0-20260813175119-abd0f529636b
	github.com/gen2brain/raylib-go/raygui v0.0.0-20260815042312-80156a59a482
	github.com/gen2brain/raylib-go/raylib v0.60.1
)

require (
	github.com/BrownNPC/Raylib-Go-Wasm/wasm-runtime v0.0.0-20260813175119-abd0f529636b // indirect
	github.com/BrownNPC/wasm-ffi-go v1.3.0 // indirect
)

replace (
	github.com/BrownNPC/Raylib-Go-Wasm/wasm-runtime => ./Raylib-Go-Wasm/wasm-runtime
	github.com/gen2brain/raylib-go/raygui => ./Raylib-Go-Wasm/raygui
	github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib
)
