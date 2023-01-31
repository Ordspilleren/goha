package main

import (
	"log"

	goha "github.com/Ordspilleren/goha"
)

var ha = goha.New("ws://10.3.3.32:8123/api/websocket", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI5MGRkODExYzY2ZmE0Yjk3YjUzNjJjNjU3YjBkY2Q5NiIsImlhdCI6MTY3NTE3NTY5NSwiZXhwIjoxOTkwNTM1Njk1fQ.DpL3lTAWwrYvFNt85ipeDOtDKmQShualrLRjToCl-P8")

func main() {
	ha.RegisterAutomations(Automations...)
	err := ha.Start()
	if err != nil {
		log.Panic(err)
	}
}
