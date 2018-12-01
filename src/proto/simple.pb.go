// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/simple.proto

package simplepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Simple struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple             bool     `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList           []int32  `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_4285f8f75e6ba5cf, []int{0}
}

func (m *Simple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple.Unmarshal(m, b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return xxx_messageInfo_Simple.Size(m)
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Simple) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *Simple) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Simple) GetSampleList() []int32 {
	if m != nil {
		return m.SampleList
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple)(nil), "example.simple.Simple")
}

func init() { proto.RegisterFile("proto/simple.proto", fileDescriptor_4285f8f75e6ba5cf) }

var fileDescriptor_4285f8f75e6ba5cf = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0x73, 0x84, 0xf8, 0x52, 0x2b, 0x12,
	0xc1, 0x5c, 0x88, 0xa8, 0x52, 0x16, 0x17, 0x5b, 0x30, 0x98, 0x25, 0xc4, 0xc7, 0xc5, 0x94, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xcd, 0xc5, 0x99, 0x59,
	0x1c, 0x0f, 0x51, 0x26, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x91, 0x59, 0x0c, 0x55, 0x2c,
	0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b,
	0xc9, 0x73, 0x71, 0x17, 0x83, 0xcd, 0x8e, 0xcf, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x51, 0x60, 0xd6,
	0x60, 0x0d, 0xe2, 0x82, 0x08, 0xf9, 0x64, 0x16, 0x97, 0x38, 0x71, 0x45, 0x71, 0x40, 0x8c, 0x2b,
	0x48, 0x4a, 0x62, 0x03, 0x3b, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xba, 0x0e, 0x42, 0x87,
	0xa4, 0x00, 0x00, 0x00,
}