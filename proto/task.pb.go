// Code generated by protoc-gen-go.
// source: task.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	RunRequest
	RunResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RunResponse_StatusCode int32

const (
	RunResponse_UNKNOWN RunResponse_StatusCode = 0
	RunResponse_OK      RunResponse_StatusCode = 200
	RunResponse_CREATED RunResponse_StatusCode = 201
	RunResponse_FAILED  RunResponse_StatusCode = 500
)

var RunResponse_StatusCode_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "CREATED",
	500: "FAILED",
}
var RunResponse_StatusCode_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      200,
	"CREATED": 201,
	"FAILED":  500,
}

func (x RunResponse_StatusCode) String() string {
	return proto1.EnumName(RunResponse_StatusCode_name, int32(x))
}
func (RunResponse_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type RunRequest struct {
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto1.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RunRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type RunResponse struct {
	StatusCode RunResponse_StatusCode `protobuf:"varint,1,opt,name=statusCode,enum=proto.RunResponse_StatusCode" json:"statusCode,omitempty"`
}

func (m *RunResponse) Reset()                    { *m = RunResponse{} }
func (m *RunResponse) String() string            { return proto1.CompactTextString(m) }
func (*RunResponse) ProtoMessage()               {}
func (*RunResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RunResponse) GetStatusCode() RunResponse_StatusCode {
	if m != nil {
		return m.StatusCode
	}
	return RunResponse_UNKNOWN
}

func init() {
	proto1.RegisterType((*RunRequest)(nil), "proto.RunRequest")
	proto1.RegisterType((*RunResponse)(nil), "proto.RunResponse")
	proto1.RegisterEnum("proto.RunResponse_StatusCode", RunResponse_StatusCode_name, RunResponse_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Task service

type TaskClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := grpc.Invoke(ctx, "/proto.Task/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Task service

type TaskServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Task/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Task_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto1.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9,
	0x79, 0xc5, 0x10, 0x45, 0x4a, 0x2a, 0x5c, 0x5c, 0x41, 0xa5, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xb9, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x52, 0x37, 0x23, 0x17, 0x37, 0x58, 0x59, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x2d, 0x17, 0x57, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x73, 0x7e, 0x4a, 0x2a,
	0x58, 0x2d, 0x9f, 0x91, 0x2c, 0xc4, 0x44, 0x3d, 0x24, 0x75, 0x7a, 0xc1, 0x70, 0x45, 0x41, 0x48,
	0x1a, 0x94, 0x6c, 0xb9, 0xb8, 0x10, 0x32, 0x42, 0xdc, 0x5c, 0xec, 0xa1, 0x7e, 0xde, 0x7e, 0xfe,
	0xe1, 0x7e, 0x02, 0x0c, 0x42, 0xec, 0x5c, 0x4c, 0xfe, 0xde, 0x02, 0x27, 0x18, 0x85, 0x78, 0xb8,
	0xd8, 0x9d, 0x83, 0x5c, 0x1d, 0x43, 0x5c, 0x5d, 0x04, 0x4e, 0x32, 0x0a, 0x71, 0x73, 0xb1, 0xb9,
	0x39, 0x7a, 0xfa, 0xb8, 0xba, 0x08, 0x7c, 0x61, 0x36, 0xf2, 0xe5, 0x62, 0x09, 0x49, 0x2c, 0xce,
	0x16, 0x72, 0xe5, 0x62, 0x0e, 0x2a, 0xcd, 0x13, 0x12, 0x44, 0xb6, 0x18, 0xec, 0x0f, 0x29, 0x21,
	0x4c, 0xb7, 0x28, 0x89, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x50, 0x89, 0x47, 0xbf, 0xcc, 0x50,
	0x1f, 0x14, 0x4c, 0xfa, 0x45, 0xa5, 0x79, 0x56, 0x8c, 0x5a, 0x49, 0x6c, 0x60, 0xb5, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x7d, 0x77, 0x1f, 0x3c, 0x01, 0x00, 0x00,
}