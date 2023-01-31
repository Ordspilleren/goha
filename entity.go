package goha

import "github.com/Ordspilleren/goha/wsclient"

type Entity interface {
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
}

type HAEntity struct {
	wsClient *wsclient.Client
	EntityID string
	State    State
}

func (e *HAEntity) GetEntityID() string {
	return e.EntityID
}

func (e *HAEntity) SetEntityID(entityID string) {
	e.EntityID = entityID
}

func (e *HAEntity) GetState() State {
	return e.State
}

func (e *HAEntity) SetState(state State) {
	e.State = state
}

func (ha *HomeAutomation) AddLight(entityId string) *Light {
	entity := &Light{}
	entity.wsClient = &ha.wsClient
	entity.SetEntityID(entityId)
	ha.Entities = append(ha.Entities, entity)
	return entity
}

func (ha *HomeAutomation) AddBinarySensor(entityId string) *BinarySensor {
	entity := &BinarySensor{}
	entity.wsClient = &ha.wsClient
	entity.SetEntityID(entityId)
	ha.Entities = append(ha.Entities, entity)
	return entity
}

func (ha *HomeAutomation) AddSensor(entityId string) *Sensor {
	entity := &Sensor{}
	entity.wsClient = &ha.wsClient
	entity.SetEntityID(entityId)
	ha.Entities = append(ha.Entities, entity)
	return entity
}
