// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: inx.proto

package inx

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// INXClient is the client API for INX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type INXClient interface {
	// Node
	ReadNodeStatus(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*NodeStatus, error)
	ListenToNodeStatus(ctx context.Context, in *NodeStatusRequest, opts ...grpc.CallOption) (INX_ListenToNodeStatusClient, error)
	ReadNodeConfiguration(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*NodeConfiguration, error)
	// Commitments
	ReadCommitment(ctx context.Context, in *CommitmentRequest, opts ...grpc.CallOption) (*Commitment, error)
	// Blocks
	ListenToBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToBlocksClient, error)
	ListenToAcceptedBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToAcceptedBlocksClient, error)
	ListenToConfirmedBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToConfirmedBlocksClient, error)
	SubmitBlock(ctx context.Context, in *RawBlock, opts ...grpc.CallOption) (*BlockId, error)
	ReadBlock(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*RawBlock, error)
	ReadBlockMetadata(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*BlockMetadata, error)
	// Block Issuance
	RequestTips(ctx context.Context, in *TipsRequest, opts ...grpc.CallOption) (*TipsResponse, error)
	ValidatePayload(ctx context.Context, in *RawPayload, opts ...grpc.CallOption) (*PayloadValidationResponse, error)
	// UTXO
	ReadUnspentOutputs(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ReadUnspentOutputsClient, error)
	// Committee
	ReadIsCommitteeMember(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	ReadIsCandidate(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	ReadIsAccountValidator(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	// A stream that yields updates to the ledger. A `LedgerUpdate` represents a batch to be applied to the ledger.
	// It first sends a `BEGIN`, then all the consumed outputs, then all the created outputs and finally an `END`.
	// `BEGIN` and `END` will also be sent for slots that did not mutate the ledger.
	// The counts in the batch markers can be used to sanity check that everything arrived and to pre-allocate space if needed.
	ListenToLedgerUpdates(ctx context.Context, in *SlotRangeRequest, opts ...grpc.CallOption) (INX_ListenToLedgerUpdatesClient, error)
	ListenToAcceptedTransactions(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToAcceptedTransactionsClient, error)
	ReadOutput(ctx context.Context, in *OutputId, opts ...grpc.CallOption) (*OutputResponse, error)
	// REST API
	RegisterAPIRoute(ctx context.Context, in *APIRouteRequest, opts ...grpc.CallOption) (*NoParams, error)
	UnregisterAPIRoute(ctx context.Context, in *APIRouteRequest, opts ...grpc.CallOption) (*NoParams, error)
	PerformAPIRequest(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*APIResponse, error)
}

type iNXClient struct {
	cc grpc.ClientConnInterface
}

func NewINXClient(cc grpc.ClientConnInterface) INXClient {
	return &iNXClient{cc}
}

func (c *iNXClient) ReadNodeStatus(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*NodeStatus, error) {
	out := new(NodeStatus)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ListenToNodeStatus(ctx context.Context, in *NodeStatusRequest, opts ...grpc.CallOption) (INX_ListenToNodeStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[0], "/inx.INX/ListenToNodeStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToNodeStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToNodeStatusClient interface {
	Recv() (*NodeStatus, error)
	grpc.ClientStream
}

type iNXListenToNodeStatusClient struct {
	grpc.ClientStream
}

func (x *iNXListenToNodeStatusClient) Recv() (*NodeStatus, error) {
	m := new(NodeStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ReadNodeConfiguration(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*NodeConfiguration, error) {
	out := new(NodeConfiguration)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadNodeConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadCommitment(ctx context.Context, in *CommitmentRequest, opts ...grpc.CallOption) (*Commitment, error) {
	out := new(Commitment)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadCommitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ListenToBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[1], "/inx.INX/ListenToBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type iNXListenToBlocksClient struct {
	grpc.ClientStream
}

func (x *iNXListenToBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ListenToAcceptedBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToAcceptedBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[2], "/inx.INX/ListenToAcceptedBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToAcceptedBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToAcceptedBlocksClient interface {
	Recv() (*BlockMetadata, error)
	grpc.ClientStream
}

type iNXListenToAcceptedBlocksClient struct {
	grpc.ClientStream
}

func (x *iNXListenToAcceptedBlocksClient) Recv() (*BlockMetadata, error) {
	m := new(BlockMetadata)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ListenToConfirmedBlocks(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToConfirmedBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[3], "/inx.INX/ListenToConfirmedBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToConfirmedBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToConfirmedBlocksClient interface {
	Recv() (*BlockMetadata, error)
	grpc.ClientStream
}

type iNXListenToConfirmedBlocksClient struct {
	grpc.ClientStream
}

func (x *iNXListenToConfirmedBlocksClient) Recv() (*BlockMetadata, error) {
	m := new(BlockMetadata)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) SubmitBlock(ctx context.Context, in *RawBlock, opts ...grpc.CallOption) (*BlockId, error) {
	out := new(BlockId)
	err := c.cc.Invoke(ctx, "/inx.INX/SubmitBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadBlock(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*RawBlock, error) {
	out := new(RawBlock)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadBlockMetadata(ctx context.Context, in *BlockId, opts ...grpc.CallOption) (*BlockMetadata, error) {
	out := new(BlockMetadata)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadBlockMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) RequestTips(ctx context.Context, in *TipsRequest, opts ...grpc.CallOption) (*TipsResponse, error) {
	out := new(TipsResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/RequestTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ValidatePayload(ctx context.Context, in *RawPayload, opts ...grpc.CallOption) (*PayloadValidationResponse, error) {
	out := new(PayloadValidationResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/ValidatePayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadUnspentOutputs(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ReadUnspentOutputsClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[4], "/inx.INX/ReadUnspentOutputs", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXReadUnspentOutputsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ReadUnspentOutputsClient interface {
	Recv() (*UnspentOutput, error)
	grpc.ClientStream
}

type iNXReadUnspentOutputsClient struct {
	grpc.ClientStream
}

func (x *iNXReadUnspentOutputsClient) Recv() (*UnspentOutput, error) {
	m := new(UnspentOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ReadIsCommitteeMember(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadIsCommitteeMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadIsCandidate(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadIsCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ReadIsAccountValidator(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadIsAccountValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) ListenToLedgerUpdates(ctx context.Context, in *SlotRangeRequest, opts ...grpc.CallOption) (INX_ListenToLedgerUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[5], "/inx.INX/ListenToLedgerUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToLedgerUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToLedgerUpdatesClient interface {
	Recv() (*LedgerUpdate, error)
	grpc.ClientStream
}

type iNXListenToLedgerUpdatesClient struct {
	grpc.ClientStream
}

func (x *iNXListenToLedgerUpdatesClient) Recv() (*LedgerUpdate, error) {
	m := new(LedgerUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ListenToAcceptedTransactions(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (INX_ListenToAcceptedTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &INX_ServiceDesc.Streams[6], "/inx.INX/ListenToAcceptedTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &iNXListenToAcceptedTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type INX_ListenToAcceptedTransactionsClient interface {
	Recv() (*AcceptedTransaction, error)
	grpc.ClientStream
}

type iNXListenToAcceptedTransactionsClient struct {
	grpc.ClientStream
}

func (x *iNXListenToAcceptedTransactionsClient) Recv() (*AcceptedTransaction, error) {
	m := new(AcceptedTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iNXClient) ReadOutput(ctx context.Context, in *OutputId, opts ...grpc.CallOption) (*OutputResponse, error) {
	out := new(OutputResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/ReadOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) RegisterAPIRoute(ctx context.Context, in *APIRouteRequest, opts ...grpc.CallOption) (*NoParams, error) {
	out := new(NoParams)
	err := c.cc.Invoke(ctx, "/inx.INX/RegisterAPIRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) UnregisterAPIRoute(ctx context.Context, in *APIRouteRequest, opts ...grpc.CallOption) (*NoParams, error) {
	out := new(NoParams)
	err := c.cc.Invoke(ctx, "/inx.INX/UnregisterAPIRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iNXClient) PerformAPIRequest(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, "/inx.INX/PerformAPIRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// INXServer is the server API for INX service.
// All implementations must embed UnimplementedINXServer
// for forward compatibility
type INXServer interface {
	// Node
	ReadNodeStatus(context.Context, *NoParams) (*NodeStatus, error)
	ListenToNodeStatus(*NodeStatusRequest, INX_ListenToNodeStatusServer) error
	ReadNodeConfiguration(context.Context, *NoParams) (*NodeConfiguration, error)
	// Commitments
	ReadCommitment(context.Context, *CommitmentRequest) (*Commitment, error)
	// Blocks
	ListenToBlocks(*NoParams, INX_ListenToBlocksServer) error
	ListenToAcceptedBlocks(*NoParams, INX_ListenToAcceptedBlocksServer) error
	ListenToConfirmedBlocks(*NoParams, INX_ListenToConfirmedBlocksServer) error
	SubmitBlock(context.Context, *RawBlock) (*BlockId, error)
	ReadBlock(context.Context, *BlockId) (*RawBlock, error)
	ReadBlockMetadata(context.Context, *BlockId) (*BlockMetadata, error)
	// Block Issuance
	RequestTips(context.Context, *TipsRequest) (*TipsResponse, error)
	ValidatePayload(context.Context, *RawPayload) (*PayloadValidationResponse, error)
	// UTXO
	ReadUnspentOutputs(*NoParams, INX_ReadUnspentOutputsServer) error
	// Committee
	ReadIsCommitteeMember(context.Context, *AccountInfoRequest) (*BoolResponse, error)
	ReadIsCandidate(context.Context, *AccountInfoRequest) (*BoolResponse, error)
	ReadIsAccountValidator(context.Context, *AccountInfoRequest) (*BoolResponse, error)
	// A stream that yields updates to the ledger. A `LedgerUpdate` represents a batch to be applied to the ledger.
	// It first sends a `BEGIN`, then all the consumed outputs, then all the created outputs and finally an `END`.
	// `BEGIN` and `END` will also be sent for slots that did not mutate the ledger.
	// The counts in the batch markers can be used to sanity check that everything arrived and to pre-allocate space if needed.
	ListenToLedgerUpdates(*SlotRangeRequest, INX_ListenToLedgerUpdatesServer) error
	ListenToAcceptedTransactions(*NoParams, INX_ListenToAcceptedTransactionsServer) error
	ReadOutput(context.Context, *OutputId) (*OutputResponse, error)
	// REST API
	RegisterAPIRoute(context.Context, *APIRouteRequest) (*NoParams, error)
	UnregisterAPIRoute(context.Context, *APIRouteRequest) (*NoParams, error)
	PerformAPIRequest(context.Context, *APIRequest) (*APIResponse, error)
	mustEmbedUnimplementedINXServer()
}

// UnimplementedINXServer must be embedded to have forward compatible implementations.
type UnimplementedINXServer struct {
}

func (UnimplementedINXServer) ReadNodeStatus(context.Context, *NoParams) (*NodeStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNodeStatus not implemented")
}
func (UnimplementedINXServer) ListenToNodeStatus(*NodeStatusRequest, INX_ListenToNodeStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToNodeStatus not implemented")
}
func (UnimplementedINXServer) ReadNodeConfiguration(context.Context, *NoParams) (*NodeConfiguration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNodeConfiguration not implemented")
}
func (UnimplementedINXServer) ReadCommitment(context.Context, *CommitmentRequest) (*Commitment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCommitment not implemented")
}
func (UnimplementedINXServer) ListenToBlocks(*NoParams, INX_ListenToBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToBlocks not implemented")
}
func (UnimplementedINXServer) ListenToAcceptedBlocks(*NoParams, INX_ListenToAcceptedBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToAcceptedBlocks not implemented")
}
func (UnimplementedINXServer) ListenToConfirmedBlocks(*NoParams, INX_ListenToConfirmedBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToConfirmedBlocks not implemented")
}
func (UnimplementedINXServer) SubmitBlock(context.Context, *RawBlock) (*BlockId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBlock not implemented")
}
func (UnimplementedINXServer) ReadBlock(context.Context, *BlockId) (*RawBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlock not implemented")
}
func (UnimplementedINXServer) ReadBlockMetadata(context.Context, *BlockId) (*BlockMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlockMetadata not implemented")
}
func (UnimplementedINXServer) RequestTips(context.Context, *TipsRequest) (*TipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTips not implemented")
}
func (UnimplementedINXServer) ValidatePayload(context.Context, *RawPayload) (*PayloadValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePayload not implemented")
}
func (UnimplementedINXServer) ReadUnspentOutputs(*NoParams, INX_ReadUnspentOutputsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadUnspentOutputs not implemented")
}
func (UnimplementedINXServer) ReadIsCommitteeMember(context.Context, *AccountInfoRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadIsCommitteeMember not implemented")
}
func (UnimplementedINXServer) ReadIsCandidate(context.Context, *AccountInfoRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadIsCandidate not implemented")
}
func (UnimplementedINXServer) ReadIsAccountValidator(context.Context, *AccountInfoRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadIsAccountValidator not implemented")
}
func (UnimplementedINXServer) ListenToLedgerUpdates(*SlotRangeRequest, INX_ListenToLedgerUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToLedgerUpdates not implemented")
}
func (UnimplementedINXServer) ListenToAcceptedTransactions(*NoParams, INX_ListenToAcceptedTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenToAcceptedTransactions not implemented")
}
func (UnimplementedINXServer) ReadOutput(context.Context, *OutputId) (*OutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOutput not implemented")
}
func (UnimplementedINXServer) RegisterAPIRoute(context.Context, *APIRouteRequest) (*NoParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAPIRoute not implemented")
}
func (UnimplementedINXServer) UnregisterAPIRoute(context.Context, *APIRouteRequest) (*NoParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterAPIRoute not implemented")
}
func (UnimplementedINXServer) PerformAPIRequest(context.Context, *APIRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformAPIRequest not implemented")
}
func (UnimplementedINXServer) mustEmbedUnimplementedINXServer() {}

// UnsafeINXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to INXServer will
// result in compilation errors.
type UnsafeINXServer interface {
	mustEmbedUnimplementedINXServer()
}

func RegisterINXServer(s grpc.ServiceRegistrar, srv INXServer) {
	s.RegisterService(&INX_ServiceDesc, srv)
}

func _INX_ReadNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadNodeStatus(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ListenToNodeStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToNodeStatus(m, &iNXListenToNodeStatusServer{stream})
}

type INX_ListenToNodeStatusServer interface {
	Send(*NodeStatus) error
	grpc.ServerStream
}

type iNXListenToNodeStatusServer struct {
	grpc.ServerStream
}

func (x *iNXListenToNodeStatusServer) Send(m *NodeStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ReadNodeConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadNodeConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadNodeConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadNodeConfiguration(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadCommitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadCommitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadCommitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadCommitment(ctx, req.(*CommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ListenToBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToBlocks(m, &iNXListenToBlocksServer{stream})
}

type INX_ListenToBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type iNXListenToBlocksServer struct {
	grpc.ServerStream
}

func (x *iNXListenToBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ListenToAcceptedBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToAcceptedBlocks(m, &iNXListenToAcceptedBlocksServer{stream})
}

type INX_ListenToAcceptedBlocksServer interface {
	Send(*BlockMetadata) error
	grpc.ServerStream
}

type iNXListenToAcceptedBlocksServer struct {
	grpc.ServerStream
}

func (x *iNXListenToAcceptedBlocksServer) Send(m *BlockMetadata) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ListenToConfirmedBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToConfirmedBlocks(m, &iNXListenToConfirmedBlocksServer{stream})
}

type INX_ListenToConfirmedBlocksServer interface {
	Send(*BlockMetadata) error
	grpc.ServerStream
}

type iNXListenToConfirmedBlocksServer struct {
	grpc.ServerStream
}

func (x *iNXListenToConfirmedBlocksServer) Send(m *BlockMetadata) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_SubmitBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).SubmitBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/SubmitBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).SubmitBlock(ctx, req.(*RawBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadBlock(ctx, req.(*BlockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadBlockMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadBlockMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadBlockMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadBlockMetadata(ctx, req.(*BlockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_RequestTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).RequestTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/RequestTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).RequestTips(ctx, req.(*TipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ValidatePayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ValidatePayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ValidatePayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ValidatePayload(ctx, req.(*RawPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadUnspentOutputs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ReadUnspentOutputs(m, &iNXReadUnspentOutputsServer{stream})
}

type INX_ReadUnspentOutputsServer interface {
	Send(*UnspentOutput) error
	grpc.ServerStream
}

type iNXReadUnspentOutputsServer struct {
	grpc.ServerStream
}

func (x *iNXReadUnspentOutputsServer) Send(m *UnspentOutput) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ReadIsCommitteeMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadIsCommitteeMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadIsCommitteeMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadIsCommitteeMember(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadIsCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadIsCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadIsCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadIsCandidate(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ReadIsAccountValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadIsAccountValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadIsAccountValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadIsAccountValidator(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_ListenToLedgerUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SlotRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToLedgerUpdates(m, &iNXListenToLedgerUpdatesServer{stream})
}

type INX_ListenToLedgerUpdatesServer interface {
	Send(*LedgerUpdate) error
	grpc.ServerStream
}

type iNXListenToLedgerUpdatesServer struct {
	grpc.ServerStream
}

func (x *iNXListenToLedgerUpdatesServer) Send(m *LedgerUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ListenToAcceptedTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(INXServer).ListenToAcceptedTransactions(m, &iNXListenToAcceptedTransactionsServer{stream})
}

type INX_ListenToAcceptedTransactionsServer interface {
	Send(*AcceptedTransaction) error
	grpc.ServerStream
}

type iNXListenToAcceptedTransactionsServer struct {
	grpc.ServerStream
}

func (x *iNXListenToAcceptedTransactionsServer) Send(m *AcceptedTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _INX_ReadOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutputId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).ReadOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/ReadOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).ReadOutput(ctx, req.(*OutputId))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_RegisterAPIRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).RegisterAPIRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/RegisterAPIRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).RegisterAPIRoute(ctx, req.(*APIRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_UnregisterAPIRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).UnregisterAPIRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/UnregisterAPIRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).UnregisterAPIRoute(ctx, req.(*APIRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _INX_PerformAPIRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INXServer).PerformAPIRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inx.INX/PerformAPIRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INXServer).PerformAPIRequest(ctx, req.(*APIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// INX_ServiceDesc is the grpc.ServiceDesc for INX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var INX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inx.INX",
	HandlerType: (*INXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadNodeStatus",
			Handler:    _INX_ReadNodeStatus_Handler,
		},
		{
			MethodName: "ReadNodeConfiguration",
			Handler:    _INX_ReadNodeConfiguration_Handler,
		},
		{
			MethodName: "ReadCommitment",
			Handler:    _INX_ReadCommitment_Handler,
		},
		{
			MethodName: "SubmitBlock",
			Handler:    _INX_SubmitBlock_Handler,
		},
		{
			MethodName: "ReadBlock",
			Handler:    _INX_ReadBlock_Handler,
		},
		{
			MethodName: "ReadBlockMetadata",
			Handler:    _INX_ReadBlockMetadata_Handler,
		},
		{
			MethodName: "RequestTips",
			Handler:    _INX_RequestTips_Handler,
		},
		{
			MethodName: "ValidatePayload",
			Handler:    _INX_ValidatePayload_Handler,
		},
		{
			MethodName: "ReadIsCommitteeMember",
			Handler:    _INX_ReadIsCommitteeMember_Handler,
		},
		{
			MethodName: "ReadIsCandidate",
			Handler:    _INX_ReadIsCandidate_Handler,
		},
		{
			MethodName: "ReadIsAccountValidator",
			Handler:    _INX_ReadIsAccountValidator_Handler,
		},
		{
			MethodName: "ReadOutput",
			Handler:    _INX_ReadOutput_Handler,
		},
		{
			MethodName: "RegisterAPIRoute",
			Handler:    _INX_RegisterAPIRoute_Handler,
		},
		{
			MethodName: "UnregisterAPIRoute",
			Handler:    _INX_UnregisterAPIRoute_Handler,
		},
		{
			MethodName: "PerformAPIRequest",
			Handler:    _INX_PerformAPIRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenToNodeStatus",
			Handler:       _INX_ListenToNodeStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenToBlocks",
			Handler:       _INX_ListenToBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenToAcceptedBlocks",
			Handler:       _INX_ListenToAcceptedBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenToConfirmedBlocks",
			Handler:       _INX_ListenToConfirmedBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadUnspentOutputs",
			Handler:       _INX_ReadUnspentOutputs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenToLedgerUpdates",
			Handler:       _INX_ListenToLedgerUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenToAcceptedTransactions",
			Handler:       _INX_ListenToAcceptedTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inx.proto",
}
