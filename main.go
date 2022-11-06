package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Team-Triage/triage-client-go/grpcServer/server"
)

// func main() {
// 	OnMessage(messageHandler)
// 	Listen()
// }

func messageHandler(message string) int {
	fmt.Println(message)
	if len(message) > 4 {
		return -1
	}
	return 1
}

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
