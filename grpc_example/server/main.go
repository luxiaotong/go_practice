package main

import (
	"context"
	"log"
	"net"

	"github.com/luxiaotong/go_practice/grpc_example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	// a := in.GetA()
	// b := in.GetB()
	a, b := in.GetA(), in.GetB()
	res := a + b
	return &proto.Response{Result: res}, nil
}

func (s *server) Multiply(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	// a := in.GetA()
	// b := in.GetB()
	a, b := in.GetA(), in.GetB()
	res := a * b
	return &proto.Response{Result: res}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatal("Serve error: ", err)
	}
}
