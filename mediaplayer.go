package goha

type MediaPlayer struct {
	HAEntity
}

func (m *MediaPlayer) Playing() bool {
	return m.GetState().State.OrZero() == "playing"
}

func (m *MediaPlayer) Paused() bool {
	return m.GetState().State.OrZero() == "paused"
}
