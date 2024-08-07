package goha

import "time"

type Sun struct {
	HAEntity
}

func (s *Sun) AboveHorizon() bool {
	return s.State().State == "above_horizon"
}

func (s *Sun) BelowHorizon() bool {
	return s.State().State == "below_horizon"
}

func (s *Sun) NextSetting() time.Time {
	return s.State().Attributes.NextSetting
}

func (s *Sun) NextRising() time.Time {
	return s.State().Attributes.NextRising
}

func (s *Sun) Elevation() float64 {
	return s.State().Attributes.Elevation
}
