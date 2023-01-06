package main

var (
	bedLight         = Devices.AddLight("light.bed_light")
	movementBackyard = Devices.AddBinarySensor("binary_sensor.movement_backyard")
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
