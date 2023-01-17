package main

import (
	"log"

	haautomations "github.com/Ordspilleren/ha-automations"
)

func main() {
	err := haautomations.Start(Entities, Automations)
	if err != nil {
		log.Panic(err)
	}
}
