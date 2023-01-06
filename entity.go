package main

type Entity interface {
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
}

type Entities map[string]Entity

func (e Entities) AddLight(entityId string) *Light {
	entity := &Light{}
	entity.SetEntityID(entityId)
	e[entityId] = entity
	return entity
}

func (e Entities) AddBinarySensor(entityId string) *BinarySensor {
	entity := &BinarySensor{}
	entity.SetEntityID(entityId)
	e[entityId] = entity
	return entity
}
