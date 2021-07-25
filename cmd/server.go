package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	sc "github.com/sonr-io/core/internal/client"
	snet "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/grpc"
)

type NodeServer struct {
	md.NodeServiceServer
	ctx context.Context

	// Client
	client sc.Client
	state  md.LifecycleState
	user   *md.User

	// Groups
	local  *snet.TopicManager
	topics map[string]*snet.TopicManager
}

func main() {
	// Create a new gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("- Sonr RPC Server Active -")

	// Set GRPC Server
	chatServer := NodeServer{}
	grpcServer := grpc.NewServer()

	// Register the gRPC service
	md.RegisterNodeServiceServer(grpcServer, &chatServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (s *NodeServer) Ping(ctx context.Context, message *md.PingRequest) (*md.PingResponse, error) {
	log.Println("Ping Called")
	return &md.PingResponse{Body: fmt.Sprintf("New Message: %s", message.Body)}, nil
}

func (s *NodeServer) Initialize(ctx context.Context, req *md.InitializeRequest) (*md.InitializeResponse, error) {
	log.Println("Initialize Called")
	return &md.InitializeResponse{
		Body:    "Initialize Success",
		Success: true}, nil
}

func (s *NodeServer) Connect(ctx context.Context, req *md.ConnectionRequest) (*md.ConnectionResponse, error) {
	log.Println("Connect Called")
	return &md.ConnectionResponse{}, nil
}

func (s *NodeServer) Sign(ctx context.Context, req *md.AuthRequest) (*md.AuthResponse, error) {
	log.Println("Sign Called")
	return &md.AuthResponse{}, nil
}

func (s *NodeServer) Verify(ctx context.Context, req *md.VerifyRequest) (*md.VerifyResponse, error) {
	log.Println("Verify Called")
	return &md.VerifyResponse{}, nil
}

func (s *NodeServer) Invite(ctx context.Context, req *md.InviteRequest) (*md.InviteResponse, error) {
	log.Println("Invite Called")
	return &md.InviteResponse{}, nil
}

func (s *NodeServer) Respond(ctx context.Context, req *md.RespondRequest) (*md.RespondResponse, error) {
	log.Println("Respond Called")
	return &md.RespondResponse{}, nil
}

func (s *NodeServer) Mail(ctx context.Context, req *md.MailboxRequest) (*md.MailboxResponse, error) {
	log.Println("Mail Called")
	return &md.MailboxResponse{}, nil
}

func (s *NodeServer) OnComplete(req *md.InitializeRequest, stream md.NodeService_OnCompleteServer) error {
	// Create Test Events
	events := []*md.CompleteEvent{
		{
			Direction: md.CompleteEvent_Incoming,
			Transfer:  &md.Transfer{},
		},
		{
			Direction: md.CompleteEvent_Default,
			Transfer:  &md.Transfer{},
		},
		{
			Direction: md.CompleteEvent_Outgoing,
			Transfer:  &md.Transfer{},
		},
	}

	// Send Sample Events
	for _, ev := range events {
		if err := stream.Send(ev); err != nil {
			return err
		}
		time.Sleep(time.Second * 4)
	}
	return nil
}
