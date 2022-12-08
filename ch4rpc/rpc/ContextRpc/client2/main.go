package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func doClientWork(client *rpc.Client) {
	helloCall := client.Go("HelloService.Hello", "hel44lo", new(string), nil)

	// do some thing

	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)
}
