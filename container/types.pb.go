// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: container/types.proto

package container

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	netmap "github.com/nspcc-dev/netmap"
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

// The Container service definition.
type Container struct {
	OwnerID              OwnerID              `protobuf:"bytes,1,opt,name=OwnerID,proto3,customtype=OwnerID" json:"OwnerID"`
	Salt                 UUID                 `protobuf:"bytes,2,opt,name=Salt,proto3,customtype=UUID" json:"Salt"`
	Capacity             uint64               `protobuf:"varint,3,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
	Rules                netmap.PlacementRule `protobuf:"bytes,4,opt,name=Rules,proto3" json:"Rules"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_1432e52ab0b53e3e, []int{0}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return m.Size()
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Container) GetRules() netmap.PlacementRule {
	if m != nil {
		return m.Rules
	}
	return netmap.PlacementRule{}
}

func init() {
	proto.RegisterType((*Container)(nil), "container.Container")
}

func init() { proto.RegisterFile("container/types.proto", fileDescriptor_1432e52ab0b53e3e) }

var fileDescriptor_1432e52ab0b53e3e = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x4b, 0x69, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xe7, 0x15, 0x17, 0x24, 0x27, 0xeb, 0xa6, 0xa4, 0x96, 0xe9, 0xe7, 0xa5, 0x96, 0xe4,
	0x26, 0x16, 0xe8, 0x17, 0xa7, 0xe6, 0xa4, 0x26, 0x97, 0xe4, 0x17, 0x41, 0xb4, 0x49, 0xe9, 0x22,
	0xa9, 0x4d, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98,
	0x05, 0x51, 0xae, 0xb4, 0x9c, 0x91, 0x8b, 0xd3, 0x19, 0x66, 0x91, 0x90, 0x26, 0x17, 0xbb, 0x7f,
	0x79, 0x5e, 0x6a, 0x91, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x13, 0xff, 0x89, 0x7b,
	0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xc3, 0x84, 0x83, 0x60, 0x0c, 0x21, 0x05, 0x2e, 0x96, 0xe0, 0xc4,
	0x9c, 0x12, 0x09, 0x26, 0xb0, 0x3a, 0x1e, 0xa8, 0x3a, 0x96, 0xd0, 0x50, 0x4f, 0x97, 0x20, 0xb0,
	0x8c, 0x90, 0x14, 0x17, 0x87, 0x73, 0x62, 0x41, 0x62, 0x72, 0x66, 0x49, 0xa5, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x4b, 0x10, 0x9c, 0x2f, 0x64, 0xc8, 0xc5, 0x1a, 0x54, 0x9a, 0x93, 0x5a, 0x2c, 0xc1,
	0xa2, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xaa, 0x07, 0xf1, 0x8c, 0x5e, 0x40, 0x4e, 0x62, 0x72, 0x6a,
	0x6e, 0x6a, 0x5e, 0x09, 0x48, 0xd6, 0x89, 0x05, 0x64, 0x6a, 0x10, 0x44, 0xa5, 0x93, 0xc3, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x8c, 0xc7, 0x72, 0x0c, 0x51, 0xb8, 0x82, 0x26, 0x3f, 0xad, 0x58, 0x17, 0xe2, 0x59, 0x78, 0x30,
	0x26, 0xb1, 0x81, 0x05, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0x6b, 0x4d, 0x08, 0x71,
	0x01, 0x00, 0x00,
}

func (m *Container) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Container) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Container) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.Rules.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Capacity != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Capacity))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Salt.Size()
		i -= size
		if _, err := m.Salt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.OwnerID.Size()
		i -= size
		if _, err := m.OwnerID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Container) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OwnerID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Salt.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Capacity != 0 {
		n += 1 + sovTypes(uint64(m.Capacity))
	}
	l = m.Rules.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Container) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Container: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Container: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Salt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capacity", wireType)
			}
			m.Capacity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Capacity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
