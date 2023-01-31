package goha

type Light struct {
	HAEntity
}

func (l *Light) IsOn() bool {
	if l.GetState().State == "on" {
		return true
	} else {
		return false
	}
}

func (l *Light) On() error {
	l.CallService(
		"light",
		"turn_on",
		nil,
		l.GetEntityID())
	return nil
}

func (l *Light) Off() error {
	l.CallService(
		"light",
		"turn_off",
		nil,
		l.GetEntityID())
	return nil
}
