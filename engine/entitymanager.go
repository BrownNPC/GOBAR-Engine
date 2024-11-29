package engine

import "container/list"

type EntityManager struct {
	entities     *list.List            // Linked list of entities
	mapEntities  map[string]*list.List // Map of tags to lists of entities
	nextEntityID int
}

func newEntityManager() *EntityManager {
	return &EntityManager{
		entities:    list.New(),
		mapEntities: make(map[string]*list.List),
	}
}

// CreateEntity creates a new entity with the given tag and components and adds it to the EntityManager.
// It retrieves the next available entity ID, creates a new entity with the given tag and components,
// adds the entity to the main list of entities,
// and adds a reference to the entity in the map of tags to entities.
// The new entity is then returned.
func (em *EntityManager) CreateEntity(tag string, components ...Component) *Entity {
	e := newEntity(em.nextEntityID, tag, components...)
	em.nextEntityID++
	element := em.entities.PushBack(e)
	if _, exists := em.mapEntities[tag]; !exists {
		em.mapEntities[tag] = list.New()
	}
	em.mapEntities[tag].PushBack(element)
	return e
}

// GetFirstEntity retrieves the first entity matching the given tag. If no entities are found, it returns nil.
func (em *EntityManager) GetFirstEntity(tag string) *Entity {
	for e := em.mapEntities[tag].Front(); e != nil; e = e.Next() {
		return e.Value.(*Entity)
	}
	return nil
}

// returns list of entities with a specific component
func (em *EntityManager) GetEntities() *list.List {
	return em.entities
}

// returns list of entities with a specific tag
func (em *EntityManager) GetEntitiesByTag(tag string) *list.List {
	return em.mapEntities[tag]
}

func (em *EntityManager) Update() {
	for e := em.entities.Front(); e != nil; {
		entity := e.Value.(*Entity)
		if entity.Alive {
			// Update components of alive entities
			e = e.Next()
		} else {
			// Remove dead entity from mapEntities
			em.removeFromMap(entity.Tag, e)

			// Remove dead entity from the main list
			next := e.Next()
			em.entities.Remove(e)
			e = next
		}
	}
}

// removeFromMap removes an entity from the mapEntities.
func (em *EntityManager) removeFromMap(tag string, element *list.Element) {
	tagList, exists := em.mapEntities[tag]
	if !exists {
		return
	}

	for e := tagList.Front(); e != nil; e = e.Next() {
		if e.Value.(*list.Element) == element {
			tagList.Remove(e)
			break
		}
	}

	if tagList.Len() == 0 {
		delete(em.mapEntities, tag)
	}
}

func (em *EntityManager) Destroy() {
	em.entities = nil
	em.mapEntities = nil
}
