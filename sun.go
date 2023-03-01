package goha

import "time"

type Sun struct {
	HAEntity
}

func (s *Sun) AboveHorizon() bool {
	if s.GetState().State.OrZero() == "above_horizon" {
		return true
	} else {
		return false
	}
}

func (s *Sun) BelowHorizon() bool {
	if s.GetState().State.OrZero() == "below_horizon" {
		return true
	} else {
		return false
	}
}

func (s *Sun) NextSetting() time.Time {
	return s.GetState().Attributes.NextSetting.OrZero()
}

func (s *Sun) NextRising() time.Time {
	return s.GetState().Attributes.NextRising.OrZero()
}

func (s *Sun) Elevation() float64 {
	return s.GetState().Attributes.Elevation.OrZero()
}
