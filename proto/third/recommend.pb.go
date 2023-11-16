// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/recommend.proto

package third

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

type RecommendInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Quote                uint32   `protobuf:"varint,10,opt,name=quote,proto3" json:"quote,omitempty"`
	Targets              []string `protobuf:"bytes,31,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendInfo) Reset()         { *m = RecommendInfo{} }
func (m *RecommendInfo) String() string { return proto.CompactTextString(m) }
func (*RecommendInfo) ProtoMessage()    {}
func (*RecommendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{0}
}

func (m *RecommendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendInfo.Unmarshal(m, b)
}
func (m *RecommendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendInfo.Marshal(b, m, deterministic)
}
func (m *RecommendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendInfo.Merge(m, src)
}
func (m *RecommendInfo) XXX_Size() int {
	return xxx_messageInfo_RecommendInfo.Size(m)
}
func (m *RecommendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendInfo proto.InternalMessageInfo

func (m *RecommendInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecommendInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecommendInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RecommendInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *RecommendInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RecommendInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RecommendInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecommendInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RecommendInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RecommendInfo) GetQuote() uint32 {
	if m != nil {
		return m.Quote
	}
	return 0
}

func (m *RecommendInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqRecommendAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Quote                string   `protobuf:"bytes,6,opt,name=quote,proto3" json:"quote,omitempty"`
	Targets              []string `protobuf:"bytes,21,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRecommendAdd) Reset()         { *m = ReqRecommendAdd{} }
func (m *ReqRecommendAdd) String() string { return proto.CompactTextString(m) }
func (*ReqRecommendAdd) ProtoMessage()    {}
func (*ReqRecommendAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{1}
}

func (m *ReqRecommendAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRecommendAdd.Unmarshal(m, b)
}
func (m *ReqRecommendAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRecommendAdd.Marshal(b, m, deterministic)
}
func (m *ReqRecommendAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRecommendAdd.Merge(m, src)
}
func (m *ReqRecommendAdd) XXX_Size() int {
	return xxx_messageInfo_ReqRecommendAdd.Size(m)
}
func (m *ReqRecommendAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRecommendAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRecommendAdd proto.InternalMessageInfo

func (m *ReqRecommendAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRecommendAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqRecommendAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqRecommendAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRecommendAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqRecommendAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqRecommendAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqRecommendUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Targets              []string `protobuf:"bytes,21,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRecommendUpdate) Reset()         { *m = ReqRecommendUpdate{} }
func (m *ReqRecommendUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqRecommendUpdate) ProtoMessage()    {}
func (*ReqRecommendUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{2}
}

func (m *ReqRecommendUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRecommendUpdate.Unmarshal(m, b)
}
func (m *ReqRecommendUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRecommendUpdate.Marshal(b, m, deterministic)
}
func (m *ReqRecommendUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRecommendUpdate.Merge(m, src)
}
func (m *ReqRecommendUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqRecommendUpdate.Size(m)
}
func (m *ReqRecommendUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRecommendUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRecommendUpdate proto.InternalMessageInfo

func (m *ReqRecommendUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqRecommendUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRecommendUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqRecommendUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqRecommendUpdate) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReplyRecommendInfo struct {
	Info                 *RecommendInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyRecommendInfo) Reset()         { *m = ReplyRecommendInfo{} }
func (m *ReplyRecommendInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyRecommendInfo) ProtoMessage()    {}
func (*ReplyRecommendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{3}
}

func (m *ReplyRecommendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRecommendInfo.Unmarshal(m, b)
}
func (m *ReplyRecommendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRecommendInfo.Marshal(b, m, deterministic)
}
func (m *ReplyRecommendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRecommendInfo.Merge(m, src)
}
func (m *ReplyRecommendInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyRecommendInfo.Size(m)
}
func (m *ReplyRecommendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRecommendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRecommendInfo proto.InternalMessageInfo

func (m *ReplyRecommendInfo) GetInfo() *RecommendInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyRecommendInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyRecommendList struct {
	Status               *ReplyStatus     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32           `protobuf:"varint,3,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32           `protobuf:"varint,4,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*RecommendInfo `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyRecommendList) Reset()         { *m = ReplyRecommendList{} }
func (m *ReplyRecommendList) String() string { return proto.CompactTextString(m) }
func (*ReplyRecommendList) ProtoMessage()    {}
func (*ReplyRecommendList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cef07d596f5e8559, []int{4}
}

func (m *ReplyRecommendList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRecommendList.Unmarshal(m, b)
}
func (m *ReplyRecommendList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRecommendList.Marshal(b, m, deterministic)
}
func (m *ReplyRecommendList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRecommendList.Merge(m, src)
}
func (m *ReplyRecommendList) XXX_Size() int {
	return xxx_messageInfo_ReplyRecommendList.Size(m)
}
func (m *ReplyRecommendList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRecommendList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRecommendList proto.InternalMessageInfo

func (m *ReplyRecommendList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyRecommendList) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyRecommendList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyRecommendList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyRecommendList) GetList() []*RecommendInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendInfo)(nil), "omo.msp.third.RecommendInfo")
	proto.RegisterType((*ReqRecommendAdd)(nil), "omo.msp.third.ReqRecommendAdd")
	proto.RegisterType((*ReqRecommendUpdate)(nil), "omo.msp.third.ReqRecommendUpdate")
	proto.RegisterType((*ReplyRecommendInfo)(nil), "omo.msp.third.ReplyRecommendInfo")
	proto.RegisterType((*ReplyRecommendList)(nil), "omo.msp.third.ReplyRecommendList")
}

func init() {
	proto.RegisterFile("proto/third/recommend.proto", fileDescriptor_cef07d596f5e8559)
}

var fileDescriptor_cef07d596f5e8559 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0x6b, 0xc7, 0x75, 0x9b, 0xc9, 0x3f, 0xfd, 0x57, 0x2b, 0x40, 0xab, 0x50, 0x20, 0xcd,
	0x29, 0xa7, 0x14, 0x85, 0x27, 0x68, 0x91, 0x88, 0x8a, 0x08, 0x54, 0x8e, 0xb8, 0x70, 0x33, 0xd9,
	0x69, 0x58, 0x11, 0x7b, 0x9d, 0xf5, 0xa6, 0x25, 0x9c, 0xb9, 0xf0, 0x3c, 0xbc, 0x02, 0xef, 0x05,
	0xda, 0xb1, 0x1d, 0xd9, 0xc4, 0x69, 0x72, 0xf3, 0xb7, 0xf3, 0xed, 0x97, 0xdf, 0xcc, 0x8e, 0x14,
	0x78, 0x9a, 0x68, 0x65, 0xd4, 0x85, 0xf9, 0x22, 0xb5, 0xb8, 0xd0, 0x38, 0x55, 0x51, 0x84, 0xb1,
	0x18, 0xd0, 0x29, 0x6b, 0xab, 0x48, 0x0d, 0xa2, 0x34, 0x19, 0x50, 0xb9, 0xc3, 0xcb, 0x5e, 0xeb,
	0x54, 0x71, 0x66, 0xec, 0xfd, 0x74, 0xa1, 0x1d, 0x14, 0x97, 0xaf, 0xe3, 0x5b, 0xc5, 0x4e, 0xa1,
	0xb1, 0x94, 0x82, 0x3b, 0x5d, 0xa7, 0xdf, 0x0c, 0xec, 0x27, 0x3b, 0x01, 0x57, 0x0a, 0xee, 0x76,
	0x9d, 0xbe, 0x17, 0xb8, 0x52, 0x30, 0x0e, 0x47, 0x53, 0x8d, 0xa1, 0x41, 0xc1, 0x1b, 0x5d, 0xa7,
	0xdf, 0x08, 0x0a, 0x69, 0x2b, 0xcb, 0x44, 0x50, 0xc5, 0xcb, 0x2a, 0xb9, 0x5c, 0xdf, 0x51, 0x9a,
	0x1f, 0x52, 0x72, 0x21, 0x59, 0x07, 0x8e, 0x55, 0x82, 0x9a, 0x4a, 0x3e, 0x95, 0xd6, 0x9a, 0x31,
	0xf0, 0xe2, 0x30, 0x42, 0x7e, 0x44, 0xe7, 0xf4, 0xcd, 0x1e, 0xc1, 0xa1, 0xba, 0x8f, 0x51, 0xf3,
	0x63, 0x3a, 0xcc, 0x84, 0x75, 0x9a, 0x55, 0x82, 0xbc, 0xd9, 0x75, 0xfa, 0xed, 0x80, 0xbe, 0xad,
	0x73, 0xb1, 0x54, 0x06, 0x39, 0xd0, 0x61, 0x26, 0x2c, 0x89, 0x09, 0xf5, 0x0c, 0x4d, 0xca, 0x5f,
	0x74, 0x1b, 0x96, 0x24, 0x97, 0xbd, 0x5f, 0x0e, 0xfc, 0x1f, 0xe0, 0x62, 0x3d, 0x8e, 0x4b, 0x21,
	0xd6, 0x04, 0x4e, 0x1d, 0x81, 0x5b, 0x26, 0x78, 0x02, 0xbe, 0xc6, 0x28, 0xd4, 0x5f, 0x69, 0x28,
	0xcd, 0x20, 0x57, 0x95, 0xfe, 0xbc, 0xcd, 0xfe, 0x88, 0xfa, 0xb0, 0x8e, 0x3a, 0x1b, 0xc6, 0x26,
	0xf5, 0xe3, 0x2a, 0xf5, 0x0f, 0x07, 0x58, 0x99, 0xfa, 0x23, 0x4d, 0xbc, 0xe6, 0x19, 0x8b, 0x56,
	0xdc, 0x52, 0x2b, 0x65, 0xb8, 0xc6, 0x16, 0x38, 0xaf, 0x04, 0xb7, 0x1d, 0xe3, 0xbb, 0xa5, 0x48,
	0xe6, 0xab, 0xea, 0x32, 0xbd, 0x04, 0x4f, 0xc6, 0xb7, 0x8a, 0x30, 0x5a, 0xc3, 0xb3, 0x41, 0x65,
	0x2d, 0x07, 0x15, 0x6f, 0x40, 0x4e, 0x36, 0x04, 0x3f, 0x35, 0xa1, 0x59, 0xa6, 0xc4, 0xd9, 0x1a,
	0x76, 0x36, 0xee, 0x24, 0xf3, 0xd5, 0x84, 0x1c, 0x41, 0xee, 0xec, 0xfd, 0x76, 0xfe, 0xfd, 0xf1,
	0x77, 0x32, 0x35, 0xa5, 0x28, 0x67, 0xdf, 0x28, 0x3b, 0x7d, 0xa3, 0x4c, 0x38, 0xcf, 0xd7, 0x3d,
	0x13, 0xb6, 0xed, 0x24, 0x9c, 0xe1, 0x38, 0xfc, 0x46, 0x53, 0x6a, 0x07, 0x85, 0x2c, 0x2a, 0xef,
	0xd5, 0x7d, 0x3e, 0xa7, 0x42, 0xda, 0xd6, 0xe7, 0x32, 0x35, 0x34, 0xa7, 0x9d, 0xad, 0x5b, 0xe7,
	0xf0, 0x8f, 0x07, 0xa7, 0xeb, 0xf3, 0x09, 0xea, 0x3b, 0x39, 0x45, 0x36, 0x06, 0xff, 0x52, 0x88,
	0x0f, 0x31, 0xb2, 0xe7, 0x1b, 0x11, 0x95, 0x55, 0xed, 0x9c, 0xd7, 0xb5, 0x57, 0xf9, 0x9d, 0xde,
	0x01, 0x9b, 0x40, 0x33, 0x5b, 0x10, 0x9b, 0x78, 0xfe, 0x40, 0x62, 0xe6, 0xda, 0x2f, 0xf4, 0x1a,
	0xfc, 0x11, 0x1a, 0x9b, 0xb8, 0x39, 0xe2, 0xc5, 0x12, 0x53, 0x63, 0x7d, 0xfb, 0x45, 0xbd, 0x85,
	0xa3, 0x11, 0x1a, 0x7a, 0xbe, 0x2d, 0x59, 0x37, 0xe1, 0x6c, 0x17, 0x96, 0xbd, 0xde, 0x3b, 0x60,
	0xaf, 0xa1, 0x19, 0x60, 0xa4, 0xee, 0x70, 0x17, 0x19, 0xaf, 0x4b, 0xcb, 0x81, 0x6e, 0xa0, 0x35,
	0x42, 0x73, 0xb5, 0x7a, 0x23, 0xe7, 0x06, 0x35, 0x3b, 0xab, 0x8f, 0xc9, 0xaa, 0xfb, 0x61, 0x8d,
	0xe1, 0xbf, 0x11, 0x1a, 0xbb, 0x77, 0x32, 0x35, 0x72, 0xba, 0x23, 0xf2, 0xd9, 0xb6, 0xa5, 0xa5,
	0xcb, 0x34, 0xb1, 0x93, 0xec, 0xad, 0x76, 0x31, 0xe6, 0x2f, 0xfa, 0x40, 0xb3, 0x57, 0xed, 0x4f,
	0xad, 0xd2, 0x3f, 0xc5, 0x67, 0x9f, 0xc4, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x08,
	0x04, 0x92, 0x6b, 0x06, 0x00, 0x00,
}
