package engine

type BaseComponent struct {
	exists bool
}

func (c BaseComponent) Exists() bool {
	return c.exists
}
