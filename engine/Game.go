package engine

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene interface {
	Update(virtualWidth float32, virtualHeight float32)
	Render()
	Init()
	Unload()
	IsLoaded() bool
	NextScene() string
	//whether scene should be switched to the next scene
	IsDone() bool
	ReceiveMemoryPool(memoryPool *memoryPool)
	InitBaseScene()
	UnloadBaseScene()
}

type Game struct {
	conf Config

	currentScene string
	Scenes       map[string]Scene

	virtualWidth  float32
	virtualHeight float32
	// the user gives the game a memory pool, after registering components in the memory pool
	//it passes estimate max entities to the memory pool by reading the config passed by the user
	// the base scene has an entity manager, the game sends the memory pool reference to the base scene,
	// the base scene gives the memory pool to the entity manager
	// the entity manager then passes the memory pool reference to the entity when it is created
	memoryPool *memoryPool

	RenderScale int
}

// NewGame creates a new game with the default configuration (config.toml)
func NewGame(sceneCFG []SceneConfig, gameCFG Config, memoryPool *memoryPool) Game {
	g := Game{conf: gameCFG}
	g.currentScene = g.conf.DefaultScene
	g.Scenes = make(map[string]Scene)
	g.virtualWidth, g.virtualHeight = g.conf.VirtualResolution[0], g.conf.VirtualResolution[1]
	g.memoryPool = memoryPool

	// Register the main scene and the menu scene
	for _, scene := range sceneCFG {
		g._registerScene(scene.Name, scene.Scene)
	}

	// Return the game
	return g
}

func (g *Game) _registerScene(name string, scene Scene) {
	g.Scenes[name] = scene
}

// _changeScene changes the current scene and unloads the previous one.
// It makes sure that the previous scene is unloaded and the new one is loaded.
// If the current scene is not done or the new scene is not loaded, it will log a fatal error.
func (g *Game) _changeScene(name string) {
	if g.Scenes[g.currentScene].IsLoaded() {
		// Unload the current scene
		g.Scenes[g.currentScene].UnloadBaseScene()
		g.Scenes[g.currentScene].Unload()
		// Check that the current scene is unloaded and done
	}

	// Change to the new scene
	g.currentScene = name
	// Load the new scene
	g.Scenes[g.currentScene].ReceiveMemoryPool(g.memoryPool)

	g.Scenes[g.currentScene].InitBaseScene()
	g.Scenes[g.currentScene].Init()
}

// GetScaleAndOffset returns the scale and offset values for the virtual resolution.
// The scale is the minimum of the screen width divided by the virtual width and the screen height divided by the virtual height.
// The offset is the difference between the screen width and the virtual width multiplied by the scale value, divided by 2.
// The same applies for the offset in the Y-axis.
func (g Game) GetScaleAndOffset() (scale float32, offsetX float32, offsetY float32) {
	scaleX := float32(rl.GetScreenWidth()) / g.virtualWidth
	scaleY := float32(rl.GetScreenHeight()) / g.virtualHeight
	scale = float32(math.Min(float64(scaleX), float64(scaleY)))
	// Calculate the offset for the virtual resolution.
	offsetX = (float32(rl.GetScreenWidth()) - g.virtualWidth*scale) / 2
	offsetY = (float32(rl.GetScreenHeight()) - g.virtualHeight*scale) / 2
	return scale, offsetX, offsetY
}

// Run runs the game loop.
// It creates a window with the dimensions of the virtual resolution,
// creates a render texture for the virtual resolution,
// initializes the current scene,
// and then runs the game loop.
// In the game loop it checks if the current scene is done,
// and if so, changes to the next scene.
// It then updates the current scene,
// renders the current scene to the render texture,
// and then draws the render texture to the screen.
// It also adds some debug information to the screen.
func (g Game) Run() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.conf.Resolution[0], g.conf.Resolution[1], "Suika Game")
	rl.SetTargetFPS(g.conf.FPS)
	defer rl.CloseWindow()

	// Create a render texture for the virtual resolution.
	RenderTexture := rl.LoadRenderTexture(int32(g.virtualWidth), int32(g.virtualHeight))
	defer rl.UnloadRenderTexture(RenderTexture)
	if g.Scenes[g.currentScene] == nil {
		panic(fmt.Sprint("The scene ", g.currentScene, " is not registered"))
	}
	g.Scenes[g.currentScene].ReceiveMemoryPool(g.memoryPool)
	g.Scenes[g.currentScene].InitBaseScene()
	g.Scenes[g.currentScene].Init()
	for !rl.WindowShouldClose() {
		if g.Scenes[g.currentScene].IsDone() {
			g._changeScene(g.Scenes[g.currentScene].NextScene())
		}
		// Calculate the scale factor for the virtual resolution.
		scale, offsetX, offsetY := g.GetScaleAndOffset()

		// Update the current scene.
		g.Scenes[g.currentScene].Update(g.virtualWidth, g.virtualHeight)

		// Begin the texture mode.
		rl.BeginTextureMode(RenderTexture)
		// Clear the background of the render texture.
		rl.ClearBackground(rl.Black)
		// Render the current scene to the virtual resolution
		g.Scenes[g.currentScene].Render()
		rl.EndTextureMode()
		// Begin drawing to the screen.
		rl.BeginDrawing()
		// Clear the background of the screen.
		rl.ClearBackground(rl.Black)

		// Draw the virtual resolution texture.
		src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.virtualWidth), Height: -float32(g.virtualHeight)} // Flip Y-axis
		dest := rl.Rectangle{
			X:      offsetX,
			Y:      offsetY,
			Width:  float32(g.virtualWidth) * scale,
			Height: float32(g.virtualHeight) * scale,
		}

		rl.DrawTexturePro(RenderTexture.Texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

		if g.conf.Debug {
			rl.DrawText(fmt.Sprintf("Scale: %.2f", scale), 10, 10, 20, rl.DarkGreen)
			rl.DrawText(fmt.Sprintf("OffsetX: %.2f OffsetY: %.2f", offsetX, offsetY), 10, 40, 20, rl.DarkGreen)
			rl.DrawText(fmt.Sprint("FPS ", rl.GetFPS()), 10, 70, 20, rl.DarkGreen)
		}
		// Debug information

		// End drawing to the screen.
		rl.EndDrawing()

	}
}
