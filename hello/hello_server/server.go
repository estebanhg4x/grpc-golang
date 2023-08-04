package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	"net"
	hello "test/hello/proto"
)

type Server struct{
	hello.UnimplementedHelloServiceServer;
}

func (*Server) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customerHello := "Welcome ! " + prefix + " " + firstName

	res := &hello.HelloResponse{
		CustomerHello: customerHello,
	}

	return res, nil
}

func (*Server) HelloManyLenguagues(req *hello.HelloManyLenguagesRequest,stream hello.HelloService_HelloManyLenguaguesServer)error{

	fmt.Println("Hello Many times fuction was invoked with %v\n",req)

	langs := [5]string{"Saludos", "Hello", "Aló", "Ni hao", "Hola"}

	firstName := req.GetHelloLenguages().GetFirstName()
	prefix := req.GetHelloLenguages().GetPrefix()

	for _, helloLang := range langs{
		helloLagauge := helloLang + prefix + " " + firstName

		res := &hello.HelloManyLenguagesResponse{
			HelloLenguages: helloLagauge,
		}

		stream.Send(res)

		time.Sleep(2000 + time.Millisecond) // one second
	}
	return nil
}

func main() {
	fmt.Println("Hello, Go Server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	hello.RegisterHelloServiceServer(s, &Server{})


	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
