package main

import (
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

	c := hello.NewHelloServiceClient(cc)
}