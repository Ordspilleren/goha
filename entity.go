package goha

import "github.com/Ordspilleren/goha/wsclient"

type Entity interface {
	SetClient(*wsclient.Client)
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
	MergeState(State)
}

type HAEntity struct {
	wsClient *wsclient.Client
	EntityID string
	State    State
}

func (e *HAEntity) SetClient(client *wsclient.Client) {
	e.wsClient = client
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

func (e *HAEntity) MergeState(state State) {
	e.State.Merge(state)
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
