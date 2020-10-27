// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: v2/storagegroup/grpc/types.proto

package storagegroup

import (
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
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

// StorageGroup keeps verification information for Data Audit sessions. Objects
// that require payed storage guaranties are gathered in `StorageGroups` with
// additional information used for proof of storage. `StorageGroup` only
// contains objects from the same container.
type StorageGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total size of the payloads of objects in the storage group
	ValidationDataSize uint64 `protobuf:"varint,1,opt,name=validation_data_size,json=validationDataSize,proto3" json:"validation_data_size,omitempty"`
	// Homomorphic hash from the concatenation of the payloads of the storage
	// group members. The order of concatenation is the same as the order of the
	// members in the `members` field.
	ValidationHash *grpc.Checksum `protobuf:"bytes,2,opt,name=validation_hash,json=validationHash,proto3" json:"validation_hash,omitempty"`
	// Last NeoFS epoch number of the storage group lifetime
	ExpirationEpoch uint64 `protobuf:"varint,3,opt,name=expiration_epoch,json=expirationEpoch,proto3" json:"expiration_epoch,omitempty"`
	// Strictly ordered list of storage group member objects
	Members []*grpc.ObjectID `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *StorageGroup) Reset() {
	*x = StorageGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_storagegroup_grpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageGroup) ProtoMessage() {}

func (x *StorageGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_storagegroup_grpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageGroup.ProtoReflect.Descriptor instead.
func (*StorageGroup) Descriptor() ([]byte, []int) {
	return file_v2_storagegroup_grpc_types_proto_rawDescGZIP(), []int{0}
}

func (x *StorageGroup) GetValidationDataSize() uint64 {
	if x != nil {
		return x.ValidationDataSize
	}
	return 0
}

func (x *StorageGroup) GetValidationHash() *grpc.Checksum {
	if x != nil {
		return x.ValidationHash
	}
	return nil
}

func (x *StorageGroup) GetExpirationEpoch() uint64 {
	if x != nil {
		return x.ExpirationEpoch
	}
	return 0
}

func (x *StorageGroup) GetMembers() []*grpc.ObjectID {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_v2_storagegroup_grpc_types_proto protoreflect.FileDescriptor

var file_v2_storagegroup_grpc_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x18, 0x76, 0x32, 0x2f, 0x72,
	0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66,
	0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x61, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0xaa, 0x02, 0x19, 0x4e, 0x65, 0x6f, 0x46, 0x53, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_storagegroup_grpc_types_proto_rawDescOnce sync.Once
	file_v2_storagegroup_grpc_types_proto_rawDescData = file_v2_storagegroup_grpc_types_proto_rawDesc
)

func file_v2_storagegroup_grpc_types_proto_rawDescGZIP() []byte {
	file_v2_storagegroup_grpc_types_proto_rawDescOnce.Do(func() {
		file_v2_storagegroup_grpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_storagegroup_grpc_types_proto_rawDescData)
	})
	return file_v2_storagegroup_grpc_types_proto_rawDescData
}

var file_v2_storagegroup_grpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v2_storagegroup_grpc_types_proto_goTypes = []interface{}{
	(*StorageGroup)(nil),  // 0: neo.fs.v2.storagegroup.StorageGroup
	(*grpc.Checksum)(nil), // 1: neo.fs.v2.refs.Checksum
	(*grpc.ObjectID)(nil), // 2: neo.fs.v2.refs.ObjectID
}
var file_v2_storagegroup_grpc_types_proto_depIdxs = []int32{
	1, // 0: neo.fs.v2.storagegroup.StorageGroup.validation_hash:type_name -> neo.fs.v2.refs.Checksum
	2, // 1: neo.fs.v2.storagegroup.StorageGroup.members:type_name -> neo.fs.v2.refs.ObjectID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v2_storagegroup_grpc_types_proto_init() }
func file_v2_storagegroup_grpc_types_proto_init() {
	if File_v2_storagegroup_grpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_storagegroup_grpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageGroup); i {
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
			RawDescriptor: file_v2_storagegroup_grpc_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_storagegroup_grpc_types_proto_goTypes,
		DependencyIndexes: file_v2_storagegroup_grpc_types_proto_depIdxs,
		MessageInfos:      file_v2_storagegroup_grpc_types_proto_msgTypes,
	}.Build()
	File_v2_storagegroup_grpc_types_proto = out.File
	file_v2_storagegroup_grpc_types_proto_rawDesc = nil
	file_v2_storagegroup_grpc_types_proto_goTypes = nil
	file_v2_storagegroup_grpc_types_proto_depIdxs = nil
}
