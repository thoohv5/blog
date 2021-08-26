// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package command

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

type CommandHTTPServer interface {
	Execute(context.Context, *ExecuteReq) (*ExecuteResp, error)
}

func RegisterCommandHTTPServer(s *http.Server, srv CommandHTTPServer) {
	r := s.Route("/")
	r.GET("/common/v1/command/{command}", _Command_Execute0_HTTP_Handler(srv))
}

func _Command_Execute0_HTTP_Handler(srv CommandHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExecuteReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/common.v1.command.Command/Execute")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Execute(ctx, req.(*ExecuteReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecuteResp)
		return ctx.Result(200, reply)
	}
}

type CommandHTTPClient interface {
	Execute(ctx context.Context, req *ExecuteReq, opts ...http.CallOption) (rsp *ExecuteResp, err error)
}

type CommandHTTPClientImpl struct {
	cc *http.Client
}

func NewCommandHTTPClient(client *http.Client) CommandHTTPClient {
	return &CommandHTTPClientImpl{client}
}

func (c *CommandHTTPClientImpl) Execute(ctx context.Context, in *ExecuteReq, opts ...http.CallOption) (*ExecuteResp, error) {
	var out ExecuteResp
	pattern := "/common/v1/command/{command}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/common.v1.command.Command/Execute"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}