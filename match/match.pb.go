// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: Protocol/match/match.proto

package match

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "protocol/common"
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
		mi := &file_Protocol_match_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_match_match_proto_msgTypes[0]
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
	return file_Protocol_match_match_proto_rawDescGZIP(), []int{0}
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
		mi := &file_Protocol_match_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchReply) ProtoMessage() {}

func (x *MatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_match_match_proto_msgTypes[1]
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
	return file_Protocol_match_match_proto_rawDescGZIP(), []int{1}
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
		mi := &file_Protocol_match_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchRequest) ProtoMessage() {}

func (x *CancelMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_match_match_proto_msgTypes[2]
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
	return file_Protocol_match_match_proto_rawDescGZIP(), []int{2}
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
		mi := &file_Protocol_match_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMatchReply) ProtoMessage() {}

func (x *CancelMatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_Protocol_match_match_proto_msgTypes[3]
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
	return file_Protocol_match_match_proto_rawDescGZIP(), []int{3}
}

func (x *CancelMatchReply) GetSuc() bool {
	if x != nil {
		return x.Suc
	}
	return false
}

var File_Protocol_match_match_proto protoreflect.FileDescriptor

var file_Protocol_match_match_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x1c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x24, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x73, 0x75, 0x63, 0x32, 0x9f, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x41, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Protocol_match_match_proto_rawDescOnce sync.Once
	file_Protocol_match_match_proto_rawDescData = file_Protocol_match_match_proto_rawDesc
)

func file_Protocol_match_match_proto_rawDescGZIP() []byte {
	file_Protocol_match_match_proto_rawDescOnce.Do(func() {
		file_Protocol_match_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_Protocol_match_match_proto_rawDescData)
	})
	return file_Protocol_match_match_proto_rawDescData
}

var file_Protocol_match_match_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Protocol_match_match_proto_goTypes = []interface{}{
	(*MatchRequest)(nil),       // 0: protocol.match.MatchRequest
	(*MatchReply)(nil),         // 1: protocol.match.MatchReply
	(*CancelMatchRequest)(nil), // 2: protocol.match.CancelMatchRequest
	(*CancelMatchReply)(nil),   // 3: protocol.match.CancelMatchReply
	(common.RoomMode)(0),       // 4: protocol.common.RoomMode
}
var file_Protocol_match_match_proto_depIdxs = []int32{
	4, // 0: protocol.match.MatchRequest.mode:type_name -> protocol.common.RoomMode
	0, // 1: protocol.match.Match.Match:input_type -> protocol.match.MatchRequest
	2, // 2: protocol.match.Match.CancelMatch:input_type -> protocol.match.CancelMatchRequest
	1, // 3: protocol.match.Match.Match:output_type -> protocol.match.MatchReply
	3, // 4: protocol.match.Match.CancelMatch:output_type -> protocol.match.CancelMatchReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Protocol_match_match_proto_init() }
func file_Protocol_match_match_proto_init() {
	if File_Protocol_match_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Protocol_match_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protocol_match_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protocol_match_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protocol_match_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Protocol_match_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Protocol_match_match_proto_goTypes,
		DependencyIndexes: file_Protocol_match_match_proto_depIdxs,
		MessageInfos:      file_Protocol_match_match_proto_msgTypes,
	}.Build()
	File_Protocol_match_match_proto = out.File
	file_Protocol_match_match_proto_rawDesc = nil
	file_Protocol_match_match_proto_goTypes = nil
	file_Protocol_match_match_proto_depIdxs = nil
}
