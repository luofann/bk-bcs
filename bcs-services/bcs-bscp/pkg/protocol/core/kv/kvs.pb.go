// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: kvs.proto

package pbkv

import (
	base "bscp.io/pkg/protocol/core/base"
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

// Kv source resource reference: pkg/dal/table/kvs.go
type Kv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	KvState    string         `protobuf:"bytes,2,opt,name=kv_state,json=kvState,proto3" json:"kv_state,omitempty"`
	Spec       *KvSpec        `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *KvAttachment  `protobuf:"bytes,4,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Kv) Reset() {
	*x = Kv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kv) ProtoMessage() {}

func (x *Kv) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kv.ProtoReflect.Descriptor instead.
func (*Kv) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{0}
}

func (x *Kv) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Kv) GetKvState() string {
	if x != nil {
		return x.KvState
	}
	return ""
}

func (x *Kv) GetSpec() *KvSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Kv) GetAttachment() *KvAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Kv) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// KvSpec source resource reference: pkg/dal/table/kvs.go
type KvSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KvType string `protobuf:"bytes,2,opt,name=kv_type,json=kvType,proto3" json:"kv_type,omitempty"`
	Value  string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KvSpec) Reset() {
	*x = KvSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvSpec) ProtoMessage() {}

func (x *KvSpec) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvSpec.ProtoReflect.Descriptor instead.
func (*KvSpec) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{1}
}

func (x *KvSpec) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvSpec) GetKvType() string {
	if x != nil {
		return x.KvType
	}
	return ""
}

func (x *KvSpec) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// KvAttachment source resource reference: pkg/dal/table/kvs.go
type KvAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *KvAttachment) Reset() {
	*x = KvAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvAttachment) ProtoMessage() {}

func (x *KvAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvAttachment.ProtoReflect.Descriptor instead.
func (*KvAttachment) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{2}
}

func (x *KvAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *KvAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

var File_kvs_proto protoreflect.FileDescriptor

var file_kvs_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6b, 0x76, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x6b,
	0x76, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x02, 0x4b, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6b,
	0x76, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x76, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x6b, 0x76, 0x2e, 0x4b, 0x76, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x6b, 0x76, 0x2e, 0x4b, 0x76, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x06, 0x4b, 0x76,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x76, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x4b, 0x76, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x42, 0x23, 0x5a, 0x21, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x6b, 0x76, 0x3b, 0x70, 0x62, 0x6b, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kvs_proto_rawDescOnce sync.Once
	file_kvs_proto_rawDescData = file_kvs_proto_rawDesc
)

func file_kvs_proto_rawDescGZIP() []byte {
	file_kvs_proto_rawDescOnce.Do(func() {
		file_kvs_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvs_proto_rawDescData)
	})
	return file_kvs_proto_rawDescData
}

var file_kvs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kvs_proto_goTypes = []interface{}{
	(*Kv)(nil),            // 0: pbkv.Kv
	(*KvSpec)(nil),        // 1: pbkv.KvSpec
	(*KvAttachment)(nil),  // 2: pbkv.KvAttachment
	(*base.Revision)(nil), // 3: pbbase.Revision
}
var file_kvs_proto_depIdxs = []int32{
	1, // 0: pbkv.Kv.spec:type_name -> pbkv.KvSpec
	2, // 1: pbkv.Kv.attachment:type_name -> pbkv.KvAttachment
	3, // 2: pbkv.Kv.revision:type_name -> pbbase.Revision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kvs_proto_init() }
func file_kvs_proto_init() {
	if File_kvs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kvs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kv); i {
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
		file_kvs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvSpec); i {
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
		file_kvs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvAttachment); i {
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
			RawDescriptor: file_kvs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kvs_proto_goTypes,
		DependencyIndexes: file_kvs_proto_depIdxs,
		MessageInfos:      file_kvs_proto_msgTypes,
	}.Build()
	File_kvs_proto = out.File
	file_kvs_proto_rawDesc = nil
	file_kvs_proto_goTypes = nil
	file_kvs_proto_depIdxs = nil
}
