package goha

type Entity interface {
	SetIntegration(Integration)
	GetEntityID() string
	SetEntityID(string)
	GetState() State
	SetState(State)
	GetPreviousState() State
	SetAutomations(...Automation) *HAEntity
	GetAutomations() []Automation
}

type HAEntity struct {
	integration   Integration
	entityID      string
	state         State
	previousState State
	automations   []Automation
}

func (e *HAEntity) SetIntegration(integration Integration) {
	e.integration = integration
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

func (e *HAEntity) SetAutomations(automations ...Automation) *HAEntity {
	e.automations = append(e.automations, automations...)
	return e
}

func (e *HAEntity) GetAutomations() []Automation {
	return e.automations
}

/* TODO: Can we make this work?

func (e *HAEntity) Light() *Light {
	return &Light{e}
}

func (e *HAEntity) BinarySensor() *BinarySensor {
	return &BinarySensor{e}
}

func (e *HAEntity) Sensor() *Sensor {
	return &Sensor{e}
}

func (e *HAEntity) Person() *Person {
	return &Person{e}
}

func (e *HAEntity) Sun() *Sun {
	return &Sun{e}
}

*/

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
