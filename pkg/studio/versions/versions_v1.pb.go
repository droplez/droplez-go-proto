/// This file has messages for describing droplez-projects service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: studio/versions/versions_v1.proto

package versions

import (
	common "github.com/droplez/droplez-go-proto/pkg/common"
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

type VersionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VersionId) Reset() {
	*x = VersionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_studio_versions_versions_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionId) ProtoMessage() {}

func (x *VersionId) ProtoReflect() protoreflect.Message {
	mi := &file_studio_versions_versions_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionId.ProtoReflect.Descriptor instead.
func (*VersionId) Descriptor() ([]byte, []int) {
	return file_studio_versions_versions_v1_proto_rawDescGZIP(), []int{0}
}

func (x *VersionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VersionMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ProjectId  string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ObjectName string `protobuf:"bytes,3,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	Message    string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	UploadedAt string `protobuf:"bytes,6,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
}

func (x *VersionMeta) Reset() {
	*x = VersionMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_studio_versions_versions_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionMeta) ProtoMessage() {}

func (x *VersionMeta) ProtoReflect() protoreflect.Message {
	mi := &file_studio_versions_versions_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionMeta.ProtoReflect.Descriptor instead.
func (*VersionMeta) Descriptor() ([]byte, []int) {
	return file_studio_versions_versions_v1_proto_rawDescGZIP(), []int{1}
}

func (x *VersionMeta) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *VersionMeta) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *VersionMeta) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *VersionMeta) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VersionMeta) GetUploadedAt() string {
	if x != nil {
		return x.UploadedAt
	}
	return ""
}

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *VersionId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata *VersionMeta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_studio_versions_versions_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_studio_versions_versions_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_studio_versions_versions_v1_proto_rawDescGZIP(), []int{2}
}

func (x *VersionInfo) GetId() *VersionId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *VersionInfo) GetMetadata() *VersionMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ListOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *common.Paging `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListOptions) Reset() {
	*x = ListOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_studio_versions_versions_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOptions) ProtoMessage() {}

func (x *ListOptions) ProtoReflect() protoreflect.Message {
	mi := &file_studio_versions_versions_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOptions.ProtoReflect.Descriptor instead.
func (*ListOptions) Descriptor() ([]byte, []int) {
	return file_studio_versions_versions_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ListOptions) GetPaging() *common.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

var File_studio_versions_versions_v1_proto protoreflect.FileDescriptor

var file_studio_versions_versions_v1_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a,
	0x09, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x75, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x72, 0x6f,
	0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x32, 0xa5, 0x02,
	0x0a, 0x08, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1b, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x1d,
	0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1d, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x1d,
	0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65,
	0x7a, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x2f, 0x64, 0x72, 0x6f, 0x70,
	0x6c, 0x65, 0x7a, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_studio_versions_versions_v1_proto_rawDescOnce sync.Once
	file_studio_versions_versions_v1_proto_rawDescData = file_studio_versions_versions_v1_proto_rawDesc
)

func file_studio_versions_versions_v1_proto_rawDescGZIP() []byte {
	file_studio_versions_versions_v1_proto_rawDescOnce.Do(func() {
		file_studio_versions_versions_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_studio_versions_versions_v1_proto_rawDescData)
	})
	return file_studio_versions_versions_v1_proto_rawDescData
}

var file_studio_versions_versions_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_studio_versions_versions_v1_proto_goTypes = []interface{}{
	(*VersionId)(nil),     // 0: droplez_versions.VersionId
	(*VersionMeta)(nil),   // 1: droplez_versions.VersionMeta
	(*VersionInfo)(nil),   // 2: droplez_versions.VersionInfo
	(*ListOptions)(nil),   // 3: droplez_versions.ListOptions
	(*common.Paging)(nil), // 4: common.Paging
}
var file_studio_versions_versions_v1_proto_depIdxs = []int32{
	0, // 0: droplez_versions.VersionInfo.id:type_name -> droplez_versions.VersionId
	1, // 1: droplez_versions.VersionInfo.metadata:type_name -> droplez_versions.VersionMeta
	4, // 2: droplez_versions.ListOptions.paging:type_name -> common.Paging
	0, // 3: droplez_versions.Versions.Get:input_type -> droplez_versions.VersionId
	3, // 4: droplez_versions.Versions.List:input_type -> droplez_versions.ListOptions
	1, // 5: droplez_versions.Versions.Create:input_type -> droplez_versions.VersionMeta
	2, // 6: droplez_versions.Versions.Update:input_type -> droplez_versions.VersionInfo
	2, // 7: droplez_versions.Versions.Get:output_type -> droplez_versions.VersionInfo
	2, // 8: droplez_versions.Versions.List:output_type -> droplez_versions.VersionInfo
	2, // 9: droplez_versions.Versions.Create:output_type -> droplez_versions.VersionInfo
	2, // 10: droplez_versions.Versions.Update:output_type -> droplez_versions.VersionInfo
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_studio_versions_versions_v1_proto_init() }
func file_studio_versions_versions_v1_proto_init() {
	if File_studio_versions_versions_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_studio_versions_versions_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionId); i {
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
		file_studio_versions_versions_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionMeta); i {
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
		file_studio_versions_versions_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
		file_studio_versions_versions_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOptions); i {
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
			RawDescriptor: file_studio_versions_versions_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_studio_versions_versions_v1_proto_goTypes,
		DependencyIndexes: file_studio_versions_versions_v1_proto_depIdxs,
		MessageInfos:      file_studio_versions_versions_v1_proto_msgTypes,
	}.Build()
	File_studio_versions_versions_v1_proto = out.File
	file_studio_versions_versions_v1_proto_rawDesc = nil
	file_studio_versions_versions_v1_proto_goTypes = nil
	file_studio_versions_versions_v1_proto_depIdxs = nil
}
