package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 10, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Call(
		"KVStoreService.Set", [2]string{"abc", "abcf-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Call(
		"KVStoreService.Set", [2]string{"xyz", "abcf-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
	select {}
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)
}
