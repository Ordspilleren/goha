package goha

type MediaPlayer struct {
	HAEntity
}

func (m *MediaPlayer) Playing() bool {
	return m.State().State == "playing"
}

func (m *MediaPlayer) Paused() bool {
	return m.State().State == "paused"
}

func (m *MediaPlayer) Stopped() bool {
	return m.State().State == "off" || m.State().State == "idle" || m.State().State == "standby"
}
