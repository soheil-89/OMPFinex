// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/chunk.proto

package ChunkPb

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

type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha256    string `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Id        int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Size      int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ChunkSize int32  `protobuf:"varint,4,opt,name=chunkSize,proto3" json:"chunkSize,omitempty"`
	Data      string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_proto_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkRequest) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *ChunkRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChunkRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ChunkRequest) GetChunkSize() int32 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *ChunkRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receive bool `protobuf:"varint,1,opt,name=receive,proto3" json:"receive,omitempty"`
}

func (x *ChunkResponse) Reset() {
	*x = ChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkResponse) ProtoMessage() {}

func (x *ChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkResponse.ProtoReflect.Descriptor instead.
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return file_proto_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkResponse) GetReceive() bool {
	if x != nil {
		return x.Receive
	}
	return false
}

var File_proto_chunk_proto protoreflect.FileDescriptor

var file_proto_chunk_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x7c, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a,
	0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x32, 0x46, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x3d, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x50, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chunk_proto_rawDescOnce sync.Once
	file_proto_chunk_proto_rawDescData = file_proto_chunk_proto_rawDesc
)

func file_proto_chunk_proto_rawDescGZIP() []byte {
	file_proto_chunk_proto_rawDescOnce.Do(func() {
		file_proto_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chunk_proto_rawDescData)
	})
	return file_proto_chunk_proto_rawDescData
}

var file_proto_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chunk_proto_goTypes = []interface{}{
	(*ChunkRequest)(nil),  // 0: grpcServer.ChunkRequest
	(*ChunkResponse)(nil), // 1: grpcServer.ChunkResponse
}
var file_proto_chunk_proto_depIdxs = []int32{
	0, // 0: grpcServer.Chunk.send:input_type -> grpcServer.ChunkRequest
	1, // 1: grpcServer.Chunk.send:output_type -> grpcServer.ChunkResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chunk_proto_init() }
func file_proto_chunk_proto_init() {
	if File_proto_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequest); i {
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
		file_proto_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkResponse); i {
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
			RawDescriptor: file_proto_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chunk_proto_goTypes,
		DependencyIndexes: file_proto_chunk_proto_depIdxs,
		MessageInfos:      file_proto_chunk_proto_msgTypes,
	}.Build()
	File_proto_chunk_proto = out.File
	file_proto_chunk_proto_rawDesc = nil
	file_proto_chunk_proto_goTypes = nil
	file_proto_chunk_proto_depIdxs = nil
}
