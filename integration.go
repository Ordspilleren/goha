package goha

type Integration interface {
	SendCommand(Entity, string, any) error
}
