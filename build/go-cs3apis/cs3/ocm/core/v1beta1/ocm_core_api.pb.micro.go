// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cs3/ocm/core/v1beta1/ocm_core_api.proto

package corev1beta1

import (
	fmt "fmt"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/identity/user/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OcmCoreAPI service

func NewOcmCoreAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OcmCoreAPI service

type OcmCoreAPIService interface {
	// Creates a new ocm share.
	CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...client.CallOption) (*CreateOCMCoreShareResponse, error)
}

type ocmCoreAPIService struct {
	c    client.Client
	name string
}

func NewOcmCoreAPIService(name string, c client.Client) OcmCoreAPIService {
	return &ocmCoreAPIService{
		c:    c,
		name: name,
	}
}

func (c *ocmCoreAPIService) CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, opts ...client.CallOption) (*CreateOCMCoreShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmCoreAPI.CreateOCMCoreShare", in)
	out := new(CreateOCMCoreShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OcmCoreAPI service

type OcmCoreAPIHandler interface {
	// Creates a new ocm share.
	CreateOCMCoreShare(context.Context, *CreateOCMCoreShareRequest, *CreateOCMCoreShareResponse) error
}

func RegisterOcmCoreAPIHandler(s server.Server, hdlr OcmCoreAPIHandler, opts ...server.HandlerOption) error {
	type ocmCoreAPI interface {
		CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, out *CreateOCMCoreShareResponse) error
	}
	type OcmCoreAPI struct {
		ocmCoreAPI
	}
	h := &ocmCoreAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&OcmCoreAPI{h}, opts...))
}

type ocmCoreAPIHandler struct {
	OcmCoreAPIHandler
}

func (h *ocmCoreAPIHandler) CreateOCMCoreShare(ctx context.Context, in *CreateOCMCoreShareRequest, out *CreateOCMCoreShareResponse) error {
	return h.OcmCoreAPIHandler.CreateOCMCoreShare(ctx, in, out)
}