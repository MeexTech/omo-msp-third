// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/honor.proto

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

// Client API for HonorService service

type HonorService interface {
	AddOne(ctx context.Context, in *ReqHonorAdd, opts ...client.CallOption) (*ReplyHonorInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyHonorInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyHonorList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateBase(ctx context.Context, in *ReqHonorUpdate, opts ...client.CallOption) (*ReplyHonorInfo, error)
	UpdateStatus(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateContents(ctx context.Context, in *ReqHonorContents, opts ...client.CallOption) (*ReplyInfo, error)
}

type honorService struct {
	c    client.Client
	name string
}

func NewHonorService(name string, c client.Client) HonorService {
	return &honorService{
		c:    c,
		name: name,
	}
}

func (c *honorService) AddOne(ctx context.Context, in *ReqHonorAdd, opts ...client.CallOption) (*ReplyHonorInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.AddOne", in)
	out := new(ReplyHonorInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyHonorInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.GetOne", in)
	out := new(ReplyHonorInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyHonorList, error) {
	req := c.c.NewRequest(c.name, "HonorService.GetByFilter", in)
	out := new(ReplyHonorList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "HonorService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) UpdateBase(ctx context.Context, in *ReqHonorUpdate, opts ...client.CallOption) (*ReplyHonorInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.UpdateBase", in)
	out := new(ReplyHonorInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) UpdateStatus(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honorService) UpdateContents(ctx context.Context, in *ReqHonorContents, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HonorService.UpdateContents", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HonorService service

type HonorServiceHandler interface {
	AddOne(context.Context, *ReqHonorAdd, *ReplyHonorInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyHonorInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyHonorList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	UpdateBase(context.Context, *ReqHonorUpdate, *ReplyHonorInfo) error
	UpdateStatus(context.Context, *RequestFlag, *ReplyInfo) error
	UpdateContents(context.Context, *ReqHonorContents, *ReplyInfo) error
}

func RegisterHonorServiceHandler(s server.Server, hdlr HonorServiceHandler, opts ...server.HandlerOption) error {
	type honorService interface {
		AddOne(ctx context.Context, in *ReqHonorAdd, out *ReplyHonorInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyHonorInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyHonorList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		UpdateBase(ctx context.Context, in *ReqHonorUpdate, out *ReplyHonorInfo) error
		UpdateStatus(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		UpdateContents(ctx context.Context, in *ReqHonorContents, out *ReplyInfo) error
	}
	type HonorService struct {
		honorService
	}
	h := &honorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HonorService{h}, opts...))
}

type honorServiceHandler struct {
	HonorServiceHandler
}

func (h *honorServiceHandler) AddOne(ctx context.Context, in *ReqHonorAdd, out *ReplyHonorInfo) error {
	return h.HonorServiceHandler.AddOne(ctx, in, out)
}

func (h *honorServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyHonorInfo) error {
	return h.HonorServiceHandler.GetOne(ctx, in, out)
}

func (h *honorServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyHonorList) error {
	return h.HonorServiceHandler.GetByFilter(ctx, in, out)
}

func (h *honorServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.HonorServiceHandler.GetStatistic(ctx, in, out)
}

func (h *honorServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.HonorServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *honorServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.HonorServiceHandler.RemoveOne(ctx, in, out)
}

func (h *honorServiceHandler) UpdateBase(ctx context.Context, in *ReqHonorUpdate, out *ReplyHonorInfo) error {
	return h.HonorServiceHandler.UpdateBase(ctx, in, out)
}

func (h *honorServiceHandler) UpdateStatus(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.HonorServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *honorServiceHandler) UpdateContents(ctx context.Context, in *ReqHonorContents, out *ReplyInfo) error {
	return h.HonorServiceHandler.UpdateContents(ctx, in, out)
}
