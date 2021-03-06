// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.1

package login

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

type LoginHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterLoginHTTPServer(s *http.Server, srv LoginHTTPServer) {
	r := s.Route("/")
	r.GET("/login/", _Login_Login0_HTTP_Handler(srv))
}

func _Login_Login0_HTTP_Handler(srv LoginHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/protocol.login.Login/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type LoginHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type LoginHTTPClientImpl struct {
	cc *http.Client
}

func NewLoginHTTPClient(client *http.Client) LoginHTTPClient {
	return &LoginHTTPClientImpl{client}
}

func (c *LoginHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/login/"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/protocol.login.Login/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
