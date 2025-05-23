// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/link.proto

package third

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for LinkService service

type LinkService interface {
	AddOne(ctx context.Context, in *ReqLinkAdd, opts ...client.CallOption) (*ReplyLinkInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyLinkInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyLinkList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type linkService struct {
	c    client.Client
	name string
}

func NewLinkService(name string, c client.Client) LinkService {
	return &linkService{
		c:    c,
		name: name,
	}
}

func (c *linkService) AddOne(ctx context.Context, in *ReqLinkAdd, opts ...client.CallOption) (*ReplyLinkInfo, error) {
	req := c.c.NewRequest(c.name, "LinkService.AddOne", in)
	out := new(ReplyLinkInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyLinkInfo, error) {
	req := c.c.NewRequest(c.name, "LinkService.GetOne", in)
	out := new(ReplyLinkInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "LinkService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyLinkList, error) {
	req := c.c.NewRequest(c.name, "LinkService.GetByFilter", in)
	out := new(ReplyLinkList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "LinkService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "LinkService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LinkService service

type LinkServiceHandler interface {
	AddOne(context.Context, *ReqLinkAdd, *ReplyLinkInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyLinkInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyLinkList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterLinkServiceHandler(s server.Server, hdlr LinkServiceHandler, opts ...server.HandlerOption) error {
	type linkService interface {
		AddOne(ctx context.Context, in *ReqLinkAdd, out *ReplyLinkInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyLinkInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyLinkList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type LinkService struct {
		linkService
	}
	h := &linkServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LinkService{h}, opts...))
}

type linkServiceHandler struct {
	LinkServiceHandler
}

func (h *linkServiceHandler) AddOne(ctx context.Context, in *ReqLinkAdd, out *ReplyLinkInfo) error {
	return h.LinkServiceHandler.AddOne(ctx, in, out)
}

func (h *linkServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyLinkInfo) error {
	return h.LinkServiceHandler.GetOne(ctx, in, out)
}

func (h *linkServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.LinkServiceHandler.RemoveOne(ctx, in, out)
}

func (h *linkServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyLinkList) error {
	return h.LinkServiceHandler.GetByFilter(ctx, in, out)
}

func (h *linkServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.LinkServiceHandler.GetStatistic(ctx, in, out)
}

func (h *linkServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.LinkServiceHandler.UpdateByFilter(ctx, in, out)
}
