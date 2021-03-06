// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: contractservice.proto

package contractservice

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddContractReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ContractId string `protobuf:"bytes,2,opt,name=contractId,proto3" json:"contractId,omitempty"`
}

func (x *AddContractReq) Reset() {
	*x = AddContractReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContractReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContractReq) ProtoMessage() {}

func (x *AddContractReq) ProtoReflect() protoreflect.Message {
	mi := &file_contractservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContractReq.ProtoReflect.Descriptor instead.
func (*AddContractReq) Descriptor() ([]byte, []int) {
	return file_contractservice_proto_rawDescGZIP(), []int{0}
}

func (x *AddContractReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AddContractReq) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

type AddContractRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddContractRes) Reset() {
	*x = AddContractRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContractRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContractRes) ProtoMessage() {}

func (x *AddContractRes) ProtoReflect() protoreflect.Message {
	mi := &file_contractservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContractRes.ProtoReflect.Descriptor instead.
func (*AddContractRes) Descriptor() ([]byte, []int) {
	return file_contractservice_proto_rawDescGZIP(), []int{1}
}

func (x *AddContractRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_contractservice_proto protoreflect.FileDescriptor

var file_contractservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x62, 0x0a, 0x0f, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contractservice_proto_rawDescOnce sync.Once
	file_contractservice_proto_rawDescData = file_contractservice_proto_rawDesc
)

func file_contractservice_proto_rawDescGZIP() []byte {
	file_contractservice_proto_rawDescOnce.Do(func() {
		file_contractservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_contractservice_proto_rawDescData)
	})
	return file_contractservice_proto_rawDescData
}

var file_contractservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contractservice_proto_goTypes = []interface{}{
	(*AddContractReq)(nil), // 0: contractservice.AddContractReq
	(*AddContractRes)(nil), // 1: contractservice.AddContractRes
}
var file_contractservice_proto_depIdxs = []int32{
	0, // 0: contractservice.ContractService.AddContract:input_type -> contractservice.AddContractReq
	1, // 1: contractservice.ContractService.AddContract:output_type -> contractservice.AddContractRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contractservice_proto_init() }
func file_contractservice_proto_init() {
	if File_contractservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contractservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddContractReq); i {
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
		file_contractservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddContractRes); i {
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
			RawDescriptor: file_contractservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contractservice_proto_goTypes,
		DependencyIndexes: file_contractservice_proto_depIdxs,
		MessageInfos:      file_contractservice_proto_msgTypes,
	}.Build()
	File_contractservice_proto = out.File
	file_contractservice_proto_rawDesc = nil
	file_contractservice_proto_goTypes = nil
	file_contractservice_proto_depIdxs = nil
}
