package haautomations

type Entity interface {
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
}

type EntityData struct {
	State State
}

func (e *EntityData) GetEntityID() string {
	return e.State.EntityID
}

func (e *EntityData) SetEntityID(entityID string) {
	e.State.EntityID = entityID
}

func (e *EntityData) GetState() State {
	return e.State
}

func (e *EntityData) SetState(state State) {
	e.State = state
}

type EntityList map[string]Entity

func (e EntityList) AddLight(entityId string) *Light {
	entity := &Light{}
	entity.SetEntityID(entityId)
	e[entityId] = entity
	return entity
}

func (e EntityList) AddBinarySensor(entityId string) *BinarySensor {
	entity := &BinarySensor{}
	entity.SetEntityID(entityId)
	e[entityId] = entity
	return entity
}
