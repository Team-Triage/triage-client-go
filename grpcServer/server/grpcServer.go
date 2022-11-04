package server

import (
	"context"
	"log"
	"net"

	"github.com/Team-Triage/triage-client-go/grpcServer/pb" // import protobuf module

	"google.golang.org/grpc"
)

type MessageHandlerServer struct { // to register this type with grpc, embed unimpl. server inside the type
	pb.UnimplementedMessageHandlerServer
}

var messageProcessor func(string) int

func OnMessage(messageHandler func(string) int) {
	messageProcessor = messageHandler
}

func (s *MessageHandlerServer) SendMessage(ctx context.Context, in *pb.Message) (*pb.MessageResponse, error) {
	status := messageProcessor(in.GetBody())

	// log.Printf("Received: %v, Status: %s", in.GetBody(), status)
	return &pb.MessageResponse{Body: in.GetBody(), Status: int32(status)}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                                       // init new server
	pb.RegisterMessageHandlerServer(s, &MessageHandlerServer{}) // register server as a new gRPC service!
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
