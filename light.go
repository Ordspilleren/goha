package goha

type Light struct {
	HAEntity
}

func (l *Light) On() bool {
	if l.State().State == "on" {
		return true
	} else {
		return false
	}
}

func (l *Light) TurnOn() error {
	l.integration.SendCommand(l, "turn_on", nil)
	return nil
}

func (l *Light) TurnOff() error {
	l.integration.SendCommand(l, "turn_off", nil)
	return nil
}

func (l *Light) Fade(transitionTime int, brightness int) error {
	data := struct {
		Transition    int `json:"transition,omitempty"`
		BrightnessPct int `json:"brightness_pct,omitempty"`
	}{
		Transition:    transitionTime,
		BrightnessPct: brightness,
	}
	l.integration.SendCommand(l, "turn_on", data)
	return nil
}
