package engine

// Component interface with an ID method instead of Mask.
type Component interface {
	ID() int
}

type BaseComponent struct {
}
