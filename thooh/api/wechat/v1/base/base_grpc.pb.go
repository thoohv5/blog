// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package base

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

// WeChatClient is the client API for WeChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeChatClient interface {
	// 二维码.
	QRCode(ctx context.Context, in *QRCodeReq, opts ...grpc.CallOption) (*QRCodeResp, error)
	// 二维码结果
	CheckQRCode(ctx context.Context, in *CheckQRCodeReq, opts ...grpc.CallOption) (*CheckQRCodeResp, error)
}

type weChatClient struct {
	cc grpc.ClientConnInterface
}

func NewWeChatClient(cc grpc.ClientConnInterface) WeChatClient {
	return &weChatClient{cc}
}

func (c *weChatClient) QRCode(ctx context.Context, in *QRCodeReq, opts ...grpc.CallOption) (*QRCodeResp, error) {
	out := new(QRCodeResp)
	err := c.cc.Invoke(ctx, "/wechat.v1.base.WeChat/QRCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weChatClient) CheckQRCode(ctx context.Context, in *CheckQRCodeReq, opts ...grpc.CallOption) (*CheckQRCodeResp, error) {
	out := new(CheckQRCodeResp)
	err := c.cc.Invoke(ctx, "/wechat.v1.base.WeChat/CheckQRCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeChatServer is the server API for WeChat service.
// All implementations must embed UnimplementedWeChatServer
// for forward compatibility
type WeChatServer interface {
	// 二维码.
	QRCode(context.Context, *QRCodeReq) (*QRCodeResp, error)
	// 二维码结果
	CheckQRCode(context.Context, *CheckQRCodeReq) (*CheckQRCodeResp, error)
	mustEmbedUnimplementedWeChatServer()
}

// UnimplementedWeChatServer must be embedded to have forward compatible implementations.
type UnimplementedWeChatServer struct {
}

func (UnimplementedWeChatServer) QRCode(context.Context, *QRCodeReq) (*QRCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QRCode not implemented")
}
func (UnimplementedWeChatServer) CheckQRCode(context.Context, *CheckQRCodeReq) (*CheckQRCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckQRCode not implemented")
}
func (UnimplementedWeChatServer) mustEmbedUnimplementedWeChatServer() {}

// UnsafeWeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeChatServer will
// result in compilation errors.
type UnsafeWeChatServer interface {
	mustEmbedUnimplementedWeChatServer()
}

func RegisterWeChatServer(s grpc.ServiceRegistrar, srv WeChatServer) {
	s.RegisterService(&WeChat_ServiceDesc, srv)
}

func _WeChat_QRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QRCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeChatServer).QRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wechat.v1.base.WeChat/QRCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeChatServer).QRCode(ctx, req.(*QRCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeChat_CheckQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckQRCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeChatServer).CheckQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wechat.v1.base.WeChat/CheckQRCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeChatServer).CheckQRCode(ctx, req.(*CheckQRCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WeChat_ServiceDesc is the grpc.ServiceDesc for WeChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wechat.v1.base.WeChat",
	HandlerType: (*WeChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QRCode",
			Handler:    _WeChat_QRCode_Handler,
		},
		{
			MethodName: "CheckQRCode",
			Handler:    _WeChat_CheckQRCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/wechat/v1/base/base.proto",
}
