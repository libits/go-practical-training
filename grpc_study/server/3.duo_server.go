package main

import (
	"context"
	"fmt"
	"grpc_study/service_proto/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type VideoServer struct {
}

func (VideoServer) Look(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("video:", request)
	return &proto.Response{
		Name: "xiaoli",
	}, nil
}

type OrderServer struct {
}

func (OrderServer) Look(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("order:", request)
	return &proto.Response{
		Name: "xiaoli",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterVideoServiceServer(s, &VideoServer{})
	proto.RegisterOrderServiceServer(s, &OrderServer{})
	fmt.Println("grpc server程序运行在：8080")
	err = s.Serve(listen)
}
