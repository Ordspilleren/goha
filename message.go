package main

import "time"

type Message struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	HaVersion string `json:"ha_version"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Event     Event  `json:"event"`
	Error     Error  `json:"error"`
}
type Attributes struct {
	RgbColor          []int     `json:"rgb_color"`
	ColorTemp         int       `json:"color_temp"`
	SupportedFeatures int       `json:"supported_features"`
	XyColor           []float64 `json:"xy_color"`
	Brightness        int       `json:"brightness"`
	WhiteValue        int       `json:"white_value"`
	FriendlyName      string    `json:"friendly_name"`
}
type Context struct {
	ID       string      `json:"id"`
	ParentID interface{} `json:"parent_id"`
	UserID   string      `json:"user_id"`
}
type State struct {
	EntityID    string     `json:"entity_id"`
	LastChanged time.Time  `json:"last_changed"`
	State       string     `json:"state"`
	Attributes  Attributes `json:"attributes"`
	LastUpdated time.Time  `json:"last_updated"`
	Context     Context    `json:"context"`
}
type Data struct {
	EntityID string `json:"entity_id"`
	NewState State  `json:"new_state"`
	OldState State  `json:"old_state"`
}
type Event struct {
	Data      Data      `json:"data"`
	EventType string    `json:"event_type"`
	TimeFired time.Time `json:"time_fired"`
	Origin    string    `json:"origin"`
	Context   Context   `json:"context"`
}
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
