package engine

// The `GenericComponent` interface is a common base for all components in the system.
// Components represent pieces of data or functionality attached to entities.
// The `Exists` and `SetExists` methods provide basic functionality for determining
// whether a GenericComponent is active or not. This interface ensures uniformity
// for all components in the memory pool.
type GenericComponent interface {
	Exists() bool
	Default() GenericComponent
}

// `GenericComponentSlice` is a wrapper for a slice of components of a specific type.
// Each slice corresponds to all components of that type in the system. For example,
// if Position is a component, this slice would hold all Position components for every entity.
// Wrapping the slice allows us to organize and handle slices in a generic way without duplicating code.
type GenericComponentSlice[T GenericComponent] struct {
	Items []T // A dynamically resizable array holding components of type T.
}

// `memoryPool` is the central structure for managing components in the system.
// It maps component types to their respective data storage (slices).
// Entities are represented as indices, so their components are stored at the same index in their respective slices.
// This design makes accessing components efficient and avoids redundant storage.
type memoryPool struct {
	EstimateMaxEntities int                 // A rough guess of the maximum number of entities the game will have.
	ComponentsMap       map[int]interface{} // Maps component indices (types) to slices of components.
	Active              []bool              // track which entities are active (columns)
	Tags                []string            // track
}

// `NewMemoryPool` initializes a new `memoryPool` instance.
// It prepares the system to store component data for multiple entities.
// The estimate helps preallocate storage for efficiency.
func NewMemoryPool(estimateMaxEntities int) *memoryPool {
	return &memoryPool{
		EstimateMaxEntities: estimateMaxEntities,       // Store the maximum number of entities for preallocation.
		ComponentsMap:       make(map[int]interface{}), // Initialize an empty map to hold components.
		Active:              make([]bool, estimateMaxEntities),
		Tags:                make([]string, estimateMaxEntities),
	}
}

// `memoryPoolRegisterComponent` registers a new component type in the memory pool at the specified index.
// This function prepares a slice to store all components of this type. The index acts as an ID
// for the component type, ensuring we can distinguish between different component types.
// Generics are used here to allow handling components of any type that implements `component`.
//
// Example: Registering Position components would create a slice to store Position instances for all entities.
func MemoryPoolRegisterComponent[T GenericComponent](mp *memoryPool, index int) {
	// Create a new slice for the component type, preallocated with space for all entities.
	mp.ComponentsMap[index] = GenericComponentSlice[T]{
		Items: make([]T, mp.EstimateMaxEntities), // Preallocate space for efficient access.
	}
}

// `memoryPoolGetComponent` retrieves the component of type T for a specific entity and index.
// It looks up the appropriate slice (based on the index), and then fetches the component
// at the position corresponding to the entity's ID.
//
// This function ensures type safety by using generics to confirm that the requested component
// type matches the one stored at the index.
//
// Returns: The component of type T, or a zero value if the component isn't found or an error occurs.
func memoryPoolGetComponent[T GenericComponent](mp *memoryPool, entityID, componentIndex int) (T, bool) {
	// Look up the slice for the requested component type.
	wrapper, ok := mp.ComponentsMap[componentIndex].(GenericComponentSlice[T])
	if !ok { // If the slice doesn't exist or the type doesn't match, return a default value.
		var zero T
		return zero, false
	}
	// Ensure the entity ID is within bounds of the slice.
	if entityID < 0 || entityID >= len(wrapper.Items) {
		var zero T
		return zero, false
	}
	return wrapper.Items[entityID], true // Return the requested component.
}

// `memoryPoolSetComponent` assigns a component of type T to a specific entity at a specific index.
// This function allows adding or updating a component for an entity.
//
// Example: Setting a Position component for entity 5 would involve updating the Position slice
// at index 5 with the new component data.
//
// Returns: An error if the index isn't registered or the entity ID is invalid.
// func memoryPoolSetComponent[T GenericComponent](mp *memoryPool, entityID, componentIndex int, component T) error {
// 	// Look up the slice for the requested component type.
// 	wrapper, ok := mp.ComponentsMap[componentIndex].(GenericComponentSlice[T])
// 	if !ok { // If the slice doesn't exist, return an error.
// 		return fmt.Errorf("component type not registered at index %d", componentIndex)
// 	}

// 	// Update the slice with the new component data for the entity.
// 	wrapper.Items[entityID] = component
// 	mp.ComponentsMap[componentIndex] = wrapper // Save the updated slice back into the map.
// 	return nil
// }

func getNextAvailableEntityID(mp *memoryPool) int {
	for i, active := range mp.Active {
		if !active {
			mp.Active[i] = true
			return i
		}
	}
	return -1
}

func memoryPoolAddEntity(mp *memoryPool, tag string) entity {
	entityID := getNextAvailableEntityID(mp)
	mp.Active[entityID] = true
	mp.Tags[entityID] = tag
	//set all the components to default
	for _, v := range mp.ComponentsMap {
		if wrapper, ok := v.(GenericComponentSlice[GenericComponent]); ok {
			wrapper.Items[entityID] = v.(GenericComponentSlice[GenericComponent]).Items[0]
		}
	}

	return newEntity(entityID, mp)
}

func memoryPoolGetTag(mp *memoryPool, entityID int) string {
	return mp.Tags[entityID]
}
