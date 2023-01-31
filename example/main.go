package main

import (
	"log"

	goha "github.com/Ordspilleren/goha"
)

var ha = goha.New("ws://ha_ip:8123/api/websocket", "access_token")

func main() {
	ha.RegisterAutomations(Automations...)
	err := ha.Start()
	if err != nil {
		log.Panic(err)
	}
}
