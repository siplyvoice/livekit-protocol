// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/sip.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"

var _ = version.PsrpcVersion_0_6

// ============================
// SIPInternal Client Interface
// ============================

type SIPInternalClient interface {
	CreateSIPParticipant(ctx context.Context, topic string, req *InternalCreateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalCreateSIPParticipantResponse, error)

	TransferSIPParticipant(ctx context.Context, sipCallId string, req *InternalTransferSIPParticipantRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ================================
// SIPInternal ServerImpl Interface
// ================================

type SIPInternalServerImpl interface {
	CreateSIPParticipant(context.Context, *InternalCreateSIPParticipantRequest) (*InternalCreateSIPParticipantResponse, error)
	CreateSIPParticipantAffinity(context.Context, *InternalCreateSIPParticipantRequest) float32

	TransferSIPParticipant(context.Context, *InternalTransferSIPParticipantRequest) (*google_protobuf.Empty, error)
}

// ============================
// SIPInternal Server Interface
// ============================

type SIPInternalServer interface {
	RegisterCreateSIPParticipantTopic(topic string) error
	DeregisterCreateSIPParticipantTopic(topic string)
	RegisterTransferSIPParticipantTopic(sipCallId string) error
	DeregisterTransferSIPParticipantTopic(sipCallId string)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ==================
// SIPInternal Client
// ==================

type sIPInternalClient struct {
	client *client.RPCClient
}

// NewSIPInternalClient creates a psrpc client that implements the SIPInternalClient interface.
func NewSIPInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SIPInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "SIPInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateSIPParticipant", true, false, true, false)
	sd.RegisterMethod("TransferSIPParticipant", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &sIPInternalClient{
		client: rpcClient,
	}, nil
}

func (c *sIPInternalClient) CreateSIPParticipant(ctx context.Context, topic string, req *InternalCreateSIPParticipantRequest, opts ...psrpc.RequestOption) (*InternalCreateSIPParticipantResponse, error) {
	return client.RequestSingle[*InternalCreateSIPParticipantResponse](ctx, c.client, "CreateSIPParticipant", []string{topic}, req, opts...)
}

func (c *sIPInternalClient) TransferSIPParticipant(ctx context.Context, sipCallId string, req *InternalTransferSIPParticipantRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "TransferSIPParticipant", []string{sipCallId}, req, opts...)
}

func (s *sIPInternalClient) Close() {
	s.client.Close()
}

// ==================
// SIPInternal Server
// ==================

type sIPInternalServer struct {
	svc SIPInternalServerImpl
	rpc *server.RPCServer
}

// NewSIPInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSIPInternalServer(svc SIPInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SIPInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "SIPInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateSIPParticipant", true, false, true, false)
	sd.RegisterMethod("TransferSIPParticipant", false, false, true, true)
	return &sIPInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *sIPInternalServer) RegisterCreateSIPParticipantTopic(topic string) error {
	return server.RegisterHandler(s.rpc, "CreateSIPParticipant", []string{topic}, s.svc.CreateSIPParticipant, s.svc.CreateSIPParticipantAffinity)
}

func (s *sIPInternalServer) DeregisterCreateSIPParticipantTopic(topic string) {
	s.rpc.DeregisterHandler("CreateSIPParticipant", []string{topic})
}

func (s *sIPInternalServer) RegisterTransferSIPParticipantTopic(sipCallId string) error {
	return server.RegisterHandler(s.rpc, "TransferSIPParticipant", []string{sipCallId}, s.svc.TransferSIPParticipant, nil)
}

func (s *sIPInternalServer) DeregisterTransferSIPParticipantTopic(sipCallId string) {
	s.rpc.DeregisterHandler("TransferSIPParticipant", []string{sipCallId})
}

func (s *sIPInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *sIPInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor10 = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0x1e, 0x5a, 0x89, 0x6c, 0xad, 0xac, 0x1f, 0xc3, 0x96, 0x82, 0x28, 0xd3, 0x44, 0x55, 0x9a,
	0x19, 0xa5, 0x07, 0xaa, 0x55, 0xa6, 0x33, 0x9d, 0x1c, 0x3a, 0x4d, 0x1c, 0x67, 0xa2, 0x43, 0x5b,
	0x8f, 0xac, 0x5e, 0x7a, 0xe1, 0x40, 0x24, 0x2c, 0xa3, 0x26, 0x09, 0x14, 0x00, 0xe3, 0xa8, 0x6f,
	0x90, 0x43, 0xdf, 0xa1, 0xcf, 0x90, 0xd7, 0xe8, 0x23, 0xf5, 0xd2, 0x21, 0x00, 0x3a, 0xb4, 0x24,
	0xbb, 0xf6, 0x8d, 0xfb, 0x7d, 0xbb, 0x9f, 0x16, 0x0b, 0xec, 0x27, 0x68, 0x48, 0x11, 0x8e, 0x14,
	0x13, 0xbe, 0x90, 0x5c, 0x73, 0x54, 0x91, 0x22, 0xec, 0x3d, 0x5a, 0x70, 0xbe, 0x88, 0xe9, 0xc8,
	0x40, 0xf3, 0xec, 0x74, 0x44, 0x13, 0xa1, 0x97, 0x36, 0xa3, 0xf7, 0x78, 0x95, 0x8c, 0x32, 0x49,
	0x34, 0xe3, 0xa9, 0xe3, 0x1b, 0x5c, 0xe4, 0x91, 0x72, 0xe1, 0x5e, 0xcc, 0xde, 0xd3, 0x73, 0xa6,
	0x83, 0xcb, 0xdf, 0x18, 0xfc, 0x03, 0xf0, 0x74, 0x92, 0x6a, 0x2a, 0x53, 0x12, 0x1f, 0x4a, 0x4a,
	0x34, 0x3d, 0x99, 0x1c, 0x1f, 0x13, 0xa9, 0x59, 0xc8, 0x04, 0x49, 0xf5, 0x94, 0xfe, 0x91, 0x51,
	0xa5, 0xd1, 0x17, 0x00, 0x42, 0xf2, 0xdf, 0x69, 0xa8, 0x03, 0x16, 0x61, 0xd4, 0xf7, 0x86, 0xb5,
	0x69, 0xcd, 0x21, 0x93, 0x08, 0x3d, 0x86, 0xba, 0x62, 0x22, 0x08, 0x49, 0x1c, 0xe7, 0x7c, 0xc3,
	0xf2, 0x8a, 0x89, 0x43, 0x12, 0xc7, 0x93, 0x08, 0xf5, 0x61, 0x37, 0xe7, 0xb5, 0xcc, 0xd2, 0xf3,
	0x3c, 0x61, 0xdf, 0x24, 0x80, 0x62, 0x62, 0x96, 0x43, 0x93, 0x08, 0x61, 0xd8, 0x26, 0x51, 0x24,
	0xa9, 0x52, 0x78, 0xcb, 0x90, 0x45, 0x88, 0x7a, 0xb0, 0x73, 0xc6, 0x95, 0x4e, 0x49, 0x42, 0xf1,
	0x81, 0xa1, 0x2e, 0x63, 0xf4, 0x02, 0x6a, 0x5a, 0x92, 0x54, 0x09, 0x2e, 0x35, 0x6e, 0xf7, 0xbd,
	0x61, 0x73, 0xdc, 0xf1, 0xdd, 0x29, 0xfd, 0x93, 0xc9, 0xf1, 0xac, 0x20, 0xa7, 0x9f, 0xf3, 0x50,
	0x17, 0xaa, 0x69, 0x96, 0xcc, 0xa9, 0xc4, 0x15, 0x23, 0xe7, 0x22, 0xf4, 0x00, 0xb6, 0xcd, 0x01,
	0x34, 0xc7, 0xf7, 0x2c, 0x91, 0x87, 0x33, 0x9e, 0x77, 0x90, 0xa9, 0x7c, 0x44, 0x09, 0xc5, 0xf7,
	0x6d, 0x07, 0x45, 0x9c, 0x73, 0x82, 0x28, 0x75, 0xc1, 0x65, 0x84, 0xab, 0x96, 0x2b, 0x62, 0xf4,
	0x08, 0x6a, 0x92, 0xf3, 0x24, 0x30, 0x85, 0xdb, 0x96, 0xcc, 0x81, 0x9f, 0xf3, 0xc2, 0x6f, 0xe1,
	0x40, 0x7c, 0x9e, 0x73, 0xc0, 0x22, 0x9a, 0x6a, 0xa6, 0x97, 0x78, 0xc7, 0xe4, 0xed, 0x97, 0xb8,
	0x89, 0xa3, 0xd0, 0x73, 0x68, 0x97, 0x4b, 0x8c, 0x6c, 0xd3, 0xa4, 0xb7, 0x4a, 0xf8, 0x26, 0xf5,
	0x84, 0x6a, 0x12, 0x11, 0x4d, 0x70, 0x6b, 0x4d, 0xfd, 0x27, 0x47, 0xa1, 0x3f, 0xa1, 0x5b, 0x2e,
	0x21, 0x5a, 0x4b, 0x36, 0xcf, 0x34, 0x55, 0x78, 0xaf, 0x5f, 0x19, 0xd6, 0xc7, 0x87, 0xbe, 0x14,
	0xa1, 0x7f, 0x8b, 0xc7, 0xe2, 0x97, 0xa0, 0x57, 0x97, 0x2a, 0x47, 0xa9, 0x96, 0xcb, 0x69, 0x47,
	0x6c, 0xe2, 0xd0, 0x01, 0xdc, 0xd7, 0xfc, 0x9c, 0xa6, 0xb8, 0x66, 0xfa, 0xb3, 0x01, 0xea, 0x40,
	0xf5, 0x42, 0x05, 0x99, 0x8c, 0x31, 0x58, 0xf8, 0x42, 0xfd, 0x2a, 0x63, 0x84, 0xe0, 0x5e, 0xa4,
	0x93, 0x53, 0x5c, 0x37, 0xa0, 0xf9, 0x46, 0x4f, 0xa1, 0x21, 0x62, 0xb2, 0x0c, 0x22, 0x46, 0x62,
	0xcd, 0x53, 0x8a, 0x77, 0xfb, 0xde, 0x70, 0x67, 0xba, 0x9b, 0x83, 0x6f, 0x1c, 0x86, 0x7e, 0x81,
	0xed, 0x33, 0x4a, 0x22, 0x2a, 0x15, 0xee, 0x98, 0x23, 0x7d, 0x77, 0xeb, 0x23, 0xbd, 0xb3, 0x75,
	0xf6, 0x10, 0x85, 0x0a, 0xca, 0xa0, 0xe3, 0x3e, 0x03, 0xcd, 0xcb, 0x13, 0xeb, 0x1a, 0xf9, 0x57,
	0x77, 0x95, 0x9f, 0xf1, 0xd5, 0x79, 0xed, 0x9f, 0xad, 0x33, 0xe8, 0x07, 0x68, 0xd3, 0x94, 0xcc,
	0x63, 0x1a, 0x05, 0xa7, 0x94, 0xe8, 0x4c, 0x52, 0x85, 0x1f, 0xf6, 0x2b, 0xc3, 0xe6, 0x78, 0xbf,
	0xfc, 0xf8, 0xdf, 0x5a, 0x6e, 0xda, 0x72, 0xc9, 0x2e, 0x56, 0xe8, 0x35, 0xb4, 0x24, 0x4b, 0x17,
	0x2c, 0x5d, 0x04, 0x9a, 0x25, 0x94, 0x67, 0x1a, 0x3f, 0xe8, 0x7b, 0xc3, 0xfa, 0xf8, 0xa1, 0x6f,
	0x0d, 0xc5, 0x2f, 0x0c, 0xc5, 0x7f, 0xe3, 0x0c, 0x65, 0xda, 0x74, 0x15, 0x33, 0x5b, 0x80, 0x8e,
	0x60, 0x2f, 0x21, 0x1f, 0xec, 0xc6, 0x17, 0xae, 0x83, 0xf1, 0xff, 0xa9, 0xb4, 0x12, 0xf2, 0x21,
	0xb7, 0x84, 0x02, 0xe8, 0xbd, 0x83, 0xde, 0xf5, 0xaf, 0x05, 0xb5, 0xa1, 0x72, 0x4e, 0x97, 0xd8,
	0x33, 0x17, 0x9d, 0x7f, 0xe6, 0x0f, 0xe5, 0x3d, 0x89, 0x33, 0xea, 0x4c, 0xc2, 0x06, 0x2f, 0xb7,
	0xbe, 0xf7, 0x7a, 0x2f, 0x61, 0xb7, 0x7c, 0x49, 0x77, 0xaa, 0x7d, 0x0b, 0xf8, 0xba, 0x1b, 0xb8,
	0x8b, 0xce, 0xe0, 0x6f, 0x0f, 0xbe, 0xba, 0xf9, 0xba, 0x95, 0xe0, 0xa9, 0xa2, 0xe8, 0x19, 0x34,
	0xaf, 0x2e, 0xbf, 0xd3, 0x6f, 0x5c, 0x59, 0xfb, 0x6b, 0x3d, 0x62, 0xeb, 0x7a, 0x8f, 0x58, 0x71,
	0xe2, 0xca, 0x8a, 0x13, 0x0f, 0xfe, 0xf2, 0xe0, 0x59, 0xd1, 0xa2, 0x71, 0xc7, 0x53, 0x2a, 0x37,
	0x5b, 0xfe, 0x8a, 0x92, 0xb7, 0xea, 0xe9, 0x4f, 0xa0, 0xae, 0x9d, 0x40, 0x6e, 0x99, 0xb6, 0x27,
	0x28, 0xa0, 0x19, 0x5f, 0xdf, 0xc9, 0xca, 0xfa, 0x4e, 0x8e, 0xff, 0xf5, 0xa0, 0x7e, 0x32, 0x39,
	0x2e, 0x5a, 0x42, 0x17, 0x70, 0xb0, 0x69, 0x72, 0x68, 0x78, 0xdb, 0x5d, 0xea, 0x3d, 0xbf, 0x45,
	0xa6, 0xbd, 0x86, 0x01, 0x7c, 0xfa, 0xe8, 0x55, 0xdb, 0xde, 0x8f, 0xde, 0x37, 0x1e, 0x52, 0xd0,
	0xdd, 0x3c, 0x0f, 0xf4, 0xf5, 0x15, 0xc1, 0x1b, 0x87, 0xd6, 0xeb, 0xae, 0xbd, 0xfd, 0xa3, 0xfc,
	0xff, 0x7a, 0xd0, 0xf9, 0xf4, 0xd1, 0xdb, 0x6b, 0x7b, 0xbd, 0x06, 0x2a, 0x0f, 0xf5, 0xf5, 0x97,
	0xbf, 0x3d, 0x59, 0x30, 0x7d, 0x96, 0xcd, 0xfd, 0x90, 0x27, 0x23, 0xb7, 0xbb, 0xf6, 0xef, 0x3c,
	0xe4, 0xf1, 0x48, 0x8a, 0x70, 0x5e, 0x35, 0xd1, 0x8b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0xdc, 0x1d, 0x35, 0x1d, 0x08, 0x00, 0x00,
}
