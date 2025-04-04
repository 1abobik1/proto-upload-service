// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/upload_service/v1/upload_service.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileService_Upload_FullMethodName          = "/upload_service.v1.FileService/Upload"
	FileService_UpdateFile_FullMethodName      = "/upload_service.v1.FileService/UpdateFile"
	FileService_ListFiles_FullMethodName       = "/upload_service.v1.FileService/ListFiles"
	FileService_GetDownloadLink_FullMethodName = "/upload_service.v1.FileService/GetDownloadLink"
	FileService_DownloadZip_FullMethodName     = "/upload_service.v1.FileService/DownloadZip"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	// загрузка файла на сервер
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error)
	// обновление файла, можно обновить название файла и сами данные
	UpdateFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpdateFileRequest, UpdateFileResponse], error)
	// получение списка файлов
	ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// генерация Presigned URL с помощью Minio для скачивания одного файла
	GetDownloadLink(ctx context.Context, in *DownloadLinkRequest, opts ...grpc.CallOption) (*DownloadLinkResponse, error)
	// скачивание нескольких файлов в zip-архиве
	DownloadZip(ctx context.Context, in *DownloadZipRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadZipResponse], error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadRequest, UploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadClient = grpc.ClientStreamingClient[UploadRequest, UploadResponse]

func (c *fileServiceClient) UpdateFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpdateFileRequest, UpdateFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], FileService_UpdateFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UpdateFileRequest, UpdateFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UpdateFileClient = grpc.ClientStreamingClient[UpdateFileRequest, UpdateFileResponse]

func (c *fileServiceClient) ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, FileService_ListFiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetDownloadLink(ctx context.Context, in *DownloadLinkRequest, opts ...grpc.CallOption) (*DownloadLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadLinkResponse)
	err := c.cc.Invoke(ctx, FileService_GetDownloadLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DownloadZip(ctx context.Context, in *DownloadZipRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadZipResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[2], FileService_DownloadZip_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadZipRequest, DownloadZipResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownloadZipClient = grpc.ServerStreamingClient[DownloadZipResponse]

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
type FileServiceServer interface {
	// загрузка файла на сервер
	Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error
	// обновление файла, можно обновить название файла и сами данные
	UpdateFile(grpc.ClientStreamingServer[UpdateFileRequest, UpdateFileResponse]) error
	// получение списка файлов
	ListFiles(context.Context, *ListRequest) (*ListResponse, error)
	// генерация Presigned URL с помощью Minio для скачивания одного файла
	GetDownloadLink(context.Context, *DownloadLinkRequest) (*DownloadLinkResponse, error)
	// скачивание нескольких файлов в zip-архиве
	DownloadZip(*DownloadZipRequest, grpc.ServerStreamingServer[DownloadZipResponse]) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileServiceServer) UpdateFile(grpc.ClientStreamingServer[UpdateFileRequest, UpdateFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileServiceServer) ListFiles(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileServiceServer) GetDownloadLink(context.Context, *DownloadLinkRequest) (*DownloadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadLink not implemented")
}
func (UnimplementedFileServiceServer) DownloadZip(*DownloadZipRequest, grpc.ServerStreamingServer[DownloadZipResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadZip not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).Upload(&grpc.GenericServerStream[UploadRequest, UploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadServer = grpc.ClientStreamingServer[UploadRequest, UploadResponse]

func _FileService_UpdateFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UpdateFile(&grpc.GenericServerStream[UpdateFileRequest, UpdateFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UpdateFileServer = grpc.ClientStreamingServer[UpdateFileRequest, UpdateFileResponse]

func _FileService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFiles(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetDownloadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetDownloadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetDownloadLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetDownloadLink(ctx, req.(*DownloadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DownloadZip_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadZipRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadZip(m, &grpc.GenericServerStream[DownloadZipRequest, DownloadZipResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownloadZipServer = grpc.ServerStreamingServer[DownloadZipResponse]

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upload_service.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileService_ListFiles_Handler,
		},
		{
			MethodName: "GetDownloadLink",
			Handler:    _FileService_GetDownloadLink_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateFile",
			Handler:       _FileService_UpdateFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadZip",
			Handler:       _FileService_DownloadZip_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/upload_service/v1/upload_service.proto",
}
