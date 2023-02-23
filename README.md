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
var testAutomation = goha.Automation{
	Condition: goha.DefaultCondition,
	Action: func(e goha.Entity) error {
		if officeButton.Triggered() {
			officeLight.TurnOn()
		} else {
			officeLight.TurnOff()
		}
		return nil
	},
}
```
Within the `Condition` and `Action` functions you can do anything you would normally do in Go, including performing actions on the previously defined device variables.

### Unit tests for automations
One of the primary benefits of writing automations as code is that we get the ability to properly test them.

In the [`example/`](./example) folder you can find an example of how to do this.

### Running in Docker
Since Go programs compile into a single binary, the automations can be run on any machine by just copying it and executing. However, if you prefer to run it as a Docker container, this can faily easily be achieved.

The below `docker-compose` example runs a Go binary from a mounted volume, and as such it can be easily replaced without rebuilding the image:

```yaml
homeautomation:
  container_name: homeautomation
  image: gcr.io/distroless/static-debian11
  volumes:
    - /srv/appdata/homeautomation:/app
  entrypoint: ["/app/homeautomation"]
  restart: unless-stopped
```

Since the above example just runs a binary from a volume, we can very easily write a script to replace it when new automations have been added:

```bash
#!/bin/bash
CGO_ENABLED=0 go build
ssh server 'rm /srv/appdata/homeautomation/homeautomation'
scp homeautomation server:/srv/appdata/homeautomation
```
