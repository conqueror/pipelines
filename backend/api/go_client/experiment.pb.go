// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: experiment.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

type CreateExperimentRequest struct {
	Experiment           *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateExperimentRequest) Reset()         { *m = CreateExperimentRequest{} }
func (m *CreateExperimentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateExperimentRequest) ProtoMessage()    {}
func (*CreateExperimentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{0}
}

func (m *CreateExperimentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExperimentRequest.Unmarshal(m, b)
}
func (m *CreateExperimentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExperimentRequest.Marshal(b, m, deterministic)
}
func (m *CreateExperimentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExperimentRequest.Merge(m, src)
}
func (m *CreateExperimentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateExperimentRequest.Size(m)
}
func (m *CreateExperimentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExperimentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExperimentRequest proto.InternalMessageInfo

func (m *CreateExperimentRequest) GetExperiment() *Experiment {
	if m != nil {
		return m.Experiment
	}
	return nil
}

type GetExperimentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExperimentRequest) Reset()         { *m = GetExperimentRequest{} }
func (m *GetExperimentRequest) String() string { return proto.CompactTextString(m) }
func (*GetExperimentRequest) ProtoMessage()    {}
func (*GetExperimentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{1}
}

func (m *GetExperimentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExperimentRequest.Unmarshal(m, b)
}
func (m *GetExperimentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExperimentRequest.Marshal(b, m, deterministic)
}
func (m *GetExperimentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExperimentRequest.Merge(m, src)
}
func (m *GetExperimentRequest) XXX_Size() int {
	return xxx_messageInfo_GetExperimentRequest.Size(m)
}
func (m *GetExperimentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExperimentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExperimentRequest proto.InternalMessageInfo

func (m *GetExperimentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListExperimentsRequest struct {
	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Can be format of "field_name", "field_name asc" or "field_name des"
	// Ascending by default.
	SortBy string `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	// A base-64 encoded, JSON-serialized Filter protocol buffer (see
	// filter.proto).
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListExperimentsRequest) Reset()         { *m = ListExperimentsRequest{} }
func (m *ListExperimentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListExperimentsRequest) ProtoMessage()    {}
func (*ListExperimentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{2}
}

func (m *ListExperimentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExperimentsRequest.Unmarshal(m, b)
}
func (m *ListExperimentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExperimentsRequest.Marshal(b, m, deterministic)
}
func (m *ListExperimentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExperimentsRequest.Merge(m, src)
}
func (m *ListExperimentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListExperimentsRequest.Size(m)
}
func (m *ListExperimentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExperimentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListExperimentsRequest proto.InternalMessageInfo

func (m *ListExperimentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListExperimentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListExperimentsRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *ListExperimentsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListExperimentsResponse struct {
	Experiments          []*Experiment `protobuf:"bytes,1,rep,name=experiments,proto3" json:"experiments,omitempty"`
	NextPageToken        string        `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListExperimentsResponse) Reset()         { *m = ListExperimentsResponse{} }
func (m *ListExperimentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListExperimentsResponse) ProtoMessage()    {}
func (*ListExperimentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{3}
}

func (m *ListExperimentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExperimentsResponse.Unmarshal(m, b)
}
func (m *ListExperimentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExperimentsResponse.Marshal(b, m, deterministic)
}
func (m *ListExperimentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExperimentsResponse.Merge(m, src)
}
func (m *ListExperimentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListExperimentsResponse.Size(m)
}
func (m *ListExperimentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExperimentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListExperimentsResponse proto.InternalMessageInfo

func (m *ListExperimentsResponse) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

func (m *ListExperimentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type DeleteExperimentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExperimentRequest) Reset()         { *m = DeleteExperimentRequest{} }
func (m *DeleteExperimentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteExperimentRequest) ProtoMessage()    {}
func (*DeleteExperimentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{4}
}

func (m *DeleteExperimentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExperimentRequest.Unmarshal(m, b)
}
func (m *DeleteExperimentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExperimentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteExperimentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExperimentRequest.Merge(m, src)
}
func (m *DeleteExperimentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteExperimentRequest.Size(m)
}
func (m *DeleteExperimentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExperimentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExperimentRequest proto.InternalMessageInfo

func (m *DeleteExperimentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Experiment struct {
	// Output. Unique experiment ID. Generated by API server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required input field. Unique experiment name provided by user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional input field. Describing the purpose of the experiment
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output. The time that the experiment created.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Experiment) Reset()         { *m = Experiment{} }
func (m *Experiment) String() string { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()    {}
func (*Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_7daedc28b4b25757, []int{5}
}

func (m *Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiment.Unmarshal(m, b)
}
func (m *Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiment.Marshal(b, m, deterministic)
}
func (m *Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiment.Merge(m, src)
}
func (m *Experiment) XXX_Size() int {
	return xxx_messageInfo_Experiment.Size(m)
}
func (m *Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_Experiment proto.InternalMessageInfo

func (m *Experiment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Experiment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Experiment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Experiment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateExperimentRequest)(nil), "api.CreateExperimentRequest")
	proto.RegisterType((*GetExperimentRequest)(nil), "api.GetExperimentRequest")
	proto.RegisterType((*ListExperimentsRequest)(nil), "api.ListExperimentsRequest")
	proto.RegisterType((*ListExperimentsResponse)(nil), "api.ListExperimentsResponse")
	proto.RegisterType((*DeleteExperimentRequest)(nil), "api.DeleteExperimentRequest")
	proto.RegisterType((*Experiment)(nil), "api.Experiment")
}

func init() { proto.RegisterFile("experiment.proto", fileDescriptor_7daedc28b4b25757) }

var fileDescriptor_7daedc28b4b25757 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xd9, 0xe9, 0x2f, 0x6d, 0x27, 0xa4, 0x7f, 0x16, 0x94, 0xa4, 0x4e, 0x4a, 0x8d, 0x0f,
	0xa1, 0x20, 0x62, 0x2b, 0xe1, 0x04, 0xb7, 0x06, 0x2a, 0x24, 0x04, 0x52, 0x95, 0xf4, 0xc4, 0x25,
	0xda, 0xc4, 0x13, 0xb3, 0x22, 0xb1, 0x8d, 0x77, 0x52, 0x9a, 0x20, 0x2e, 0x48, 0x1c, 0xb9, 0xc0,
	0xd7, 0xe2, 0xc6, 0x57, 0xe0, 0x83, 0x20, 0x6f, 0x9c, 0xc6, 0x75, 0x12, 0x71, 0x4a, 0x66, 0xe6,
	0xed, 0xbc, 0x99, 0xb7, 0xcf, 0x0b, 0x07, 0x78, 0x1d, 0x62, 0x24, 0xc6, 0xe8, 0x93, 0x1d, 0x46,
	0x01, 0x05, 0x2c, 0xc7, 0x43, 0x61, 0xd4, 0xbc, 0x20, 0xf0, 0x46, 0xe8, 0xf0, 0x50, 0x38, 0xdc,
	0xf7, 0x03, 0xe2, 0x24, 0x02, 0x5f, 0xce, 0x21, 0x46, 0x35, 0xa9, 0xaa, 0xa8, 0x3f, 0x19, 0x3a,
	0x38, 0x0e, 0x69, 0x9a, 0x14, 0x4f, 0xb2, 0x45, 0x12, 0x63, 0x94, 0xc4, 0xc7, 0x61, 0x02, 0x78,
	0xa2, 0x7e, 0x06, 0x0d, 0x0f, 0xfd, 0x86, 0xfc, 0xc4, 0x3d, 0x0f, 0x23, 0x27, 0x08, 0x55, 0xff,
	0x55, 0x2e, 0xeb, 0x35, 0x94, 0x5f, 0x44, 0xc8, 0x09, 0xcf, 0x6f, 0x06, 0xed, 0xe0, 0xc7, 0x09,
	0x4a, 0x62, 0x0e, 0xc0, 0x72, 0xfa, 0x8a, 0x66, 0x6a, 0xa7, 0x85, 0xd6, 0xbe, 0xcd, 0x43, 0x61,
	0xa7, 0xb0, 0x29, 0x88, 0x55, 0x87, 0x7b, 0xaf, 0x90, 0x56, 0x1b, 0xed, 0x81, 0x2e, 0x5c, 0xd5,
	0x60, 0xb7, 0xa3, 0x0b, 0xd7, 0xfa, 0xa6, 0x41, 0xe9, 0x8d, 0x90, 0x29, 0xa4, 0x5c, 0x40, 0x8f,
	0x01, 0x42, 0xee, 0x61, 0x8f, 0x82, 0x0f, 0xe8, 0x27, 0x47, 0x76, 0xe3, 0xcc, 0x65, 0x9c, 0x60,
	0x55, 0x50, 0x41, 0x4f, 0x8a, 0x19, 0x56, 0x74, 0x53, 0x3b, 0xfd, 0xbf, 0xb3, 0x13, 0x27, 0xba,
	0x62, 0x86, 0xac, 0x0c, 0xdb, 0x32, 0x88, 0xa8, 0xd7, 0x9f, 0x56, 0x72, 0xea, 0x60, 0x3e, 0x0e,
	0xdb, 0x53, 0x56, 0x82, 0xfc, 0x50, 0x8c, 0x08, 0xa3, 0xca, 0xd6, 0x3c, 0x3f, 0x8f, 0x2c, 0x82,
	0xf2, 0xca, 0x18, 0x32, 0x0c, 0x7c, 0x89, 0xac, 0x09, 0x85, 0xe5, 0x62, 0xb2, 0xa2, 0x99, 0xb9,
	0x75, 0xcb, 0xa7, 0x31, 0xac, 0x0e, 0xfb, 0x3e, 0x5e, 0x53, 0x2f, 0x35, 0xbf, 0xae, 0xe8, 0x8a,
	0x71, 0xfa, 0x62, 0xb1, 0x83, 0xf5, 0x08, 0xca, 0x2f, 0x71, 0x84, 0xeb, 0x14, 0xcf, 0x0a, 0xf5,
	0x5d, 0x03, 0x58, 0xa2, 0xb2, 0x65, 0xc6, 0x60, 0xcb, 0xe7, 0x63, 0x4c, 0x68, 0xd4, 0x7f, 0x66,
	0x42, 0xc1, 0x45, 0x39, 0x88, 0x84, 0xba, 0xf1, 0x44, 0x88, 0x74, 0x8a, 0x3d, 0x03, 0x18, 0xa8,
	0x1b, 0x77, 0x7b, 0x9c, 0x94, 0x22, 0x85, 0x96, 0x61, 0xcf, 0x5d, 0x65, 0x2f, 0x5c, 0x65, 0x5f,
	0x2e, 0x5c, 0xd5, 0xd9, 0x4d, 0xd0, 0x67, 0xd4, 0xfa, 0x95, 0x83, 0xc3, 0xe5, 0x3c, 0x5d, 0x8c,
	0xae, 0xc4, 0x00, 0x59, 0x08, 0x07, 0x59, 0x0b, 0xb1, 0x9a, 0x92, 0x6a, 0x83, 0xb3, 0x8c, 0xac,
	0x90, 0x56, 0xe3, 0xeb, 0xef, 0x3f, 0x3f, 0xf5, 0x87, 0xd6, 0x51, 0xfc, 0x45, 0x48, 0xe7, 0xaa,
	0xd9, 0x47, 0xe2, 0x4d, 0x27, 0x25, 0xef, 0xf3, 0x94, 0xd1, 0xd8, 0x00, 0x8a, 0xb7, 0x8c, 0xc6,
	0x8e, 0x54, 0xc3, 0x75, 0xe6, 0x5b, 0xe5, 0xaa, 0x2b, 0x2e, 0x93, 0xdd, 0xdf, 0xc8, 0xe5, 0x7c,
	0x16, 0xee, 0x17, 0xe6, 0xc3, 0xde, 0x6d, 0x77, 0xb0, 0xaa, 0x6a, 0xb5, 0xde, 0xb9, 0x46, 0x6d,
	0x7d, 0x71, 0xee, 0x27, 0xeb, 0x81, 0x22, 0xad, 0xb2, 0xcd, 0x0b, 0xc6, 0x32, 0x66, 0x7d, 0x91,
	0xc8, 0xb8, 0xc1, 0x2e, 0x46, 0x69, 0xe5, 0xd6, 0xce, 0xe3, 0x87, 0x62, 0xb1, 0xe1, 0xe3, 0x7f,
	0x6c, 0xd8, 0xbe, 0xf8, 0x71, 0xf6, 0xb6, 0x53, 0x83, 0x6d, 0x17, 0x87, 0x7c, 0x32, 0x22, 0x76,
	0xc8, 0xf6, 0xa1, 0x68, 0x14, 0x14, 0x67, 0x97, 0x38, 0x4d, 0xe4, 0xbb, 0x13, 0x38, 0x86, 0x7c,
	0x1b, 0x79, 0x84, 0x11, 0xbb, 0xbb, 0xa3, 0x1b, 0x45, 0x3e, 0xa1, 0xf7, 0x41, 0x24, 0x66, 0xea,
	0x11, 0x31, 0xf5, 0xfe, 0x1d, 0x80, 0x1b, 0xc0, 0x7f, 0xfd, 0xbc, 0x9a, 0xe4, 0xe9, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xd6, 0xbe, 0xa2, 0xf7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExperimentServiceClient is the client API for ExperimentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExperimentServiceClient interface {
	CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*Experiment, error)
	GetExperiment(ctx context.Context, in *GetExperimentRequest, opts ...grpc.CallOption) (*Experiment, error)
	ListExperiment(ctx context.Context, in *ListExperimentsRequest, opts ...grpc.CallOption) (*ListExperimentsResponse, error)
	DeleteExperiment(ctx context.Context, in *DeleteExperimentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type experimentServiceClient struct {
	cc *grpc.ClientConn
}

func NewExperimentServiceClient(cc *grpc.ClientConn) ExperimentServiceClient {
	return &experimentServiceClient{cc}
}

func (c *experimentServiceClient) CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*Experiment, error) {
	out := new(Experiment)
	err := c.cc.Invoke(ctx, "/api.ExperimentService/CreateExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) GetExperiment(ctx context.Context, in *GetExperimentRequest, opts ...grpc.CallOption) (*Experiment, error) {
	out := new(Experiment)
	err := c.cc.Invoke(ctx, "/api.ExperimentService/GetExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) ListExperiment(ctx context.Context, in *ListExperimentsRequest, opts ...grpc.CallOption) (*ListExperimentsResponse, error) {
	out := new(ListExperimentsResponse)
	err := c.cc.Invoke(ctx, "/api.ExperimentService/ListExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) DeleteExperiment(ctx context.Context, in *DeleteExperimentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ExperimentService/DeleteExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentServiceServer is the server API for ExperimentService service.
type ExperimentServiceServer interface {
	CreateExperiment(context.Context, *CreateExperimentRequest) (*Experiment, error)
	GetExperiment(context.Context, *GetExperimentRequest) (*Experiment, error)
	ListExperiment(context.Context, *ListExperimentsRequest) (*ListExperimentsResponse, error)
	DeleteExperiment(context.Context, *DeleteExperimentRequest) (*empty.Empty, error)
}

func RegisterExperimentServiceServer(s *grpc.Server, srv ExperimentServiceServer) {
	s.RegisterService(&_ExperimentService_serviceDesc, srv)
}

func _ExperimentService_CreateExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).CreateExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExperimentService/CreateExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).CreateExperiment(ctx, req.(*CreateExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_GetExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).GetExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExperimentService/GetExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).GetExperiment(ctx, req.(*GetExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_ListExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).ListExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExperimentService/ListExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).ListExperiment(ctx, req.(*ListExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_DeleteExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).DeleteExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ExperimentService/DeleteExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).DeleteExperiment(ctx, req.(*DeleteExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExperimentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ExperimentService",
	HandlerType: (*ExperimentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExperiment",
			Handler:    _ExperimentService_CreateExperiment_Handler,
		},
		{
			MethodName: "GetExperiment",
			Handler:    _ExperimentService_GetExperiment_Handler,
		},
		{
			MethodName: "ListExperiment",
			Handler:    _ExperimentService_ListExperiment_Handler,
		},
		{
			MethodName: "DeleteExperiment",
			Handler:    _ExperimentService_DeleteExperiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "experiment.proto",
}
