package components

import (
	"game/engine"

	"github.com/jakecoffman/cp"
)

const TransformComponentId = 1

// transform component
type Transform struct {
	engine.Component
	Position cp.Vector
	Angle    float64
}

func (c Transform) ID() int {
	return TransformComponentId
}
