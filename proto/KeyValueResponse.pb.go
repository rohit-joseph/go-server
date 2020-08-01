// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: KeyValueResponse.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type KVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode          int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Pid              int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Version          int32  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	OverloadWaitTime int32  `protobuf:"varint,5,opt,name=overloadWaitTime,proto3" json:"overloadWaitTime,omitempty"`
	MembershipCount  int32  `protobuf:"varint,6,opt,name=membershipCount,proto3" json:"membershipCount,omitempty"`
}

func (x *KVResponse) Reset() {
	*x = KVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeyValueResponse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVResponse) ProtoMessage() {}

func (x *KVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_KeyValueResponse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVResponse.ProtoReflect.Descriptor instead.
func (*KVResponse) Descriptor() ([]byte, []int) {
	return file_KeyValueResponse_proto_rawDescGZIP(), []int{0}
}

func (x *KVResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *KVResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KVResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *KVResponse) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *KVResponse) GetOverloadWaitTime() int32 {
	if x != nil {
		return x.OverloadWaitTime
	}
	return 0
}

func (x *KVResponse) GetMembershipCount() int32 {
	if x != nil {
		return x.MembershipCount
	}
	return 0
}

var File_KeyValueResponse_proto protoreflect.FileDescriptor

var file_KeyValueResponse_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbe, 0x01, 0x0a, 0x0a, 0x4b, 0x56, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x76,
	0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x61,
	0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KeyValueResponse_proto_rawDescOnce sync.Once
	file_KeyValueResponse_proto_rawDescData = file_KeyValueResponse_proto_rawDesc
)

func file_KeyValueResponse_proto_rawDescGZIP() []byte {
	file_KeyValueResponse_proto_rawDescOnce.Do(func() {
		file_KeyValueResponse_proto_rawDescData = protoimpl.X.CompressGZIP(file_KeyValueResponse_proto_rawDescData)
	})
	return file_KeyValueResponse_proto_rawDescData
}

var file_KeyValueResponse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KeyValueResponse_proto_goTypes = []interface{}{
	(*KVResponse)(nil), // 0: proto.KVResponse
}
var file_KeyValueResponse_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KeyValueResponse_proto_init() }
func file_KeyValueResponse_proto_init() {
	if File_KeyValueResponse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KeyValueResponse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVResponse); i {
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
			RawDescriptor: file_KeyValueResponse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KeyValueResponse_proto_goTypes,
		DependencyIndexes: file_KeyValueResponse_proto_depIdxs,
		MessageInfos:      file_KeyValueResponse_proto_msgTypes,
	}.Build()
	File_KeyValueResponse_proto = out.File
	file_KeyValueResponse_proto_rawDesc = nil
	file_KeyValueResponse_proto_goTypes = nil
	file_KeyValueResponse_proto_depIdxs = nil
}
