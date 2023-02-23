package goha

type Integration interface {
	SendCommand(Entity, string) error
}
