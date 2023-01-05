package main

var Automations = []Automation{
	testAutomation,
}

var testAutomation = Automation{
	Trigger: Trigger{
		Entity: Devices["binary_sensor.movement_backyard"].(*BinarySensor),
		State:  "on",
		Function: func() bool {
			return true
		},
	},
	Action: func() error {
		Devices["light.bed_light"].(*Light).On()
		return nil
	},
}
