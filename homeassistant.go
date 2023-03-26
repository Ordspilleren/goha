package goha

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/Ordspilleren/goha/wsclient"
)

var interactionId int

func InteractionID() int {
	interactionId += 1
	return interactionId
}

type HomeAssistant struct {
	wsClient      wsclient.Client
	HAEndpoint    string
	HAAccessToken string
	Entities      []Entity
}

func New(endpoint string, accessToken string) *HomeAssistant {
	return &HomeAssistant{
		HAEndpoint:    endpoint,
		HAAccessToken: accessToken,
	}
}

func (ha *HomeAssistant) Start(waitGroup *sync.WaitGroup) error {
	ha.wsClient = *wsclient.New(ha.HAEndpoint, ha.stateChanger)
	ha.wsClient.Start()
	ha.sendAuth()

	waitGroup.Add(1)

	return nil
}

func (ha *HomeAssistant) RegisterEntities(entities ...Entity) error {
	ha.Entities = append(ha.Entities, entities...)

	return nil
}

func (ha *HomeAssistant) AddEntity(entity Entity, entityId string) Entity {
	entity.SetIntegration(ha)
	entity.SetEntityID(entityId)
	ha.RegisterEntities(entity)

	return entity
}

func (ha *HomeAssistant) AddLight(entityId string) *Light {
	return ha.AddEntity(&Light{}, entityId).(*Light)
}

func (ha *HomeAssistant) AddBinarySensor(entityId string) *BinarySensor {
	return ha.AddEntity(&BinarySensor{}, entityId).(*BinarySensor)
}

func (ha *HomeAssistant) AddSensor(entityId string) *Sensor {
	return ha.AddEntity(&Sensor{}, entityId).(*Sensor)
}

func (ha *HomeAssistant) AddPerson(entityId string) *Person {
	return ha.AddEntity(&Person{}, entityId).(*Person)
}

func (ha *HomeAssistant) AddSun(entityId string) *Sun {
	return ha.AddEntity(&Sun{}, entityId).(*Sun)
}

func (ha *HomeAssistant) AddMediaPlayer(entityId string) *MediaPlayer {
	return ha.AddEntity(&MediaPlayer{}, entityId).(*MediaPlayer)
}

func (ha *HomeAssistant) AddVacuum(entityId string) *Vacuum {
	return ha.AddEntity(&Vacuum{}, entityId).(*Vacuum)
}

func (ha *HomeAssistant) AddSchedule(entityId string) *Schedule {
	return ha.AddEntity(&Schedule{}, entityId).(*Schedule)
}

func (ha *HomeAssistant) sendAuth() {
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

func (ha *HomeAssistant) stateChanger(wsMessage []byte) {
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
				err := json.Unmarshal(state.Additions, ha.Entities[entityIndex].GetStatePtr())
				if err != nil {
					log.Printf("failed unmarshaling state change: %s", err)
				}
				for automationIndex := range ha.Entities[entityIndex].GetAutomations() {
					go ha.Entities[entityIndex].GetAutomations()[automationIndex].Evaluate(ha.Entities[entityIndex])
				}
			}
		}
	}
}

func (ha *HomeAssistant) SendCommand(entity Entity, action string, data any) error {
	var domain string
	switch t := entity.(type) {
	case *Light:
		domain = "light"
	case *Vacuum:
		domain = "vacuum"
	default:
		return fmt.Errorf("unknown entity type: %v", t)
	}

	message := Message{
		ID:          InteractionID(),
		Type:        "call_service",
		Domain:      domain,
		Service:     action,
		ServiceData: data,
		Target: &Target{
			EntityID: entity.GetEntityID(),
		},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Panic(err)
	}

	log.Print(string(payload))

	ha.wsClient.SendCommand(payload)
	return nil
}
