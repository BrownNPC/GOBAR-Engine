package components

import (
	"game/engine"

	"github.com/jakecoffman/cp"
)

const TransformComponentId = 1

// transform component
type Transform struct {
	Position cp.Vector
	Angle    float64
	engine.BaseComponent
}

func (c Transform) ID() int {
	return TransformComponentId
}
