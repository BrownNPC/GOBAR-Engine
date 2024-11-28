package components

import (
	"game/engine"

	"github.com/jakecoffman/cp"
)

// transform component
type CTransform struct {
	Position cp.Vector
	Angle    float64
	engine.BaseComponent
}

func (c CTransform) Default() engine.GenericComponent {
	return CTransform{}
}
