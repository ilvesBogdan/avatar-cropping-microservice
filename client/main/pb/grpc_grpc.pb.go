// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: grpc.proto

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (ImageService_UploadImageClient, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (ImageService_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[0], "/ImageService.ImageService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceUploadImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageService_UploadImageClient interface {
	Recv() (*ImageResponse, error)
	grpc.ClientStream
}

type imageServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *imageServiceUploadImageClient) Recv() (*ImageResponse, error) {
	m := new(ImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	UploadImage(*ImageRequest, ImageService_UploadImageServer) error
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) UploadImage(*ImageRequest, ImageService_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageServiceServer).UploadImage(m, &imageServiceUploadImageServer{stream})
}

type ImageService_UploadImageServer interface {
	Send(*ImageResponse) error
	grpc.ServerStream
}

type imageServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *imageServiceUploadImageServer) Send(m *ImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageService.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _ImageService_UploadImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
