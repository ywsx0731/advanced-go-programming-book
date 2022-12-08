package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(
		"../server.crt", "localhost",
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hefdfllo"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
