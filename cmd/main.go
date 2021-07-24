package main

import (
	"fmt"
	"log"
	"net"

	md "github.com/sonr-io/core/pkg/models/js"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	if err != nil {
		log.Fatal(err)
	}

	chatServer := md.Server{}

	grpcServer := grpc.NewServer()

	md.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
