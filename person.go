package goha

type Person struct {
	HAEntity
}

func (p *Person) IsHome() bool {
	if p.GetState().State.OrZero() == "home" {
		return true
	} else {
		return false
	}
}
