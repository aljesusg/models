// Code generated by protoc-gen-go.
// source: requests.proto
// DO NOT EDIT!

/*
Package requests is a generated protocol buffer package.

It is generated from these files:
	requests.proto

It has these top-level messages:
	Empty
	JobRequest
	JobResponse
*/
package job

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Mensaje vacío ya qu eno esperamos ningún parametro de entrada
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Tipo requests para la respuesta
type JobRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Number int32  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *JobRequest) Reset()                    { *m = JobRequest{} }
func (m *JobRequest) String() string            { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()               {}
func (*JobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type JobResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *JobResponse) Reset()                    { *m = JobResponse{} }
func (m *JobResponse) String() string            { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()               {}
func (*JobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*JobRequest)(nil), "JobRequest")
	proto.RegisterType((*JobResponse)(nil), "JobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Job service

type JobClient interface {
	GetJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
}

type jobClient struct {
	cc *grpc.ClientConn
}

func NewJobClient(cc *grpc.ClientConn) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) GetJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := grpc.Invoke(ctx, "/Job/GetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobServer interface {
	GetJob(context.Context, *JobRequest) (*JobResponse, error)
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Job/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _Job_GetJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "requests.proto",
}

func init() { proto.RegisterFile("requests.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xca, 0x4f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0xf2,
	0xe0, 0xe2, 0xf2, 0xca, 0x4f, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0xa5,
	0xb9, 0x49, 0xa9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x92, 0x2c, 0x17,
	0x37, 0xd8, 0xa4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x74, 0xa3, 0x8c, 0x74, 0xb8, 0x98, 0xbd,
	0xf2, 0x93, 0x84, 0x54, 0xb9, 0xd8, 0xdc, 0x53, 0x4b, 0x40, 0x2c, 0x6e, 0x3d, 0x84, 0xc5, 0x52,
	0x3c, 0x7a, 0x48, 0x7a, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0xe0, 0x13, 0xa4, 0xb3, 0x00, 0x00, 0x00,
}
