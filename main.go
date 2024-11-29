package main

import (
	"game/engine"
)

func main() {
	sceneConfig := []engine.SceneConfig{
		engine.ConfigureNewScene("menu", &SceneMenu{}),
		engine.ConfigureNewScene("main", &SceneMain{}),
	}

	g := engine.NewGame(sceneConfig, engine.LoadConfig())
	g.Run()
}
