package main

import goha "github.com/Ordspilleren/goha"

func SetupAutomations() {
	officeButton.SetAutomations(testAutomation)
}

var (
	officeLight  = ha.AddLight("light.office")
	officeButton = ha.AddBinarySensor("sensor.office_button")
)

var testAutomation = goha.Automation{
	Condition: goha.DefaultCondition,
	Action: func() error {
		if officeButton.Triggered() {
			officeLight.TurnOn()
		} else {
			officeLight.TurnOff()
		}
		return nil
	},
}
