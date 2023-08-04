package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	hello "test/hello/proto"
)

type server struct{
	hello.UnimplementedHelloServiceServer
}

func (*server) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customerHello := "Welcome ! " + prefix + " " + firstName

	res := &hello.HelloResponse{
		CustomerHello: customerHello,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello, Go Server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	hello.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
