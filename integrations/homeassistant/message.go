package homeassistant

import (
	"time"
)

type Message struct {
	ID          int      `json:"id,omitempty"`
	Type        string   `json:"type,omitempty"`
	AccessToken string   `json:"access_token,omitempty"`
	EntityIDs   []string `json:"entity_ids,omitempty"`
	EventType   string   `json:"event_type,omitempty"`
	HaVersion   string   `json:"ha_version,omitempty"`
	Message     string   `json:"message,omitempty"`
	Success     bool     `json:"success,omitempty"`
	Domain      string   `json:"domain,omitempty"`
	Service     string   `json:"service,omitempty"`
	ServiceData any      `json:"service_data,omitempty"`
	Target      *Target  `json:"target,omitempty"`
	Event       *Event   `json:"event,omitempty"`
	Error       *Error   `json:"error,omitempty"`
}
type Attributes struct {
	RgbColor          []int      `json:"rgb_color,omitempty"`
	ColorTemp         *int       `json:"color_temp,omitempty"`
	SupportedFeatures *int       `json:"supported_features,omitempty"`
	XyColor           []float64  `json:"xy_color,omitempty"`
	Transition        *int       `json:"transition,omitempty"`
	Brightness        *int       `json:"brightness,omitempty"`
	BrightnessPct     *int       `json:"brightness_pct,omitempty"`
	WhiteValue        *int       `json:"white_value,omitempty"`
	NextDawn          *time.Time `json:"next_dawn,omitempty"`
	NextDusk          *time.Time `json:"next_dusk,omitempty"`
	NextMidnight      *time.Time `json:"next_midnight,omitempty"`
	NextNoon          *time.Time `json:"next_noon,omitempty"`
	NextRising        *time.Time `json:"next_rising,omitempty"`
	NextSetting       *time.Time `json:"next_setting,omitempty"`
	Elevation         *float64   `json:"elevation,omitempty"`
	Azimuth           *float64   `json:"azimuth,omitempty"`
	Rising            *bool      `json:"rising,omitempty"`
	FriendlyName      *string    `json:"friendly_name,omitempty"`
	Source            *string    `json:"source,omitempty"`
}
type Context struct {
	ID       string      `json:"id,omitempty"`
	ParentID interface{} `json:"parent_id,omitempty"`
	UserID   string      `json:"user_id,omitempty"`
}
type State struct {
	LastChanged *float64   `json:"lc,omitempty"`
	LastUpdated *float64   `json:"lu,omitempty"`
	State       *string    `json:"s,omitempty"`
	Attributes  Attributes `json:"a,omitempty"`
	Context     any        `json:"c,omitempty"`
}
type Diff struct {
	Additions State         `json:"+,omitempty"`
	Removals  StateRemovals `json:"-,omitempty"`
}
type StateRemovals struct {
	LastChanged float64  `json:"lc,omitempty"`
	State       string   `json:"s,omitempty"`
	Attributes  []string `json:"a,omitempty"`
	LastUpdated float64  `json:"lu,omitempty"`
	Context     any      `json:"c,omitempty"`
}
type Data struct {
	EntityID string `json:"entity_id,omitempty"`
	NewState State  `json:"new_state,omitempty"`
	OldState State  `json:"old_state,omitempty"`
}
type Event struct {
	EventAdd    map[string]State `json:"a,omitempty"`
	EventRemove []string         `json:"r,omitempty"`
	EventChange map[string]Diff  `json:"c,omitempty"`
}
type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
type Target struct {
	EntityID string `json:"entity_id,omitempty"`
}
type Notification struct {
	Title   string           `json:"title,omitempty"`
	Message string           `json:"message,omitempty"`
	Data    NotificationData `json:"data,omitempty"`
}
type NotificationData struct {
	Actions []NotificationAction `json:"actions,omitempty"`
}
type NotificationAction struct {
	Action string `json:"action,omitempty"`
	Title  string `json:"title,omitempty"`
	URI    string `json:"uri,omitempty"`
}
