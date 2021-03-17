// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protopb

import (
	context "context"
	"final-SA-Golang/music_microservice/cmd/web"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SongServiceClient is the client API for SongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongServiceClient interface {
	CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*CreateSongResponse, error)
	UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*UpdateSongResponse, error)
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error)
	GetSong(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*GetSongResponse, error)
	GetAllSongs(ctx context.Context, in *GetAllSongsRequest, opts ...grpc.CallOption) (SongService_GetAllSongsClient, error)
}

type songServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSongServiceClient(cc grpc.ClientConnInterface) SongServiceClient {
	return &songServiceClient{cc}
}

func (c *songServiceClient) CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*CreateSongResponse, error) {
	out := new(CreateSongResponse)
	err := c.cc.Invoke(ctx, "/protopb.SongService/CreateSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*UpdateSongResponse, error) {
	out := new(UpdateSongResponse)
	err := c.cc.Invoke(ctx, "/protopb.SongService/UpdateSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error) {
	out := new(DeleteSongResponse)
	err := c.cc.Invoke(ctx, "/protopb.SongService/DeleteSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) GetSong(ctx context.Context, in *GetSongRequest, opts ...grpc.CallOption) (*GetSongResponse, error) {
	out := new(GetSongResponse)
	err := c.cc.Invoke(ctx, "/protopb.SongService/GetSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) GetAllSongs(ctx context.Context, in *GetAllSongsRequest, opts ...grpc.CallOption) (SongService_GetAllSongsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SongService_ServiceDesc.Streams[0], "/protopb.SongService/GetAllSongs", opts...)
	if err != nil {
		return nil, err
	}
	x := &songServiceGetAllSongsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SongService_GetAllSongsClient interface {
	Recv() (*GetAllSongsResponse, error)
	grpc.ClientStream
}

type songServiceGetAllSongsClient struct {
	grpc.ClientStream
}

func (x *songServiceGetAllSongsClient) Recv() (*GetAllSongsResponse, error) {
	m := new(GetAllSongsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SongServiceServer is the server API for SongService service.
// All implementations must embed UnimplementedSongServiceServer
// for forward compatibility
type SongServiceServer interface {
	CreateSong(context.Context, *CreateSongRequest) (*CreateSongResponse, error)
	UpdateSong(context.Context, *UpdateSongRequest) (*UpdateSongResponse, error)
	DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error)
	GetSong(context.Context, *GetSongRequest) (*GetSongResponse, error)
	GetAllSongs(*GetAllSongsRequest, SongService_GetAllSongsServer) error
	mustEmbedUnimplementedSongServiceServer()
}

// UnimplementedSongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSongServiceServer struct {
}

func (UnimplementedSongServiceServer) CreateSong(context.Context, *CreateSongRequest) (*CreateSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSong not implemented")
}
func (UnimplementedSongServiceServer) UpdateSong(context.Context, *UpdateSongRequest) (*UpdateSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedSongServiceServer) DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedSongServiceServer) GetSong(context.Context, *GetSongRequest) (*GetSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSong not implemented")
}
func (UnimplementedSongServiceServer) GetAllSongs(*GetAllSongsRequest, SongService_GetAllSongsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllSongs not implemented")
}
func (UnimplementedSongServiceServer) mustEmbedUnimplementedSongServiceServer() {}

// UnsafeSongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongServiceServer will
// result in compilation errors.
type UnsafeSongServiceServer interface {
	mustEmbedUnimplementedSongServiceServer()
}

func RegisterSongServiceServer(s grpc.ServiceRegistrar, srv *main.Server) {
	s.RegisterService(&SongService_ServiceDesc, srv)
}

func _SongService_CreateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).CreateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopb.SongService/CreateSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).CreateSong(ctx, req.(*CreateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopb.SongService/UpdateSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).UpdateSong(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopb.SongService/DeleteSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_GetSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).GetSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopb.SongService/GetSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).GetSong(ctx, req.(*GetSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_GetAllSongs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllSongsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SongServiceServer).GetAllSongs(m, &songServiceGetAllSongsServer{stream})
}

type SongService_GetAllSongsServer interface {
	Send(*GetAllSongsResponse) error
	grpc.ServerStream
}

type songServiceGetAllSongsServer struct {
	grpc.ServerStream
}

func (x *songServiceGetAllSongsServer) Send(m *GetAllSongsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SongService_ServiceDesc is the grpc.ServiceDesc for SongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protopb.SongService",
	HandlerType: (*SongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSong",
			Handler:    _SongService_CreateSong_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _SongService_UpdateSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _SongService_DeleteSong_Handler,
		},
		{
			MethodName: "GetSong",
			Handler:    _SongService_GetSong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllSongs",
			Handler:       _SongService_GetAllSongs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "songpb/song.proto",
}
