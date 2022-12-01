
# triage-client-go

Thin Go client for subscribing to the Triage service for Kafka.

<h2>How it works</h2>

The `triage-client-go` library interacts with a Triage deployment by running a gRPC server that listens for a incoming messages (gRPC method calls). Users pass a message handler function to the library, which acts as middleware. Each message received is passed to the message handler, and `triage-client-go` uses the return of the message handler to determine the response sent back to Triage.

<h2>How to get it</h2>

To get the latest version, use `go 1.19*` and fetch using the `go get` command. For example:

    go get github.com/team-triage/triage-client-go

<h2>How to use it</h2>
This client provides 3 utility functions used for interacting with Triage.

<h4>OnMessage</h4>

The `OnMessage(messageHandler func(string) int)` function accepts a `messageHandler` function. This function should accept a string (the value of the Kafka message) and return either integers `1` or `-1` , representing the successful or unsuccessful processing of a message.


The function passed to `OnMessage` acts as middleware - when the gRPC server receives a message from Triage, it will call the `messageHandler` passed to it to determine the response sent back to Triage.


A `1` sends a positive acknowledgment to Triage, letting it know that it is safe to commit this message's offset back to Kafka.


A `-1` sends a negative acknowledgment to Triage. Triage will store this message in a DynamoDB table for later processing, before committing it's offset back to Kafka.



<h4>Listen</h4>

The `Listen(grpcPort  string)` function starts a gRPC server on the specified port. This port should be the same that is sent to Triage by `RequestConnection`. This should be called before `RequestConnection` (see below).


<h4>RequestConnection</h4>

The `RequestConnection(triageNetworkAddress string, grpcPort string, authToken string)` function makes an initial request to Triage, informing it that the consumer application is ready to begin accepting connections. It accepts 3 arguments:
  - `triageNetworkAddress`: The network address provided during Triage deployment. This is typically an AWS internet-facing application load balancer address.
  - `grpcPort`: The port used by the consumer application to listen for incoming messages.
  - `authToken`: The authentication key provided during Triage deployment.


<h2>Example Consumer Application</h2>

```go
package main

import (
  client "github.com/team-triage/triage-client-go"
  "sync"
)

var wg sync.WaitGroup

func main() {
  grpcPort := "9000"
  triageNetworkAddress := "Triag-Triage-123456.us-west-1.elb.amazonaws.com"
  authKey := "ABADAUTHKEY"

  client.OnMessage(dummyMessageHandler)
  
  wg.Add(1)
  client.Listen(grpcPort)
  
  res := client.RequestConnection(triageNetworkAddress, grpcPort, authKey)
  
  if res.StatusCode == http.StatusOK {
		fmt.Printf("Status: %v \n ready to receive messages!\n", res.StatusCode)
	} else {
		fmt.Printf("Status Code: %v\n There was a problem connecting to Triage :(\n", res.StatusCode)
	}
  
  wg.Wait()
}

func dummyMessageHandler(msg string) int {
  if len(msg) > 4 {
    return 1
  }
  return -1
}

```

NEED TO TEAM's GH PROFILES
