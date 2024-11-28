package components

import "game/engine"

type Health struct {
	engine.BaseComponent
	HP int
}

func (c Health) Default() engine.GenericComponent {
	return Health{}
}
