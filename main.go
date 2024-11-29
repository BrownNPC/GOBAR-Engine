package main

import (
	"game/engine"
)

func main() {
	sceneConfig := []engine.SceneConfig{
		engine.ConfigureNewScene("menu", &SceneMenu{}),
		engine.ConfigureNewScene("main", &SceneMain{}),
	}

	// var MaxEntities = 1000

	g := engine.NewGame(sceneConfig, engine.LoadConfig())
	g.Run()
}

// package main

// import (
// 	"fmt"
// 	"game/engine"
// )

// // Example component
// type Position struct {
// 	X, Y float64
// }

// func main() {
// 	// Create a MemoryPool with 10 entities and Position component type
// 	mp := engine.NewMemoryPool(10, make([]Position, 10))

// 	// Activate entities
// 	mp.Alive[0] = true
// 	mp.Alive[1] = true

// 	// Add a Position slice to the tuple
// 	positions, _ := engine.GetComponentsByType[Position](mp)
// 	positions[0] = Position{X: 1, Y: 2}
// 	positions[1] = Position{X: 3, Y: 4}

// 	// Retrieve a component for an entity
// 	pos := engine.GetComponentFromEntity[Position](0, mp)
// 	fmt.Println(pos)
// }
