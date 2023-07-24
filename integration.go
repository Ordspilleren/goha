package goha

import "sync"

type Integration interface {
	Start(*sync.WaitGroup) error
	SendCommand(Entity, string, any) error
	RegisterEntities(...Entity) error
	RegisterAutomations(...Automation) error
}
