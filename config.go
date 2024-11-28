package main

// this is defined here but should be used by scenes
type TextureInfo struct {
	FilePath string    `toml:"filePath"`
	Size     []float32 `toml:"size"`
	Offset   []float32 `toml:"offset"`
	Tag      string    `toml:"tag"`
}
