package goha

import (
	"sync"
)

type Goha struct {
	PrimaryIntegration Integration
}

func (goha *Goha) Start(waitGroup *sync.WaitGroup) error {
	goha.PrimaryIntegration.Start(waitGroup)

	return nil
}

func (goha *Goha) AddEntity(entity Entity, entityId string) Entity {
	entity.SetIntegration(goha.PrimaryIntegration)
	entity.SetEntityID(entityId)
	goha.PrimaryIntegration.RegisterEntities(entity)

	return entity
}

func (goha *Goha) AddAutomations(automations ...Automation) error {
	goha.PrimaryIntegration.RegisterAutomations(automations...)
	return nil
}

func (goha *Goha) AddLight(entityId string) *Light {
	return goha.AddEntity(&Light{}, entityId).(*Light)
}

func (goha *Goha) AddBinarySensor(entityId string) *BinarySensor {
	return goha.AddEntity(&BinarySensor{}, entityId).(*BinarySensor)
}

func (goha *Goha) AddSensor(entityId string) *Sensor {
	return goha.AddEntity(&Sensor{}, entityId).(*Sensor)
}

func (goha *Goha) AddPerson(entityId string) *Person {
	return goha.AddEntity(&Person{}, entityId).(*Person)
}

func (goha *Goha) AddSun(entityId string) *Sun {
	return goha.AddEntity(&Sun{}, entityId).(*Sun)
}

func (goha *Goha) AddMediaPlayer(entityId string) *MediaPlayer {
	return goha.AddEntity(&MediaPlayer{}, entityId).(*MediaPlayer)
}

func (goha *Goha) AddVacuum(entityId string) *Vacuum {
	return goha.AddEntity(&Vacuum{}, entityId).(*Vacuum)
}

func (goha *Goha) AddSchedule(entityId string) *Schedule {
	return goha.AddEntity(&Schedule{}, entityId).(*Schedule)
}

func (goha *Goha) AddInputDatetime(entityId string) *InputDatetime {
	return goha.AddEntity(&InputDatetime{}, entityId).(*InputDatetime)
}
