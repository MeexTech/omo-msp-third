// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/partner.proto

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

// Client API for PartnerService service

type PartnerService interface {
	AddOne(ctx context.Context, in *ReqPartnerAdd, opts ...client.CallOption) (*ReplyPartnerInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPartnerInfo, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyPartnerList, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	CreateSecret(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPartnerSecret, error)
}

type partnerService struct {
	c    client.Client
	name string
}

func NewPartnerService(name string, c client.Client) PartnerService {
	return &partnerService{
		c:    c,
		name: name,
	}
}

func (c *partnerService) AddOne(ctx context.Context, in *ReqPartnerAdd, opts ...client.CallOption) (*ReplyPartnerInfo, error) {
	req := c.c.NewRequest(c.name, "PartnerService.AddOne", in)
	out := new(ReplyPartnerInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPartnerInfo, error) {
	req := c.c.NewRequest(c.name, "PartnerService.GetOne", in)
	out := new(ReplyPartnerInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyPartnerList, error) {
	req := c.c.NewRequest(c.name, "PartnerService.GetList", in)
	out := new(ReplyPartnerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PartnerService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerService) CreateSecret(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPartnerSecret, error) {
	req := c.c.NewRequest(c.name, "PartnerService.CreateSecret", in)
	out := new(ReplyPartnerSecret)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PartnerService service

type PartnerServiceHandler interface {
	AddOne(context.Context, *ReqPartnerAdd, *ReplyPartnerInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyPartnerInfo) error
	GetList(context.Context, *RequestPage, *ReplyPartnerList) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	CreateSecret(context.Context, *RequestInfo, *ReplyPartnerSecret) error
}

func RegisterPartnerServiceHandler(s server.Server, hdlr PartnerServiceHandler, opts ...server.HandlerOption) error {
	type partnerService interface {
		AddOne(ctx context.Context, in *ReqPartnerAdd, out *ReplyPartnerInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyPartnerInfo) error
		GetList(ctx context.Context, in *RequestPage, out *ReplyPartnerList) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		CreateSecret(ctx context.Context, in *RequestInfo, out *ReplyPartnerSecret) error
	}
	type PartnerService struct {
		partnerService
	}
	h := &partnerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PartnerService{h}, opts...))
}

type partnerServiceHandler struct {
	PartnerServiceHandler
}

func (h *partnerServiceHandler) AddOne(ctx context.Context, in *ReqPartnerAdd, out *ReplyPartnerInfo) error {
	return h.PartnerServiceHandler.AddOne(ctx, in, out)
}

func (h *partnerServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyPartnerInfo) error {
	return h.PartnerServiceHandler.GetOne(ctx, in, out)
}

func (h *partnerServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplyPartnerList) error {
	return h.PartnerServiceHandler.GetList(ctx, in, out)
}

func (h *partnerServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.PartnerServiceHandler.RemoveOne(ctx, in, out)
}

func (h *partnerServiceHandler) CreateSecret(ctx context.Context, in *RequestInfo, out *ReplyPartnerSecret) error {
	return h.PartnerServiceHandler.CreateSecret(ctx, in, out)
}
