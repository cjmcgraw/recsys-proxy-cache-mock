// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package recsys_proxy_cache

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

// RecsysProxyCacheClient is the client API for RecsysProxyCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecsysProxyCacheClient interface {
	//*
	//Get scores retrieves a score for every item given,
	//from the associated context
	GetScores(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*ScoreResponse, error)
}

type recsysProxyCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewRecsysProxyCacheClient(cc grpc.ClientConnInterface) RecsysProxyCacheClient {
	return &recsysProxyCacheClient{cc}
}

func (c *recsysProxyCacheClient) GetScores(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*ScoreResponse, error) {
	out := new(ScoreResponse)
	err := c.cc.Invoke(ctx, "/recsys.RecsysProxyCache/GetScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecsysProxyCacheServer is the server API for RecsysProxyCache service.
// All implementations must embed UnimplementedRecsysProxyCacheServer
// for forward compatibility
type RecsysProxyCacheServer interface {
	//*
	//Get scores retrieves a score for every item given,
	//from the associated context
	GetScores(context.Context, *ScoreRequest) (*ScoreResponse, error)
	mustEmbedUnimplementedRecsysProxyCacheServer()
}

// UnimplementedRecsysProxyCacheServer must be embedded to have forward compatible implementations.
type UnimplementedRecsysProxyCacheServer struct {
}

func (UnimplementedRecsysProxyCacheServer) GetScores(context.Context, *ScoreRequest) (*ScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScores not implemented")
}
func (UnimplementedRecsysProxyCacheServer) mustEmbedUnimplementedRecsysProxyCacheServer() {}

// UnsafeRecsysProxyCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecsysProxyCacheServer will
// result in compilation errors.
type UnsafeRecsysProxyCacheServer interface {
	mustEmbedUnimplementedRecsysProxyCacheServer()
}

func RegisterRecsysProxyCacheServer(s grpc.ServiceRegistrar, srv RecsysProxyCacheServer) {
	s.RegisterService(&RecsysProxyCache_ServiceDesc, srv)
}

func _RecsysProxyCache_GetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecsysProxyCacheServer).GetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recsys.RecsysProxyCache/GetScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecsysProxyCacheServer).GetScores(ctx, req.(*ScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecsysProxyCache_ServiceDesc is the grpc.ServiceDesc for RecsysProxyCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecsysProxyCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recsys.RecsysProxyCache",
	HandlerType: (*RecsysProxyCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScores",
			Handler:    _RecsysProxyCache_GetScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recsys.proto",
}