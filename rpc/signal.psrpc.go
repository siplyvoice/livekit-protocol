// Code generated by protoc-gen-psrpc v0.3.0, DO NOT EDIT.
// source: rpc/signal.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)

var _ = version.PsrpcVersion_0_3_0

// =======================
// Signal Client Interface
// =======================

type SignalClient[NodeIdTopicType ~string] interface {
	RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error)
}

// ===========================
// Signal ServerImpl Interface
// ===========================

type SignalServerImpl interface {
	RelaySignal(psrpc.ServerStream[*RelaySignalResponse, *RelaySignalRequest]) error
}

// =======================
// Signal Server Interface
// =======================

type SignalServer[NodeIdTopicType ~string] interface {
	RegisterRelaySignalTopic(nodeId NodeIdTopicType) error
	DeregisterRelaySignalTopic(nodeId NodeIdTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// Signal Client
// =============

type signalClient[NodeIdTopicType ~string] struct {
	client *client.RPCClient
}

// NewSignalClient creates a psrpc client that implements the SignalClient interface.
func NewSignalClient[NodeIdTopicType ~string](clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SignalClient[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   clientID,
	}

	sd.RegisterMethod("RelaySignal", false, false, false)

	rpcClient, err := client.NewRPCClientWithStreams(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &signalClient[NodeIdTopicType]{
		client: rpcClient,
	}, nil
}

func (c *signalClient[NodeIdTopicType]) RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error) {
	return client.OpenStream[*RelaySignalRequest, *RelaySignalResponse](ctx, c.client, "RelaySignal", []string{string(nodeId)}, opts...)
}

// =============
// Signal Server
// =============

type signalServer[NodeIdTopicType ~string] struct {
	svc SignalServerImpl
	rpc *server.RPCServer
}

// NewSignalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSignalServer[NodeIdTopicType ~string](serverID string, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SignalServer[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RelaySignal", false, false, false)
	return &signalServer[NodeIdTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *signalServer[NodeIdTopicType]) RegisterRelaySignalTopic(nodeId NodeIdTopicType) error {
	return server.RegisterStreamHandler(s.rpc, "RelaySignal", []string{string(nodeId)}, s.svc.RelaySignal, nil)
}

func (s *signalServer[NodeIdTopicType]) DeregisterRelaySignalTopic(nodeId NodeIdTopicType) {
	s.rpc.DeregisterHandler("RelaySignal", []string{string(nodeId)})
}

func (s *signalServer[NodeIdTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *signalServer[NodeIdTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0xe5, 0x2f, 0x6d, 0xbf, 0xe2, 0x52, 0xa9, 0xb8, 0xd0, 0x5a, 0x59, 0x45, 0x59, 0x45,
	0x08, 0x25, 0x28, 0x88, 0x0d, 0x4b, 0x8e, 0xe0, 0xae, 0x60, 0x53, 0xb5, 0xee, 0xa8, 0x58, 0x0d,
	0xb6, 0xeb, 0x71, 0x2b, 0x71, 0x04, 0xee, 0xc2, 0x8a, 0x13, 0xa2, 0xd6, 0xfd, 0x13, 0x14, 0x56,
	0x99, 0x99, 0xf7, 0xf2, 0x7e, 0xa3, 0x31, 0x1d, 0x38, 0x2b, 0x0b, 0x54, 0x4b, 0x3d, 0xab, 0x72,
	0xeb, 0x8c, 0x37, 0x2c, 0x72, 0x56, 0xc6, 0x7d, 0x63, 0xbd, 0x32, 0x1a, 0xc3, 0x2c, 0x1e, 0x55,
	0x6a, 0x0b, 0x2b, 0xe5, 0xa7, 0x4a, 0x7b, 0x70, 0x27, 0x6f, 0x7c, 0x75, 0x9c, 0x3b, 0x2f, 0xc3,
	0x28, 0xfd, 0x22, 0x94, 0x09, 0xa8, 0x66, 0x1f, 0x93, 0x7d, 0xa8, 0x80, 0xf5, 0x06, 0xd0, 0xb3,
	0x27, 0xda, 0x47, 0x3f, 0x73, 0x7e, 0x8a, 0x80, 0xa8, 0x8c, 0xe6, 0x24, 0x21, 0x59, 0xaf, 0xbc,
	0xc9, 0x0f, 0x09, 0xf9, 0x64, 0xa7, 0x4e, 0x82, 0x28, 0x2e, 0xb1, 0xd6, 0xb1, 0x92, 0x76, 0x5d,
	0x88, 0x41, 0x1e, 0x25, 0x51, 0xd6, 0x2b, 0x47, 0xe7, 0xdf, 0xea, 0x14, 0x71, 0xf2, 0xb1, 0x01,
	0x8d, 0x10, 0xd6, 0xbc, 0x95, 0x90, 0xac, 0x25, 0x76, 0x25, 0xbb, 0xa6, 0x6d, 0x59, 0x19, 0x04,
	0xde, 0x4e, 0x48, 0xd6, 0x15, 0xa1, 0x49, 0x3d, 0x1d, 0xfe, 0xda, 0x16, 0xad, 0xd1, 0x08, 0xec,
	0x91, 0x5e, 0xb8, 0x43, 0x8d, 0xfc, 0xdf, 0x9e, 0x39, 0x6e, 0x30, 0x83, 0x2e, 0xce, 0xce, 0x23,
	0x35, 0xfa, 0x83, 0xda, 0xaa, 0x51, 0x4b, 0x49, 0x3b, 0x21, 0x84, 0xbd, 0xd0, 0x5e, 0x8d, 0xcf,
	0xc6, 0xb9, 0xb3, 0x32, 0x6f, 0xde, 0x2f, 0xe6, 0x4d, 0x21, 0x40, 0xd3, 0xf1, 0xf7, 0x27, 0x19,
	0x72, 0x92, 0xf6, 0xd9, 0x7f, 0x6d, 0x16, 0x30, 0x55, 0x8b, 0xdd, 0x6d, 0xef, 0xc9, 0xf3, 0xdd,
	0xeb, 0xed, 0x52, 0xf9, 0xb7, 0xcd, 0x3c, 0x97, 0xe6, 0xbd, 0x38, 0x2c, 0x7f, 0xfa, 0xda, 0xd5,
	0xb2, 0x40, 0x70, 0x5b, 0x25, 0xa1, 0x70, 0x56, 0xce, 0x3b, 0xfb, 0xe7, 0x7b, 0xf8, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x73, 0x86, 0x5c, 0x11, 0x02, 0x00, 0x00,
}
