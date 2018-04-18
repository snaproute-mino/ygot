// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openconfig/yang_protobuf/protoc-gen-yanggo/yang/pkg/testproto/test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	github.com/openconfig/yang_protobuf/protoc-gen-yanggo/yang/pkg/testproto/test.proto

It has these top-level messages:
	AMessage
	BMessage
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/openconfig/ygot/proto/yext"

import (
	"reflect"
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

type AMessage struct {
	Set    *BMessage          `protobuf:"bytes,1,opt,name=set" json:"set,omitempty"`
	Unset  *BMessage          `protobuf:"bytes,2,opt,name=unset" json:"unset,omitempty"`
	Nested *AMessage_CMessage `protobuf:"bytes,3,opt,name=nested" json:"nested,omitempty"`
}

func (m *AMessage) Reset()                    { *m = AMessage{} }
func (m *AMessage) String() string            { return proto.CompactTextString(m) }
func (*AMessage) ProtoMessage()               {}
func (*AMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AMessage) GetSet() *BMessage {
	if m != nil {
		return m.Set
	}
	return nil
}

func (m *AMessage) GetUnset() *BMessage {
	if m != nil {
		return m.Unset
	}
	return nil
}

func (m *AMessage) GetNested() *AMessage_CMessage {
	if m != nil {
		return m.Nested
	}
	return nil
}

type AMessage_CMessage struct {
	Nested *AMessage_CMessage_DMessage `protobuf:"bytes,4,opt,name=nested" json:"nested,omitempty"`
}

func (m *AMessage_CMessage) Reset()                    { *m = AMessage_CMessage{} }
func (m *AMessage_CMessage) String() string            { return proto.CompactTextString(m) }
func (*AMessage_CMessage) ProtoMessage()               {}
func (*AMessage_CMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AMessage_CMessage) GetNested() *AMessage_CMessage_DMessage {
	if m != nil {
		return m.Nested
	}
	return nil
}

type AMessage_CMessage_DMessage struct {
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
}

func (m *AMessage_CMessage_DMessage) Reset()         { *m = AMessage_CMessage_DMessage{} }
func (m *AMessage_CMessage_DMessage) String() string { return proto.CompactTextString(m) }
func (*AMessage_CMessage_DMessage) ProtoMessage()    {}
func (*AMessage_CMessage_DMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *AMessage_CMessage_DMessage) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type BMessage struct {
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
}

func (m *BMessage) Reset()                    { *m = BMessage{} }
func (m *BMessage) String() string            { return proto.CompactTextString(m) }
func (*BMessage) ProtoMessage()               {}
func (*BMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BMessage) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func init() {
	proto.RegisterType((*AMessage)(nil), "test.AMessage")
	proto.RegisterType((*AMessage_CMessage)(nil), "test.AMessage.CMessage")
	proto.RegisterType((*AMessage_CMessage_DMessage)(nil), "test.AMessage.CMessage.DMessage")
	proto.RegisterType((*BMessage)(nil), "test.BMessage")
}

var (
	YANGPathToProtoGoStruct = map[string]reflect.Type{
		"/b":     reflect.TypeOf(BMessage{}),
		"/b/c/d": reflect.TypeOf(AMessage_CMessage_DMessage{}),
		"/b/c":   reflect.TypeOf(AMessage_CMessage{}),
	}

	ProtoGoStructPathToFieldName = map[string]map[string]string{
		"AMessage": map[string]string{
			"/b":   "Set",
			"/b/c": "Nested",
		},
		"AMessage_CMessage": map[string]string{
			"/b/c/d": "Nested",
		},
		"AMessage_CMessage_DMessage": map[string]string{
			"/b/c/d/field": "Field",
		},
		"BMessage": map[string]string{},
	}
)

func init() {
	proto.RegisterFile("github.com/openconfig/yang_protobuf/protoc-gen-yanggo/yang/pkg/testproto/test.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0x84, 0x40,
	0x18, 0x87, 0xd1, 0x55, 0x73, 0xa7, 0x28, 0x18, 0x82, 0x96, 0x3d, 0xc9, 0x52, 0xd4, 0x65, 0x7d,
	0xfb, 0x73, 0xe9, 0xaa, 0xdb, 0xb5, 0x8b, 0x7d, 0x80, 0x58, 0xf5, 0x75, 0x92, 0x6a, 0x46, 0x72,
	0x84, 0xf6, 0xea, 0xa7, 0xe8, 0xe3, 0xc6, 0xbc, 0xe3, 0x2c, 0x04, 0xed, 0x45, 0x7e, 0xf8, 0x3c,
	0xf3, 0xc8, 0xc8, 0x5e, 0x44, 0xab, 0xdf, 0x86, 0x32, 0xad, 0xd4, 0x27, 0xa8, 0x0e, 0x65, 0xa5,
	0x64, 0xd3, 0x0a, 0xd8, 0x6d, 0xa5, 0x78, 0xed, 0xbe, 0x94, 0x56, 0xe5, 0xd0, 0x00, 0x8d, 0x6a,
	0x2d, 0x50, 0xae, 0x0d, 0x10, 0x8a, 0x38, 0x74, 0xef, 0x02, 0x34, 0xf6, 0x9a, 0x30, 0xad, 0x94,
	0x26, 0x0f, 0xcc, 0x5e, 0xde, 0x1e, 0x48, 0x0b, 0xa5, 0x6d, 0x11, 0x76, 0xf8, 0xad, 0xe9, 0x61,
	0xcf, 0xad, 0x7e, 0x7c, 0x16, 0x67, 0xcf, 0xd8, 0xf7, 0x5b, 0x81, 0xfc, 0x9a, 0xcd, 0x7a, 0xd4,
	0x0b, 0x2f, 0xf1, 0x6e, 0x8e, 0xef, 0x4f, 0x53, 0xca, 0xe7, 0x13, 0xcc, 0xc3, 0x31, 0xf3, 0xa1,
	0x2c, 0x8c, 0xc1, 0x2f, 0x59, 0x38, 0x48, 0xa3, 0xfa, 0xff, 0xa9, 0x85, 0x85, 0xfc, 0x91, 0x45,
	0x12, 0x7b, 0x8d, 0xf5, 0x62, 0x46, 0xda, 0x85, 0xd5, 0xdc, 0xe7, 0xd2, 0x8d, 0x4b, 0x1f, 0x8d,
	0x59, 0x00, 0x25, 0x54, 0xc5, 0xe4, 0x2f, 0x47, 0x8f, 0xc5, 0x8e, 0xf2, 0xcd, 0x3e, 0x13, 0x50,
	0x26, 0x39, 0x90, 0x49, 0x9f, 0x5c, 0x6f, 0x3e, 0x66, 0x91, 0xe9, 0x41, 0xbd, 0x2f, 0xde, 0xb1,
	0xd8, 0x61, 0x7e, 0xc5, 0xc2, 0xa6, 0xc5, 0x8f, 0x9a, 0x2e, 0x3a, 0xcf, 0xcf, 0xc6, 0xec, 0xc4,
	0xda, 0x40, 0xaf, 0x0b, 0x4b, 0x57, 0x09, 0x8b, 0xdd, 0x8d, 0xf8, 0xf9, 0x9f, 0x23, 0x93, 0x51,
	0x46, 0xf4, 0x0f, 0x1f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x80, 0x55, 0x05, 0xd2, 0x01,
	0x00, 0x00,
}
