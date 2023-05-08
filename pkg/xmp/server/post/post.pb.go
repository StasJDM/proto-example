// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: xmp/server/post/post.proto

package post

import (
	id "github.com/StasJDM/proto-example/pkg/common/id"
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

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xmp_server_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmp_server_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_xmp_server_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_xmp_server_post_post_proto protoreflect.FileDescriptor

var file_xmp_server_post_post_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x78, 0x6d, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x78, 0x6d,
	0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x13, 0x78,
	0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x46, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x78, 0x6d, 0x70, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x78, 0x6d, 0x70, 0x2e, 0x49, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74,
	0x61, 0x73, 0x4a, 0x44, 0x4d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xmp_server_post_post_proto_rawDescOnce sync.Once
	file_xmp_server_post_post_proto_rawDescData = file_xmp_server_post_post_proto_rawDesc
)

func file_xmp_server_post_post_proto_rawDescGZIP() []byte {
	file_xmp_server_post_post_proto_rawDescOnce.Do(func() {
		file_xmp_server_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_xmp_server_post_post_proto_rawDescData)
	})
	return file_xmp_server_post_post_proto_rawDescData
}

var file_xmp_server_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xmp_server_post_post_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil), // 0: xmp.server.post.CreatePostRequest
	(*id.IdMessage)(nil),      // 1: xmp.IdMessage
}
var file_xmp_server_post_post_proto_depIdxs = []int32{
	0, // 0: xmp.server.post.Post.Create:input_type -> xmp.server.post.CreatePostRequest
	1, // 1: xmp.server.post.Post.Create:output_type -> xmp.IdMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_xmp_server_post_post_proto_init() }
func file_xmp_server_post_post_proto_init() {
	if File_xmp_server_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xmp_server_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
			RawDescriptor: file_xmp_server_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xmp_server_post_post_proto_goTypes,
		DependencyIndexes: file_xmp_server_post_post_proto_depIdxs,
		MessageInfos:      file_xmp_server_post_post_proto_msgTypes,
	}.Build()
	File_xmp_server_post_post_proto = out.File
	file_xmp_server_post_post_proto_rawDesc = nil
	file_xmp_server_post_post_proto_goTypes = nil
	file_xmp_server_post_post_proto_depIdxs = nil
}