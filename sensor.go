package goha

import "time"

type Sensor struct {
	HAEntity
}

func (s *Sensor) On() bool {
	return s.GetState().State == "on"
}

func (s *Sensor) Off() bool {
	return s.GetState().State == "off"
}

func (s *Sensor) Timestamp() time.Time {
	time, _ := time.Parse(time.RFC3339, s.GetState().State)
	return time
}
