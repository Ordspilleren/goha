package main

import (
	"log"
	"sync"

	goha "github.com/Ordspilleren/goha"
)

var wg sync.WaitGroup
var ha = goha.New("ws://ha_ip:8123/api/websocket", "access_token")

func main() {
	SetupAutomations()
	err := ha.Start(&wg)
	if err != nil {
		log.Panic(err)
	}

	wg.Wait()
}
