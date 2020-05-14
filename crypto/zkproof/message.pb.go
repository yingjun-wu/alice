// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/getamis/alice/crypto/zkproof/message.proto

package zkproof

import (
	fmt "fmt"
	ecpointgrouplaw "github.com/getamis/alice/crypto/ecpointgrouplaw"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IntegerFactorizationProofMessage struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Challenge            []byte   `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Proof                []byte   `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegerFactorizationProofMessage) Reset()         { *m = IntegerFactorizationProofMessage{} }
func (m *IntegerFactorizationProofMessage) String() string { return proto.CompactTextString(m) }
func (*IntegerFactorizationProofMessage) ProtoMessage()    {}
func (*IntegerFactorizationProofMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7463df78901cfd5c, []int{0}
}

func (m *IntegerFactorizationProofMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Unmarshal(m, b)
}
func (m *IntegerFactorizationProofMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Marshal(b, m, deterministic)
}
func (m *IntegerFactorizationProofMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerFactorizationProofMessage.Merge(m, src)
}
func (m *IntegerFactorizationProofMessage) XXX_Size() int {
	return xxx_messageInfo_IntegerFactorizationProofMessage.Size(m)
}
func (m *IntegerFactorizationProofMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerFactorizationProofMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerFactorizationProofMessage proto.InternalMessageInfo

func (m *IntegerFactorizationProofMessage) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *IntegerFactorizationProofMessage) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *IntegerFactorizationProofMessage) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type SchnorrProofMessage struct {
	Salt                 []byte                          `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
	V                    *ecpointgrouplaw.EcPointMessage `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
	Alpha                *ecpointgrouplaw.EcPointMessage `protobuf:"bytes,3,opt,name=alpha,proto3" json:"alpha,omitempty"`
	U                    []byte                          `protobuf:"bytes,4,opt,name=u,proto3" json:"u,omitempty"`
	T                    []byte                          `protobuf:"bytes,5,opt,name=t,proto3" json:"t,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SchnorrProofMessage) Reset()         { *m = SchnorrProofMessage{} }
func (m *SchnorrProofMessage) String() string { return proto.CompactTextString(m) }
func (*SchnorrProofMessage) ProtoMessage()    {}
func (*SchnorrProofMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7463df78901cfd5c, []int{1}
}

func (m *SchnorrProofMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchnorrProofMessage.Unmarshal(m, b)
}
func (m *SchnorrProofMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchnorrProofMessage.Marshal(b, m, deterministic)
}
func (m *SchnorrProofMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchnorrProofMessage.Merge(m, src)
}
func (m *SchnorrProofMessage) XXX_Size() int {
	return xxx_messageInfo_SchnorrProofMessage.Size(m)
}
func (m *SchnorrProofMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SchnorrProofMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SchnorrProofMessage proto.InternalMessageInfo

func (m *SchnorrProofMessage) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *SchnorrProofMessage) GetV() *ecpointgrouplaw.EcPointMessage {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *SchnorrProofMessage) GetAlpha() *ecpointgrouplaw.EcPointMessage {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *SchnorrProofMessage) GetU() []byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *SchnorrProofMessage) GetT() []byte {
	if m != nil {
		return m.T
	}
	return nil
}

func init() {
	proto.RegisterType((*IntegerFactorizationProofMessage)(nil), "zkproof.IntegerFactorizationProofMessage")
	proto.RegisterType((*SchnorrProofMessage)(nil), "zkproof.SchnorrProofMessage")
}

func init() {
	proto.RegisterFile("github.com/getamis/alice/crypto/zkproof/message.proto", fileDescriptor_7463df78901cfd5c)
}

var fileDescriptor_7463df78901cfd5c = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xfc, 0x30,
	0x14, 0xc4, 0xc9, 0xff, 0xbf, 0x55, 0x8c, 0x7b, 0x8a, 0x1e, 0x8a, 0x08, 0x2e, 0x7b, 0xf2, 0x62,
	0x03, 0xca, 0x9e, 0x3c, 0x2b, 0x88, 0x08, 0xcb, 0x0a, 0x7b, 0x7f, 0x0d, 0xcf, 0x34, 0x98, 0xf6,
	0x85, 0xf4, 0x15, 0xd9, 0xfd, 0x4a, 0x7e, 0x49, 0x69, 0x52, 0x14, 0xbd, 0xe8, 0xad, 0x33, 0xd3,
	0x99, 0x5f, 0x12, 0xb9, 0xb2, 0x8e, 0x9b, 0xa1, 0xae, 0x0c, 0xb5, 0xda, 0x22, 0x43, 0xeb, 0x7a,
	0x0d, 0xde, 0x19, 0xd4, 0x26, 0xee, 0x02, 0x93, 0xde, 0xbf, 0x86, 0x48, 0xf4, 0xa2, 0x5b, 0xec,
	0x7b, 0xb0, 0x58, 0x85, 0x48, 0x4c, 0xea, 0x70, 0xb2, 0xcf, 0x6e, 0x7f, 0xeb, 0xa3, 0x09, 0xe4,
	0x3a, 0xb6, 0x91, 0x86, 0xe0, 0xe1, 0x4d, 0x27, 0x95, 0x57, 0x96, 0x2c, 0x17, 0x0f, 0x1d, 0xa3,
	0xc5, 0x78, 0x0f, 0x86, 0x29, 0xba, 0x3d, 0xb0, 0xa3, 0x6e, 0x3d, 0x2e, 0x3f, 0x65, 0x9e, 0x3a,
	0x97, 0x47, 0x61, 0xa8, 0xbd, 0x33, 0x8f, 0xb8, 0x2b, 0xc5, 0x42, 0x5c, 0xce, 0x37, 0x5f, 0xc6,
	0x98, 0x9a, 0x06, 0xbc, 0xc7, 0xce, 0x62, 0xf9, 0x2f, 0xa7, 0x9f, 0x86, 0x3a, 0x95, 0x45, 0x3a,
	0x65, 0xf9, 0x3f, 0x25, 0x59, 0x2c, 0xdf, 0x85, 0x3c, 0x79, 0x36, 0x4d, 0x47, 0x31, 0x7e, 0x23,
	0x29, 0x39, 0xeb, 0xc1, 0xf3, 0x04, 0x49, 0xdf, 0xea, 0x4a, 0x8a, 0x6d, 0xda, 0x3d, 0xbe, 0xbe,
	0xa8, 0x7e, 0x5c, 0xa5, 0xba, 0x33, 0xeb, 0x51, 0x4f, 0xfd, 0x8d, 0xd8, 0xaa, 0x95, 0x2c, 0xc0,
	0x87, 0x06, 0x12, 0xf0, 0x0f, 0x95, 0xfc, 0xb7, 0x9a, 0x4b, 0x31, 0x94, 0xb3, 0x84, 0x15, 0xc3,
	0xa8, 0xb8, 0x2c, 0xb2, 0xe2, 0xfa, 0x20, 0x3d, 0xd5, 0xcd, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0xbe, 0xa2, 0xa3, 0xa9, 0x01, 0x00, 0x00,
}
