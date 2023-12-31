// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/kv.proto

package umamaheswari76

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

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_kv_proto_rawDescGZIP(), []int{0}
}

func (x *GetFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID string `protobuf:"bytes,1,opt,name=FileID,proto3" json:"FileID,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_kv_proto_rawDescGZIP(), []int{1}
}

func (x *GetFileResponse) GetFileID() string {
	if x != nil {
		return x.FileID
	}
	return ""
}

var File_proto_kv_proto protoreflect.FileDescriptor

var file_proto_kv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6b, 0x76, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x32, 0x4d, 0x0a, 0x0d, 0x4b, 0x76, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x6b, 0x76, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x76, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x6d, 0x61, 0x6d, 0x61, 0x68, 0x65, 0x73, 0x77, 0x61, 0x72, 0x69, 0x37, 0x36, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kv_proto_rawDescOnce sync.Once
	file_proto_kv_proto_rawDescData = file_proto_kv_proto_rawDesc
)

func file_proto_kv_proto_rawDescGZIP() []byte {
	file_proto_kv_proto_rawDescOnce.Do(func() {
		file_proto_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kv_proto_rawDescData)
	})
	return file_proto_kv_proto_rawDescData
}

var file_proto_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_kv_proto_goTypes = []interface{}{
	(*GetFileRequest)(nil),  // 0: kv_grpc.GetFileRequest
	(*GetFileResponse)(nil), // 1: kv_grpc.GetFileResponse
}
var file_proto_kv_proto_depIdxs = []int32{
	0, // 0: kv_grpc.KvGrpcService.GetFile:input_type -> kv_grpc.GetFileRequest
	1, // 1: kv_grpc.KvGrpcService.GetFile:output_type -> kv_grpc.GetFileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kv_proto_init() }
func file_proto_kv_proto_init() {
	if File_proto_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_proto_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
			RawDescriptor: file_proto_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kv_proto_goTypes,
		DependencyIndexes: file_proto_kv_proto_depIdxs,
		MessageInfos:      file_proto_kv_proto_msgTypes,
	}.Build()
	File_proto_kv_proto = out.File
	file_proto_kv_proto_rawDesc = nil
	file_proto_kv_proto_goTypes = nil
	file_proto_kv_proto_depIdxs = nil
}
