package main

type Entity interface {
	GetEntityID() string
	GetState() State
	SetState(State)
}

var Entities map[string]Entity
