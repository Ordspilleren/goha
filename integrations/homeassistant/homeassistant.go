package homeassistant

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/Ordspilleren/goha"
	"github.com/Ordspilleren/goha/wsclient"
)

type HomeAssistant struct {
	wsClient      wsclient.Client
	HAEndpoint    string
	HAAccessToken string
	Entities      []goha.Entity
}

var interactionId int

func InteractionID() int {
	interactionId += 1
	return interactionId
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

func (ha *HomeAssistant) RegisterEntities(entities ...goha.Entity) error {
	ha.Entities = append(ha.Entities, entities...)

	return nil
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
			entityIds = append(entityIds, entity.EntityID())
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
			if haState, ok := message.Event.EventAdd[ha.Entities[entityIndex].EntityID()]; ok {
				state := goha.State{}
				mapState(&state, &haState)
				ha.Entities[entityIndex].SetState(state)
				log.Printf("%s added! state changed to %s", ha.Entities[entityIndex].EntityID(), ha.Entities[entityIndex].State().State)
			}
			if haState, ok := message.Event.EventChange[ha.Entities[entityIndex].EntityID()]; ok {
				state := ha.Entities[entityIndex].State()
				mapState(&state, &haState.Additions)
				ha.Entities[entityIndex].SetState(state)
				log.Printf("%s changed! previous state: %s, current state: %s", ha.Entities[entityIndex].EntityID(), ha.Entities[entityIndex].PreviousState().State, ha.Entities[entityIndex].State().State)
				for automationIndex := range ha.Entities[entityIndex].Automations() {
					go ha.Entities[entityIndex].Automations()[automationIndex].Evaluate(ha.Entities[entityIndex])
				}
			}
		}
	}
}

func mergef[T comparable](a, b *T) {
	if b != nil {
		*a = *b
	}
}

func mapState(state *goha.State, haState *State) {
	mergef(&state.LastChanged, haState.LastChanged)
	mergef(&state.LastUpdated, haState.LastUpdated)
	mergef(&state.State, haState.State)

	if len(haState.Attributes.RgbColor) > 0 {
		state.Attributes.RgbColor = haState.Attributes.RgbColor
	}

	mergef(&state.Attributes.ColorTemp, haState.Attributes.ColorTemp)
	mergef(&state.Attributes.SupportedFeatures, haState.Attributes.SupportedFeatures)

	if len(haState.Attributes.XyColor) > 0 {
		state.Attributes.XyColor = haState.Attributes.XyColor
	}

	mergef(&state.Attributes.Brightness, haState.Attributes.Brightness)
	mergef(&state.Attributes.BrightnessPct, haState.Attributes.BrightnessPct)
	mergef(&state.Attributes.WhiteValue, haState.Attributes.WhiteValue)
	mergef(&state.Attributes.NextDawn, haState.Attributes.NextDawn)
	mergef(&state.Attributes.NextDusk, haState.Attributes.NextDusk)
	mergef(&state.Attributes.NextMidnight, haState.Attributes.NextMidnight)
	mergef(&state.Attributes.NextNoon, haState.Attributes.NextNoon)
	mergef(&state.Attributes.NextRising, haState.Attributes.NextRising)
	mergef(&state.Attributes.NextSetting, haState.Attributes.NextSetting)
	mergef(&state.Attributes.Elevation, haState.Attributes.Elevation)
	mergef(&state.Attributes.Azimuth, haState.Attributes.Azimuth)
	mergef(&state.Attributes.Rising, haState.Attributes.Rising)
	mergef(&state.Attributes.FriendlyName, haState.Attributes.FriendlyName)
	mergef(&state.Attributes.Source, haState.Attributes.Source)
	mergef(&state.Attributes.Transition, haState.Attributes.Transition)
}

func (ha *HomeAssistant) SendCommand(entity goha.Entity, action string, data any) error {
	var domain string
	switch t := entity.(type) {
	case *goha.Light:
		domain = "light"
	case *goha.Vacuum:
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
			EntityID: entity.EntityID(),
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

// WIP
func (ha *HomeAssistant) SendNotification(device string, title string, body string) error {
	message := Message{
		ID:      InteractionID(),
		Type:    "call_service",
		Domain:  "notify",
		Service: device,
		ServiceData: Notification{
			Title:   title,
			Message: body,
			Data: NotificationData{
				Actions: []NotificationAction{
					{
						Action: "URI",
						Title:  "Open Link",
						URI:    "https://google.dk",
					},
				},
			},
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
