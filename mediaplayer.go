package goha

type MediaPlayer struct {
	HAEntity
}

func (m *MediaPlayer) Playing() bool {
	return m.GetState().String() == "playing"
}

func (m *MediaPlayer) Paused() bool {
	return m.GetState().String() == "paused"
}

func (m *MediaPlayer) Stopped() bool {
	return m.GetState().String() == "off" || m.GetState().String() == "idle" || m.GetState().String() == "standby"
}
