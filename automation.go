package goha

type Automation struct {
	Trigger   Trigger
	Condition func() bool
	Action    func() error
}

type Trigger struct {
	Entity Entity
	State  string
}

type AutomationList []Automation

func (a *Automation) Evaluate(entityId string, state string) {
	if a.Trigger.Entity.GetEntityID() == entityId && (a.Trigger.State == state || a.Trigger.State == "") {
		if a.Condition() {
			a.Action()
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
