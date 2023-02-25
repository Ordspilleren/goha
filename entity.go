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

// TODO: Maybe having this as part of the entity doesn't make sense.
// We could move it to the automation with syntax like automation.Triggers(...Entity).
func (e *HAEntity) SetAutomations(automations ...Automation) *HAEntity {
	e.automations = append(e.automations, automations...)
	return e
}

func (e *HAEntity) GetAutomations() []Automation {
	return e.automations
}
