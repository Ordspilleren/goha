package main

import homeautomation "github.com/Ordspilleren/goha"

var Automations = []homeautomation.Automation{
	testAutomation,
}

var (
	officeLight  = ha.AddLight("light.office")
	officeButton = ha.AddBinarySensor("sensor.office_button")
)

var testAutomation = homeautomation.Automation{
	Trigger: homeautomation.Trigger{
		Entity: officeButton,
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
