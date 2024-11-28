package main

import (
	"game/components"
	"game/engine"
)

func main() {
	sceneConfig := []engine.SceneConfig{
		engine.ConfigureNewScene("menu", &SceneMenu{}),
		engine.ConfigureNewScene("main", &SceneMain{}),
	}
	pool := engine.NewMemoryPool(1000)
	engine.MemoryPoolRegisterComponent[components.CTransform](pool, 0)

	g := engine.NewGame(sceneConfig, engine.LoadConfig(), pool)
	g.Run()
}
