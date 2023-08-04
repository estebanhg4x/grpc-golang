package main

import (
	"context"
	"fmt"
	"log"
	hello "test/hello/proto"
	"io"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

	//Se ejecuta al finalizar el ciclo de vida de la
	defer cc.Close()

	client := hello.NewHelloServiceClient(cc)

	fmt.Println("Starting unary RPC Hello")

	//helloUnary(client)
	helloServerStreaming(client)
}


func helloUnary(client hello.HelloServiceClient){
	//llamado envio peticion meto hello
	req := &hello.HelloRequest{
		Hello: &hello.Hello{
			FirstName: "Test",
			Prefix: "Sr",
		},
	}

	res, err := client.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: %v\n",err)
	}

	log.Printf("Response Hello: %v",res.CustomerHello)
}

func helloServerStreaming(client hello.HelloServiceClient){
	fmt.Println("Starting server streaming RPC Hello")

	req := &hello.HelloManyLenguagesRequest{
		HelloLenguages: &hello.Hello{
			FirstName: "Test stream",
			Prefix: "Joven",
		},
	}

	restSream, err := client.HelloManyLenguagues(context.Background(),req)
	if err != nil {
		log.Printf("Error calling hello many lenguages %v",err)
	}

	for {
		msg, err := restSream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream %v", err)
		}
		log.Println("Res for HML %v\n", msg.GetHelloLenguages())
	}
}