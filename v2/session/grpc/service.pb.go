// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: v2/session/grpc/service.proto

package session

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	grpc1 "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Information necessary for opening a session.
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of create session token request message.
	Body *CreateRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_session_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_session_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_v2_session_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetBody() *CreateRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CreateRequest) GetMetaHeader() *RequestMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *CreateRequest) GetVerifyHeader() *RequestVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Information about the opened session.
type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body of create session token response message.
	Body *CreateResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader *ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_session_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_session_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_v2_session_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetBody() *CreateResponse_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CreateResponse) GetMetaHeader() *ResponseMetaHeader {
	if x != nil {
		return x.MetaHeader
	}
	return nil
}

func (x *CreateResponse) GetVerifyHeader() *ResponseVerificationHeader {
	if x != nil {
		return x.VerifyHeader
	}
	return nil
}

// Session creation request body
type CreateRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dession initiating user's or node's key derived `OwnerID`.
	OwnerId *grpc.OwnerID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Session expiration `Epoch`
	Expiration uint64 `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *CreateRequest_Body) Reset() {
	*x = CreateRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_session_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest_Body) ProtoMessage() {}

func (x *CreateRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_v2_session_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest_Body.ProtoReflect.Descriptor instead.
func (*CreateRequest_Body) Descriptor() ([]byte, []int) {
	return file_v2_session_grpc_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateRequest_Body) GetOwnerId() *grpc.OwnerID {
	if x != nil {
		return x.OwnerId
	}
	return nil
}

func (x *CreateRequest_Body) GetExpiration() uint64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

// Session creation response body
type CreateResponse_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of a newly created session
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Public key used for session
	SessionKey []byte `protobuf:"bytes,2,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
}

func (x *CreateResponse_Body) Reset() {
	*x = CreateResponse_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_session_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse_Body) ProtoMessage() {}

func (x *CreateResponse_Body) ProtoReflect() protoreflect.Message {
	mi := &file_v2_session_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse_Body.ProtoReflect.Descriptor instead.
func (*CreateResponse_Body) Descriptor() ([]byte, []int) {
	return file_v2_session_grpc_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateResponse_Body) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CreateResponse_Body) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

var File_v2_session_grpc_service_proto protoreflect.FileDescriptor

var file_v2_session_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x18, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x66, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x32,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65,
	0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x51, 0x0a,
	0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x1a, 0x5a, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x02, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x6f,
	0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x37, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x32, 0x5f, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e,
	0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x66, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x52, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x73, 0x70, 0x63, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0xaa, 0x02,
	0x14, 0x4e, 0x65, 0x6f, 0x46, 0x53, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_session_grpc_service_proto_rawDescOnce sync.Once
	file_v2_session_grpc_service_proto_rawDescData = file_v2_session_grpc_service_proto_rawDesc
)

func file_v2_session_grpc_service_proto_rawDescGZIP() []byte {
	file_v2_session_grpc_service_proto_rawDescOnce.Do(func() {
		file_v2_session_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_session_grpc_service_proto_rawDescData)
	})
	return file_v2_session_grpc_service_proto_rawDescData
}

var file_v2_session_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_session_grpc_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),              // 0: neo.fs.v2.session.CreateRequest
	(*CreateResponse)(nil),             // 1: neo.fs.v2.session.CreateResponse
	(*CreateRequest_Body)(nil),         // 2: neo.fs.v2.session.CreateRequest.Body
	(*CreateResponse_Body)(nil),        // 3: neo.fs.v2.session.CreateResponse.Body
	(*RequestMetaHeader)(nil),          // 4: neo.fs.v2.session.RequestMetaHeader
	(*RequestVerificationHeader)(nil),  // 5: neo.fs.v2.session.RequestVerificationHeader
	(*ResponseMetaHeader)(nil),         // 6: neo.fs.v2.session.ResponseMetaHeader
	(*ResponseVerificationHeader)(nil), // 7: neo.fs.v2.session.ResponseVerificationHeader
	(*grpc.OwnerID)(nil),               // 8: neo.fs.v2.refs.OwnerID
}
var file_v2_session_grpc_service_proto_depIdxs = []int32{
	2, // 0: neo.fs.v2.session.CreateRequest.body:type_name -> neo.fs.v2.session.CreateRequest.Body
	4, // 1: neo.fs.v2.session.CreateRequest.meta_header:type_name -> neo.fs.v2.session.RequestMetaHeader
	5, // 2: neo.fs.v2.session.CreateRequest.verify_header:type_name -> neo.fs.v2.session.RequestVerificationHeader
	3, // 3: neo.fs.v2.session.CreateResponse.body:type_name -> neo.fs.v2.session.CreateResponse.Body
	6, // 4: neo.fs.v2.session.CreateResponse.meta_header:type_name -> neo.fs.v2.session.ResponseMetaHeader
	7, // 5: neo.fs.v2.session.CreateResponse.verify_header:type_name -> neo.fs.v2.session.ResponseVerificationHeader
	8, // 6: neo.fs.v2.session.CreateRequest.Body.owner_id:type_name -> neo.fs.v2.refs.OwnerID
	0, // 7: neo.fs.v2.session.SessionService.Create:input_type -> neo.fs.v2.session.CreateRequest
	1, // 8: neo.fs.v2.session.SessionService.Create:output_type -> neo.fs.v2.session.CreateResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_v2_session_grpc_service_proto_init() }
func file_v2_session_grpc_service_proto_init() {
	if File_v2_session_grpc_service_proto != nil {
		return
	}
	file_v2_session_grpc_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v2_session_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_v2_session_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_v2_session_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest_Body); i {
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
		file_v2_session_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse_Body); i {
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
			RawDescriptor: file_v2_session_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_session_grpc_service_proto_goTypes,
		DependencyIndexes: file_v2_session_grpc_service_proto_depIdxs,
		MessageInfos:      file_v2_session_grpc_service_proto_msgTypes,
	}.Build()
	File_v2_session_grpc_service_proto = out.File
	file_v2_session_grpc_service_proto_rawDesc = nil
	file_v2_session_grpc_service_proto_goTypes = nil
	file_v2_session_grpc_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion6

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	// Opens a new session between two peers.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc1.CallOption) (*CreateResponse, error)
}

type sessionServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewSessionServiceClient(cc grpc1.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc1.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.session.SessionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	// Opens a new session between two peers.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedSessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (*UnimplementedSessionServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterSessionServiceServer(s *grpc1.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Create(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.session.SessionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "neo.fs.v2.session.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SessionService_Create_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "v2/session/grpc/service.proto",
}
