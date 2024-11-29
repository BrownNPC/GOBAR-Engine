package engine

// Entity is simply a composition of one or more Components with an Id.
type Entity struct {
	ID         int
	Alive      bool
	Components map[int]Component // Maps component type IDs to their data
	Tag        string
}

func newEntity(Id int, Tag string, components ...Component) *Entity {
	e := &Entity{
		Components: make(map[int]Component, len(components)),
		Tag:        Tag,
		Alive:      true,
		ID:         Id,
	}
	e.setComponentIds(components...)
	return e
}

func (e *Entity) Kill() {
	e.Alive = false
}

func (e *Entity) setComponentIds(components ...Component) {
	for i, c := range components {
		e.Components[i] = c
	}
}

func (e *Entity) AddComponent(c Component) {
	e.Components[c.ID()] = c
}

func (e *Entity) RemoveComponent(c Component) {
	delete(e.Components, c.ID())
}

func (e *Entity) GetComponent(ComponentId int) (Component, bool) {
	c, ok := e.Components[ComponentId]
	return c, ok
}

func (e *Entity) GetTag() string {
	return e.Tag
}

func (e *Entity) HasComponent(ComponentId int) bool {
	_, ok := e.Components[ComponentId]
	return ok
}
