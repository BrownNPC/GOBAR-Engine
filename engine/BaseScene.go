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
	EntityManager *EntityManager
	done          bool
}

func (s *BaseScene) unloadBaseScene() {
	s.done = true
	s.Loaded = false
	s.EntityManager.Destroy()
}
func (s *BaseScene) initBaseScene() {
	s.done = false
	s.Loaded = true
	s.EntityManager = newEntityManager()
}

// updates virtual resolution variables and the entity manager
func (s *BaseScene) UpdateBaseScene(virtualWidth float32, virtualHeight float32) {
	s.VirtualWidth, s.VirtualHeight = virtualWidth, virtualHeight
	s.EntityManager.Update()
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

func (s *BaseScene) isDone() bool {
	return s.done
}

func (s BaseScene) isLoaded() bool {
	return s.Loaded
}

// helper method
func (s *BaseScene) ForEachEntity(fn func(*Entity)) {
	entities := s.EntityManager.GetEntities()
	for e := entities.Front(); e != nil; e = e.Next() {
		entity, ok := e.Value.(*Entity)
		if !ok {
			continue
		}
		fn(entity)
	}
}
