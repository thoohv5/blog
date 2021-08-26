// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package category

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

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryClient interface {
	// 创建.
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	// 列表.
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/blog.v1.category.Category/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/blog.v1.category.Category/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
// All implementations must embed UnimplementedCategoryServer
// for forward compatibility
type CategoryServer interface {
	// 创建.
	Create(context.Context, *CreateReq) (*CreateResp, error)
	// 列表.
	List(context.Context, *ListReq) (*ListResp, error)
	mustEmbedUnimplementedCategoryServer()
}

// UnimplementedCategoryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (UnimplementedCategoryServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCategoryServer) mustEmbedUnimplementedCategoryServer() {}

// UnsafeCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServer will
// result in compilation errors.
type UnsafeCategoryServer interface {
	mustEmbedUnimplementedCategoryServer()
}

func RegisterCategoryServer(s grpc.ServiceRegistrar, srv CategoryServer) {
	s.RegisterService(&Category_ServiceDesc, srv)
}

func _Category_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.category.Category/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.category.Category/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Category_ServiceDesc is the grpc.ServiceDesc for Category service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Category_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.v1.category.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Category_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Category_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog/v1/category/category.proto",
}