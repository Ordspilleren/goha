package main

type Entity interface {
	GetEntityID() string
	GetState() State
	SetState(State)
}

type Entities map[string]Entity

var Devices Entities
