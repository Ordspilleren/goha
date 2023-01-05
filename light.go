package main

import "log"

type Light struct {
	state State
}

func (l *Light) GetEntityID() string {
	return l.state.EntityID
}

func (l *Light) GetState() State {
	return l.state
}

func (l *Light) SetState(state State) {
	l.state = state
}

func (l *Light) IsOn() bool {
	if l.GetState().State == "on" {
		return true
	} else {
		return false
	}
}

func (l *Light) On() error {
	log.Print("light turned on")
	return nil
}
