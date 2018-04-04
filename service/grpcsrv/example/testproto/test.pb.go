// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package testproto is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	LocationHeader
	CalculateStrLenRequest
	CalculateStrLenResponse
	AddRequest
	AddResponse
	AccountHeader
	DepositRequest
	DepositResponse
	ReportRequest
	DeleteUserRequest
*/
package testproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocationHeader struct {
	City string `protobuf:"bytes,1,opt,name=city" json:"city,omitempty"`
}

func (m *LocationHeader) Reset()                    { *m = LocationHeader{} }
func (m *LocationHeader) String() string            { return proto.CompactTextString(m) }
func (*LocationHeader) ProtoMessage()               {}
func (*LocationHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocationHeader) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type CalculateStrLenRequest struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
}

func (m *CalculateStrLenRequest) Reset()                    { *m = CalculateStrLenRequest{} }
func (m *CalculateStrLenRequest) String() string            { return proto.CompactTextString(m) }
func (*CalculateStrLenRequest) ProtoMessage()               {}
func (*CalculateStrLenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CalculateStrLenRequest) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type CalculateStrLenResponse struct {
	Len int32 `protobuf:"varint,1,opt,name=len" json:"len,omitempty"`
}

func (m *CalculateStrLenResponse) Reset()                    { *m = CalculateStrLenResponse{} }
func (m *CalculateStrLenResponse) String() string            { return proto.CompactTextString(m) }
func (*CalculateStrLenResponse) ProtoMessage()               {}
func (*CalculateStrLenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CalculateStrLenResponse) GetLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}

type AddRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Sum int32 `protobuf:"varint,1,opt,name=sum" json:"sum,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddResponse) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type AccountHeader struct {
	Account  string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AccountHeader) Reset()                    { *m = AccountHeader{} }
func (m *AccountHeader) String() string            { return proto.CompactTextString(m) }
func (*AccountHeader) ProtoMessage()               {}
func (*AccountHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AccountHeader) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountHeader) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 入金请求，需要附带鉴权头部AccountHeader
type DepositRequest struct {
	Money float64 `protobuf:"fixed64,1,opt,name=money" json:"money,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DepositRequest) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

// 入金响应
type DepositResponse struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DepositResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ReportRequest struct {
}

func (m *ReportRequest) Reset()                    { *m = ReportRequest{} }
func (m *ReportRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()               {}
func (*ReportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DeleteUserRequest struct {
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*LocationHeader)(nil), "testproto.LocationHeader")
	proto.RegisterType((*CalculateStrLenRequest)(nil), "testproto.CalculateStrLenRequest")
	proto.RegisterType((*CalculateStrLenResponse)(nil), "testproto.CalculateStrLenResponse")
	proto.RegisterType((*AddRequest)(nil), "testproto.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "testproto.AddResponse")
	proto.RegisterType((*AccountHeader)(nil), "testproto.AccountHeader")
	proto.RegisterType((*DepositRequest)(nil), "testproto.DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "testproto.DepositResponse")
	proto.RegisterType((*ReportRequest)(nil), "testproto.ReportRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "testproto.DeleteUserRequest")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xa9, 0x73, 0x6a, 0xaf, 0x6e, 0xd3, 0x2a, 0x5a, 0x44, 0x50, 0x82, 0xc8, 0x50, 0xd0,
	0x07, 0x7f, 0xc1, 0x70, 0x82, 0x0f, 0x7b, 0x8a, 0xf8, 0x03, 0xd2, 0xf6, 0x3e, 0x14, 0xda, 0xdc,
	0x98, 0x7b, 0x8b, 0xf8, 0xef, 0xa5, 0x69, 0xba, 0x3d, 0xf8, 0x76, 0xce, 0xe9, 0xc7, 0xe9, 0x49,
	0x02, 0x20, 0xc8, 0xf2, 0xec, 0x3c, 0x09, 0x65, 0x69, 0xaf, 0x83, 0x54, 0xf7, 0x30, 0xdf, 0x50,
	0x69, 0xa4, 0x26, 0xfb, 0x81, 0xa6, 0x42, 0x9f, 0x65, 0xb0, 0x5f, 0xd6, 0xf2, 0x9b, 0x27, 0x77,
	0xc9, 0x32, 0xd5, 0x41, 0xab, 0x47, 0xb8, 0x7c, 0x33, 0x4d, 0xd9, 0x35, 0x46, 0xf0, 0x53, 0xfc,
	0x06, 0xad, 0xc6, 0xef, 0x0e, 0x59, 0xb2, 0x53, 0x98, 0xb0, 0xf8, 0x08, 0xf7, 0x52, 0x3d, 0xc1,
	0xd5, 0x3f, 0x96, 0x1d, 0x59, 0xc6, 0x1e, 0x6e, 0xd0, 0x06, 0x78, 0xaa, 0x7b, 0xa9, 0x96, 0x00,
	0xab, 0xaa, 0x1a, 0xcb, 0x4e, 0x20, 0x31, 0xf1, 0x6b, 0x62, 0x7a, 0x57, 0xe4, 0x7b, 0x83, 0x2b,
	0xd4, 0x2d, 0x1c, 0x07, 0x72, 0x57, 0xc5, 0x5d, 0x3b, 0x56, 0x71, 0xd7, 0xaa, 0x77, 0x98, 0xad,
	0xca, 0x92, 0x3a, 0x2b, 0xf1, 0x20, 0x39, 0x1c, 0x9a, 0x21, 0x88, 0xf3, 0x46, 0x9b, 0x5d, 0xc3,
	0x91, 0x33, 0xcc, 0x3f, 0xe4, 0xab, 0xf0, 0x83, 0x54, 0x6f, 0xbd, 0x7a, 0x80, 0xf9, 0x1a, 0x1d,
	0x71, 0x2d, 0xe3, 0xaa, 0x0b, 0x98, 0xb6, 0x64, 0x71, 0xb8, 0x91, 0x44, 0x0f, 0x46, 0xbd, 0xc0,
	0x62, 0xcb, 0xc5, 0x4d, 0x37, 0x90, 0x4a, 0xdd, 0x22, 0x8b, 0x69, 0x5d, 0x80, 0x27, 0x7a, 0x17,
	0xa8, 0x05, 0xcc, 0x34, 0x3a, 0xf2, 0x63, 0xaf, 0x3a, 0x87, 0xb3, 0x35, 0x36, 0x28, 0xf8, 0xc5,
	0xe8, 0x63, 0x58, 0x1c, 0x84, 0x67, 0x79, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x55, 0xfc, 0x19,
	0x7e, 0xaf, 0x01, 0x00, 0x00,
}