package goha

import "github.com/Southclaws/opt"

type Light struct {
	HAEntity
}

func (l *Light) On() bool {
	if l.GetState().State.OrZero() == "on" {
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

func (l *Light) Dim(transitionTime int, brightness int) error {
	data := Attributes{
		Transition: opt.New(transitionTime),
		Brightness: opt.New(brightness),
	}
	l.integration.SendCommand(l, "turn_on", data)
	return nil
}
