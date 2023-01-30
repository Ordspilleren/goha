package main

import homeautomation "github.com/Ordspilleren/homeautomation"

var Entities = make(homeautomation.EntityList)

var Automations = homeautomation.AutomationList{
	testAutomation,
}

var (
	bedLight         = Entities.AddLight("light.bed_light")
	movementBackyard = Entities.AddBinarySensor("binary_sensor.movement_backyard")
)

var testAutomation = homeautomation.Automation{
	Trigger: homeautomation.Trigger{
		Entity: movementBackyard,
	},
	Condition: func() bool {
		return true
	},
	Action: func() error {
		if movementBackyard.Triggered() {
			bedLight.On()
		} else {
			bedLight.Off()
		}
		return nil
	},
}
