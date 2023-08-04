package main

import (
	"context"
	"fmt"
	"log"
	hello "test/hello/proto"

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

	helloUnary(client)
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
		log.Fatalf("Error, calling Hello RPC: \n *v",err)
	}

	log.Printf("Response Hello: %v",res.CustomerHello)
}