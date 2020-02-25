package main

import (
	pb "platform-sample-productowner/srv"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayName (ctx context.Context, in *pb.ProductOwnerRequest) (*pb.OwnerMessage, error) {
	return &pb.OwnerMessage {
		Message: "Hello, " + in.Name + ", you're our productowner!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductOwnerServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
