// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bootstrap/service.proto

package bootstrap

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	service "github.com/nspcc-dev/neofs-proto/service"
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

// Request message to communicate between DHT nodes
type Request struct {
	Type                              NodeType `protobuf:"varint,1,opt,name=type,proto3,customtype=NodeType" json:"type"`
	Info                              NodeInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	service.RequestMetaHeader         `protobuf:"bytes,98,opt,name=Meta,proto3,embedded=Meta" json:"Meta"`
	service.RequestVerificationHeader `protobuf:"bytes,99,opt,name=Verify,proto3,embedded=Verify" json:"Verify"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bce759c9d8eb63, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetInfo() NodeInfo {
	if m != nil {
		return m.Info
	}
	return NodeInfo{}
}

func init() {
	proto.RegisterType((*Request)(nil), "bootstrap.Request")
}

func init() { proto.RegisterFile("bootstrap/service.proto", fileDescriptor_21bce759c9d8eb63) }

var fileDescriptor_21bce759c9d8eb63 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xae, 0xaf, 0x72, 0xfb, 0xe3, 0xbb, 0xf9, 0x16, 0x11, 0x65, 0x48, 0xaa, 0x4e, 0x15, 0x22,
	0x89, 0xd4, 0x2e, 0x8c, 0x55, 0xc4, 0x00, 0x43, 0x11, 0x0a, 0x88, 0x81, 0xcd, 0x71, 0x4e, 0x4a,
	0x86, 0xc6, 0xc6, 0x76, 0x2b, 0xf5, 0x4d, 0x78, 0xa4, 0x8e, 0x1d, 0x11, 0x43, 0x84, 0x82, 0xc4,
	0x73, 0xa0, 0xb8, 0x69, 0xa9, 0x60, 0xb2, 0xcf, 0xf7, 0xe7, 0xe3, 0x0f, 0x9f, 0x26, 0x9c, 0x6b,
	0xa5, 0x25, 0x15, 0xa1, 0x02, 0xb9, 0xca, 0x19, 0x04, 0x42, 0x72, 0xcd, 0x49, 0xef, 0x40, 0x38,
	0xa4, 0x61, 0xc2, 0x05, 0x68, 0xba, 0xa3, 0x9d, 0xfe, 0x1e, 0x5b, 0x81, 0xcc, 0xb3, 0x75, 0x83,
	0x9e, 0x7c, 0xa7, 0xe9, 0xb5, 0x00, 0xd5, 0xc0, 0xfe, 0x3c, 0xd7, 0x4f, 0xcb, 0x24, 0x60, 0x7c,
	0x11, 0xce, 0xf9, 0x9c, 0x87, 0x06, 0x4e, 0x96, 0x99, 0x99, 0xcc, 0x60, 0x6e, 0x3b, 0xf9, 0xf0,
	0x13, 0xe1, 0x4e, 0x0c, 0xcf, 0x4b, 0x50, 0x9a, 0x9c, 0x63, 0xab, 0x4e, 0xb2, 0xd1, 0x00, 0x8d,
	0xfe, 0x46, 0xf6, 0xa6, 0xf4, 0x5a, 0x6f, 0xa5, 0xd7, 0xbd, 0xe1, 0x29, 0xdc, 0xaf, 0x05, 0x54,
	0xa5, 0x67, 0xd5, 0x67, 0x6c, 0x54, 0xc4, 0xc7, 0x56, 0x5e, 0x64, 0xdc, 0xfe, 0x33, 0x40, 0xa3,
	0x7f, 0xe3, 0xff, 0xc1, 0x61, 0x9d, 0xa0, 0x36, 0x5c, 0x17, 0x19, 0x8f, 0xac, 0x3a, 0x22, 0x36,
	0x32, 0x72, 0x81, 0xad, 0x19, 0x68, 0x6a, 0x27, 0x46, 0xee, 0x04, 0xfb, 0x06, 0x9a, 0xc7, 0x6b,
	0xee, 0x0a, 0x68, 0x0a, 0x32, 0xea, 0xd6, 0xae, 0x6d, 0xe9, 0xa1, 0xd8, 0x38, 0xc8, 0x25, 0x6e,
	0x3f, 0x98, 0x8f, 0xdb, 0xcc, 0x78, 0x87, 0x3f, 0xbd, 0x86, 0xcd, 0x19, 0xd5, 0x39, 0x2f, 0x7e,
	0x65, 0x34, 0xde, 0xf1, 0x14, 0xf7, 0xa2, 0xfd, 0x86, 0x64, 0x82, 0x3b, 0xb7, 0x92, 0x33, 0x50,
	0x8a, 0x90, 0xa3, 0xc5, 0x9b, 0x3c, 0xa7, 0x7f, 0x84, 0xdd, 0x09, 0x09, 0x34, 0x9d, 0x51, 0x11,
	0x4d, 0x37, 0x95, 0x8b, 0xb6, 0x95, 0x8b, 0x5e, 0x2b, 0x17, 0xbd, 0x57, 0x2e, 0x7a, 0xf9, 0x70,
	0x5b, 0x8f, 0x67, 0x47, 0x7d, 0x17, 0x4a, 0x30, 0xe6, 0xa7, 0xb0, 0x0a, 0x0b, 0xe0, 0x99, 0xf2,
	0x77, 0x6d, 0x1f, 0xb2, 0x92, 0xb6, 0x01, 0x26, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xcb,
	0xf7, 0x92, 0x09, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BootstrapClient is the client API for Bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootstrapClient interface {
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SpreadMap, error)
}

type bootstrapClient struct {
	cc *grpc.ClientConn
}

func NewBootstrapClient(cc *grpc.ClientConn) BootstrapClient {
	return &bootstrapClient{cc}
}

func (c *bootstrapClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SpreadMap, error) {
	out := new(SpreadMap)
	err := c.cc.Invoke(ctx, "/bootstrap.Bootstrap/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapServer is the server API for Bootstrap service.
type BootstrapServer interface {
	Process(context.Context, *Request) (*SpreadMap, error)
}

// UnimplementedBootstrapServer can be embedded to have forward compatible implementations.
type UnimplementedBootstrapServer struct {
}

func (*UnimplementedBootstrapServer) Process(ctx context.Context, req *Request) (*SpreadMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterBootstrapServer(s *grpc.Server, srv BootstrapServer) {
	s.RegisterService(&_Bootstrap_serviceDesc, srv)
}

func _Bootstrap_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bootstrap.Bootstrap/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).Process(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootstrap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bootstrap.Bootstrap",
	HandlerType: (*BootstrapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Bootstrap_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bootstrap/service.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.RequestVerificationHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6
	i--
	dAtA[i] = 0x9a
	{
		size, err := m.RequestMetaHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6
	i--
	dAtA[i] = 0x92
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
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
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovService(uint64(m.Type))
	}
	l = m.Info.Size()
	n += 1 + l + sovService(uint64(l))
	l = m.RequestMetaHeader.Size()
	n += 2 + l + sovService(uint64(l))
	l = m.RequestVerificationHeader.Size()
	n += 2 + l + sovService(uint64(l))
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
func (m *Request) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= NodeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 98:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetaHeader", wireType)
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
			if err := m.RequestMetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 99:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestVerificationHeader", wireType)
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
			if err := m.RequestVerificationHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
