// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/delivery/proto/delivery.proto

package proto

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

type DeliveryChatroomMessageCommand struct {
	RoomID               string   `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
	SenderID             string   `protobuf:"bytes,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	SenderFirstName      string   `protobuf:"bytes,3,opt,name=SenderFirstName,proto3" json:"SenderFirstName,omitempty"`
	SenderLastName       string   `protobuf:"bytes,4,opt,name=SenderLastName,proto3" json:"SenderLastName,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryChatroomMessageCommand) Reset()         { *m = DeliveryChatroomMessageCommand{} }
func (m *DeliveryChatroomMessageCommand) String() string { return proto.CompactTextString(m) }
func (*DeliveryChatroomMessageCommand) ProtoMessage()    {}
func (*DeliveryChatroomMessageCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_delivery_7c4d5b8e9471c36b, []int{0}
}
func (m *DeliveryChatroomMessageCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryChatroomMessageCommand.Unmarshal(m, b)
}
func (m *DeliveryChatroomMessageCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryChatroomMessageCommand.Marshal(b, m, deterministic)
}
func (dst *DeliveryChatroomMessageCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryChatroomMessageCommand.Merge(dst, src)
}
func (m *DeliveryChatroomMessageCommand) XXX_Size() int {
	return xxx_messageInfo_DeliveryChatroomMessageCommand.Size(m)
}
func (m *DeliveryChatroomMessageCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryChatroomMessageCommand.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryChatroomMessageCommand proto.InternalMessageInfo

func (m *DeliveryChatroomMessageCommand) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *DeliveryChatroomMessageCommand) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *DeliveryChatroomMessageCommand) GetSenderFirstName() string {
	if m != nil {
		return m.SenderFirstName
	}
	return ""
}

func (m *DeliveryChatroomMessageCommand) GetSenderLastName() string {
	if m != nil {
		return m.SenderLastName
	}
	return ""
}

func (m *DeliveryChatroomMessageCommand) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeliveryChatroomMessageCommand)(nil), "proto.DeliveryChatroomMessageCommand")
}

func init() {
	proto.RegisterFile("pkg/delivery/proto/delivery.proto", fileDescriptor_delivery_7c4d5b8e9471c36b)
}

var fileDescriptor_delivery_7c4d5b8e9471c36b = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0xc8, 0x4e, 0xd7,
	0x4f, 0x49, 0xcd, 0xc9, 0x2c, 0x4b, 0x2d, 0xaa, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x73,
	0xf5, 0xc0, 0x5c, 0x21, 0x56, 0x30, 0xa5, 0xb4, 0x8b, 0x91, 0x4b, 0xce, 0x05, 0x2a, 0xe3, 0x9c,
	0x91, 0x58, 0x52, 0x94, 0x9f, 0x9f, 0xeb, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0xea, 0x9c, 0x9f,
	0x9b, 0x9b, 0x98, 0x97, 0x22, 0x24, 0xc6, 0xc5, 0x16, 0x94, 0x9f, 0x9f, 0xeb, 0xe9, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x49, 0x71, 0x71, 0x04, 0xa7, 0xe6, 0xa5, 0xa4,
	0x16, 0x79, 0xba, 0x48, 0x30, 0x81, 0x65, 0xe0, 0x7c, 0x21, 0x0d, 0x2e, 0x7e, 0x08, 0xdb, 0x2d,
	0xb3, 0xa8, 0xb8, 0xc4, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x19, 0xac, 0x04, 0x5d, 0x58, 0x48, 0x8d,
	0x8b, 0x0f, 0x22, 0xe4, 0x93, 0x08, 0x55, 0xc8, 0x02, 0x56, 0x88, 0x26, 0x2a, 0x24, 0xc4, 0xc5,
	0xe2, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x27, 0xb1, 0x81,
	0xfd, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x94, 0x2d, 0xb6, 0x42, 0xef, 0x00, 0x00, 0x00,
}
