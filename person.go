package goha

type Person struct {
	HAEntity
}

func (p *Person) IsHome() bool {
	if p.State().State == "home" {
		return true
	} else {
		return false
	}
}
