package goha

type Vacuum struct {
	HAEntity
}

func (v *Vacuum) Docked() bool {
	return v.State().State == "docked"
}

func (v *Vacuum) StartCleaning() error {
	v.integration.SendCommand(v, "start", nil)
	return nil
}
