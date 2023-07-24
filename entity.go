package goha

type Entity interface {
	SetIntegration(Integration)
	EntityID() string
	SetEntityID(string)
	State() State
	SetState(State)
	PreviousState() State
}

type HAEntity struct {
	integration   Integration
	entityID      string
	state         State
	previousState State
}

func (e *HAEntity) SetIntegration(integration Integration) {
	e.integration = integration
}

func (e *HAEntity) EntityID() string {
	return e.entityID
}

func (e *HAEntity) SetEntityID(entityID string) {
	e.entityID = entityID
}

func (e *HAEntity) State() State {
	return e.state
}

func (e *HAEntity) SetState(state State) {
	e.previousState = e.state
	e.state = state
}

func (e *HAEntity) PreviousState() State {
	return e.previousState
}
