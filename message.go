package goha

import (
	"time"

	"github.com/Southclaws/opt"
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
	RgbColor          []opt.Optional[int]     `json:"rgb_color,omitempty"`
	ColorTemp         opt.Optional[int]       `json:"color_temp,omitempty"`
	SupportedFeatures opt.Optional[int]       `json:"supported_features,omitempty"`
	XyColor           []opt.Optional[float64] `json:"xy_color,omitempty"`
	Transition        opt.Optional[int]       `json:"transition,omitempty"`
	Brightness        opt.Optional[int]       `json:"brightness,omitempty"`
	WhiteValue        opt.Optional[int]       `json:"white_value,omitempty"`
	NextDawn          opt.Optional[time.Time] `json:"next_dawn,omitempty"`
	NextDusk          opt.Optional[time.Time] `json:"next_dusk,omitempty"`
	NextMidnight      opt.Optional[time.Time] `json:"next_midnight,omitempty"`
	NextNoon          opt.Optional[time.Time] `json:"next_noon,omitempty"`
	NextRising        opt.Optional[time.Time] `json:"next_rising,omitempty"`
	NextSetting       opt.Optional[time.Time] `json:"next_setting,omitempty"`
	Elevation         opt.Optional[float64]   `json:"elevation,omitempty"`
	Azimuth           opt.Optional[float64]   `json:"azimuth,omitempty"`
	Rising            opt.Optional[bool]      `json:"rising,omitempty"`
	FriendlyName      opt.Optional[string]    `json:"friendly_name,omitempty"`
	Source            opt.Optional[string]    `json:"source,omitempty"`
}
type Context struct {
	ID       string      `json:"id,omitempty"`
	ParentID interface{} `json:"parent_id,omitempty"`
	UserID   string      `json:"user_id,omitempty"`
}
type State struct {
	LastChanged opt.Optional[float64] `json:"lc,omitempty"`
	LastUpdated opt.Optional[float64] `json:"lu,omitempty"`
	State       opt.Optional[string]  `json:"s,omitempty"`
	Attributes  Attributes            `json:"a,omitempty"`
	Context     any                   `json:"c,omitempty"`
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
