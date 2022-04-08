// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: dongtai.proto

package DontaiGRPC

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dongtai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_dongtai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_dongtai_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   int64  `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dongtai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_dongtai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_dongtai_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_dongtai_proto protoreflect.FileDescriptor

var file_dongtai_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x6f, 0x6e, 0x67, 0x74, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x44, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x47, 0x52, 0x50, 0x43, 0x22, 0x1d, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x40, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e,
	0x44, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x47, 0x52, 0x50, 0x43, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x62, 0x3b, 0x44, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x47, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dongtai_proto_rawDescOnce sync.Once
	file_dongtai_proto_rawDescData = file_dongtai_proto_rawDesc
)

func file_dongtai_proto_rawDescGZIP() []byte {
	file_dongtai_proto_rawDescOnce.Do(func() {
		file_dongtai_proto_rawDescData = protoimpl.X.CompressGZIP(file_dongtai_proto_rawDescData)
	})
	return file_dongtai_proto_rawDescData
}

var file_dongtai_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dongtai_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: DontaiGRPC.Request
	(*Response)(nil), // 1: DontaiGRPC.Response
}
var file_dongtai_proto_depIdxs = []int32{
	0, // 0: DontaiGRPC.GRPCService.Send:input_type -> DontaiGRPC.Request
	1, // 1: DontaiGRPC.GRPCService.Send:output_type -> DontaiGRPC.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dongtai_proto_init() }
func file_dongtai_proto_init() {
	if File_dongtai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dongtai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_dongtai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_dongtai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dongtai_proto_goTypes,
		DependencyIndexes: file_dongtai_proto_depIdxs,
		MessageInfos:      file_dongtai_proto_msgTypes,
	}.Build()
	File_dongtai_proto = out.File
	file_dongtai_proto_rawDesc = nil
	file_dongtai_proto_goTypes = nil
	file_dongtai_proto_depIdxs = nil
}
