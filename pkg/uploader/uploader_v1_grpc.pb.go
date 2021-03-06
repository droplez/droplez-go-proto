// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package uploader

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

// UploaderClient is the client API for Uploader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploaderClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error)
	GetDownloadableLink(ctx context.Context, in *UploadedFileData, opts ...grpc.CallOption) (*DownloadableLink, error)
}

type uploaderClient struct {
	cc grpc.ClientConnInterface
}

func NewUploaderClient(cc grpc.ClientConnInterface) UploaderClient {
	return &uploaderClient{cc}
}

func (c *uploaderClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Uploader_ServiceDesc.Streams[0], "/uploader.Uploader/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploaderUploadClient{stream}
	return x, nil
}

type Uploader_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadedFileData, error)
	grpc.ClientStream
}

type uploaderUploadClient struct {
	grpc.ClientStream
}

func (x *uploaderUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploaderUploadClient) CloseAndRecv() (*UploadedFileData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadedFileData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploaderClient) GetDownloadableLink(ctx context.Context, in *UploadedFileData, opts ...grpc.CallOption) (*DownloadableLink, error) {
	out := new(DownloadableLink)
	err := c.cc.Invoke(ctx, "/uploader.Uploader/GetDownloadableLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploaderServer is the server API for Uploader service.
// All implementations must embed UnimplementedUploaderServer
// for forward compatibility
type UploaderServer interface {
	Upload(Uploader_UploadServer) error
	GetDownloadableLink(context.Context, *UploadedFileData) (*DownloadableLink, error)
	mustEmbedUnimplementedUploaderServer()
}

// UnimplementedUploaderServer must be embedded to have forward compatible implementations.
type UnimplementedUploaderServer struct {
}

func (UnimplementedUploaderServer) Upload(Uploader_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploaderServer) GetDownloadableLink(context.Context, *UploadedFileData) (*DownloadableLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadableLink not implemented")
}
func (UnimplementedUploaderServer) mustEmbedUnimplementedUploaderServer() {}

// UnsafeUploaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploaderServer will
// result in compilation errors.
type UnsafeUploaderServer interface {
	mustEmbedUnimplementedUploaderServer()
}

func RegisterUploaderServer(s grpc.ServiceRegistrar, srv UploaderServer) {
	s.RegisterService(&Uploader_ServiceDesc, srv)
}

func _Uploader_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploaderServer).Upload(&uploaderUploadServer{stream})
}

type Uploader_UploadServer interface {
	SendAndClose(*UploadedFileData) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type uploaderUploadServer struct {
	grpc.ServerStream
}

func (x *uploaderUploadServer) SendAndClose(m *UploadedFileData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploaderUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Uploader_GetDownloadableLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadedFileData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploaderServer).GetDownloadableLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uploader.Uploader/GetDownloadableLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploaderServer).GetDownloadableLink(ctx, req.(*UploadedFileData))
	}
	return interceptor(ctx, in, info, handler)
}

// Uploader_ServiceDesc is the grpc.ServiceDesc for Uploader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Uploader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uploader.Uploader",
	HandlerType: (*UploaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDownloadableLink",
			Handler:    _Uploader_GetDownloadableLink_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Uploader_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "uploader/uploader_v1.proto",
}
