package components

import "game/engine"

const HealthComponentID = 0

type Health struct {
	engine.Component
	HP int
}

func (c Health) ID() int {
	return HealthComponentID
}
