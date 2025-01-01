package main

import (
	"fmt"
	"game/components"
	"game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp"
)

type SceneMain struct {
	engine.BaseScene
}

func (s *SceneMain) Init() {
	s.EntityManager.CreateEntity("player",
		&components.Health{},
		&components.Transform{
			Position: cp.Vector{X: 1280 / 2, Y: 720 / 2},
		},
	)

}
func (s *SceneMain) Render() {
	s.ForEachEntity(func(e *engine.Entity) {
		health, ok := e.GetComponent(components.HealthComponentID)
		if ok {
			health := health.(*components.Health)
			rl.DrawText(fmt.Sprint("Health is ", health.HP), int32(s.VirtualWidth/2), int32(s.VirtualHeight/2), 20, rl.White)
		}
	})
	rl.DrawText("This is the Main scene! or.. Is it?, \npress backsapce to go to the menu scene ", 300, 200, 30, rl.Green)
}
func (s *SceneMain) Update(virtualWidth float32, virtualHeight float32) {
	s.UpdateBaseScene(virtualWidth, virtualHeight)
	if rl.IsKeyPressed(rl.KeyBackspace) {
		s.GoToNextScene()
	}
	entities := s.EntityManager.GetEntities()
	for element := entities.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		entity := element.Value.(*engine.Entity)
		health, ok := entity.GetComponent(components.HealthComponentID)
		if ok {
			health := health.(*components.Health)
			health.HP--
			if health.HP <= 0 {
				health.HP = 100
			}
		}
	}

}

func (s *SceneMain) NextScene() string {
	return "menu"
}
func (s *SceneMain) Unload() {
}
