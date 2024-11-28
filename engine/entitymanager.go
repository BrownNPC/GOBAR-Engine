package engine

type EntityManager struct {
	entities   []entity
	MemoryPool *memoryPool
	// how many entities have been created since the start of this entity manager

}

func newEntityManager(memoryPool *memoryPool) EntityManager {
	em := EntityManager{
		MemoryPool: memoryPool,
	}
	return em
}

func (m *EntityManager) Update() {
	// if new entities were created on the previous frame

}

func (m *EntityManager) CreateEntity(tag string) entity {
	e := memoryPoolAddEntity(m.MemoryPool, tag)
	return e
}

func (m *EntityManager) GetAliveEntities() []entity {
	return m.entities
}

// free all the memory, textures etc. and destroy all the entities
func (m *EntityManager) Destroy() {
}
