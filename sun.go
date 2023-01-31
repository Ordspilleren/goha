package goha

import "time"

type Sun struct {
	EntityData
}

func (s *Sun) AboveHorizon() bool {
	if s.GetState().State == "above_horizon" {
		return true
	} else {
		return false
	}
}

func (s *Sun) NextSetting() time.Time {
	return s.GetState().Attributes.NextSetting
}

func (s *Sun) NextRising() time.Time {
	return s.GetState().Attributes.NextRising
}
