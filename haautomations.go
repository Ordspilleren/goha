package haautomations

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

var Entities = EntityList{}
var Automations = AutomationList{}

func Start(entities EntityList, automations AutomationList) error {
	Entities = entities
	Automations = automations

	wsClient = wsclient.StartClient()
	wsClient.OnMessage(stateChanger)

	return nil
}

func stateChanger(wsMessage []byte) {
	var message Message

	err := json.Unmarshal(wsMessage, &message)
	if err != nil {
		log.Print(err)
	}

	if message.Type == "auth_required" {
		log.Print("Authentication required. Authenticating...")
		auth := Message{
			Type:        "auth",
			AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjOGVlZGU5NmQxNTI0NzVhYmJiMzc5OGMxZjI0Y2VkZSIsImlhdCI6MTY3Mzk5Mzg5NSwiZXhwIjoxOTg5MzUzODk1fQ.8g4EP1p0k_8vHKRvv7BxLUkuNSjtsDRBsa4EqKjLH00",
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
		if _, ok := Entities[message.Event.Data.EntityID]; ok {
			Entities[message.Event.Data.EntityID].SetState(message.Event.Data.NewState)
			log.Print(Entities)
			for _, automation := range Automations {
				go automation.Evaluate(message.Event.Data.EntityID, message.Event.Data.NewState.State)
			}
		}
	}
}

func CallService(domain string, service string, serviceData any, targetEntityID string) {
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
