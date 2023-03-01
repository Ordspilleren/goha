package goha

type BinarySensor struct {
	HAEntity
}

func (b *BinarySensor) Triggered() bool {
	if b.GetState().String() == "on" {
		return true
	} else {
		return false
	}
}
