package goha

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Ordspilleren/goha/wsclient"
)

var interactionId int

func InteractionID() int {
	interactionId += 1
	return interactionId
}

type HomeAutomation struct {
	wsClient      wsclient.Client
	HAEndpoint    string
	HAAccessToken string
	Entities      []Entity
}

func New(endpoint string, accessToken string) *HomeAutomation {
	return &HomeAutomation{
		HAEndpoint:    endpoint,
		HAAccessToken: accessToken,
	}
}

func (ha *HomeAutomation) Start(waitGroup *sync.WaitGroup) error {
	ha.wsClient = *wsclient.New(ha.HAEndpoint, ha.stateChanger)
	ha.wsClient.Start()
	ha.sendAuth()

	waitGroup.Add(1)

	return nil
}

func (ha *HomeAutomation) RegisterEntities(entities ...Entity) error {
	ha.Entities = append(ha.Entities, entities...)

	return nil
}

func (ha *HomeAutomation) AddEntity(entity Entity, entityId string) Entity {
	entity.SetClient(&ha.wsClient)
	entity.SetEntityID(entityId)
	ha.RegisterEntities(entity)

	return entity
}

func (ha *HomeAutomation) sendAuth() {
	auth := Message{
		Type:        "auth",
		AccessToken: ha.HAAccessToken,
	}
	payload, err := json.Marshal(auth)
	if err != nil {
		log.Panic(err)
	}
	ha.wsClient.SendCommand(payload)
}

func (ha *HomeAutomation) stateChanger(wsMessage []byte) {
	var message Message

	err := json.Unmarshal(wsMessage, &message)
	if err != nil {
		log.Print(err)
	}

	if message.Type == "auth_required" {
		log.Print("Authentication required. Authenticating...")
		ha.sendAuth()
	}

	if message.Type == "auth_ok" {
		log.Print("Authentication OK. Subscribing to events.")
		var entityIds []string
		for _, entity := range ha.Entities {
			entityIds = append(entityIds, entity.GetEntityID())
		}
		subscribe := Message{
			ID:        InteractionID(),
			Type:      "subscribe_entities",
			EntityIDs: entityIds,
		}
		payload, err := json.Marshal(subscribe)
		if err != nil {
			log.Panic(err)
		}
		ha.wsClient.SendCommand(payload)
	}

	if message.Type == "event" {
		log.Print(string(wsMessage))
		for entityIndex := range ha.Entities {
			if state, ok := message.Event.EventAdd[ha.Entities[entityIndex].GetEntityID()]; ok {
				ha.Entities[entityIndex].SetState(state)
				log.Print("device added, state changed")
			}
			if state, ok := message.Event.EventChange[ha.Entities[entityIndex].GetEntityID()]; ok {
				ha.Entities[entityIndex].SetState(state.Additions)
				for automationIndex := range ha.Entities[entityIndex].GetAutomations() {
					go ha.Entities[entityIndex].GetAutomations()[automationIndex].Evaluate(ha.Entities[entityIndex])
				}
			}
		}
	}
}

func (ha *HAEntity) CallService(domain string, service string, serviceData any, targetEntityID string) {
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

	ha.wsClient.SendCommand(payload)
}
