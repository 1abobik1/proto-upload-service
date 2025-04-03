// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/upload_service/v1/upload_service.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// --- Общие сообщения ---
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId    string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename  string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"` // Оригинальное имя файла
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Size      uint64                 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"` // Размер в байтах
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FileInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// --- Upload ---
type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadRequest_Filename
	//	*UploadRequest_Chunk
	Data isUploadRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{1}
}

func (m *UploadRequest) GetData() isUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadRequest) GetFilename() string {
	if x, ok := x.GetData().(*UploadRequest_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*UploadRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isUploadRequest_Data interface {
	isUploadRequest_Data()
}

type UploadRequest_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3,oneof"` // Только в первом chunk
}

type UploadRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadRequest_Filename) isUploadRequest_Data() {}

func (*UploadRequest_Chunk) isUploadRequest_Data() {}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Size   uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UploadResponse) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// --- List Files ---
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{3}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

// --- Presigned URL ---
type DownloadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"` // инфа о файле
}

func (x *DownloadLinkRequest) Reset() {
	*x = DownloadLinkRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLinkRequest) ProtoMessage() {}

func (x *DownloadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadLinkRequest.ProtoReflect.Descriptor instead.
func (*DownloadLinkRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadLinkRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"` // Presigned URL от MinIO
}

func (x *DownloadLinkResponse) Reset() {
	*x = DownloadLinkResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLinkResponse) ProtoMessage() {}

func (x *DownloadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadLinkResponse.ProtoReflect.Descriptor instead.
func (*DownloadLinkResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// --- Zip Archive ---
type DownloadZipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIds []string `protobuf:"bytes,1,rep,name=file_ids,json=fileIds,proto3" json:"file_ids,omitempty"` // ID файлов для архивации
}

func (x *DownloadZipRequest) Reset() {
	*x = DownloadZipRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadZipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadZipRequest) ProtoMessage() {}

func (x *DownloadZipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadZipRequest.ProtoReflect.Descriptor instead.
func (*DownloadZipRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadZipRequest) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

type DownloadZipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"` // Порция zip-архива
}

func (x *DownloadZipResponse) Reset() {
	*x = DownloadZipResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadZipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadZipResponse) ProtoMessage() {}

func (x *DownloadZipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadZipResponse.ProtoReflect.Descriptor instead.
func (*DownloadZipResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadZipResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_proto_upload_service_v1_upload_service_proto protoreflect.FileDescriptor

var file_proto_upload_service_v1_upload_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x3d, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x41, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2f, 0x0a,
	0x12, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x2b,
	0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xf0, 0x02, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x26, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x69, 0x70, 0x12, 0x25, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x03,
	0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_upload_service_v1_upload_service_proto_rawDescOnce sync.Once
	file_proto_upload_service_v1_upload_service_proto_rawDescData = file_proto_upload_service_v1_upload_service_proto_rawDesc
)

func file_proto_upload_service_v1_upload_service_proto_rawDescGZIP() []byte {
	file_proto_upload_service_v1_upload_service_proto_rawDescOnce.Do(func() {
		file_proto_upload_service_v1_upload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_upload_service_v1_upload_service_proto_rawDescData)
	})
	return file_proto_upload_service_v1_upload_service_proto_rawDescData
}

var file_proto_upload_service_v1_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_upload_service_v1_upload_service_proto_goTypes = []any{
	(*FileInfo)(nil),              // 0: upload_service.v1.FileInfo
	(*UploadRequest)(nil),         // 1: upload_service.v1.UploadRequest
	(*UploadResponse)(nil),        // 2: upload_service.v1.UploadResponse
	(*ListRequest)(nil),           // 3: upload_service.v1.ListRequest
	(*ListResponse)(nil),          // 4: upload_service.v1.ListResponse
	(*DownloadLinkRequest)(nil),   // 5: upload_service.v1.DownloadLinkRequest
	(*DownloadLinkResponse)(nil),  // 6: upload_service.v1.DownloadLinkResponse
	(*DownloadZipRequest)(nil),    // 7: upload_service.v1.DownloadZipRequest
	(*DownloadZipResponse)(nil),   // 8: upload_service.v1.DownloadZipResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_proto_upload_service_v1_upload_service_proto_depIdxs = []int32{
	9, // 0: upload_service.v1.FileInfo.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: upload_service.v1.ListResponse.files:type_name -> upload_service.v1.FileInfo
	1, // 2: upload_service.v1.FileService.Upload:input_type -> upload_service.v1.UploadRequest
	3, // 3: upload_service.v1.FileService.ListFiles:input_type -> upload_service.v1.ListRequest
	5, // 4: upload_service.v1.FileService.GetDownloadLink:input_type -> upload_service.v1.DownloadLinkRequest
	7, // 5: upload_service.v1.FileService.DownloadZip:input_type -> upload_service.v1.DownloadZipRequest
	2, // 6: upload_service.v1.FileService.Upload:output_type -> upload_service.v1.UploadResponse
	4, // 7: upload_service.v1.FileService.ListFiles:output_type -> upload_service.v1.ListResponse
	6, // 8: upload_service.v1.FileService.GetDownloadLink:output_type -> upload_service.v1.DownloadLinkResponse
	8, // 9: upload_service.v1.FileService.DownloadZip:output_type -> upload_service.v1.DownloadZipResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_upload_service_v1_upload_service_proto_init() }
func file_proto_upload_service_v1_upload_service_proto_init() {
	if File_proto_upload_service_v1_upload_service_proto != nil {
		return
	}
	file_proto_upload_service_v1_upload_service_proto_msgTypes[1].OneofWrappers = []any{
		(*UploadRequest_Filename)(nil),
		(*UploadRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_upload_service_v1_upload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_upload_service_v1_upload_service_proto_goTypes,
		DependencyIndexes: file_proto_upload_service_v1_upload_service_proto_depIdxs,
		MessageInfos:      file_proto_upload_service_v1_upload_service_proto_msgTypes,
	}.Build()
	File_proto_upload_service_v1_upload_service_proto = out.File
	file_proto_upload_service_v1_upload_service_proto_rawDesc = nil
	file_proto_upload_service_v1_upload_service_proto_goTypes = nil
	file_proto_upload_service_v1_upload_service_proto_depIdxs = nil
}
