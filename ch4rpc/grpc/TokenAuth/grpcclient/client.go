package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}

	port := ":1234"
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
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
