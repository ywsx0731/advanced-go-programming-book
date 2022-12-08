package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	err := RegisterHelloService(new(HelloService))
	if err != nil {
		log.Fatal("register service error:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
		return
	}
loop:
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			continue loop
		}

		go rpc.ServeConn(conn)
	}
}
