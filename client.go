package client

import (
	"github.com/team-triage/triage-client-go/grpcServer/server"
)

func OnMessage(messageHandler func(string) int) {
	server.OnMessage(messageHandler)
}

func Listen(grpcPort string) {
	server.StartServer(grpcPort)
}
