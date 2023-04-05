package goha

type Schedule struct {
	HAEntity
}

func (s *Schedule) Active() bool {
	return s.State().State == "on"
}
