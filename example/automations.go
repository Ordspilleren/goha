package main

import goha "github.com/Ordspilleren/goha"

func SetupAutomations() {
	officeButton.SetAutomations(testAutomation)
}

var (
	officeLight  = homeautomation.AddLight("light.office")
	officeButton = homeautomation.AddBinarySensor("sensor.office_button")
)

var testAutomation = goha.Automation{
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
