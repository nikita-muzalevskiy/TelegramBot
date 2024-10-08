// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.0--rc1
// source: protos/pb.proto

package pb

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

// MangaClient is the client API for Manga service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MangaClient interface {
	Channel1(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackReply, error)
}

type mangaClient struct {
	cc grpc.ClientConnInterface
}

func NewMangaClient(cc grpc.ClientConnInterface) MangaClient {
	return &mangaClient{cc}
}

func (c *mangaClient) Channel1(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackReply, error) {
	out := new(CallbackReply)
	err := c.cc.Invoke(ctx, "/manga.Manga/Channel1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MangaServer is the server API for Manga service.
// All implementations must embed UnimplementedMangaServer
// for forward compatibility
type MangaServer interface {
	Channel1(context.Context, *CallbackRequest) (*CallbackReply, error)
	mustEmbedUnimplementedMangaServer()
}

// UnimplementedMangaServer must be embedded to have forward compatible implementations.
type UnimplementedMangaServer struct {
}

func (UnimplementedMangaServer) Channel1(context.Context, *CallbackRequest) (*CallbackReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Channel1 not implemented")
}
func (UnimplementedMangaServer) mustEmbedUnimplementedMangaServer() {}

// UnsafeMangaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MangaServer will
// result in compilation errors.
type UnsafeMangaServer interface {
	mustEmbedUnimplementedMangaServer()
}

func RegisterMangaServer(s grpc.ServiceRegistrar, srv MangaServer) {
	s.RegisterService(&Manga_ServiceDesc, srv)
}

func _Manga_Channel1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServer).Channel1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manga.Manga/Channel1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServer).Channel1(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manga_ServiceDesc is the grpc.ServiceDesc for Manga service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manga_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manga.Manga",
	HandlerType: (*MangaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Channel1",
			Handler:    _Manga_Channel1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pb.proto",
}
