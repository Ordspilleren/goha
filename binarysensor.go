package main

type BinarySensor struct {
	state State
}

func (l *BinarySensor) GetEntityID() string {
	return l.state.EntityID
}

func (l *BinarySensor) SetEntityID(entityID string) {
	l.state.EntityID = entityID
}

func (l *BinarySensor) GetState() State {
	return l.state
}

func (l *BinarySensor) SetState(state State) {
	l.state = state
}

func (l *BinarySensor) Triggered() bool {
	if l.state.State == "on" {
		return true
	} else {
		return false
	}
}
