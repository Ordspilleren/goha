package goha

import "log"

type Condition func(Entity) bool
type Action func(Entity) error

type Automation struct {
	Triggers  []Entity
	Condition Condition
	Action    Action
}

func DefaultCondition(e Entity) bool {
	return e.State().State != e.PreviousState().State
}

func (a *Automation) Evaluate(entity Entity) {
	for i := range a.Triggers {
		if a.Triggers[i].EntityID() == entity.EntityID() {
			if a.Condition(entity) {
				log.Printf("%s changed and condition met! previous state: %s, current state: %s", entity.EntityID(), entity.PreviousState().State, entity.State().State)
				a.Action(entity)
				return
			}
		}
	}
}

func NewAutomation(condition Condition, action Action, triggers ...Entity) Automation {
	return Automation{
		Triggers:  triggers,
		Condition: condition,
		Action:    action,
	}
}
