package goha

type Automation struct {
	Triggers  []Entity
	Condition func() bool
	Action    func() error
}

func (a *Automation) Evaluate(entityId string, previousState State) {
	for _, trigger := range a.Triggers {
		if trigger.GetEntityID() == entityId {
			if a.Condition() && trigger.GetState().State != previousState.State {
				a.Action()
			}
		}
	}
}

/*
type Automation struct {
	Trigger   Trigger
	Condition func() bool
	Action    func() error
}

type Trigger struct {
	Entity Entity
	State  string
}

func (a *Automation) SetTrigger(entity Entity, state string) {
	a.Trigger = Trigger{
		Entity: entity,
		State:  state,
	}
}

func (a *Automation) SetCondition(condition func() bool) {
	a.Condition = condition
}

func (a *Automation) SetAction(action func() error) {
	a.Action = action
}

func (a *Automation) GetTriggerEntityID() string {
	return a.Trigger.Entity.GetEntityID()
}
*/
