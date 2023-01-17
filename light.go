package haautomations

type Light struct {
	EntityData
}

func (l *Light) IsOn() bool {
	if l.GetState().State == "on" {
		return true
	} else {
		return false
	}
}

func (l *Light) On() error {
	CallService(
		"light",
		"turn_on",
		nil,
		l.GetEntityID())
	return nil
}

func (l *Light) Off() error {
	CallService(
		"light",
		"turn_off",
		nil,
		l.GetEntityID())
	return nil
}
