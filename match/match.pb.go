// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: match/match.proto

package match

import (
	common "github.com/Waylon1991/KingdomProto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode   common.RoomMode `protobuf:"varint,1,opt,name=mode,proto3,enum=protocol.common.RoomMode" json:"mode,omitempty"` //模式
	UserId int32           `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             //用户id
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRequest) GetMode() common.RoomMode {
	if x != nil {
		return x.Mode
	}
	return common.RoomMode(0)
}

func (x *MatchRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type MatchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`                    // 匹配人数
	PreTime int32 `protobuf:"varint,2,opt,name=pre_time,json=preTime,proto3" json:"pre_time,omitempty"` //预计时间
}

func (x *MatchReply) Reset() {
	*x = MatchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchReply) ProtoMessage() {}

func (x *MatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchReply.ProtoReflect.Descriptor instead.
func (*MatchReply) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{1}
}

func (x *MatchReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MatchReply) GetPreTime() int32 {
	if x != nil {
		return x.PreTime
	}
	return 0
}

type CancelMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CancelMatchRequest) Reset() {
	*x = CancelMatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchRequest) ProtoMessage() {}

func (x *CancelMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMatchRequest.ProtoReflect.Descriptor instead.
func (*CancelMatchRequest) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{2}
}

func (x *CancelMatchRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CancelMatchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suc bool `protobuf:"varint,1,opt,name=suc,proto3" json:"suc,omitempty"`
}

func (x *CancelMatchReply) Reset() {
	*x = CancelMatchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchReply) ProtoMessage() {}

func (x *CancelMatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMatchReply.ProtoReflect.Descriptor instead.
func (*CancelMatchReply) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{3}
}

func (x *CancelMatchReply) GetSuc() bool {
	if x != nil {
		return x.Suc
	}
	return false
}

type SelectHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int32 `protobuf:"varint,1,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SelectHeroRequest) Reset() {
	*x = SelectHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectHeroRequest) ProtoMessage() {}

func (x *SelectHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectHeroRequest.ProtoReflect.Descriptor instead.
func (*SelectHeroRequest) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{4}
}

func (x *SelectHeroRequest) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *SelectHeroRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SelectHeroReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*common.SelectInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"` //选择信息
	Heros []*common.Hero       `protobuf:"bytes,2,rep,name=heros,proto3" json:"heros,omitempty"` //英雄列表
}

func (x *SelectHeroReply) Reset() {
	*x = SelectHeroReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_match_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectHeroReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectHeroReply) ProtoMessage() {}

func (x *SelectHeroReply) ProtoReflect() protoreflect.Message {
	mi := &file_match_match_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectHeroReply.ProtoReflect.Descriptor instead.
func (*SelectHeroReply) Descriptor() ([]byte, []int) {
	return file_match_match_proto_rawDescGZIP(), []int{5}
}

func (x *SelectHeroReply) GetInfos() []*common.SelectInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *SelectHeroReply) GetHeros() []*common.Hero {
	if x != nil {
		return x.Heros
	}
	return nil
}

var File_match_match_proto protoreflect.FileDescriptor

var file_match_match_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x2d, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24,
	0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x73, 0x75, 0x63, 0x22, 0x45, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x65,
	0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x65, 0x72,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x0f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x2b, 0x0a, 0x05, 0x68, 0x65, 0x72, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x05, 0x68, 0x65, 0x72, 0x6f, 0x73, 0x32, 0xf1,
	0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x41, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0b, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x50, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x57, 0x61, 0x79, 0x6c, 0x6f, 0x6e, 0x31, 0x39, 0x39, 0x31, 0x2f, 0x4b, 0x69, 0x6e, 0x67,
	0x64, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x3b, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_match_proto_rawDescOnce sync.Once
	file_match_match_proto_rawDescData = file_match_match_proto_rawDesc
)

func file_match_match_proto_rawDescGZIP() []byte {
	file_match_match_proto_rawDescOnce.Do(func() {
		file_match_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_match_proto_rawDescData)
	})
	return file_match_match_proto_rawDescData
}

var file_match_match_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_match_match_proto_goTypes = []interface{}{
	(*MatchRequest)(nil),       // 0: protocol.match.MatchRequest
	(*MatchReply)(nil),         // 1: protocol.match.MatchReply
	(*CancelMatchRequest)(nil), // 2: protocol.match.CancelMatchRequest
	(*CancelMatchReply)(nil),   // 3: protocol.match.CancelMatchReply
	(*SelectHeroRequest)(nil),  // 4: protocol.match.SelectHeroRequest
	(*SelectHeroReply)(nil),    // 5: protocol.match.SelectHeroReply
	(common.RoomMode)(0),       // 6: protocol.common.RoomMode
	(*common.SelectInfo)(nil),  // 7: protocol.common.SelectInfo
	(*common.Hero)(nil),        // 8: protocol.common.Hero
}
var file_match_match_proto_depIdxs = []int32{
	6, // 0: protocol.match.MatchRequest.mode:type_name -> protocol.common.RoomMode
	7, // 1: protocol.match.SelectHeroReply.infos:type_name -> protocol.common.SelectInfo
	8, // 2: protocol.match.SelectHeroReply.heros:type_name -> protocol.common.Hero
	0, // 3: protocol.match.Match.Match:input_type -> protocol.match.MatchRequest
	2, // 4: protocol.match.Match.CancelMatch:input_type -> protocol.match.CancelMatchRequest
	4, // 5: protocol.match.Match.SelectHero:input_type -> protocol.match.SelectHeroRequest
	1, // 6: protocol.match.Match.Match:output_type -> protocol.match.MatchReply
	3, // 7: protocol.match.Match.CancelMatch:output_type -> protocol.match.CancelMatchReply
	5, // 8: protocol.match.Match.SelectHero:output_type -> protocol.match.SelectHeroReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_match_match_proto_init() }
func file_match_match_proto_init() {
	if File_match_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMatchReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_match_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectHeroRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_match_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectHeroReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_match_proto_goTypes,
		DependencyIndexes: file_match_match_proto_depIdxs,
		MessageInfos:      file_match_match_proto_msgTypes,
	}.Build()
	File_match_match_proto = out.File
	file_match_match_proto_rawDesc = nil
	file_match_match_proto_goTypes = nil
	file_match_match_proto_depIdxs = nil
}
