// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/echo/echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	service/echo/echo.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
*/
package echo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type EchoResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "echo.EchoResponse")
}

func init() { proto.RegisterFile("service/echo/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x4d, 0xce, 0xc8, 0x07, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x2c, 0x20, 0xb6, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6,
	0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x8d, 0x94, 0x6e,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e,
	0x58, 0x38, 0xa9, 0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0xa8, 0x72, 0x69, 0xa8, 0x61, 0x70,
	0x55, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x50, 0x49, 0x79, 0x74, 0xc9, 0x92, 0xcc, 0xdc, 0xd4, 0xe2,
	0x92, 0xc4, 0xdc, 0x02, 0x88, 0x02, 0x25, 0x45, 0x2e, 0x6e, 0xd7, 0xe4, 0x8c, 0xfc, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x89, 0x8b, 0x07, 0xa2, 0xa4, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x15, 0x9b, 0x1a, 0xa3, 0xa9, 0x8c, 0x5c, 0x2c, 0x20, 0x45, 0x42, 0x3e, 0x5c, 0x2c, 0x01,
	0x99, 0x79, 0xe9, 0x42, 0x62, 0x7a, 0x10, 0x9b, 0xf5, 0x60, 0x36, 0xeb, 0xb9, 0x82, 0x9c, 0x25,
	0x25, 0x85, 0x21, 0x1e, 0x02, 0x73, 0x91, 0x92, 0x40, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb8, 0x94,
	0x58, 0xf5, 0x0b, 0x32, 0xf3, 0xd2, 0xad, 0x18, 0xb5, 0x84, 0xec, 0xa1, 0xa6, 0x0a, 0xea, 0x81,
	0xc3, 0x10, 0xc9, 0xa5, 0x52, 0x42, 0xc8, 0x42, 0x10, 0x97, 0x21, 0x19, 0x00, 0x92, 0xb3, 0x62,
	0xd4, 0x72, 0xe2, 0x39, 0xf1, 0x48, 0x8e, 0xe1, 0xc2, 0x23, 0x39, 0x86, 0x07, 0x8f, 0xe4, 0x18,
	0x92, 0xd8, 0xc0, 0x96, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xde, 0x9c, 0x80, 0x02, 0x9f,
	0x01, 0x00, 0x00,
}
