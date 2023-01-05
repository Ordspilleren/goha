package main

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

func (l *Light) On() error {
	return nil
}
