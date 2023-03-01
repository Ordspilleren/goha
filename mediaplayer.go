package goha

type MediaPlayer struct {
	HAEntity
}

func (m *MediaPlayer) Playing() bool {
	return m.GetState().State == "playing"
}

func (m *MediaPlayer) Paused() bool {
	return m.GetState().State == "paused"
}

func (m *MediaPlayer) Stopped() bool {
	return m.GetState().State == "off" || m.GetState().State == "idle" || m.GetState().State == "standby"
}
