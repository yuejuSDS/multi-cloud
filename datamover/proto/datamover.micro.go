// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: datamover.proto

/*
Package datamover is a generated protocol buffer package.

It is generated from these files:
	datamover.proto

It has these top-level messages:
	KV
	Filter
	Connector
	AsistInfo
	RunJobRequest
	RunJobResponse
*/
package datamover

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Datamover service

type DatamoverService interface {
	Runjob(ctx context.Context, in *RunJobRequest, opts ...client.CallOption) (*RunJobResponse, error)
}

type datamoverService struct {
	c    client.Client
	name string
}

func NewDatamoverService(name string, c client.Client) DatamoverService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "datamover"
	}
	return &datamoverService{
		c:    c,
		name: name,
	}
}

func (c *datamoverService) Runjob(ctx context.Context, in *RunJobRequest, opts ...client.CallOption) (*RunJobResponse, error) {
	req := c.c.NewRequest(c.name, "Datamover.Runjob", in)
	out := new(RunJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Datamover service

type DatamoverHandler interface {
	Runjob(context.Context, *RunJobRequest, *RunJobResponse) error
}

func RegisterDatamoverHandler(s server.Server, hdlr DatamoverHandler, opts ...server.HandlerOption) error {
	type datamover interface {
		Runjob(ctx context.Context, in *RunJobRequest, out *RunJobResponse) error
	}
	type Datamover struct {
		datamover
	}
	h := &datamoverHandler{hdlr}
	return s.Handle(s.NewHandler(&Datamover{h}, opts...))
}

type datamoverHandler struct {
	DatamoverHandler
}

func (h *datamoverHandler) Runjob(ctx context.Context, in *RunJobRequest, out *RunJobResponse) error {
	return h.DatamoverHandler.Runjob(ctx, in, out)
}
