package main

import (
	"context"
	"net"
	"time"

	pb "github.com/AndrewShear/proto_example/pb/protos/greeter"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8083"
)

func main() {
	conn, err := grpc.NewClient(net.JoinHostPort("", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Andrew"})
	if err != nil {
		logrus.Fatalf("could not greet: %v", err)
	}

	logrus.Infof("Greetings: %s", r.GetMessage())
}
