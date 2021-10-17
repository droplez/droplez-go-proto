/// This file has messages for describing droplez-uploader service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: uploader/uploader_v1.proto

package proto_uploader

import (
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

type Metadata_ContentType int32

const (
	Metadata_CONTENT_TYPE_UNSPECIFIED Metadata_ContentType = 0
	Metadata_CONTENT_TYPE_ARCHIVE     Metadata_ContentType = 1
	Metadata_CONTENT_TYPE_PICTURE     Metadata_ContentType = 2
	Metadata_CONTENT_TYPE_AUDIO       Metadata_ContentType = 3
)

// Enum value maps for Metadata_ContentType.
var (
	Metadata_ContentType_name = map[int32]string{
		0: "CONTENT_TYPE_UNSPECIFIED",
		1: "CONTENT_TYPE_ARCHIVE",
		2: "CONTENT_TYPE_PICTURE",
		3: "CONTENT_TYPE_AUDIO",
	}
	Metadata_ContentType_value = map[string]int32{
		"CONTENT_TYPE_UNSPECIFIED": 0,
		"CONTENT_TYPE_ARCHIVE":     1,
		"CONTENT_TYPE_PICTURE":     2,
		"CONTENT_TYPE_AUDIO":       3,
	}
)

func (x Metadata_ContentType) Enum() *Metadata_ContentType {
	p := new(Metadata_ContentType)
	*p = x
	return p
}

func (x Metadata_ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_uploader_uploader_v1_proto_enumTypes[0].Descriptor()
}

func (Metadata_ContentType) Type() protoreflect.EnumType {
	return &file_uploader_uploader_v1_proto_enumTypes[0]
}

func (x Metadata_ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Metadata_ContentType.Descriptor instead.
func (Metadata_ContentType) EnumDescriptor() ([]byte, []int) {
	return file_uploader_uploader_v1_proto_rawDescGZIP(), []int{1, 0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content      []byte    `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileMetadata *Metadata `protobuf:"bytes,4,opt,name=file_metadata,json=fileMetadata,proto3" json:"file_metadata,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetFileMetadata() *Metadata {
	if x != nil {
		return x.FileMetadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContentType Metadata_ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3,enum=uploader.Metadata_ContentType" json:"content_type,omitempty"`
	LocalName   string               `protobuf:"bytes,3,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	FileSize    int64                `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	UserId      string               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_v1_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetContentType() Metadata_ContentType {
	if x != nil {
		return x.ContentType
	}
	return Metadata_CONTENT_TYPE_UNSPECIFIED
}

func (x *Metadata) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Metadata) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Metadata) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UploadedFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object  string `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UploadedFileData) Reset() {
	*x = UploadedFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadedFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadedFileData) ProtoMessage() {}

func (x *UploadedFileData) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadedFileData.ProtoReflect.Descriptor instead.
func (*UploadedFileData) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UploadedFileData) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *UploadedFileData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DownloadableLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DownloadableLink) Reset() {
	*x = DownloadableLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadableLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadableLink) ProtoMessage() {}

func (x *DownloadableLink) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadableLink.ProtoReflect.Descriptor instead.
func (*DownloadableLink) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_v1_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadableLink) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_uploader_uploader_v1_proto protoreflect.FileDescriptor

var file_uploader_uploader_v1_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xaf, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x49, 0x43, 0x54, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x44,
	0x49, 0x4f, 0x10, 0x03, 0x22, 0x44, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x10, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x32, 0x92, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x37, 0x0a,
	0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x6c, 0x65, 0x7a, 0x2f, 0x64, 0x72, 0x6f, 0x70,
	0x6c, 0x65, 0x7a, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uploader_uploader_v1_proto_rawDescOnce sync.Once
	file_uploader_uploader_v1_proto_rawDescData = file_uploader_uploader_v1_proto_rawDesc
)

func file_uploader_uploader_v1_proto_rawDescGZIP() []byte {
	file_uploader_uploader_v1_proto_rawDescOnce.Do(func() {
		file_uploader_uploader_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_uploader_uploader_v1_proto_rawDescData)
	})
	return file_uploader_uploader_v1_proto_rawDescData
}

var file_uploader_uploader_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_uploader_uploader_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_uploader_uploader_v1_proto_goTypes = []interface{}{
	(Metadata_ContentType)(0), // 0: uploader.Metadata.ContentType
	(*Chunk)(nil),             // 1: uploader.Chunk
	(*Metadata)(nil),          // 2: uploader.Metadata
	(*UploadedFileData)(nil),  // 3: uploader.UploadedFileData
	(*DownloadableLink)(nil),  // 4: uploader.DownloadableLink
}
var file_uploader_uploader_v1_proto_depIdxs = []int32{
	2, // 0: uploader.Chunk.file_metadata:type_name -> uploader.Metadata
	0, // 1: uploader.Metadata.content_type:type_name -> uploader.Metadata.ContentType
	1, // 2: uploader.Uploader.Upload:input_type -> uploader.Chunk
	3, // 3: uploader.Uploader.GetDownloadableLink:input_type -> uploader.UploadedFileData
	3, // 4: uploader.Uploader.Upload:output_type -> uploader.UploadedFileData
	4, // 5: uploader.Uploader.GetDownloadableLink:output_type -> uploader.DownloadableLink
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_uploader_uploader_v1_proto_init() }
func file_uploader_uploader_v1_proto_init() {
	if File_uploader_uploader_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uploader_uploader_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_uploader_uploader_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_uploader_uploader_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadedFileData); i {
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
		file_uploader_uploader_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadableLink); i {
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
			RawDescriptor: file_uploader_uploader_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uploader_uploader_v1_proto_goTypes,
		DependencyIndexes: file_uploader_uploader_v1_proto_depIdxs,
		EnumInfos:         file_uploader_uploader_v1_proto_enumTypes,
		MessageInfos:      file_uploader_uploader_v1_proto_msgTypes,
	}.Build()
	File_uploader_uploader_v1_proto = out.File
	file_uploader_uploader_v1_proto_rawDesc = nil
	file_uploader_uploader_v1_proto_goTypes = nil
	file_uploader_uploader_v1_proto_depIdxs = nil
}
