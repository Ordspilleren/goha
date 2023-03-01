package goha

type Sensor struct {
	HAEntity
}

func (b *Sensor) On() bool {
	return b.GetState().State.OrZero() == "on"
}

func (b *Sensor) Off() bool {
	return b.GetState().State.OrZero() == "off"
}
