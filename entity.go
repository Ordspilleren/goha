package goha

type Entity interface {
	SetIntegration(Integration)
	EntityID() string
	SetEntityID(string)
	State() State
	SetState(State)
	PreviousState() State
	SetAutomations(...Automation) *HAEntity
	Automations() []Automation
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

// TODO: Maybe having this as part of the entity doesn't make sense.
// We could move it to the automation with syntax like automation.Triggers(...Entity).
func (e *HAEntity) SetAutomations(automations ...Automation) *HAEntity {
	e.automations = append(e.automations, automations...)
	return e
}

func (e *HAEntity) Automations() []Automation {
	return e.automations
}
