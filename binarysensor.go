package goha

type BinarySensor struct {
	HAEntity
}

func (b *BinarySensor) Triggered() bool {
	return b.State().State == "on"
}
