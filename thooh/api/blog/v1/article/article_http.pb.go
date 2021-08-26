// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package article

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

type ArticleHTTPServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Detail(context.Context, *DetailReq) (*DetailResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.POST("/blog/v1/article/create", _Article_Create1_HTTP_Handler(srv))
	r.GET("/blog/v1/article/list", _Article_List1_HTTP_Handler(srv))
	r.GET("/blog/v1/article/detail", _Article_Detail0_HTTP_Handler(srv))
}

func _Article_Create1_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.article.Article/Create")
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

func _Article_List1_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.article.Article/List")
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

func _Article_Detail0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.article.Article/Detail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*DetailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailResp)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	Create(ctx context.Context, req *CreateReq, opts ...http.CallOption) (rsp *CreateResp, err error)
	Detail(ctx context.Context, req *DetailReq, opts ...http.CallOption) (rsp *DetailResp, err error)
	List(ctx context.Context, req *ListReq, opts ...http.CallOption) (rsp *ListResp, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) Create(ctx context.Context, in *CreateReq, opts ...http.CallOption) (*CreateResp, error) {
	var out CreateResp
	pattern := "/blog/v1/article/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/blog.v1.article.Article/Create"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) Detail(ctx context.Context, in *DetailReq, opts ...http.CallOption) (*DetailResp, error) {
	var out DetailResp
	pattern := "/blog/v1/article/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.article.Article/Detail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) List(ctx context.Context, in *ListReq, opts ...http.CallOption) (*ListResp, error) {
	var out ListResp
	pattern := "/blog/v1/article/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.article.Article/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}