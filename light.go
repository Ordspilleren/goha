package goha

type Light struct {
	HAEntity
}

func (l *Light) On() bool {
	if l.GetState().State == "on" {
		return true
	} else {
		return false
	}
}

func (l *Light) TurnOn() error {
	l.CallService(
		"light",
		"turn_on",
		nil,
		l.GetEntityID())
	return nil
}

func (l *Light) TurnOff() error {
	l.CallService(
		"light",
		"turn_off",
		nil,
		l.GetEntityID())
	return nil
}
