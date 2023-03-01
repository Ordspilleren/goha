package main

import (
	"fmt"
	"testing"

	goha "github.com/Ordspilleren/goha"
)

var onState = "on"

type MockIntegration struct {
}

func (t *MockIntegration) SendCommand(entity goha.Entity, action string, data any) error {
	switch t := entity.(type) {
	case *goha.Light:
		if action == "turn_on" {
			t.SetState(goha.State{State: &onState})
		}
	default:
		return fmt.Errorf("unknown entity type: %v", t)
	}
	return nil
}

func TestAutomation(t *testing.T) {
	officeLight.SetIntegration(&MockIntegration{})
	officeButton.SetIntegration(&MockIntegration{})

	officeButton.SetState(goha.State{
		State: &onState,
	})

	testAutomation.Evaluate(officeButton)

	if officeLight.GetState().State != &onState {
		t.Error("light not on")
	}
}
