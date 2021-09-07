// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package category

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CategoryHTTPServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
}

func RegisterCategoryHTTPServer(s *http.Server, srv CategoryHTTPServer) {
	r := s.Route("/")
	r.POST("/blog/v1/category/create", _Category_Create0_HTTP_Handler(srv))
	r.GET("/blog/v1/category/list", _Category_List0_HTTP_Handler(srv))
}

func _Category_Create0_HTTP_Handler(srv CategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.category.Category/Create")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CreateReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResp)
		return ctx.Result(200, reply)
	}
}

func _Category_List0_HTTP_Handler(srv CategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.category.Category/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResp)
		return ctx.Result(200, reply)
	}
}

type CategoryHTTPClient interface {
	Create(ctx context.Context, req *CreateReq, opts ...http.CallOption) (rsp *CreateResp, err error)
	List(ctx context.Context, req *ListReq, opts ...http.CallOption) (rsp *ListResp, err error)
}

type CategoryHTTPClientImpl struct {
	cc *http.Client
}

func NewCategoryHTTPClient(client *http.Client) CategoryHTTPClient {
	return &CategoryHTTPClientImpl{client}
}

func (c *CategoryHTTPClientImpl) Create(ctx context.Context, in *CreateReq, opts ...http.CallOption) (*CreateResp, error) {
	var out CreateResp
	pattern := "/blog/v1/category/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/blog.v1.category.Category/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryHTTPClientImpl) List(ctx context.Context, in *ListReq, opts ...http.CallOption) (*ListResp, error) {
	var out ListResp
	pattern := "/blog/v1/category/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.category.Category/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
