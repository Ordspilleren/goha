package goha

import "time"

type Sensor struct {
	HAEntity
}

func (s *Sensor) On() bool {
	return s.GetState().String() == "on"
}

func (s *Sensor) Off() bool {
	return s.GetState().String() == "off"
}

func (s *Sensor) Timestamp() time.Time {
	time, _ := time.Parse(time.RFC3339, s.GetState().String())
	return time
}
