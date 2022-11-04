package main

import (
	"log"
	"net/http"

	"github.com/Team-Triage/triage-client-go/grpcServer/server"
)

func RequestConnection(triageNetworkAddress string) *http.Response {
	res, err := http.Get(triageNetworkAddress)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}

func OnMessage(messageHandler func(string) int) {
	server.OnMessage(messageHandler)
}

func Listen() {
	server.StartServer()
}
