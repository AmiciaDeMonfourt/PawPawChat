// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: s3.proto

package s3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	S3Service_SinglePartUpload_FullMethodName = "/s3.s3Service/SinglePartUpload"
	S3Service_StreamUpload_FullMethodName     = "/s3.s3Service/StreamUpload"
	S3Service_GetUploadURL_FullMethodName     = "/s3.s3Service/GetUploadURL"
	S3Service_GetDownloadURL_FullMethodName   = "/s3.s3Service/GetDownloadURL"
)

// S3ServiceClient is the client API for S3Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3ServiceClient interface {
	SinglePartUpload(ctx context.Context, in *SinglePartUploadRequest, opts ...grpc.CallOption) (*SinglePartUploadResponse, error)
	StreamUpload(ctx context.Context, opts ...grpc.CallOption) (S3Service_StreamUploadClient, error)
	GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error)
	GetDownloadURL(ctx context.Context, in *GetDownloadRequest, opts ...grpc.CallOption) (*GetDownloadResponse, error)
}

type s3ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewS3ServiceClient(cc grpc.ClientConnInterface) S3ServiceClient {
	return &s3ServiceClient{cc}
}

func (c *s3ServiceClient) SinglePartUpload(ctx context.Context, in *SinglePartUploadRequest, opts ...grpc.CallOption) (*SinglePartUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SinglePartUploadResponse)
	err := c.cc.Invoke(ctx, S3Service_SinglePartUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ServiceClient) StreamUpload(ctx context.Context, opts ...grpc.CallOption) (S3Service_StreamUploadClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &S3Service_ServiceDesc.Streams[0], S3Service_StreamUpload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &s3ServiceStreamUploadClient{ClientStream: stream}
	return x, nil
}

type S3Service_StreamUploadClient interface {
	Send(*StreamUploadRequest) error
	CloseAndRecv() (*StreamUploadResponse, error)
	grpc.ClientStream
}

type s3ServiceStreamUploadClient struct {
	grpc.ClientStream
}

func (x *s3ServiceStreamUploadClient) Send(m *StreamUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *s3ServiceStreamUploadClient) CloseAndRecv() (*StreamUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *s3ServiceClient) GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadURLResponse)
	err := c.cc.Invoke(ctx, S3Service_GetUploadURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3ServiceClient) GetDownloadURL(ctx context.Context, in *GetDownloadRequest, opts ...grpc.CallOption) (*GetDownloadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadResponse)
	err := c.cc.Invoke(ctx, S3Service_GetDownloadURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3ServiceServer is the server API for S3Service service.
// All implementations must embed UnimplementedS3ServiceServer
// for forward compatibility
type S3ServiceServer interface {
	SinglePartUpload(context.Context, *SinglePartUploadRequest) (*SinglePartUploadResponse, error)
	StreamUpload(S3Service_StreamUploadServer) error
	GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error)
	GetDownloadURL(context.Context, *GetDownloadRequest) (*GetDownloadResponse, error)
	mustEmbedUnimplementedS3ServiceServer()
}

// UnimplementedS3ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedS3ServiceServer struct {
}

func (UnimplementedS3ServiceServer) SinglePartUpload(context.Context, *SinglePartUploadRequest) (*SinglePartUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinglePartUpload not implemented")
}
func (UnimplementedS3ServiceServer) StreamUpload(S3Service_StreamUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUpload not implemented")
}
func (UnimplementedS3ServiceServer) GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadURL not implemented")
}
func (UnimplementedS3ServiceServer) GetDownloadURL(context.Context, *GetDownloadRequest) (*GetDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadURL not implemented")
}
func (UnimplementedS3ServiceServer) mustEmbedUnimplementedS3ServiceServer() {}

// UnsafeS3ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3ServiceServer will
// result in compilation errors.
type UnsafeS3ServiceServer interface {
	mustEmbedUnimplementedS3ServiceServer()
}

func RegisterS3ServiceServer(s grpc.ServiceRegistrar, srv S3ServiceServer) {
	s.RegisterService(&S3Service_ServiceDesc, srv)
}

func _S3Service_SinglePartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePartUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ServiceServer).SinglePartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3Service_SinglePartUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ServiceServer).SinglePartUpload(ctx, req.(*SinglePartUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Service_StreamUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(S3ServiceServer).StreamUpload(&s3ServiceStreamUploadServer{ServerStream: stream})
}

type S3Service_StreamUploadServer interface {
	SendAndClose(*StreamUploadResponse) error
	Recv() (*StreamUploadRequest, error)
	grpc.ServerStream
}

type s3ServiceStreamUploadServer struct {
	grpc.ServerStream
}

func (x *s3ServiceStreamUploadServer) SendAndClose(m *StreamUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *s3ServiceStreamUploadServer) Recv() (*StreamUploadRequest, error) {
	m := new(StreamUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _S3Service_GetUploadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ServiceServer).GetUploadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3Service_GetUploadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ServiceServer).GetUploadURL(ctx, req.(*GetUploadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Service_GetDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3ServiceServer).GetDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3Service_GetDownloadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3ServiceServer).GetDownloadURL(ctx, req.(*GetDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S3Service_ServiceDesc is the grpc.ServiceDesc for S3Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "s3.s3Service",
	HandlerType: (*S3ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SinglePartUpload",
			Handler:    _S3Service_SinglePartUpload_Handler,
		},
		{
			MethodName: "GetUploadURL",
			Handler:    _S3Service_GetUploadURL_Handler,
		},
		{
			MethodName: "GetDownloadURL",
			Handler:    _S3Service_GetDownloadURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUpload",
			Handler:       _S3Service_StreamUpload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "s3.proto",
}
