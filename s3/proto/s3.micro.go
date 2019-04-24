// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: s3.proto

/*
Package s3 is a generated protocol buffer package.

It is generated from these files:
	s3.proto

It has these top-level messages:
	ACL
	ServerSideEncryption
	VersioningConfiguration
	RedirectAllRequestsTo
	Redirect
	Condition
	RoutingRules
	WebsiteConfiguration
	CORSConfiguration
	Destination
	ReplicationRole
	ReplicationConfiguration
	Tag
	LifecycleFilter
	Action
	LifecycleRule
	ReplicationInfo
	Bucket
	Partion
	Version
	Object
	ListBucketsResponse
	BaseResponse
	BaseRequest
	ListObjectsRequest
	ListObjectResponse
	DeleteObjectInput
	GetObjectInput
	MultipartUpload
	ListParts
	StorageClassDef
	NullRequest
	TList
	Tier2ClassName
	GetTierMapResponse
	UpdateObjMetaRequest
*/
package s3

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

// Client API for S3 service

type S3Service interface {
	ListBuckets(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*ListBucketsResponse, error)
	CreateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error)
	DeleteBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error)
	GetBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*Bucket, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*ListObjectResponse, error)
	CreateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error)
	UpdateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error)
	GetObject(ctx context.Context, in *GetObjectInput, opts ...client.CallOption) (*Object, error)
	DeleteObject(ctx context.Context, in *DeleteObjectInput, opts ...client.CallOption) (*BaseResponse, error)
	GetTierMap(ctx context.Context, in *NullRequest, opts ...client.CallOption) (*GetTierMapResponse, error)
	UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type s3Service struct {
	c    client.Client
	name string
}

func NewS3Service(name string, c client.Client) S3Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "s3"
	}
	return &s3Service{
		c:    c,
		name: name,
	}
}

func (c *s3Service) ListBuckets(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*ListBucketsResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListBuckets", in)
	out := new(ListBucketsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CreateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CreateBucket", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteBucket", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*Bucket, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucket", in)
	out := new(Bucket)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*ListObjectResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListObjects", in)
	out := new(ListObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CreateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CreateObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UpdateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UpdateObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetObject(ctx context.Context, in *GetObjectInput, opts ...client.CallOption) (*Object, error) {
	req := c.c.NewRequest(c.name, "S3.GetObject", in)
	out := new(Object)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteObject(ctx context.Context, in *DeleteObjectInput, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetTierMap(ctx context.Context, in *NullRequest, opts ...client.CallOption) (*GetTierMapResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetTierMap", in)
	out := new(GetTierMapResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UpdateObjMeta", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S3 service

type S3Handler interface {
	ListBuckets(context.Context, *BaseRequest, *ListBucketsResponse) error
	CreateBucket(context.Context, *Bucket, *BaseResponse) error
	DeleteBucket(context.Context, *Bucket, *BaseResponse) error
	GetBucket(context.Context, *Bucket, *Bucket) error
	ListObjects(context.Context, *ListObjectsRequest, *ListObjectResponse) error
	CreateObject(context.Context, *Object, *BaseResponse) error
	UpdateObject(context.Context, *Object, *BaseResponse) error
	GetObject(context.Context, *GetObjectInput, *Object) error
	DeleteObject(context.Context, *DeleteObjectInput, *BaseResponse) error
	GetTierMap(context.Context, *NullRequest, *GetTierMapResponse) error
	UpdateObjMeta(context.Context, *UpdateObjMetaRequest, *BaseResponse) error
}

func RegisterS3Handler(s server.Server, hdlr S3Handler, opts ...server.HandlerOption) error {
	type s3 interface {
		ListBuckets(ctx context.Context, in *BaseRequest, out *ListBucketsResponse) error
		CreateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error
		DeleteBucket(ctx context.Context, in *Bucket, out *BaseResponse) error
		GetBucket(ctx context.Context, in *Bucket, out *Bucket) error
		ListObjects(ctx context.Context, in *ListObjectsRequest, out *ListObjectResponse) error
		CreateObject(ctx context.Context, in *Object, out *BaseResponse) error
		UpdateObject(ctx context.Context, in *Object, out *BaseResponse) error
		GetObject(ctx context.Context, in *GetObjectInput, out *Object) error
		DeleteObject(ctx context.Context, in *DeleteObjectInput, out *BaseResponse) error
		GetTierMap(ctx context.Context, in *NullRequest, out *GetTierMapResponse) error
		UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, out *BaseResponse) error
	}
	type S3 struct {
		s3
	}
	h := &s3Handler{hdlr}
	return s.Handle(s.NewHandler(&S3{h}, opts...))
}

type s3Handler struct {
	S3Handler
}

func (h *s3Handler) ListBuckets(ctx context.Context, in *BaseRequest, out *ListBucketsResponse) error {
	return h.S3Handler.ListBuckets(ctx, in, out)
}

func (h *s3Handler) CreateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error {
	return h.S3Handler.CreateBucket(ctx, in, out)
}

func (h *s3Handler) DeleteBucket(ctx context.Context, in *Bucket, out *BaseResponse) error {
	return h.S3Handler.DeleteBucket(ctx, in, out)
}

func (h *s3Handler) GetBucket(ctx context.Context, in *Bucket, out *Bucket) error {
	return h.S3Handler.GetBucket(ctx, in, out)
}

func (h *s3Handler) ListObjects(ctx context.Context, in *ListObjectsRequest, out *ListObjectResponse) error {
	return h.S3Handler.ListObjects(ctx, in, out)
}

func (h *s3Handler) CreateObject(ctx context.Context, in *Object, out *BaseResponse) error {
	return h.S3Handler.CreateObject(ctx, in, out)
}

func (h *s3Handler) UpdateObject(ctx context.Context, in *Object, out *BaseResponse) error {
	return h.S3Handler.UpdateObject(ctx, in, out)
}

func (h *s3Handler) GetObject(ctx context.Context, in *GetObjectInput, out *Object) error {
	return h.S3Handler.GetObject(ctx, in, out)
}

func (h *s3Handler) DeleteObject(ctx context.Context, in *DeleteObjectInput, out *BaseResponse) error {
	return h.S3Handler.DeleteObject(ctx, in, out)
}

func (h *s3Handler) GetTierMap(ctx context.Context, in *NullRequest, out *GetTierMapResponse) error {
	return h.S3Handler.GetTierMap(ctx, in, out)
}

func (h *s3Handler) UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, out *BaseResponse) error {
	return h.S3Handler.UpdateObjMeta(ctx, in, out)
}
