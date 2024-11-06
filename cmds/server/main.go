package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/AndrewShear/proto_example/pb/protos/greeter"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = "8083"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Infof("Received: %+v", in.GetName())
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s", in.GetName()),
	}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	lis, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		logrus.Fatalf("failed to listen on port %s", port)
	}
	logrus.Info("server listening on ", port)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
