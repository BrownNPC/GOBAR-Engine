package engine

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// this is defined here but should be used by scenes
type Video struct {
	Resolution []int32 `toml:"resolution"`
	Windowed   bool    `toml:"windowed"`
	FPS        int32   `toml:"fps"`
}

type Internal struct {
	DefaultScene      string    `toml:"defaultScene"`
	VirtualResolution []float32 `toml:"virtualResolution"`
	Debug             bool      `toml:"debug"`
	WindowTitle       string    `toml:"windowTitle"`
}

type Config struct {
	Video    `toml:"Video"`
	Internal `toml:"Internal"`
}

/*
make this `config.toml` in your project root

[Video]

resolution=[1280, 720]

fps = 60

windowed=true

[Internal]

defaultScene="menu"

virtualResolution=[1280,720]

debug=true
*/
func LoadConfig() Config {
	configPath := "./config.toml"
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	conf := &Config{}
	_, err = toml.NewDecoder(file).Decode(conf)
	if err != nil {
		panic(err)
	}
	return *conf
}

// describes which scenes should be registered
type SceneConfig struct {
	Name  string
	Scene Scene
}

func ConfigureNewScene(name string, scene Scene) SceneConfig {
	return SceneConfig{
		Name:  name,
		Scene: scene,
	}
}
