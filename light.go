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
	l.integration.SendCommand(l, "turn_on")
	return nil
}

func (l *Light) TurnOff() error {
	l.integration.SendCommand(l, "turn_off")
	return nil
}
