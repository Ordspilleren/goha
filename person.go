package goha

type Person struct {
	HAEntity
}

func (p *Person) IsHome() bool {
	return p.State().State == "home"
}
