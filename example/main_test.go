package main

import (
	"fmt"
	"sync"
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
			t.SetState(goha.State{State: onState})
		}
	default:
		return fmt.Errorf("unknown entity type: %v", t)
	}
	return nil
}

func (t *MockIntegration) Start(waitGroup *sync.WaitGroup) error {

	return nil
}

func (t *MockIntegration) RegisterEntities(entities ...goha.Entity) error {

	return nil
}

func (t *MockIntegration) RegisterAutomations(automations ...goha.Automation) error {

	return nil
}

func TestAutomation(t *testing.T) {
	officeLight.SetIntegration(&MockIntegration{})
	officeButton.SetIntegration(&MockIntegration{})

	officeButton.SetState(goha.State{
		State: onState,
	})

	testAutomation.Evaluate(officeButton)

	if officeLight.State().State != onState {
		t.Error("light not on")
	}
}
