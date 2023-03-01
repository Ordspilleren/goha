package goha

import "time"

type Sun struct {
	HAEntity
}

func (s *Sun) AboveHorizon() bool {
	if s.GetState().State == "above_horizon" {
		return true
	} else {
		return false
	}
}

func (s *Sun) BelowHorizon() bool {
	if s.GetState().State == "below_horizon" {
		return true
	} else {
		return false
	}
}

func (s *Sun) NextSetting() time.Time {
	return OrZero(s.GetState().Attributes.NextSetting)
}

func (s *Sun) NextRising() time.Time {
	return OrZero(s.GetState().Attributes.NextRising)
}

func (s *Sun) Elevation() float64 {
	return OrZero(s.GetState().Attributes.Elevation)
}
