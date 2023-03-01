package goha

type Condition func(Entity) bool
type Action func(Entity) error

type Automation struct {
	Condition Condition
	Action    Action
}

func DefaultCondition(e Entity) bool {
	return e.GetState().State != e.GetPreviousState().State
}

func (a *Automation) Evaluate(entity Entity) {
	if a.Condition(entity) {
		a.Action(entity)
	}
}
