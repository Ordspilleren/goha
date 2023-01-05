package main

import (
	"encoding/json"
	"log"

	"github.com/Ordspilleren/ha-automations/wsclient"
)

var wsClient *wsclient.Client

var Devices = Entities{
	"light.bed_light":                 &Light{},
	"binary_sensor.movement_backyard": &BinarySensor{},
}

func main() {
	wsClient = wsclient.StartClient()
	wsClient.OnMessage(StateChanger)
}

func StateChanger(wsMessage []byte) {
	var message Message

	err := json.Unmarshal(wsMessage, &message)
	if err != nil {
		log.Print(err)
	}

	if message.Type == "auth_required" {
		log.Print("Authentication required. Authenticating...")
		auth := `
		{
			"type": "auth",
			"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI5NjAyZjNiZTA3NWM0NTkzYjRhMmU2NmFlNzBmOWE1MyIsImlhdCI6MTY3Mjg3ODk3NywiZXhwIjoxOTg4MjM4OTc3fQ.wuZeXOt42fcJjkVb2awZ7ZMRfnFyOIIOcb3uIqyriz8"
		  }
		`
		wsClient.SendCommand([]byte(auth))
	}

	if message.Type == "auth_ok" {
		log.Print("Authentication OK. Subscribing to events.")
		subscribePayload := `
		{
			"id": 18,
			"type": "subscribe_events",
			"event_type": "state_changed"
		  }
		`
		wsClient.SendCommand([]byte(subscribePayload))
	}

	if message.Type == "event" {
		if _, ok := Devices[message.Event.Data.EntityID]; ok {
			Devices[message.Event.Data.EntityID].SetState(message.Event.Data.NewState)
			log.Print(Devices)
			for _, automation := range Automations {
				go automation.Evaluate(message.Event.Data.EntityID, message.Event.Data.NewState.State)
			}
		}
	}
}
