// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: hook_release.proto

package pbhr

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

// HookRelease source resource reference: pkg/dal/table/hook_release.go
type HookRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *HookReleaseSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *HookReleaseAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision         `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *HookRelease) Reset() {
	*x = HookRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_release_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookRelease) ProtoMessage() {}

func (x *HookRelease) ProtoReflect() protoreflect.Message {
	mi := &file_hook_release_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookRelease.ProtoReflect.Descriptor instead.
func (*HookRelease) Descriptor() ([]byte, []int) {
	return file_hook_release_proto_rawDescGZIP(), []int{0}
}

func (x *HookRelease) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HookRelease) GetSpec() *HookReleaseSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *HookRelease) GetAttachment() *HookReleaseAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *HookRelease) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// HookReleaseAttachment source resource reference: pkg/dal/table/hook_release.go
type HookReleaseSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	PublishNum uint32 `protobuf:"varint,3,opt,name=publish_num,json=publishNum,proto3" json:"publish_num,omitempty"`
	PubState   string `protobuf:"bytes,4,opt,name=pub_state,json=pubState,proto3" json:"pub_state,omitempty"`
	Memo       string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *HookReleaseSpec) Reset() {
	*x = HookReleaseSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_release_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookReleaseSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookReleaseSpec) ProtoMessage() {}

func (x *HookReleaseSpec) ProtoReflect() protoreflect.Message {
	mi := &file_hook_release_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookReleaseSpec.ProtoReflect.Descriptor instead.
func (*HookReleaseSpec) Descriptor() ([]byte, []int) {
	return file_hook_release_proto_rawDescGZIP(), []int{1}
}

func (x *HookReleaseSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HookReleaseSpec) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *HookReleaseSpec) GetPublishNum() uint32 {
	if x != nil {
		return x.PublishNum
	}
	return 0
}

func (x *HookReleaseSpec) GetPubState() string {
	if x != nil {
		return x.PubState
	}
	return ""
}

func (x *HookReleaseSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// HookReleaseAttachment source resource reference: pkg/dal/table/hook_release.go
type HookReleaseAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId  uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	HookId uint32 `protobuf:"varint,2,opt,name=hook_id,json=hookId,proto3" json:"hook_id,omitempty"`
}

func (x *HookReleaseAttachment) Reset() {
	*x = HookReleaseAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hook_release_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HookReleaseAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HookReleaseAttachment) ProtoMessage() {}

func (x *HookReleaseAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_hook_release_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HookReleaseAttachment.ProtoReflect.Descriptor instead.
func (*HookReleaseAttachment) Descriptor() ([]byte, []int) {
	return file_hook_release_proto_rawDescGZIP(), []int{2}
}

func (x *HookReleaseAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *HookReleaseAttachment) GetHookId() uint32 {
	if x != nil {
		return x.HookId
	}
	return 0
}

var File_hook_release_proto protoreflect.FileDescriptor

var file_hook_release_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x68, 0x72, 0x1a, 0x29, 0x62, 0x73, 0x63, 0x70,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x68, 0x72, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x68, 0x72, 0x2e, 0x48, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x0f,
	0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x75, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22,
	0x47, 0x0a, 0x15, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x62, 0x73, 0x63, 0x70,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x3b, 0x70, 0x62, 0x68, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hook_release_proto_rawDescOnce sync.Once
	file_hook_release_proto_rawDescData = file_hook_release_proto_rawDesc
)

func file_hook_release_proto_rawDescGZIP() []byte {
	file_hook_release_proto_rawDescOnce.Do(func() {
		file_hook_release_proto_rawDescData = protoimpl.X.CompressGZIP(file_hook_release_proto_rawDescData)
	})
	return file_hook_release_proto_rawDescData
}

var file_hook_release_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hook_release_proto_goTypes = []interface{}{
	(*HookRelease)(nil),           // 0: pbhr.HookRelease
	(*HookReleaseSpec)(nil),       // 1: pbhr.HookReleaseSpec
	(*HookReleaseAttachment)(nil), // 2: pbhr.HookReleaseAttachment
	(*base.Revision)(nil),         // 3: pbbase.Revision
}
var file_hook_release_proto_depIdxs = []int32{
	1, // 0: pbhr.HookRelease.spec:type_name -> pbhr.HookReleaseSpec
	2, // 1: pbhr.HookRelease.attachment:type_name -> pbhr.HookReleaseAttachment
	3, // 2: pbhr.HookRelease.revision:type_name -> pbbase.Revision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hook_release_proto_init() }
func file_hook_release_proto_init() {
	if File_hook_release_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hook_release_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookRelease); i {
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
		file_hook_release_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookReleaseSpec); i {
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
		file_hook_release_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HookReleaseAttachment); i {
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
			RawDescriptor: file_hook_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hook_release_proto_goTypes,
		DependencyIndexes: file_hook_release_proto_depIdxs,
		MessageInfos:      file_hook_release_proto_msgTypes,
	}.Build()
	File_hook_release_proto = out.File
	file_hook_release_proto_rawDesc = nil
	file_hook_release_proto_goTypes = nil
	file_hook_release_proto_depIdxs = nil
}
