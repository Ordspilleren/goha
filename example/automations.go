package main

import goha "github.com/Ordspilleren/goha"

var (
	officeLight  = homeautomation.AddLight("light.office")
	officeButton = homeautomation.AddBinarySensor("sensor.office_button")
)

var automations = []goha.Automation{
	testAutomation,
}

var testAutomation = goha.Automation{
	Triggers: []goha.Entity{
		officeButton,
	},
	Condition: goha.DefaultCondition,
	Action: func(e goha.Entity) error {
		if officeButton.Triggered() {
			officeLight.TurnOn()
		} else {
			officeLight.TurnOff()
		}
		return nil
	},
}
