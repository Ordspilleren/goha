package goha

import (
	"encoding/json"
	"log"

	"github.com/Ordspilleren/goha/wsclient"
)

var wsClient *wsclient.Client

var interactionId int

func InteractionID() int {
	interactionId += 1
	return interactionId
}

type HomeAutomation struct {
	client        *wsclient.Client
	HAEndpoint    string
	HAAccessToken string
	Entities      []Entity
	Automations   []Automation
}

func New(endpoint string, accessToken string) *HomeAutomation {
	return &HomeAutomation{
		HAEndpoint:    endpoint,
		HAAccessToken: accessToken,
	}
}

func (ha *HomeAutomation) Start() error {
	wsClient = wsclient.StartClient(ha.HAEndpoint)
	ha.client = wsClient
	ha.client.OnMessage(ha.stateChanger)

	return nil
}

func (ha *HomeAutomation) RegisterAutomations(automations ...Automation) error {
	ha.Automations = append(ha.Automations, automations...)

	return nil
}

func (ha *HomeAutomation) RegisterEntities(entities ...Entity) error {
	ha.Entities = append(ha.Entities, entities...)

	return nil
}

func (ha *HomeAutomation) stateChanger(wsMessage []byte) {
	var message Message

	err := json.Unmarshal(wsMessage, &message)
	if err != nil {
		log.Print(err)
	}

	if message.Type == "auth_required" {
		log.Print("Authentication required. Authenticating...")
		auth := Message{
			Type:        "auth",
			AccessToken: ha.HAAccessToken,
		}
		payload, err := json.Marshal(auth)
		if err != nil {
			log.Panic(err)
		}
		ha.client.SendCommand(payload)
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
		ha.client.SendCommand(payload)
	}

	if message.Type == "event" {
		for entityIndex := range ha.Entities {
			if ha.Entities[entityIndex].GetEntityID() == message.Event.Data.EntityID {
				ha.Entities[entityIndex].SetState(message.Event.Data.NewState)
				log.Print(ha.Entities)
				for automationIndex := range ha.Automations {
					go ha.Automations[automationIndex].Evaluate(message.Event.Data.EntityID, message.Event.Data.NewState.State)
				}
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
