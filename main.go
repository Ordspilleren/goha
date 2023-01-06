package main

import (
	"encoding/json"
	"log"

	"github.com/Ordspilleren/ha-automations/wsclient"
)

var wsClient *wsclient.Client

var interactionId int

func InteractionID() int {
	interactionId += 1
	return interactionId
}

var Devices = Entities{}

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
		auth := Message{
			Type:        "auth",
			AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI5NjAyZjNiZTA3NWM0NTkzYjRhMmU2NmFlNzBmOWE1MyIsImlhdCI6MTY3Mjg3ODk3NywiZXhwIjoxOTg4MjM4OTc3fQ.wuZeXOt42fcJjkVb2awZ7ZMRfnFyOIIOcb3uIqyriz8",
		}
		payload, err := json.Marshal(auth)
		if err != nil {
			log.Panic(err)
		}
		wsClient.SendCommand(payload)
	}

	if message.Type == "auth_ok" {
		log.Print("Authentication OK. Subscribing to events.")
		subscribe := Message{
			ID:        InteractionID(),
			Type:      "subscribe_events",
			EventType: "state_changed",
		}
		payload, err := json.Marshal(subscribe)
		if err != nil {
			log.Panic(err)
		}
		wsClient.SendCommand(payload)
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

func CallService(wsClient *wsclient.Client, domain string, service string, serviceData any, targetEntityID string) {
	message := Message{
		ID:          InteractionID(),
		Type:        "call_service",
		Domain:      domain,
		Service:     service,
		ServiceData: serviceData,
		Target: &Target{
			EntityID: targetEntityID,
		},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Panic(err)
	}

	wsClient.SendCommand(payload)
}
