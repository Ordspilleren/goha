# goha
Write Home Assistant automations in Go!

## Features and Design Considerations
- **Simplicity:** Only provide the bare minimum to integrate with Home Assistant.
- **Low friction:** Writing automations should be an easy task provided you know Go. No unnecessary boilerplate.
- **High performance:** Evaluate automations in separate threads to avoid blocking incoming state changes.

## Quick Start
### Installation
Set up a new (or use an existing) Go repository and import the library.
````
go get https://github.com/Ordspilleren/goha
````

### Writing Automations
In the [`example/`](./example) folder you will find an example project with some automations to use as a starting point.

Writing automations is generally very easy. You simply define the devices you want to use across your automations, both ones that will trigger, but also ones that you simply want to control, like so:
```go
var (
  officeLight  = ha.AddLight("light.office")
  officeButton = ha.AddBinarySensor("sensor.office_button")
)
```
Then, you create your automations by adding instances of the `Automation` struct:
```go
var testAutomation = homeautomation.Automation{
  Trigger: homeautomation.Trigger{
    Entity: officeButton,
    State: "on"
  },
  Condition: func() bool {},
  Action: func() error {},
}
```
Within the `Condition` and `Action` functions you can do anything you would normally do in Go, including performing actions on the previously defined device variables.
