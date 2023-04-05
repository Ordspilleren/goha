package goha

import "time"

type Sensor struct {
	HAEntity
}

func (s *Sensor) On() bool {
	return s.State().State == "on"
}

func (s *Sensor) Off() bool {
	return s.State().State == "off"
}

func (s *Sensor) Timestamp() time.Time {
	time, _ := time.Parse(time.RFC3339, s.State().State)
	return time
}
