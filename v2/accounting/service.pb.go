// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/accounting/service.proto

package accounting

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	service "github.com/nspcc-dev/neofs-api-go/v2/service"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Message defines the request body of Balance method.
//
// To indicate the account for which the balance is requested, it's identifier
// is used.
//
// To gain access to the requested information, the request body must be formed
// according to the requirements from the system specification.
type BalanceRequest struct {
	// Body of the balance request message.
	Body *BalanceRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *service.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness
	// of transmission.
	VerifyHeader         *service.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_484f6b0e24e3172f, []int{0}
}
func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetBody() *BalanceRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BalanceRequest) GetMetaHeader() *service.RequestMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *BalanceRequest) GetVerifyHeader() *service.RequestVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

//Request body
type BalanceRequest_Body struct {
	// Carries user identifier in NeoFS system for which the balance
	// is requested.
	OwnerId              *refs.OwnerID `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BalanceRequest_Body) Reset()         { *m = BalanceRequest_Body{} }
func (m *BalanceRequest_Body) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest_Body) ProtoMessage()    {}
func (*BalanceRequest_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_484f6b0e24e3172f, []int{0, 0}
}
func (m *BalanceRequest_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceRequest_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceRequest_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceRequest_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest_Body.Merge(m, src)
}
func (m *BalanceRequest_Body) XXX_Size() int {
	return m.Size()
}
func (m *BalanceRequest_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest_Body.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest_Body proto.InternalMessageInfo

func (m *BalanceRequest_Body) GetOwnerId() *refs.OwnerID {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

// Decimal represents the decimal numbers.
type Decimal struct {
	// value carries number value.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// precision carries value precision.
	Precision            uint32   `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal) Reset()         { *m = Decimal{} }
func (m *Decimal) String() string { return proto.CompactTextString(m) }
func (*Decimal) ProtoMessage()    {}
func (*Decimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_484f6b0e24e3172f, []int{1}
}
func (m *Decimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Decimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Decimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Decimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal.Merge(m, src)
}
func (m *Decimal) XXX_Size() int {
	return m.Size()
}
func (m *Decimal) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal proto.InternalMessageInfo

func (m *Decimal) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Decimal) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Message defines the response body of Balance method.
//
// The amount of funds is calculated in decimal numbers.
type BalanceResponse struct {
	// Body of the balance response message.
	Body *BalanceResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *service.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness
	// of transmission.
	VerifyHeader         *service.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_484f6b0e24e3172f, []int{2}
}
func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetBody() *BalanceResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BalanceResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *BalanceResponse) GetVerifyHeader() *service.ResponseVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

//Request body
type BalanceResponse_Body struct {
	// Carries the amount of funds on the account.
	Balance              *Decimal `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceResponse_Body) Reset()         { *m = BalanceResponse_Body{} }
func (m *BalanceResponse_Body) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse_Body) ProtoMessage()    {}
func (*BalanceResponse_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_484f6b0e24e3172f, []int{2, 0}
}
func (m *BalanceResponse_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceResponse_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceResponse_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceResponse_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse_Body.Merge(m, src)
}
func (m *BalanceResponse_Body) XXX_Size() int {
	return m.Size()
}
func (m *BalanceResponse_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse_Body.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse_Body proto.InternalMessageInfo

func (m *BalanceResponse_Body) GetBalance() *Decimal {
	if m != nil {
		return m.Balance
	}
	return nil
}

func init() {
	proto.RegisterType((*BalanceRequest)(nil), "neo.fs.v2.accounting.BalanceRequest")
	proto.RegisterType((*BalanceRequest_Body)(nil), "neo.fs.v2.accounting.BalanceRequest.Body")
	proto.RegisterType((*Decimal)(nil), "neo.fs.v2.accounting.Decimal")
	proto.RegisterType((*BalanceResponse)(nil), "neo.fs.v2.accounting.BalanceResponse")
	proto.RegisterType((*BalanceResponse_Body)(nil), "neo.fs.v2.accounting.BalanceResponse.Body")
}

func init() { proto.RegisterFile("v2/accounting/service.proto", fileDescriptor_484f6b0e24e3172f) }

var fileDescriptor_484f6b0e24e3172f = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0x37, 0x28, 0x78, 0x0c, 0x84, 0x19, 0xea, 0x54, 0xa0, 0x42, 0xd5, 0x26, 0x01,
	0xa2, 0x8e, 0x14, 0x2e, 0x90, 0x86, 0x06, 0x5a, 0x35, 0x26, 0x7a, 0xc1, 0xbf, 0x4c, 0xda, 0x05,
	0x37, 0x93, 0xeb, 0x7c, 0xe9, 0x2c, 0x5a, 0x3b, 0xc4, 0xae, 0x51, 0xdf, 0x84, 0x17, 0xe0, 0x86,
	0x0b, 0x9e, 0x83, 0x4b, 0x1e, 0x01, 0x95, 0x17, 0x41, 0xb1, 0xdd, 0x36, 0x95, 0xb2, 0xad, 0x77,
	0xb1, 0x7d, 0xce, 0xf9, 0x7c, 0x7e, 0x8a, 0xd1, 0x7d, 0x13, 0x85, 0x94, 0x31, 0x39, 0x16, 0x9a,
	0x8b, 0x41, 0xa8, 0x20, 0x37, 0x9c, 0x01, 0xc9, 0x72, 0xa9, 0x25, 0xde, 0x12, 0x20, 0x49, 0xaa,
	0x88, 0x89, 0xc8, 0x42, 0xd3, 0xbc, 0x6b, 0xa2, 0x30, 0x87, 0x54, 0x85, 0x7a, 0x92, 0x81, 0x72,
	0xd2, 0xe6, 0x3d, 0x13, 0xcd, 0xcc, 0xe1, 0x08, 0x34, 0xf5, 0xdb, 0x8d, 0xd2, 0xb6, 0x81, 0x9c,
	0xa7, 0x13, 0x77, 0xd0, 0xfe, 0x51, 0x43, 0xb7, 0xba, 0x74, 0x48, 0x05, 0x83, 0x18, 0xbe, 0x8e,
	0x41, 0x69, 0xbc, 0x8f, 0xd6, 0xfb, 0x32, 0x99, 0x6c, 0x07, 0x8f, 0x82, 0xc7, 0x1b, 0xd1, 0x13,
	0x52, 0x35, 0x9c, 0x2c, 0x7b, 0x48, 0x57, 0x26, 0x93, 0xd8, 0xda, 0xf0, 0x1b, 0xb4, 0x51, 0x0c,
	0x3e, 0x3d, 0x03, 0x9a, 0x40, 0xbe, 0x5d, 0xb3, 0x29, 0x3b, 0xa5, 0x94, 0x59, 0x37, 0xef, 0x7d,
	0x07, 0x9a, 0xbe, 0xb5, 0xda, 0x18, 0x8d, 0xe6, 0xdf, 0xf8, 0x13, 0xda, 0x74, 0x17, 0x9d, 0x05,
	0xad, 0xd9, 0xa0, 0x67, 0xe7, 0x07, 0x9d, 0x14, 0x72, 0xce, 0xa8, 0xe6, 0x52, 0xf8, 0xc0, 0x9b,
	0x2e, 0xc2, 0xad, 0x9a, 0x7b, 0x68, 0xbd, 0xb8, 0x27, 0x8e, 0xd0, 0x75, 0xf9, 0x4d, 0x40, 0x7e,
	0xca, 0x13, 0x5f, 0xb2, 0x51, 0x4a, 0x2d, 0x90, 0x92, 0x0f, 0xc5, 0x79, 0xef, 0x30, 0xae, 0x5b,
	0x61, 0x2f, 0x69, 0xef, 0xa3, 0xfa, 0x21, 0x30, 0x3e, 0xa2, 0x43, 0xbc, 0x85, 0xae, 0x1a, 0x3a,
	0x1c, 0x83, 0xf5, 0xae, 0xc5, 0x6e, 0x81, 0x1f, 0xa0, 0x1b, 0x59, 0x0e, 0x8c, 0x2b, 0x2e, 0x85,
	0x2d, 0xbd, 0x19, 0x2f, 0x36, 0xda, 0xbf, 0x6a, 0xe8, 0xf6, 0x1c, 0x99, 0xca, 0xa4, 0x50, 0x80,
	0x5f, 0x2d, 0x71, 0x7e, 0x7a, 0x09, 0x67, 0x67, 0x2a, 0x83, 0x3e, 0xaa, 0x02, 0xbd, 0x5b, 0xc9,
	0xc7, 0x99, 0xcf, 0x21, 0x1d, 0x57, 0x93, 0xee, 0x5c, 0x90, 0x74, 0x29, 0xea, 0xd7, 0x1e, 0xf5,
	0x0b, 0x54, 0xef, 0xbb, 0x06, 0xbe, 0xe6, 0xc3, 0xea, 0x9a, 0x9e, 0x6d, 0x3c, 0x53, 0x47, 0x5f,
	0xd0, 0x9d, 0x83, 0xf9, 0xf1, 0xb1, 0x1b, 0x8f, 0x4f, 0x50, 0xdd, 0xf3, 0xc0, 0x3b, 0xab, 0xfc,
	0x96, 0xcd, 0xdd, 0x95, 0xa0, 0x76, 0xd3, 0xdf, 0xd3, 0x56, 0xf0, 0x67, 0xda, 0x0a, 0xfe, 0x4e,
	0x5b, 0xc1, 0xf7, 0x7f, 0xad, 0x2b, 0x9f, 0xf7, 0x06, 0x5c, 0x9f, 0x8d, 0xfb, 0x84, 0xc9, 0x51,
	0x28, 0x54, 0xc6, 0x58, 0x27, 0x01, 0x13, 0x0a, 0x90, 0xa9, 0xea, 0xd0, 0x8c, 0x77, 0x06, 0x32,
	0x5c, 0x7a, 0xb0, 0x2f, 0x17, 0x9f, 0x3f, 0x6b, 0x8d, 0xf7, 0x20, 0x8f, 0x8e, 0xc9, 0xc1, 0xc7,
	0x5e, 0x31, 0x77, 0x51, 0xa3, 0x7f, 0xcd, 0xbe, 0xb9, 0xe7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x24, 0x48, 0xec, 0xed, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountingServiceClient is the client API for AccountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingServiceClient interface {
	// Returns the amount of funds for the requested NeoFS account.
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type accountingServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountingServiceClient(cc *grpc.ClientConn) AccountingServiceClient {
	return &accountingServiceClient{cc}
}

func (c *accountingServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.accounting.AccountingService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServiceServer is the server API for AccountingService service.
type AccountingServiceServer interface {
	// Returns the amount of funds for the requested NeoFS account.
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

// UnimplementedAccountingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServiceServer struct {
}

func (*UnimplementedAccountingServiceServer) Balance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

func RegisterAccountingServiceServer(s *grpc.Server, srv AccountingServiceServer) {
	s.RegisterService(&_AccountingService_serviceDesc, srv)
}

func _AccountingService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.accounting.AccountingService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neo.fs.v2.accounting.AccountingService",
	HandlerType: (*AccountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _AccountingService_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2/accounting/service.proto",
}

func (m *BalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceRequest_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceRequest_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceRequest_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OwnerId != nil {
		{
			size, err := m.OwnerId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Decimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Decimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Decimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Precision != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Precision))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceResponse_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceResponse_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceResponse_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Balance != nil {
		{
			size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceRequest_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnerId != nil {
		l = m.OwnerId.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Decimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovService(uint64(m.Value))
	}
	if m.Precision != 0 {
		n += 1 + sovService(uint64(m.Precision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BalanceResponse_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Balance != nil {
		l = m.Balance.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &BalanceRequest_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &service.RequestMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &service.RequestVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceRequest_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerId == nil {
				m.OwnerId = &refs.OwnerID{}
			}
			if err := m.OwnerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Decimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Decimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Decimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &BalanceResponse_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &service.ResponseMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &service.ResponseVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BalanceResponse_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Balance == nil {
				m.Balance = &Decimal{}
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
