package main

import (
	"log"
	"sync"

	goha "github.com/Ordspilleren/goha"
	"github.com/Ordspilleren/goha/integrations/homeassistant"
)

var wg sync.WaitGroup
var homeautomation = goha.Goha{
	PrimaryIntegration: homeassistant.New("ws://ha_ip:8123/api/websocket", "access_token"),
}

func main() {
	SetupAutomations()
	err := homeautomation.Start(&wg)
	if err != nil {
		log.Panic(err)
	}

	wg.Wait()
}
