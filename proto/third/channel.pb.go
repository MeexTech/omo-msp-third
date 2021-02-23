// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/channel.proto

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

type ChannelInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Phone                string   `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Secret               string   `protobuf:"bytes,10,opt,name=secret,proto3" json:"secret,omitempty"`
	Tags                 []string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{0}
}

func (m *ChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInfo.Unmarshal(m, b)
}
func (m *ChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInfo.Marshal(b, m, deterministic)
}
func (m *ChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInfo.Merge(m, src)
}
func (m *ChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ChannelInfo.Size(m)
}
func (m *ChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInfo proto.InternalMessageInfo

func (m *ChannelInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ChannelInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChannelInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ChannelInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ChannelInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ChannelInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ChannelInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChannelInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ChannelInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChannelInfo) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ChannelInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqChannelAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Passwords            string   `protobuf:"bytes,9,opt,name=passwords,proto3" json:"passwords,omitempty"`
	Entity               string   `protobuf:"bytes,10,opt,name=entity,proto3" json:"entity,omitempty"`
	Portrait             string   `protobuf:"bytes,11,opt,name=portrait,proto3" json:"portrait,omitempty"`
	Tags                 []string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelAdd) Reset()         { *m = ReqChannelAdd{} }
func (m *ReqChannelAdd) String() string { return proto.CompactTextString(m) }
func (*ReqChannelAdd) ProtoMessage()    {}
func (*ReqChannelAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{1}
}

func (m *ReqChannelAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelAdd.Unmarshal(m, b)
}
func (m *ReqChannelAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelAdd.Marshal(b, m, deterministic)
}
func (m *ReqChannelAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelAdd.Merge(m, src)
}
func (m *ReqChannelAdd) XXX_Size() int {
	return xxx_messageInfo_ReqChannelAdd.Size(m)
}
func (m *ReqChannelAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelAdd proto.InternalMessageInfo

func (m *ReqChannelAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqChannelAdd) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ReqChannelAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqChannelAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelAdd) GetPasswords() string {
	if m != nil {
		return m.Passwords
	}
	return ""
}

func (m *ReqChannelAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqChannelAdd) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

func (m *ReqChannelAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqChannelUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NickName             string   `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Job                  string   `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Portrait             string   `protobuf:"bytes,9,opt,name=portrait,proto3" json:"portrait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelUpdate) Reset()         { *m = ReqChannelUpdate{} }
func (m *ReqChannelUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqChannelUpdate) ProtoMessage()    {}
func (*ReqChannelUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{2}
}

func (m *ReqChannelUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelUpdate.Unmarshal(m, b)
}
func (m *ReqChannelUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelUpdate.Marshal(b, m, deterministic)
}
func (m *ReqChannelUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelUpdate.Merge(m, src)
}
func (m *ReqChannelUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqChannelUpdate.Size(m)
}
func (m *ReqChannelUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelUpdate proto.InternalMessageInfo

func (m *ReqChannelUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqChannelUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqChannelUpdate) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *ReqChannelUpdate) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ReqChannelUpdate) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *ReqChannelUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqChannelUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelUpdate) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

type ReqChannelTags struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelTags) Reset()         { *m = ReqChannelTags{} }
func (m *ReqChannelTags) String() string { return proto.CompactTextString(m) }
func (*ReqChannelTags) ProtoMessage()    {}
func (*ReqChannelTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{3}
}

func (m *ReqChannelTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelTags.Unmarshal(m, b)
}
func (m *ReqChannelTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelTags.Marshal(b, m, deterministic)
}
func (m *ReqChannelTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelTags.Merge(m, src)
}
func (m *ReqChannelTags) XXX_Size() int {
	return xxx_messageInfo_ReqChannelTags.Size(m)
}
func (m *ReqChannelTags) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelTags.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelTags proto.InternalMessageInfo

func (m *ReqChannelTags) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqChannelTags) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqChannelTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReplyChannelInfo struct {
	Info                 *ChannelInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyChannelInfo) Reset()         { *m = ReplyChannelInfo{} }
func (m *ReplyChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyChannelInfo) ProtoMessage()    {}
func (*ReplyChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{4}
}

func (m *ReplyChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyChannelInfo.Unmarshal(m, b)
}
func (m *ReplyChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyChannelInfo.Marshal(b, m, deterministic)
}
func (m *ReplyChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyChannelInfo.Merge(m, src)
}
func (m *ReplyChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyChannelInfo.Size(m)
}
func (m *ReplyChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyChannelInfo proto.InternalMessageInfo

func (m *ReplyChannelInfo) GetInfo() *ChannelInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyChannelInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqChannelList struct {
	List                 []string `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqChannelList) Reset()         { *m = ReqChannelList{} }
func (m *ReqChannelList) String() string { return proto.CompactTextString(m) }
func (*ReqChannelList) ProtoMessage()    {}
func (*ReqChannelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{5}
}

func (m *ReqChannelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqChannelList.Unmarshal(m, b)
}
func (m *ReqChannelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqChannelList.Marshal(b, m, deterministic)
}
func (m *ReqChannelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqChannelList.Merge(m, src)
}
func (m *ReqChannelList) XXX_Size() int {
	return xxx_messageInfo_ReqChannelList.Size(m)
}
func (m *ReqChannelList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqChannelList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqChannelList proto.InternalMessageInfo

func (m *ReqChannelList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyChannelList struct {
	Total                uint64         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32         `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32         `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*ChannelInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyChannelList) Reset()         { *m = ReplyChannelList{} }
func (m *ReplyChannelList) String() string { return proto.CompactTextString(m) }
func (*ReplyChannelList) ProtoMessage()    {}
func (*ReplyChannelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e36d889a4cd775, []int{6}
}

func (m *ReplyChannelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyChannelList.Unmarshal(m, b)
}
func (m *ReplyChannelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyChannelList.Marshal(b, m, deterministic)
}
func (m *ReplyChannelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyChannelList.Merge(m, src)
}
func (m *ReplyChannelList) XXX_Size() int {
	return xxx_messageInfo_ReplyChannelList.Size(m)
}
func (m *ReplyChannelList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyChannelList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyChannelList proto.InternalMessageInfo

func (m *ReplyChannelList) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyChannelList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyChannelList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyChannelList) GetList() []*ChannelInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyChannelList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelInfo)(nil), "omo.msp.third.ChannelInfo")
	proto.RegisterType((*ReqChannelAdd)(nil), "omo.msp.third.ReqChannelAdd")
	proto.RegisterType((*ReqChannelUpdate)(nil), "omo.msp.third.ReqChannelUpdate")
	proto.RegisterType((*ReqChannelTags)(nil), "omo.msp.third.ReqChannelTags")
	proto.RegisterType((*ReplyChannelInfo)(nil), "omo.msp.third.ReplyChannelInfo")
	proto.RegisterType((*ReqChannelList)(nil), "omo.msp.third.ReqChannelList")
	proto.RegisterType((*ReplyChannelList)(nil), "omo.msp.third.ReplyChannelList")
}

func init() {
	proto.RegisterFile("proto/third/channel.proto", fileDescriptor_76e36d889a4cd775)
}

var fileDescriptor_76e36d889a4cd775 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6e, 0xdb, 0x3c,
	0x10, 0x8d, 0x7e, 0xac, 0xd8, 0x63, 0x38, 0x08, 0x88, 0xe0, 0x03, 0x3f, 0x23, 0x40, 0x0d, 0xa1,
	0x8b, 0xac, 0x14, 0x20, 0x3d, 0x41, 0x9a, 0x45, 0x1a, 0xa0, 0x75, 0x0b, 0xa5, 0xdd, 0x74, 0xc7,
	0x58, 0x8c, 0xad, 0xc6, 0x12, 0x15, 0x92, 0xb6, 0x9b, 0x13, 0xf4, 0x40, 0xbd, 0x40, 0x0f, 0xd0,
	0x5d, 0x2f, 0x54, 0x70, 0x28, 0xc9, 0x74, 0xec, 0x36, 0x41, 0x77, 0x7c, 0x33, 0x4f, 0x6f, 0xde,
	0x1b, 0x9a, 0x86, 0xff, 0x2b, 0x29, 0xb4, 0x38, 0xd5, 0xb3, 0x5c, 0x66, 0xa7, 0x93, 0x19, 0x2b,
	0x4b, 0x3e, 0x4f, 0xb0, 0x46, 0x06, 0xa2, 0x10, 0x49, 0xa1, 0xaa, 0x04, 0x9b, 0x43, 0xba, 0xc1,
	0x14, 0x45, 0x21, 0x4a, 0x4b, 0x8c, 0xbf, 0xf9, 0xd0, 0xbf, 0xb0, 0x9f, 0x5e, 0x95, 0xb7, 0x82,
	0x1c, 0x42, 0xb0, 0xc8, 0x33, 0xea, 0x8d, 0xbc, 0x93, 0x5e, 0x6a, 0x8e, 0xe4, 0x00, 0xfc, 0x3c,
	0xa3, 0xfe, 0xc8, 0x3b, 0x09, 0x53, 0x3f, 0xcf, 0x08, 0x85, 0xfd, 0x89, 0xe4, 0x4c, 0xf3, 0x8c,
	0x06, 0x23, 0xef, 0x24, 0x48, 0x1b, 0x68, 0x3a, 0x8b, 0x2a, 0xc3, 0x4e, 0x68, 0x3b, 0x35, 0x6c,
	0xbf, 0x11, 0x92, 0x76, 0x50, 0xb9, 0x81, 0x64, 0x08, 0x5d, 0x51, 0x71, 0x89, 0xad, 0x08, 0x5b,
	0x2d, 0x26, 0x04, 0xc2, 0x92, 0x15, 0x9c, 0xee, 0x63, 0x1d, 0xcf, 0xe4, 0x3f, 0x88, 0x24, 0x2f,
	0x98, 0xbc, 0xa3, 0x5d, 0xac, 0xd6, 0x88, 0x1c, 0x41, 0xa7, 0x9a, 0x89, 0x92, 0xd3, 0x1e, 0x96,
	0x2d, 0x30, 0x6c, 0xc5, 0x27, 0x92, 0x6b, 0x0a, 0x96, 0x6d, 0x91, 0x51, 0xd6, 0x6c, 0xaa, 0x68,
	0x7f, 0x14, 0x18, 0x65, 0x73, 0x8e, 0x7f, 0x79, 0x30, 0x48, 0xf9, 0x7d, 0xbd, 0x8c, 0xf3, 0x2c,
	0x6b, 0xe7, 0x7b, 0xce, 0xfc, 0x76, 0x4e, 0xe7, 0xd1, 0x9c, 0xda, 0x55, 0xb4, 0xe1, 0xca, 0x4d,
	0xd7, 0x7d, 0x94, 0xee, 0x18, 0x7a, 0x15, 0x53, 0x6a, 0x25, 0x64, 0xa6, 0x6a, 0xd7, 0xeb, 0x82,
	0x51, 0xe4, 0xa5, 0xce, 0xf5, 0x43, 0xe3, 0xdc, 0x22, 0xa3, 0x58, 0x09, 0xa9, 0x25, 0xcb, 0x35,
	0xed, 0x5b, 0xc5, 0x06, 0xb7, 0xa9, 0x06, 0x4e, 0xaa, 0x9f, 0x1e, 0x1c, 0xae, 0x53, 0x7d, 0xc2,
	0xfb, 0xd8, 0x71, 0xc9, 0x4d, 0x54, 0xdf, 0x89, 0x3a, 0x84, 0x6e, 0x99, 0x4f, 0xee, 0xc6, 0xa6,
	0x1e, 0xd8, 0x51, 0x0d, 0x5e, 0xaf, 0x21, 0x74, 0xd7, 0x70, 0x08, 0xc1, 0x17, 0x71, 0x53, 0xaf,
	0xc6, 0x1c, 0xff, 0x69, 0x31, 0x6e, 0xc4, 0xde, 0x66, 0xc4, 0x38, 0x85, 0x83, 0x75, 0x9a, 0x8f,
	0x6c, 0xaa, 0x76, 0x64, 0x71, 0xb5, 0xfd, 0xed, 0x9f, 0x14, 0xae, 0x28, 0x70, 0x56, 0xb4, 0x34,
	0x1b, 0xaa, 0xe6, 0x0f, 0xee, 0x33, 0x48, 0x20, 0xcc, 0xcb, 0x5b, 0x81, 0xb2, 0xfd, 0xb3, 0x61,
	0xb2, 0xf1, 0x9c, 0x12, 0x87, 0x99, 0x22, 0x8f, 0x9c, 0x41, 0xa4, 0x34, 0xd3, 0x0b, 0x85, 0x13,
	0xb7, 0xbf, 0xc0, 0x01, 0xd7, 0xc8, 0x48, 0x6b, 0x66, 0xfc, 0xd2, 0xcd, 0xf2, 0x36, 0x57, 0x78,
	0x81, 0xf3, 0x5c, 0x69, 0x1a, 0x5a, 0x77, 0xe6, 0x1c, 0xff, 0xf0, 0x36, 0xed, 0x21, 0xf1, 0x08,
	0x3a, 0x5a, 0x68, 0x36, 0x47, 0x7f, 0x61, 0x6a, 0x81, 0x79, 0x65, 0x15, 0x9b, 0xf2, 0x77, 0xec,
	0x2b, 0xba, 0x18, 0xa4, 0x0d, 0x6c, 0x3a, 0x63, 0xb1, 0xc2, 0x9b, 0xac, 0x3b, 0x63, 0xb1, 0x32,
	0x41, 0xdb, 0x91, 0x4f, 0x04, 0x35, 0x3c, 0x27, 0x68, 0xe7, 0xb9, 0x41, 0xcf, 0xbe, 0xfb, 0x70,
	0x50, 0x2b, 0x5d, 0x73, 0xb9, 0xcc, 0x27, 0x9c, 0x5c, 0x41, 0x74, 0x9e, 0x65, 0xef, 0x4b, 0x4e,
	0x8e, 0xb7, 0x04, 0x9c, 0x27, 0x38, 0x7c, 0xb1, 0x4b, 0xde, 0x71, 0x15, 0xef, 0x91, 0x4b, 0x88,
	0x2e, 0xb9, 0x36, 0x52, 0xdb, 0x5e, 0xee, 0x17, 0x5c, 0x69, 0xc3, 0x7b, 0x8e, 0xd0, 0x1b, 0xd8,
	0xbf, 0xe4, 0x1a, 0xf7, 0xfb, 0x07, 0xa5, 0x0f, 0x6c, 0xca, 0xff, 0xaa, 0x64, 0x3e, 0x8e, 0xf7,
	0xc8, 0x05, 0xf4, 0x52, 0x5e, 0x88, 0x25, 0x7f, 0xca, 0x15, 0xdd, 0xa5, 0x65, 0xed, 0xbc, 0x1e,
	0x7c, 0xee, 0x3b, 0xff, 0xda, 0x37, 0x11, 0x82, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x6f, 0x72, 0x0a, 0xf5, 0x05, 0x00, 0x00,
}
