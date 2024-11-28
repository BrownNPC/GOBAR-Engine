package main

import (
	"game/components"
	"game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp"
)

type SceneMain struct {
	engine.BaseScene
}

func (s *SceneMain) Init() {

	s.EntityManager.CreateEntity("player").
		AddComponent(0, components.CTransform{Position: cp.Vector{X: 0, Y: 0}})
}
func (s *SceneMain) Render() {
}
func (s *SceneMain) Update(virtualWidth float32, virtualHeight float32) {
	s.UpdateVirtualResolution(virtualWidth, virtualHeight)
	if rl.IsKeyPressed(rl.KeyBackspace) {
		s.GoToNextScene()
	}

}
func (s *SceneMain) Unload() {
}

func (s *SceneMain) IsLoaded() bool {
	return s.Loaded
}

func (s *SceneMain) NextScene() string {
	return "menu"
}
