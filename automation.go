package goha

type Condition func(*Trigger) bool
type Action func() error

type Automation struct {
	Triggers  []Trigger
	Condition Condition
	Action    Action
}

type Trigger struct {
	Entity Entity
}

func DefaultCondition(t *Trigger) bool {
	return t.EnsureDifferentState()
}

func (t *Trigger) EnsureDifferentState() bool {
	return t.Entity.GetState().State != t.Entity.GetPreviousState().State
}

func (a *Automation) Evaluate(entityId string) {
	for _, trigger := range a.Triggers {
		if trigger.Entity.GetEntityID() == entityId {
			if a.Condition(&trigger) {
				a.Action()
			}
		}
	}
}
