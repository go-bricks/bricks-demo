// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package subworkshop

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SubWorkshopClient is the client API for SubWorkshop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubWorkshopClient interface {
	PaintCar(ctx context.Context, in *SubPaintCarRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type subWorkshopClient struct {
	cc grpc.ClientConnInterface
}

func NewSubWorkshopClient(cc grpc.ClientConnInterface) SubWorkshopClient {
	return &subWorkshopClient{cc}
}

func (c *subWorkshopClient) PaintCar(ctx context.Context, in *SubPaintCarRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/demo.subworkshop.SubWorkshop/PaintCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubWorkshopServer is the server API for SubWorkshop service.
// All implementations must embed UnimplementedSubWorkshopServer
// for forward compatibility
type SubWorkshopServer interface {
	PaintCar(context.Context, *SubPaintCarRequest) (*empty.Empty, error)
	mustEmbedUnimplementedSubWorkshopServer()
}

// UnimplementedSubWorkshopServer must be embedded to have forward compatible implementations.
type UnimplementedSubWorkshopServer struct {
}

func (UnimplementedSubWorkshopServer) PaintCar(context.Context, *SubPaintCarRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaintCar not implemented")
}
func (UnimplementedSubWorkshopServer) mustEmbedUnimplementedSubWorkshopServer() {}

// UnsafeSubWorkshopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubWorkshopServer will
// result in compilation errors.
type UnsafeSubWorkshopServer interface {
	mustEmbedUnimplementedSubWorkshopServer()
}

func RegisterSubWorkshopServer(s *grpc.Server, srv SubWorkshopServer) {
	s.RegisterService(&_SubWorkshop_serviceDesc, srv)
}

func _SubWorkshop_PaintCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubPaintCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubWorkshopServer).PaintCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.subworkshop.SubWorkshop/PaintCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubWorkshopServer).PaintCar(ctx, req.(*SubPaintCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubWorkshop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.subworkshop.SubWorkshop",
	HandlerType: (*SubWorkshopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PaintCar",
			Handler:    _SubWorkshop_PaintCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/subworkshop.proto",
}
