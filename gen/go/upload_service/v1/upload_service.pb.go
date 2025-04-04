// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.3
// source: proto/upload_service/v1/upload_service.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`                    // оригинальное имя файла
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // дата создания
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // дата обновления файла
	Size          uint64                 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *FileInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *FileInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type UploadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*UploadRequest_Filename
	//	*UploadRequest_Chunk
	Data          isUploadRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *UploadRequest) GetData() isUploadRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadRequest) GetFilename() string {
	if x != nil {
		if x, ok := x.Data.(*UploadRequest_Filename); ok {
			return x.Filename
		}
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*UploadRequest_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

type isUploadRequest_Data interface {
	isUploadRequest_Data()
}

type UploadRequest_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3,oneof"` // Только в первом chunk имя файла
}

type UploadRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadRequest_Filename) isUploadRequest_Data() {}

func (*UploadRequest_Chunk) isUploadRequest_Data() {}

type UploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Size          uint64                 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type UpdateFileRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*UpdateFileRequest_FileId
	//	*UpdateFileRequest_Chunk
	Data          isUpdateFileRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFileRequest) GetData() isUpdateFileRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateFileRequest) GetFileId() string {
	if x != nil {
		if x, ok := x.Data.(*UpdateFileRequest_FileId); ok {
			return x.FileId
		}
	}
	return ""
}

func (x *UpdateFileRequest) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*UpdateFileRequest_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

type isUpdateFileRequest_Data interface {
	isUpdateFileRequest_Data()
}

type UpdateFileRequest_FileId struct {
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3,oneof"` // ID файла для обновления (только в первом chunk)
}

type UpdateFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UpdateFileRequest_FileId) isUpdateFileRequest_Data() {}

func (*UpdateFileRequest_Chunk) isUpdateFileRequest_Data() {}

type UpdateFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	NewSize       uint64                 `protobuf:"varint,2,opt,name=new_size,json=newSize,proto3" json:"new_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFileResponse) Reset() {
	*x = UpdateFileResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileResponse) ProtoMessage() {}

func (x *UpdateFileResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateFileResponse.ProtoReflect.Descriptor instead.
func (*UpdateFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFileResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UpdateFileResponse) GetNewSize() uint64 {
	if x != nil {
		return x.NewSize
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{5}
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*FileInfo            `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type DownloadLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadLinkRequest) Reset() {
	*x = DownloadLinkRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLinkRequest) ProtoMessage() {}

func (x *DownloadLinkRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadLinkRequest.ProtoReflect.Descriptor instead.
func (*DownloadLinkRequest) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadLinkRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"` // Presigned URL от MinIO
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadLinkResponse) Reset() {
	*x = DownloadLinkResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLinkResponse) ProtoMessage() {}

func (x *DownloadLinkResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadLinkResponse.ProtoReflect.Descriptor instead.
func (*DownloadLinkResponse) Descriptor() ([]byte, []int) {
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DownloadZipRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileIds       []string               `protobuf:"bytes,1,rep,name=file_ids,json=fileIds,proto3" json:"file_ids,omitempty"` // ID файлов для архивации
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadZipRequest) Reset() {
	*x = DownloadZipRequest{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadZipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadZipRequest) ProtoMessage() {}

func (x *DownloadZipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[9]
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
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadZipRequest) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

type DownloadZipResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"` // Порция zip-архива
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadZipResponse) Reset() {
	*x = DownloadZipResponse{}
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadZipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadZipResponse) ProtoMessage() {}

func (x *DownloadZipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_upload_service_v1_upload_service_proto_msgTypes[10]
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
	return file_proto_upload_service_v1_upload_service_proto_rawDescGZIP(), []int{10}
}

func (x *DownloadZipResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_proto_upload_service_v1_upload_service_proto protoreflect.FileDescriptor

const file_proto_upload_service_v1_upload_service_proto_rawDesc = "" +
	"\n" +
	",proto/upload_service/v1/upload_service.proto\x12\x11upload_service.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc9\x01\n" +
	"\bFileInfo\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x12\n" +
	"\x04size\x18\x05 \x01(\x04R\x04size\"M\n" +
	"\rUploadRequest\x12\x1c\n" +
	"\bfilename\x18\x01 \x01(\tH\x00R\bfilename\x12\x16\n" +
	"\x05chunk\x18\x02 \x01(\fH\x00R\x05chunkB\x06\n" +
	"\x04data\"=\n" +
	"\x0eUploadResponse\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x12\n" +
	"\x04size\x18\x02 \x01(\x04R\x04size\"N\n" +
	"\x11UpdateFileRequest\x12\x19\n" +
	"\afile_id\x18\x01 \x01(\tH\x00R\x06fileId\x12\x16\n" +
	"\x05chunk\x18\x02 \x01(\fH\x00R\x05chunkB\x06\n" +
	"\x04data\"H\n" +
	"\x12UpdateFileResponse\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x19\n" +
	"\bnew_size\x18\x02 \x01(\x04R\anewSize\"\r\n" +
	"\vListRequest\"A\n" +
	"\fListResponse\x121\n" +
	"\x05files\x18\x01 \x03(\v2\x1b.upload_service.v1.FileInfoR\x05files\".\n" +
	"\x13DownloadLinkRequest\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\"(\n" +
	"\x14DownloadLinkResponse\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\"/\n" +
	"\x12DownloadZipRequest\x12\x19\n" +
	"\bfile_ids\x18\x01 \x03(\tR\afileIds\"+\n" +
	"\x13DownloadZipResponse\x12\x14\n" +
	"\x05chunk\x18\x01 \x01(\fR\x05chunk2\xcd\x03\n" +
	"\vFileService\x12O\n" +
	"\x06Upload\x12 .upload_service.v1.UploadRequest\x1a!.upload_service.v1.UploadResponse(\x01\x12[\n" +
	"\n" +
	"UpdateFile\x12$.upload_service.v1.UpdateFileRequest\x1a%.upload_service.v1.UpdateFileResponse(\x01\x12L\n" +
	"\tListFiles\x12\x1e.upload_service.v1.ListRequest\x1a\x1f.upload_service.v1.ListResponse\x12b\n" +
	"\x0fGetDownloadLink\x12&.upload_service.v1.DownloadLinkRequest\x1a'.upload_service.v1.DownloadLinkResponse\x12^\n" +
	"\vDownloadZip\x12%.upload_service.v1.DownloadZipRequest\x1a&.upload_service.v1.DownloadZipResponse0\x01B\x03Z\x01.b\x06proto3"

var (
	file_proto_upload_service_v1_upload_service_proto_rawDescOnce sync.Once
	file_proto_upload_service_v1_upload_service_proto_rawDescData []byte
)

func file_proto_upload_service_v1_upload_service_proto_rawDescGZIP() []byte {
	file_proto_upload_service_v1_upload_service_proto_rawDescOnce.Do(func() {
		file_proto_upload_service_v1_upload_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_upload_service_v1_upload_service_proto_rawDesc), len(file_proto_upload_service_v1_upload_service_proto_rawDesc)))
	})
	return file_proto_upload_service_v1_upload_service_proto_rawDescData
}

var file_proto_upload_service_v1_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_upload_service_v1_upload_service_proto_goTypes = []any{
	(*FileInfo)(nil),              // 0: upload_service.v1.FileInfo
	(*UploadRequest)(nil),         // 1: upload_service.v1.UploadRequest
	(*UploadResponse)(nil),        // 2: upload_service.v1.UploadResponse
	(*UpdateFileRequest)(nil),     // 3: upload_service.v1.UpdateFileRequest
	(*UpdateFileResponse)(nil),    // 4: upload_service.v1.UpdateFileResponse
	(*ListRequest)(nil),           // 5: upload_service.v1.ListRequest
	(*ListResponse)(nil),          // 6: upload_service.v1.ListResponse
	(*DownloadLinkRequest)(nil),   // 7: upload_service.v1.DownloadLinkRequest
	(*DownloadLinkResponse)(nil),  // 8: upload_service.v1.DownloadLinkResponse
	(*DownloadZipRequest)(nil),    // 9: upload_service.v1.DownloadZipRequest
	(*DownloadZipResponse)(nil),   // 10: upload_service.v1.DownloadZipResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_proto_upload_service_v1_upload_service_proto_depIdxs = []int32{
	11, // 0: upload_service.v1.FileInfo.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: upload_service.v1.FileInfo.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: upload_service.v1.ListResponse.files:type_name -> upload_service.v1.FileInfo
	1,  // 3: upload_service.v1.FileService.Upload:input_type -> upload_service.v1.UploadRequest
	3,  // 4: upload_service.v1.FileService.UpdateFile:input_type -> upload_service.v1.UpdateFileRequest
	5,  // 5: upload_service.v1.FileService.ListFiles:input_type -> upload_service.v1.ListRequest
	7,  // 6: upload_service.v1.FileService.GetDownloadLink:input_type -> upload_service.v1.DownloadLinkRequest
	9,  // 7: upload_service.v1.FileService.DownloadZip:input_type -> upload_service.v1.DownloadZipRequest
	2,  // 8: upload_service.v1.FileService.Upload:output_type -> upload_service.v1.UploadResponse
	4,  // 9: upload_service.v1.FileService.UpdateFile:output_type -> upload_service.v1.UpdateFileResponse
	6,  // 10: upload_service.v1.FileService.ListFiles:output_type -> upload_service.v1.ListResponse
	8,  // 11: upload_service.v1.FileService.GetDownloadLink:output_type -> upload_service.v1.DownloadLinkResponse
	10, // 12: upload_service.v1.FileService.DownloadZip:output_type -> upload_service.v1.DownloadZipResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
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
	file_proto_upload_service_v1_upload_service_proto_msgTypes[3].OneofWrappers = []any{
		(*UpdateFileRequest_FileId)(nil),
		(*UpdateFileRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_upload_service_v1_upload_service_proto_rawDesc), len(file_proto_upload_service_v1_upload_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_upload_service_v1_upload_service_proto_goTypes,
		DependencyIndexes: file_proto_upload_service_v1_upload_service_proto_depIdxs,
		MessageInfos:      file_proto_upload_service_v1_upload_service_proto_msgTypes,
	}.Build()
	File_proto_upload_service_v1_upload_service_proto = out.File
	file_proto_upload_service_v1_upload_service_proto_goTypes = nil
	file_proto_upload_service_v1_upload_service_proto_depIdxs = nil
}
