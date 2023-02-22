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
	message := Message{
		ID:          InteractionID(),
		Type:        "call_service",
		Domain:      "light",
		Service:     "turn_on",
		ServiceData: nil,
		Target: &Target{
			EntityID: l.GetEntityID(),
		},
	}
	l.ChangeState(message)
	return nil
}

func (l *Light) TurnOff() error {
	message := Message{
		ID:          InteractionID(),
		Type:        "call_service",
		Domain:      "light",
		Service:     "turn_off",
		ServiceData: nil,
		Target: &Target{
			EntityID: l.GetEntityID(),
		},
	}
	l.ChangeState(message)
	return nil
}
