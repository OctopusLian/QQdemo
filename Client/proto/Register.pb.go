// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Register.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	Register.proto

It has these top-level messages:
	Register_ToS
	Register_ToC
*/
package proto

import proto1 "QQdemo/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Register_ToS struct {
	Uid              *string `protobuf:"bytes,1,req,name=uid" json:"uid,omitempty"`
	Username         *string `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Register_ToS) Reset()                    { *m = Register_ToS{} }
func (m *Register_ToS) String() string            { return proto1.CompactTextString(m) }
func (*Register_ToS) ProtoMessage()               {}
func (*Register_ToS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Register_ToS) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *Register_ToS) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Register_ToS) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type Register_ToC struct {
	Uid              *string `protobuf:"bytes,1,req,name=uid" json:"uid,omitempty"`
	Username         *string `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Register_ToC) Reset()                    { *m = Register_ToC{} }
func (m *Register_ToC) String() string            { return proto1.CompactTextString(m) }
func (*Register_ToC) ProtoMessage()               {}
func (*Register_ToC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Register_ToC) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *Register_ToC) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Register_ToC) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func init() {
	proto1.RegisterType((*Register_ToS)(nil), "proto.Register_ToS")
	proto1.RegisterType((*Register_ToC)(nil), "proto.Register_ToC")
}

func init() { proto1.RegisterFile("Register.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x11,
	0x5c, 0x3c, 0x30, 0x89, 0xf8, 0x90, 0xfc, 0x60, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x26, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x8a, 0x8b, 0xa3, 0xb4, 0x38, 0xb5, 0x28,
	0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x0c, 0xe7, 0x83, 0xe4, 0x0a, 0x12, 0x8b, 0x8b, 0xcb,
	0xf3, 0x8b, 0x52, 0x24, 0x98, 0x21, 0x72, 0x30, 0x3e, 0x9a, 0xc9, 0xce, 0xd4, 0x33, 0x19, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x31, 0x09, 0xed, 0x4b, 0xcb, 0x00, 0x00, 0x00,
}