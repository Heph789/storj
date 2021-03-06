// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_INVALID   NodeType = 0
	NodeType_SATELLITE NodeType = 1
	NodeType_STORAGE   NodeType = 2
	NodeType_UPLINK    NodeType = 3
	NodeType_BOOTSTRAP NodeType = 4
)

var NodeType_name = map[int32]string{
	0: "INVALID",
	1: "SATELLITE",
	2: "STORAGE",
	3: "UPLINK",
	4: "BOOTSTRAP",
}

var NodeType_value = map[string]int32{
	"INVALID":   0,
	"SATELLITE": 1,
	"STORAGE":   2,
	"UPLINK":    3,
	"BOOTSTRAP": 4,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}

func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}

var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}

func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

// TODO move statdb.Update() stuff out of here
// Node represents a node in the overlay network
// Node is info for a updating a single storagenode, used in the Update rpc calls
type Node struct {
	Id                   NodeID       `protobuf:"bytes,1,opt,name=id,proto3,customtype=NodeID" json:"id"`
	Address              *NodeAddress `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	LastIp               string       `protobuf:"bytes,14,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=node.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (m *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(m, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// NodeOperator contains info about the storage node operator
type NodeOperator struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeOperator) Reset()         { *m = NodeOperator{} }
func (m *NodeOperator) String() string { return proto.CompactTextString(m) }
func (*NodeOperator) ProtoMessage()    {}
func (*NodeOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}
func (m *NodeOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeOperator.Unmarshal(m, b)
}
func (m *NodeOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeOperator.Marshal(b, m, deterministic)
}
func (m *NodeOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeOperator.Merge(m, src)
}
func (m *NodeOperator) XXX_Size() int {
	return xxx_messageInfo_NodeOperator.Size(m)
}
func (m *NodeOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeOperator.DiscardUnknown(m)
}

var xxx_messageInfo_NodeOperator proto.InternalMessageInfo

func (m *NodeOperator) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NodeOperator) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

// NodeCapacity contains all relevant data about a nodes ability to store data
type NodeCapacity struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=free_bandwidth,json=freeBandwidth,proto3" json:"free_bandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=free_disk,json=freeDisk,proto3" json:"free_disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeCapacity) Reset()         { *m = NodeCapacity{} }
func (m *NodeCapacity) String() string { return proto.CompactTextString(m) }
func (*NodeCapacity) ProtoMessage()    {}
func (*NodeCapacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}
func (m *NodeCapacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeCapacity.Unmarshal(m, b)
}
func (m *NodeCapacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeCapacity.Marshal(b, m, deterministic)
}
func (m *NodeCapacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeCapacity.Merge(m, src)
}
func (m *NodeCapacity) XXX_Size() int {
	return xxx_messageInfo_NodeCapacity.Size(m)
}
func (m *NodeCapacity) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeCapacity.DiscardUnknown(m)
}

var xxx_messageInfo_NodeCapacity proto.InternalMessageInfo

func (m *NodeCapacity) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeCapacity) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// Deprecated: use NodeOperator instead
type NodeMetadata struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{4}
}
func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (m *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(m, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NodeMetadata) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

// Deprecated: use NodeCapacity instead
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=free_bandwidth,json=freeBandwidth,proto3" json:"free_bandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=free_disk,json=freeDisk,proto3" json:"free_disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{5}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (m *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(m, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// NodeVersion contains
type NodeVersion struct {
	Version              string    `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CommitHash           string    `protobuf:"bytes,2,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	Timestamp            time.Time `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Release              bool      `protobuf:"varint,4,opt,name=release,proto3" json:"release,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NodeVersion) Reset()         { *m = NodeVersion{} }
func (m *NodeVersion) String() string { return proto.CompactTextString(m) }
func (*NodeVersion) ProtoMessage()    {}
func (*NodeVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{6}
}
func (m *NodeVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeVersion.Unmarshal(m, b)
}
func (m *NodeVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeVersion.Marshal(b, m, deterministic)
}
func (m *NodeVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeVersion.Merge(m, src)
}
func (m *NodeVersion) XXX_Size() int {
	return xxx_messageInfo_NodeVersion.Size(m)
}
func (m *NodeVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeVersion.DiscardUnknown(m)
}

var xxx_messageInfo_NodeVersion proto.InternalMessageInfo

func (m *NodeVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NodeVersion) GetCommitHash() string {
	if m != nil {
		return m.CommitHash
	}
	return ""
}

func (m *NodeVersion) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *NodeVersion) GetRelease() bool {
	if m != nil {
		return m.Release
	}
	return false
}

func init() {
	proto.RegisterEnum("node.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("node.NodeTransport", NodeTransport_name, NodeTransport_value)
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeAddress)(nil), "node.NodeAddress")
	proto.RegisterType((*NodeOperator)(nil), "node.NodeOperator")
	proto.RegisterType((*NodeCapacity)(nil), "node.NodeCapacity")
	proto.RegisterType((*NodeMetadata)(nil), "node.NodeMetadata")
	proto.RegisterType((*NodeRestrictions)(nil), "node.NodeRestrictions")
	proto.RegisterType((*NodeVersion)(nil), "node.NodeVersion")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0x1a, 0x3d,
	0x14, 0xcd, 0xc0, 0x04, 0x98, 0xcb, 0x8f, 0xe6, 0xf3, 0x17, 0xb5, 0x28, 0x95, 0x0a, 0x45, 0xaa,
	0x84, 0x52, 0x89, 0xa8, 0xe9, 0xb6, 0x1b, 0x48, 0xa2, 0x94, 0x96, 0x02, 0x32, 0xd3, 0x2c, 0xb2,
	0x19, 0x19, 0xec, 0x80, 0x95, 0x61, 0x6c, 0xd9, 0x9e, 0x46, 0xbc, 0x45, 0x9f, 0xa2, 0xcf, 0xd2,
	0x67, 0xe8, 0x22, 0x7d, 0x93, 0xaa, 0xf2, 0xfc, 0xe4, 0x67, 0x59, 0xa9, 0xbb, 0x39, 0xe7, 0x9e,
	0x6b, 0x9f, 0x39, 0xf7, 0x1a, 0x20, 0x16, 0x94, 0x0d, 0xa4, 0x12, 0x46, 0x20, 0xd7, 0x7e, 0x1f,
	0xc2, 0x5a, 0xac, 0x45, 0xc6, 0x1c, 0x76, 0xd6, 0x42, 0xac, 0x23, 0x76, 0x9c, 0xa2, 0x65, 0x72,
	0x7d, 0x6c, 0xf8, 0x96, 0x69, 0x43, 0xb6, 0x32, 0x13, 0xf4, 0x7e, 0x3b, 0xe0, 0x4e, 0x05, 0x65,
	0xe8, 0x25, 0x94, 0x38, 0x6d, 0x3b, 0x5d, 0xa7, 0xdf, 0x18, 0xb5, 0x7e, 0xdc, 0x75, 0xf6, 0x7e,
	0xde, 0x75, 0x2a, 0xb6, 0x32, 0x3e, 0xc3, 0x25, 0x4e, 0xd1, 0x1b, 0xa8, 0x12, 0x4a, 0x15, 0xd3,
	0xba, 0x5d, 0xea, 0x3a, 0xfd, 0xfa, 0xc9, 0x7f, 0x83, 0xf4, 0x66, 0x2b, 0x19, 0x66, 0x05, 0x5c,
	0x28, 0xd0, 0x73, 0xa8, 0x46, 0x44, 0x9b, 0x90, 0xcb, 0x76, 0xab, 0xeb, 0xf4, 0x3d, 0x5c, 0xb1,
	0x70, 0x2c, 0x3f, 0xba, 0xb5, 0xb2, 0xdf, 0xc2, 0xae, 0xd9, 0x49, 0x86, 0x1b, 0x8a, 0x69, 0xa3,
	0xf8, 0xca, 0x70, 0x11, 0x6b, 0x0c, 0x8a, 0xc9, 0xc4, 0x10, 0x0b, 0x70, 0x6d, 0xcb, 0x0c, 0xa1,
	0xc4, 0x10, 0xdc, 0x88, 0x88, 0x61, 0xf1, 0x6a, 0x17, 0x46, 0x5c, 0x1b, 0xdc, 0x24, 0x09, 0xe5,
	0x26, 0xd4, 0xc9, 0x6a, 0x65, 0xaf, 0xdb, 0xe7, 0x3a, 0x4c, 0x24, 0x6e, 0x25, 0x92, 0x12, 0xc3,
	0xc2, 0x5c, 0x8a, 0x0f, 0x72, 0xfc, 0x54, 0xdc, 0xcc, 0xd9, 0x44, 0xda, 0x08, 0x70, 0xf5, 0x2b,
	0x53, 0x9a, 0x8b, 0xb8, 0x77, 0x05, 0xf5, 0x47, 0xbf, 0x80, 0xde, 0x82, 0x67, 0x14, 0x89, 0xb5,
	0x14, 0xca, 0xa4, 0x69, 0xb4, 0x4e, 0xfe, 0x7f, 0xf8, 0xd1, 0xa0, 0x28, 0xe1, 0x07, 0x15, 0x6a,
	0x3f, 0x4d, 0xc6, 0xbb, 0x8f, 0xa1, 0xf7, 0x1e, 0x1a, 0xb6, 0x6b, 0x26, 0x99, 0x22, 0x46, 0x28,
	0x74, 0x00, 0xfb, 0x6c, 0x4b, 0x78, 0x94, 0x1e, 0xec, 0xe1, 0x0c, 0xa0, 0x67, 0x50, 0xb9, 0x25,
	0x51, 0xc4, 0x4c, 0xde, 0x9e, 0xa3, 0x1e, 0xce, 0xba, 0x4f, 0x89, 0x24, 0x2b, 0x6e, 0x76, 0xe8,
	0x35, 0xb4, 0xae, 0x15, 0x63, 0xe1, 0x92, 0xc4, 0xf4, 0x96, 0x53, 0xb3, 0x49, 0x8f, 0x29, 0xe3,
	0xa6, 0x65, 0x47, 0x05, 0x89, 0x5e, 0x80, 0x97, 0xca, 0x28, 0xd7, 0x37, 0xe9, 0x89, 0x65, 0x5c,
	0xb3, 0xc4, 0x19, 0xd7, 0x37, 0x85, 0xa3, 0xcf, 0x79, 0xbe, 0x7f, 0xe9, 0xe8, 0x12, 0x7c, 0xdb,
	0x8d, 0x1f, 0xcd, 0xed, 0x9f, 0xb8, 0xfa, 0xee, 0x64, 0x43, 0xb8, 0xcc, 0x66, 0x62, 0x13, 0xcd,
	0xc7, 0x93, 0xfb, 0x2a, 0x20, 0xea, 0x40, 0x7d, 0x25, 0xb6, 0x5b, 0x6e, 0xc2, 0x0d, 0xd1, 0x9b,
	0xdc, 0x1e, 0x64, 0xd4, 0x07, 0xa2, 0x37, 0x68, 0x04, 0xde, 0xfd, 0x8a, 0xb7, 0xcb, 0xe9, 0xa2,
	0x1e, 0x0e, 0xb2, 0x47, 0x30, 0x28, 0x1e, 0xc1, 0x20, 0x28, 0x14, 0xa3, 0x9a, 0xdd, 0xf4, 0x6f,
	0xbf, 0x3a, 0x0e, 0x7e, 0x68, 0xb3, 0xd7, 0x2b, 0x16, 0x31, 0xa2, 0x59, 0xdb, 0xed, 0x3a, 0xfd,
	0x1a, 0x2e, 0xe0, 0xd1, 0x14, 0x6a, 0xe9, 0x1a, 0xec, 0x24, 0x43, 0x75, 0xa8, 0x8e, 0xa7, 0x97,
	0xc3, 0xc9, 0xf8, 0xcc, 0xdf, 0x43, 0x4d, 0xf0, 0x16, 0xc3, 0xe0, 0x7c, 0x32, 0x19, 0x07, 0xe7,
	0xbe, 0x63, 0x6b, 0x8b, 0x60, 0x86, 0x87, 0x17, 0xe7, 0x7e, 0x09, 0x01, 0x54, 0xbe, 0xcc, 0x27,
	0xe3, 0xe9, 0x27, 0xbf, 0x6c, 0x75, 0xa3, 0xd9, 0x2c, 0x58, 0x04, 0x78, 0x38, 0xf7, 0xdd, 0xa3,
	0x57, 0xd0, 0x7c, 0xb2, 0x56, 0xc8, 0x87, 0x46, 0x70, 0x3a, 0x0f, 0x83, 0xc9, 0x22, 0xbc, 0xc0,
	0xf3, 0x53, 0x7f, 0x6f, 0xe4, 0x5e, 0x95, 0xe4, 0x72, 0x59, 0x49, 0xbd, 0xbf, 0xfb, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x0c, 0xb7, 0x9a, 0xee, 0x03, 0x00, 0x00,
}
