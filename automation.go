package main

type Automation struct {
	Trigger Trigger
	Action  func() error
}

type Trigger struct {
	Entity   Entity
	State    string
	Function func() bool
}

/*
func (a *Automation) SetTrigger(entity Entity, state string, condition func() bool) {
	a.trigger = Trigger{
		entity:   entity,
		state:    state,
		function: condition,
	}
}

func (a *Automation) SetAction(action func() error) {
	a.action = action
}
*/

func (a *Automation) GetTriggerEntityID() string {
	return a.Trigger.Entity.GetEntityID()
}
