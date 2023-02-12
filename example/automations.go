package main

import goha "github.com/Ordspilleren/goha"

var Automations = []goha.Automation{
	testAutomation,
}

var (
	officeLight  = ha.AddLight("light.office")
	officeButton = ha.AddBinarySensor("sensor.office_button")
)

var testAutomation = goha.Automation{
	Triggers: []goha.Entity{
		officeButton,
	},
	Condition: func() bool {
		return true
	},
	Action: func() error {
		if officeButton.Triggered() {
			officeLight.On()
		} else {
			officeLight.Off()
		}
		return nil
	},
}
