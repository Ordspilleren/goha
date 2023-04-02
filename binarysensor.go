package goha

type BinarySensor struct {
	HAEntity
}

func (b *BinarySensor) Triggered() bool {
	if b.GetState().State == "on" {
		return true
	} else {
		return false
	}
}
