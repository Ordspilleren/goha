package goha

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

type EntityList []Entity

func (ha *HomeAutomation) AddLight(entityId string) *Light {
	entity := &Light{}
	entity.SetEntityID(entityId)
	ha.Entities = append(ha.Entities, entity)
	return entity
}

func (ha *HomeAutomation) AddBinarySensor(entityId string) *BinarySensor {
	entity := &BinarySensor{}
	entity.SetEntityID(entityId)
	ha.Entities = append(ha.Entities, entity)
	return entity
}
