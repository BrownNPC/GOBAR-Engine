package main

import (
	"game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneMenu struct {
	engine.BaseScene
}

func (s *SceneMenu) Init() {

}
func (s *SceneMenu) Render() {
	rl.DrawText("This is the Menu scene! or.. Is it? 🤨", 300, 200, 30, rl.Yellow)
}
func (s *SceneMenu) Update(virtualWidth float32, virtualHeight float32) {
	s.UpdateBaseScene(virtualWidth, virtualHeight)
	if rl.IsKeyPressed(rl.KeyBackspace) {
		s.GoToNextScene()
	}
}
func (s *SceneMenu) Unload() {
}

func (s SceneMenu) IsLoaded() bool {
	return s.Loaded
}

func (s SceneMenu) NextScene() string {
	return "main"
}
