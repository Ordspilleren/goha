package main

var (
	bedLight         = Devices["light.bed_light"].(*Light)
	movementBackyard = Devices["binary_sensor.movement_backyard"].(*BinarySensor)
)

var Automations = []Automation{
	testAutomation,
}

var testAutomation = Automation{
	Trigger: Trigger{
		Entity: movementBackyard,
		State:  "on",
	},
	Condition: func() bool {
		return !bedLight.IsOn()
	},
	Action: func() error {
		bedLight.On()
		return nil
	},
}
