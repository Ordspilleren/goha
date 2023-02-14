package goha

import "github.com/Ordspilleren/goha/wsclient"

type Entity interface {
	SetClient(*wsclient.Client)
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
	GetPreviousState() State
}

type HAEntity struct {
	wsClient      *wsclient.Client
	entityID      string
	state         State
	previousState State
}

func (e *HAEntity) SetClient(client *wsclient.Client) {
	e.wsClient = client
}

func (e *HAEntity) GetEntityID() string {
	return e.entityID
}

func (e *HAEntity) SetEntityID(entityID string) {
	e.entityID = entityID
}

func (e *HAEntity) GetState() State {
	return e.state
}

func (e *HAEntity) SetState(state State) {
	e.previousState = e.state
	e.state.Merge(state)
}

func (e *HAEntity) GetPreviousState() State {
	return e.previousState
}

func (ha *HomeAutomation) AddLight(entityId string) *Light {
	return ha.AddEntity(&Light{}, entityId).(*Light)
}

func (ha *HomeAutomation) AddBinarySensor(entityId string) *BinarySensor {
	return ha.AddEntity(&BinarySensor{}, entityId).(*BinarySensor)
}

func (ha *HomeAutomation) AddSensor(entityId string) *Sensor {
	return ha.AddEntity(&Sensor{}, entityId).(*Sensor)
}

func (ha *HomeAutomation) AddPerson(entityId string) *Person {
	return ha.AddEntity(&Person{}, entityId).(*Person)
}

func (ha *HomeAutomation) AddSun(entityId string) *Sun {
	return ha.AddEntity(&Sun{}, entityId).(*Sun)
}
