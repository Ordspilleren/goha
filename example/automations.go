package main

import homeautomation "github.com/Ordspilleren/goha"

var Automations = homeautomation.AutomationList{
	testAutomation,
}

var (
	officeLight  = ha.AddLight("light.0x000b57fffe115b4f")
	officeButton = ha.AddBinarySensor("sensor.0x680ae2fffe2f3b59_action")
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
