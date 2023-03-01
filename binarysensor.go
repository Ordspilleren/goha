package goha

type BinarySensor struct {
	HAEntity
}

func (b *BinarySensor) Triggered() bool {
	if b.GetState().State.OrZero() == "on" {
		return true
	} else {
		return false
	}
}
