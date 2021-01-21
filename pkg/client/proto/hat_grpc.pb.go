// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HatClient is the client API for Hat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HatClient interface {
	Get(ctx context.Context, in *Path, opts ...grpc.CallOption) (*Data, error)
	GetBulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *Path, opts ...grpc.CallOption) (*Empty, error)
}

type hatClient struct {
	cc grpc.ClientConnInterface
}

func NewHatClient(cc grpc.ClientConnInterface) HatClient {
	return &hatClient{cc}
}

func (c *hatClient) Get(ctx context.Context, in *Path, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/Hat/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hatClient) GetBulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error) {
	out := new(BulkResponse)
	err := c.cc.Invoke(ctx, "/Hat/GetBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hatClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Hat/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hatClient) Delete(ctx context.Context, in *Path, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Hat/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HatServer is the server API for Hat service.
// All implementations must embed UnimplementedHatServer
// for forward compatibility
type HatServer interface {
	Get(context.Context, *Path) (*Data, error)
	GetBulk(context.Context, *BulkRequest) (*BulkResponse, error)
	Set(context.Context, *SetRequest) (*Empty, error)
	Delete(context.Context, *Path) (*Empty, error)
	mustEmbedUnimplementedHatServer()
}

// UnimplementedHatServer must be embedded to have forward compatible implementations.
type UnimplementedHatServer struct {
}

func (UnimplementedHatServer) Get(context.Context, *Path) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHatServer) GetBulk(context.Context, *BulkRequest) (*BulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulk not implemented")
}
func (UnimplementedHatServer) Set(context.Context, *SetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedHatServer) Delete(context.Context, *Path) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHatServer) mustEmbedUnimplementedHatServer() {}

// UnsafeHatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HatServer will
// result in compilation errors.
type UnsafeHatServer interface {
	mustEmbedUnimplementedHatServer()
}

func RegisterHatServer(s grpc.ServiceRegistrar, srv HatServer) {
	s.RegisterService(&_Hat_serviceDesc, srv)
}

func _Hat_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Path)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HatServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hat/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HatServer).Get(ctx, req.(*Path))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hat_GetBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HatServer).GetBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hat/GetBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HatServer).GetBulk(ctx, req.(*BulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hat_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HatServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hat/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HatServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hat_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Path)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HatServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hat/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HatServer).Delete(ctx, req.(*Path))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hat",
	HandlerType: (*HatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Hat_Get_Handler,
		},
		{
			MethodName: "GetBulk",
			Handler:    _Hat_GetBulk_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Hat_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Hat_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hat.proto",
}