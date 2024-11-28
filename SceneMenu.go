package main

import (
	"game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneMenu struct {
	engine.BaseScene
	background rl.Texture2D
}

func (s *SceneMenu) Init() {
	// s.InitBaseScene()
	s.background = rl.LoadTexture("assets/sprites/apple.png")

}
func (s *SceneMenu) Render() {
	s.DrawTexture(s.background, rl.NewRectangle(0, 0, 100, 100), rl.White)
}
func (s *SceneMenu) Update(virtualWidth float32, virtualHeight float32) {
	s.UpdateVirtualResolution(virtualWidth, virtualHeight)
	if rl.IsKeyPressed(rl.KeyBackspace) {
		s.GoToNextScene()
	}
}
func (s *SceneMenu) Unload() {
	rl.UnloadTexture(s.background)
}

func (s SceneMenu) IsLoaded() bool {
	return s.Loaded
}

func (s SceneMenu) NextScene() string {
	return "main"
}
