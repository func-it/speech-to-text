// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SpeechToText_Ping_FullMethodName         = "/proto.SpeechToText/Ping"
	SpeechToText_SpeechToText_FullMethodName = "/proto.SpeechToText/SpeechToText"
)

// SpeechToTextClient is the client API for SpeechToText service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeechToTextClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SpeechToText(ctx context.Context, in *SpeechToTextRequest, opts ...grpc.CallOption) (*SpeechToTextResponse, error)
}

type speechToTextClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeechToTextClient(cc grpc.ClientConnInterface) SpeechToTextClient {
	return &speechToTextClient{cc}
}

func (c *speechToTextClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SpeechToText_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *speechToTextClient) SpeechToText(ctx context.Context, in *SpeechToTextRequest, opts ...grpc.CallOption) (*SpeechToTextResponse, error) {
	out := new(SpeechToTextResponse)
	err := c.cc.Invoke(ctx, SpeechToText_SpeechToText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpeechToTextServer is the server API for SpeechToText service.
// All implementations should embed UnimplementedSpeechToTextServer
// for forward compatibility
type SpeechToTextServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	SpeechToText(context.Context, *SpeechToTextRequest) (*SpeechToTextResponse, error)
}

// UnimplementedSpeechToTextServer should be embedded to have forward compatible implementations.
type UnimplementedSpeechToTextServer struct {
}

func (UnimplementedSpeechToTextServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSpeechToTextServer) SpeechToText(context.Context, *SpeechToTextRequest) (*SpeechToTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpeechToText not implemented")
}

// UnsafeSpeechToTextServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeechToTextServer will
// result in compilation errors.
type UnsafeSpeechToTextServer interface {
	mustEmbedUnimplementedSpeechToTextServer()
}

func RegisterSpeechToTextServer(s grpc.ServiceRegistrar, srv SpeechToTextServer) {
	s.RegisterService(&SpeechToText_ServiceDesc, srv)
}

func _SpeechToText_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeechToTextServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpeechToText_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeechToTextServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpeechToText_SpeechToText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeechToTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeechToTextServer).SpeechToText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpeechToText_SpeechToText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeechToTextServer).SpeechToText(ctx, req.(*SpeechToTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpeechToText_ServiceDesc is the grpc.ServiceDesc for SpeechToText service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeechToText_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SpeechToText",
	HandlerType: (*SpeechToTextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SpeechToText_Ping_Handler,
		},
		{
			MethodName: "SpeechToText",
			Handler:    _SpeechToText_SpeechToText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
