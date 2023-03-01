package main

import (
	"fmt"
	"testing"

	goha "github.com/Ordspilleren/goha"
)

type MockIntegration struct {
}

func (t *MockIntegration) SendCommand(entity goha.Entity, action string, data any) error {
	switch t := entity.(type) {
	case *goha.Light:
		if action == "turn_on" {
			t.SetState(goha.State{State: "on"})
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
		State: "on",
	})

	testAutomation.Evaluate(officeButton)

	if officeLight.GetState().State != "on" {
		t.Error("light not on")
	}
}
