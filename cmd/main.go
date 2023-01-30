package main

import (
	"log"

	homeautomation "github.com/Ordspilleren/homeautomation"
)

func main() {
	err := homeautomation.Start(Entities, Automations)
	if err != nil {
		log.Panic(err)
	}
}
