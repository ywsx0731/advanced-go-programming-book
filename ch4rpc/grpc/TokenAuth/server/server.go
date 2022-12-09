package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type Authentication struct {
	User     string
	Password string
}

//func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
//	map[string]string, error,
//) {
//	return map[string]string{"user": a.User, "password": a.Password}, nil
//}
//func (a *Authentication) RequireTransportSecurity() bool {
//	return false
//}

type HelloServiceImpl struct {
	auth *Authentication
}

func NewHelloServiceImpl() *HelloServiceImpl {
	return &HelloServiceImpl{&Authentication{
		User:     "gopher",
		Password: "password",
	}}
}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *String,
) (*String, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}

	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, NewHelloServiceImpl())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
