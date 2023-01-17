package main

import haautomations "github.com/Ordspilleren/ha-automations"

var Entities = make(haautomations.EntityList)

var Automations = haautomations.AutomationList{
	testAutomation,
}

var (
	bedLight         = Entities.AddLight("light.bed_light")
	movementBackyard = Entities.AddBinarySensor("binary_sensor.movement_backyard")
)

var testAutomation = haautomations.Automation{
	Trigger: haautomations.Trigger{
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
