package engine

import (
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BaseScene struct {
	Loaded        bool
	VirtualWidth  float32
	VirtualHeight float32
	EntityManager EntityManager
	memoryPool    *memoryPool
	done          bool
}

func (s *BaseScene) UnloadBaseScene() {
	s.done = true
	s.Loaded = false
	// EntityManager.Destroy()
}
func (s *BaseScene) InitBaseScene() {
	s.done = false
	s.Loaded = true
}

func (s *BaseScene) UpdateVirtualResolution(virtualWidth float32, virtualHeight float32) {
	s.VirtualWidth, s.VirtualHeight = virtualWidth, virtualHeight
}

// draws texture according to virtual Resolution
func (s BaseScene) DrawTexture(texture rl.Texture2D, dest rl.Rectangle, tint color.RGBA) {
	if s.VirtualWidth == 0 || s.VirtualHeight == 0 {
		log.Fatal("please call s.UpdateVirtualResolution in your update method in your scene")
	}
	src := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))

	rl.DrawTexturePro(texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, tint)
}

// draw the texture anchored to the center, so it rotates around its center
func (s BaseScene) DrawTextureRotateCenter(texture rl.Texture2D, dest rl.Rectangle, rotation float32, tint color.RGBA) {
	if s.VirtualWidth == 0 || s.VirtualHeight == 0 {
		log.Fatal("please call s.UpdateVirtualResolution in your update method in your scene")
	}
	src := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	rl.DrawTexturePro(texture, src, dest, rl.Vector2{X: float32(texture.Width) / 2, Y: float32(texture.Height) / 2}, rotation, tint)
}

func (s *BaseScene) GoToNextScene() {
	s.done = true
}

func (s *BaseScene) IsDone() bool {
	return s.done
}

// receives memory pool and initializes entity manager with it
func (s *BaseScene) ReceiveMemoryPool(memoryPool *memoryPool) {
	s.memoryPool = memoryPool
	s.EntityManager = newEntityManager(s.memoryPool)
}
