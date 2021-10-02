// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

import (
	context "context"
	common "github.com/sonr-io/core/internal/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Stat Method returns the Node Stats
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	// Events Streams
	// Returns a stream of StatusEvents
	OnNodeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnNodeStatusClient, error)
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnLobbyRefreshClient, error)
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferAcceptedClient, error)
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferDeclinedClient, error)
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferInviteClient, error)
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferProgressClient, error)
	// Returns a stream of Completed Transfers
	OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferCompleteClient, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error) {
	out := new(RespondResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) OnNodeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnNodeStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/sonr.node.ClientService/OnNodeStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnNodeStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnNodeStatusClient interface {
	Recv() (*common.StatusEvent, error)
	grpc.ClientStream
}

type clientServiceOnNodeStatusClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnNodeStatusClient) Recv() (*common.StatusEvent, error) {
	m := new(common.StatusEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnLobbyRefreshClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[1], "/sonr.node.ClientService/OnLobbyRefresh", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnLobbyRefreshClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnLobbyRefreshClient interface {
	Recv() (*common.RefreshEvent, error)
	grpc.ClientStream
}

type clientServiceOnLobbyRefreshClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnLobbyRefreshClient) Recv() (*common.RefreshEvent, error) {
	m := new(common.RefreshEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferAcceptedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[2], "/sonr.node.ClientService/OnTransferAccepted", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferAcceptedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferAcceptedClient interface {
	Recv() (*common.DecisionEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferAcceptedClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferAcceptedClient) Recv() (*common.DecisionEvent, error) {
	m := new(common.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferDeclinedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[3], "/sonr.node.ClientService/OnTransferDeclined", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferDeclinedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferDeclinedClient interface {
	Recv() (*common.DecisionEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferDeclinedClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferDeclinedClient) Recv() (*common.DecisionEvent, error) {
	m := new(common.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferInviteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[4], "/sonr.node.ClientService/OnTransferInvite", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferInviteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferInviteClient interface {
	Recv() (*common.InviteEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferInviteClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferInviteClient) Recv() (*common.InviteEvent, error) {
	m := new(common.InviteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[5], "/sonr.node.ClientService/OnTransferProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferProgressClient interface {
	Recv() (*common.ProgressEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferProgressClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferProgressClient) Recv() (*common.ProgressEvent, error) {
	m := new(common.ProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferCompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[6], "/sonr.node.ClientService/OnTransferComplete", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferCompleteClient interface {
	Recv() (*common.CompleteEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferCompleteClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferCompleteClient) Recv() (*common.CompleteEvent, error) {
	m := new(common.CompleteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(context.Context, *SupplyRequest) (*SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(context.Context, *EditRequest) (*EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(context.Context, *ShareRequest) (*ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(context.Context, *RespondRequest) (*RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Stat Method returns the Node Stats
	Stat(context.Context, *StatRequest) (*StatResponse, error)
	// Events Streams
	// Returns a stream of StatusEvents
	OnNodeStatus(*Empty, ClientService_OnNodeStatusServer) error
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(*Empty, ClientService_OnLobbyRefreshServer) error
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(*Empty, ClientService_OnTransferAcceptedServer) error
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(*Empty, ClientService_OnTransferDeclinedServer) error
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(*Empty, ClientService_OnTransferInviteServer) error
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(*Empty, ClientService_OnTransferProgressServer) error
	// Returns a stream of Completed Transfers
	OnTransferComplete(*Empty, ClientService_OnTransferCompleteServer) error
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Supply(context.Context, *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedClientServiceServer) Edit(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedClientServiceServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedClientServiceServer) Share(context.Context, *ShareRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedClientServiceServer) Respond(context.Context, *RespondRequest) (*RespondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedClientServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedClientServiceServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedClientServiceServer) OnNodeStatus(*Empty, ClientService_OnNodeStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method OnNodeStatus not implemented")
}
func (UnimplementedClientServiceServer) OnLobbyRefresh(*Empty, ClientService_OnLobbyRefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method OnLobbyRefresh not implemented")
}
func (UnimplementedClientServiceServer) OnTransferAccepted(*Empty, ClientService_OnTransferAcceptedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferAccepted not implemented")
}
func (UnimplementedClientServiceServer) OnTransferDeclined(*Empty, ClientService_OnTransferDeclinedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferDeclined not implemented")
}
func (UnimplementedClientServiceServer) OnTransferInvite(*Empty, ClientService_OnTransferInviteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferInvite not implemented")
}
func (UnimplementedClientServiceServer) OnTransferProgress(*Empty, ClientService_OnTransferProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferProgress not implemented")
}
func (UnimplementedClientServiceServer) OnTransferComplete(*Empty, ClientService_OnTransferCompleteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferComplete not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Supply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Edit(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Respond(ctx, req.(*RespondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_OnNodeStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnNodeStatus(m, &clientServiceOnNodeStatusServer{stream})
}

type ClientService_OnNodeStatusServer interface {
	Send(*common.StatusEvent) error
	grpc.ServerStream
}

type clientServiceOnNodeStatusServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnNodeStatusServer) Send(m *common.StatusEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnLobbyRefresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnLobbyRefresh(m, &clientServiceOnLobbyRefreshServer{stream})
}

type ClientService_OnLobbyRefreshServer interface {
	Send(*common.RefreshEvent) error
	grpc.ServerStream
}

type clientServiceOnLobbyRefreshServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnLobbyRefreshServer) Send(m *common.RefreshEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferAccepted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferAccepted(m, &clientServiceOnTransferAcceptedServer{stream})
}

type ClientService_OnTransferAcceptedServer interface {
	Send(*common.DecisionEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferAcceptedServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferAcceptedServer) Send(m *common.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferDeclined_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferDeclined(m, &clientServiceOnTransferDeclinedServer{stream})
}

type ClientService_OnTransferDeclinedServer interface {
	Send(*common.DecisionEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferDeclinedServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferDeclinedServer) Send(m *common.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferInvite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferInvite(m, &clientServiceOnTransferInviteServer{stream})
}

type ClientService_OnTransferInviteServer interface {
	Send(*common.InviteEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferInviteServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferInviteServer) Send(m *common.InviteEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferProgress(m, &clientServiceOnTransferProgressServer{stream})
}

type ClientService_OnTransferProgressServer interface {
	Send(*common.ProgressEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferProgressServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferProgressServer) Send(m *common.ProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferComplete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferComplete(m, &clientServiceOnTransferCompleteServer{stream})
}

type ClientService_OnTransferCompleteServer interface {
	Send(*common.CompleteEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferCompleteServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferCompleteServer) Send(m *common.CompleteEvent) error {
	return x.ServerStream.SendMsg(m)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonr.node.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Supply",
			Handler:    _ClientService_Supply_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _ClientService_Edit_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _ClientService_Fetch_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _ClientService_Share_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _ClientService_Respond_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ClientService_Search_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _ClientService_Stat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnNodeStatus",
			Handler:       _ClientService_OnNodeStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnLobbyRefresh",
			Handler:       _ClientService_OnLobbyRefresh_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferAccepted",
			Handler:       _ClientService_OnTransferAccepted_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferDeclined",
			Handler:       _ClientService_OnTransferDeclined_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferInvite",
			Handler:       _ClientService_OnTransferInvite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferProgress",
			Handler:       _ClientService_OnTransferProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferComplete",
			Handler:       _ClientService_OnTransferComplete_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/node/client.proto",
}
