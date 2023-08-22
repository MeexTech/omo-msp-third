// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/third/holiday.proto

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

type HolidayInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	From                 int64    `protobuf:"varint,11,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,12,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HolidayInfo) Reset()         { *m = HolidayInfo{} }
func (m *HolidayInfo) String() string { return proto.CompactTextString(m) }
func (*HolidayInfo) ProtoMessage()    {}
func (*HolidayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{0}
}

func (m *HolidayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HolidayInfo.Unmarshal(m, b)
}
func (m *HolidayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HolidayInfo.Marshal(b, m, deterministic)
}
func (m *HolidayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HolidayInfo.Merge(m, src)
}
func (m *HolidayInfo) XXX_Size() int {
	return xxx_messageInfo_HolidayInfo.Size(m)
}
func (m *HolidayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HolidayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HolidayInfo proto.InternalMessageInfo

func (m *HolidayInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *HolidayInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HolidayInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *HolidayInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *HolidayInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *HolidayInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *HolidayInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *HolidayInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HolidayInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *HolidayInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *HolidayInfo) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *HolidayInfo) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ReqHolidayAdd struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	From                 int64    `protobuf:"varint,6,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,7,opt,name=end,proto3" json:"end,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHolidayAdd) Reset()         { *m = ReqHolidayAdd{} }
func (m *ReqHolidayAdd) String() string { return proto.CompactTextString(m) }
func (*ReqHolidayAdd) ProtoMessage()    {}
func (*ReqHolidayAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{1}
}

func (m *ReqHolidayAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHolidayAdd.Unmarshal(m, b)
}
func (m *ReqHolidayAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHolidayAdd.Marshal(b, m, deterministic)
}
func (m *ReqHolidayAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHolidayAdd.Merge(m, src)
}
func (m *ReqHolidayAdd) XXX_Size() int {
	return xxx_messageInfo_ReqHolidayAdd.Size(m)
}
func (m *ReqHolidayAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHolidayAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHolidayAdd proto.InternalMessageInfo

func (m *ReqHolidayAdd) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqHolidayAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqHolidayAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqHolidayAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqHolidayAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqHolidayAdd) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ReqHolidayAdd) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ReqHolidayAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqHolidayUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	From                 int64    `protobuf:"varint,6,opt,name=from,proto3" json:"from,omitempty"`
	End                  int64    `protobuf:"varint,7,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHolidayUpdate) Reset()         { *m = ReqHolidayUpdate{} }
func (m *ReqHolidayUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqHolidayUpdate) ProtoMessage()    {}
func (*ReqHolidayUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{2}
}

func (m *ReqHolidayUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHolidayUpdate.Unmarshal(m, b)
}
func (m *ReqHolidayUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHolidayUpdate.Marshal(b, m, deterministic)
}
func (m *ReqHolidayUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHolidayUpdate.Merge(m, src)
}
func (m *ReqHolidayUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqHolidayUpdate.Size(m)
}
func (m *ReqHolidayUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHolidayUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHolidayUpdate proto.InternalMessageInfo

func (m *ReqHolidayUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqHolidayUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqHolidayUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqHolidayUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqHolidayUpdate) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ReqHolidayUpdate) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type ReplyHolidayInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *HolidayInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyHolidayInfo) Reset()         { *m = ReplyHolidayInfo{} }
func (m *ReplyHolidayInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyHolidayInfo) ProtoMessage()    {}
func (*ReplyHolidayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{3}
}

func (m *ReplyHolidayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHolidayInfo.Unmarshal(m, b)
}
func (m *ReplyHolidayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHolidayInfo.Marshal(b, m, deterministic)
}
func (m *ReplyHolidayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHolidayInfo.Merge(m, src)
}
func (m *ReplyHolidayInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyHolidayInfo.Size(m)
}
func (m *ReplyHolidayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHolidayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHolidayInfo proto.InternalMessageInfo

func (m *ReplyHolidayInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyHolidayInfo) GetInfo() *HolidayInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyHolidayList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*HolidayInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyHolidayList) Reset()         { *m = ReplyHolidayList{} }
func (m *ReplyHolidayList) String() string { return proto.CompactTextString(m) }
func (*ReplyHolidayList) ProtoMessage()    {}
func (*ReplyHolidayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_933eba8bbce86885, []int{4}
}

func (m *ReplyHolidayList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHolidayList.Unmarshal(m, b)
}
func (m *ReplyHolidayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHolidayList.Marshal(b, m, deterministic)
}
func (m *ReplyHolidayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHolidayList.Merge(m, src)
}
func (m *ReplyHolidayList) XXX_Size() int {
	return xxx_messageInfo_ReplyHolidayList.Size(m)
}
func (m *ReplyHolidayList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHolidayList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHolidayList proto.InternalMessageInfo

func (m *ReplyHolidayList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyHolidayList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyHolidayList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyHolidayList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyHolidayList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyHolidayList) GetList() []*HolidayInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*HolidayInfo)(nil), "omo.msp.third.HolidayInfo")
	proto.RegisterType((*ReqHolidayAdd)(nil), "omo.msp.third.ReqHolidayAdd")
	proto.RegisterType((*ReqHolidayUpdate)(nil), "omo.msp.third.ReqHolidayUpdate")
	proto.RegisterType((*ReplyHolidayInfo)(nil), "omo.msp.third.ReplyHolidayInfo")
	proto.RegisterType((*ReplyHolidayList)(nil), "omo.msp.third.ReplyHolidayList")
}

func init() {
	proto.RegisterFile("proto/third/holiday.proto", fileDescriptor_933eba8bbce86885)
}

var fileDescriptor_933eba8bbce86885 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x60, 0x4c, 0x18, 0x62, 0x84, 0x56, 0x51, 0xb5, 0x45, 0xad, 0x82, 0xfc, 0xc4, 0x93,
	0x23, 0xd1, 0x13, 0x84, 0x4a, 0xa5, 0xa9, 0xfa, 0x27, 0x47, 0x7d, 0xe9, 0x9b, 0xc3, 0x0e, 0xcd,
	0xaa, 0xd8, 0xeb, 0xac, 0x17, 0x2a, 0xae, 0x51, 0xa9, 0x57, 0xe8, 0x1d, 0x7a, 0x89, 0x9e, 0xa9,
	0xda, 0xb1, 0x21, 0x8e, 0x70, 0x42, 0xd4, 0xbc, 0xcd, 0x37, 0x33, 0x3b, 0xf3, 0x7d, 0x33, 0x63,
	0xc3, 0xf3, 0x4c, 0x2b, 0xa3, 0xce, 0xcc, 0xb5, 0xd4, 0xe2, 0xec, 0x5a, 0x2d, 0xa5, 0x88, 0x37,
	0x21, 0xf9, 0x98, 0xaf, 0x12, 0x15, 0x26, 0x79, 0x16, 0x52, 0x70, 0xc8, 0xab, 0x99, 0x73, 0x95,
	0x24, 0x2a, 0x2d, 0x12, 0x83, 0x5f, 0x4d, 0xe8, 0xbd, 0x2d, 0x9e, 0x5e, 0xa4, 0x0b, 0xc5, 0x06,
	0xd0, 0x5a, 0x49, 0xc1, 0x9d, 0x91, 0x33, 0xee, 0x46, 0xd6, 0x64, 0x7d, 0x68, 0x4a, 0xc1, 0x9b,
	0x23, 0x67, 0xec, 0x46, 0x4d, 0x29, 0x18, 0x03, 0xd7, 0x6c, 0x32, 0xe4, 0xad, 0x91, 0x33, 0xf6,
	0x23, 0xb2, 0x19, 0x87, 0xce, 0x5c, 0x63, 0x6c, 0x50, 0x70, 0x77, 0xe4, 0x8c, 0x5b, 0xd1, 0x16,
	0xda, 0xc8, 0x2a, 0x13, 0x14, 0x69, 0x17, 0x91, 0x12, 0xee, 0xde, 0x28, 0xcd, 0x3d, 0xea, 0xb6,
	0x85, 0x6c, 0x08, 0x47, 0x2a, 0x43, 0x4d, 0xa1, 0x0e, 0x85, 0x76, 0xd8, 0x76, 0x4f, 0xe3, 0x04,
	0xf9, 0x11, 0xf9, 0xc9, 0x66, 0xcf, 0xc0, 0xd3, 0x98, 0xc4, 0xfa, 0x3b, 0xef, 0x92, 0xb7, 0x44,
	0xec, 0x04, 0xda, 0xea, 0x47, 0x8a, 0x9a, 0x03, 0xb9, 0x0b, 0x60, 0x2b, 0x2c, 0xb4, 0x4a, 0x78,
	0x8f, 0xe8, 0x90, 0x6d, 0x55, 0x63, 0x2a, 0xf8, 0x31, 0xb9, 0xac, 0x19, 0xfc, 0x71, 0xc0, 0x8f,
	0xf0, 0xa6, 0x1c, 0xcd, 0xb9, 0x10, 0x35, 0x93, 0xd9, 0x72, 0x69, 0xd6, 0x72, 0x69, 0xd5, 0x73,
	0x71, 0xab, 0x5c, 0xaa, 0x4a, 0xdb, 0xfb, 0x4a, 0x89, 0xa7, 0xb7, 0xcf, 0xb3, 0xb3, 0xe3, 0xb9,
	0xdb, 0x46, 0xf7, 0x76, 0x1b, 0xc1, 0x4f, 0x07, 0x06, 0xb7, 0xdc, 0xbf, 0xd0, 0xbc, 0x9f, 0x48,
	0xff, 0xc9, 0x44, 0x83, 0xb5, 0xe5, 0x94, 0x2d, 0x37, 0xd5, 0x63, 0x9b, 0x80, 0x97, 0x9b, 0xd8,
	0xac, 0x72, 0xa2, 0xd5, 0x9b, 0x0c, 0xc3, 0x3b, 0x67, 0x1b, 0xd2, 0x83, 0x4b, 0xca, 0x88, 0xca,
	0x4c, 0x16, 0x82, 0x2b, 0xd3, 0x85, 0x22, 0xd6, 0xfb, 0x2f, 0x2a, 0xd5, 0x23, 0xca, 0x0b, 0xfe,
	0x3a, 0x77, 0x1b, 0xbf, 0x97, 0xb9, 0xf9, 0xaf, 0xc6, 0x27, 0xd0, 0x36, 0xca, 0xc4, 0x4b, 0xea,
	0xec, 0x47, 0x05, 0xb0, 0xde, 0x2c, 0xfe, 0x86, 0x79, 0xf9, 0x39, 0x14, 0xc0, 0x8e, 0xc4, 0x1a,
	0xb4, 0x6c, 0x3f, 0x22, 0xdb, 0x8e, 0x36, 0x5d, 0x25, 0x57, 0x58, 0x0c, 0xd0, 0x8f, 0x4a, 0x64,
	0x05, 0x2d, 0x65, 0x6e, 0xf8, 0x60, 0xd4, 0x3a, 0x24, 0xc8, 0xe6, 0x4d, 0x7e, 0xbb, 0xd0, 0x2f,
	0xbd, 0x97, 0xa8, 0xd7, 0x72, 0x8e, 0xec, 0x02, 0xbc, 0x73, 0x21, 0x3e, 0xa5, 0xc8, 0x5e, 0xec,
	0x09, 0xa9, 0x9c, 0xf0, 0xf0, 0xb4, 0x4e, 0x66, 0xa5, 0x43, 0xd0, 0x60, 0x33, 0xf0, 0x66, 0x68,
	0x6c, 0xa9, 0xfd, 0x99, 0xdc, 0xac, 0x30, 0x37, 0x36, 0xef, 0x31, 0x85, 0x3e, 0x03, 0x14, 0x97,
	0x37, 0x8d, 0x73, 0x64, 0xa7, 0xf7, 0xf2, 0x2a, 0x92, 0x1e, 0x53, 0xf1, 0x23, 0xf4, 0x66, 0x68,
	0xa6, 0x9b, 0x37, 0x72, 0x69, 0x50, 0xd7, 0x49, 0xb5, 0xfc, 0x8a, 0xe8, 0x83, 0xf5, 0xec, 0x09,
	0x04, 0x0d, 0xf6, 0x01, 0x8e, 0x67, 0x68, 0xec, 0x96, 0x65, 0x6e, 0xe4, 0xfc, 0x40, 0xc1, 0x97,
	0xf7, 0x9d, 0x08, 0x3d, 0x0e, 0x1a, 0xec, 0x1d, 0xf4, 0x4b, 0xc1, 0x07, 0x18, 0x96, 0x8a, 0x79,
	0x5d, 0xc1, 0x52, 0xea, 0x6b, 0xe8, 0x46, 0x98, 0xa8, 0x35, 0x1e, 0x5a, 0xc4, 0x03, 0x45, 0xa6,
	0xfe, 0xd7, 0x5e, 0xe5, 0xb7, 0x7f, 0xe5, 0x11, 0x78, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x7a, 0xb9, 0xa4, 0x36, 0x06, 0x00, 0x00,
}
