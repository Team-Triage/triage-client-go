package main

import (
	"context"
	"log"
	"net"
	// "status"
	// "codes"

	pb "example.com/triage-grpc/triage_main" // import protobuf module
	"google.golang.org/grpc"
	// "google.golang.org/grpc/status"
)

const (
	port = ":9001"
)

type MessageHandlerServer struct { // to register this type with grpc, embed unimpl. server inside the type
	pb.UnimplementedMessageHandlerServer
}

func (s *MessageHandlerServer) GetMessage(ctx context.Context, in *pb.Message) (*pb.MessageResponse, error) {
// 	if err != nil {
//     errStatus := status.Convert(err)
//     log.Printf("SayHello return error: code: %d, msg: %s\n", errStatus.Code(), errStatus.Message())
// }

// 	s, ok := status.New()
	statusCode := "available"
	log.Printf("Received: %v, Status: %s", in.GetBody(), statusCode)
	return &pb.MessageResponse{Body: in.GetBody(), Status: statusCode}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // init new server
	pb.RegisterMessageHandlerServer(s, &MessageHandlerServer{}) // register server as a new gRPC service!
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// s, ok := status.FromError(err)
	// fmt.Println(s, ok)
}