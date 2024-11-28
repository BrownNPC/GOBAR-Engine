package engine

type entity struct {
	id int

	memoryPool *memoryPool
}

func (e entity) Kill() {

}

func newEntity(id int, memoryPool *memoryPool) entity {
	return entity{
		id:         id,
		memoryPool: memoryPool,
	}
}

func (e entity) AddComponent(ComponentIndex int, Component GenericComponent) entity {

	return e
}

func (e entity) GetComponent(ComponentIndex int, Component GenericComponent) {
	memoryPoolGetComponent[GenericComponent](e.memoryPool, e.id, ComponentIndex)
}
