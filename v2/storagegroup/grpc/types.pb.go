// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/storagegroup/grpc/types.proto

package storagegroup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
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

// StorageGroup groups the information about the NeoFS storage group.
// The storage group consists of objects from single container.
type StorageGroup struct {
	// validation_data_size carries the total size of the payloads of the storage
	// group members.
	ValidationDataSize uint64 `protobuf:"varint,1,opt,name=validation_data_size,json=validationDataSize,proto3" json:"validation_data_size,omitempty"`
	// validation_hash carries homomorphic hash from the concatenation of the
	// payloads of the storage group members
	// The order of concatenation is the same as the order of the members in the
	// Members field.
	ValidationHash []byte `protobuf:"bytes,2,opt,name=validation_hash,json=validationHash,proto3" json:"validation_hash,omitempty"`
	// expiration_epoch carries last NeoFS epoch number of the storage group
	// lifetime.
	ExpirationEpoch uint64 `protobuf:"varint,3,opt,name=expiration_epoch,json=expirationEpoch,proto3" json:"expiration_epoch,omitempty"`
	// Members carries the list of identifiers of the object storage group members.
	// The list is strictly ordered.
	Members              []*grpc.ObjectID `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StorageGroup) Reset()         { *m = StorageGroup{} }
func (m *StorageGroup) String() string { return proto.CompactTextString(m) }
func (*StorageGroup) ProtoMessage()    {}
func (*StorageGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1685cc94b670845, []int{0}
}
func (m *StorageGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StorageGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StorageGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StorageGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageGroup.Merge(m, src)
}
func (m *StorageGroup) XXX_Size() int {
	return m.Size()
}
func (m *StorageGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageGroup.DiscardUnknown(m)
}

var xxx_messageInfo_StorageGroup proto.InternalMessageInfo

func (m *StorageGroup) GetValidationDataSize() uint64 {
	if m != nil {
		return m.ValidationDataSize
	}
	return 0
}

func (m *StorageGroup) GetValidationHash() []byte {
	if m != nil {
		return m.ValidationHash
	}
	return nil
}

func (m *StorageGroup) GetExpirationEpoch() uint64 {
	if m != nil {
		return m.ExpirationEpoch
	}
	return 0
}

func (m *StorageGroup) GetMembers() []*grpc.ObjectID {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*StorageGroup)(nil), "neo.fs.v2.storagegroup.StorageGroup")
}

func init() { proto.RegisterFile("v2/storagegroup/grpc/types.proto", fileDescriptor_f1685cc94b670845) }

var fileDescriptor_f1685cc94b670845 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x33, 0x31,
	0x18, 0x85, 0xbf, 0xb4, 0xe5, 0x13, 0x62, 0xb1, 0x12, 0x44, 0xa2, 0x8b, 0x61, 0x70, 0x63, 0x5d,
	0x34, 0x91, 0x71, 0xe9, 0x4a, 0xad, 0x3f, 0xdd, 0xa8, 0xb4, 0x3b, 0x37, 0x25, 0x93, 0xbe, 0x9d,
	0x89, 0xd8, 0x49, 0x4c, 0xd2, 0x41, 0x7b, 0x25, 0x5e, 0x83, 0x57, 0xd2, 0xa5, 0x97, 0x20, 0xe3,
	0x8d, 0xc8, 0xb4, 0xc8, 0x0c, 0xe8, 0xf6, 0xc9, 0x73, 0xf2, 0x72, 0x0e, 0x0e, 0xf3, 0x88, 0x3b,
	0xaf, 0xad, 0x48, 0x20, 0xb1, 0x7a, 0x6e, 0x78, 0x62, 0x8d, 0xe4, 0xfe, 0xd5, 0x80, 0x63, 0xc6,
	0x6a, 0xaf, 0xc9, 0x6e, 0x06, 0x9a, 0x4d, 0x1d, 0xcb, 0x23, 0x56, 0x17, 0xf7, 0x69, 0x1e, 0x71,
	0x0b, 0x53, 0xf7, 0x2b, 0x71, 0xb0, 0x44, 0xb8, 0x3d, 0x5a, 0xab, 0xd7, 0xa5, 0x4a, 0x8e, 0xf1,
	0x4e, 0x2e, 0x9e, 0xd4, 0x44, 0x78, 0xa5, 0xb3, 0xf1, 0x44, 0x78, 0x31, 0x76, 0x6a, 0x01, 0x14,
	0x85, 0xa8, 0xdb, 0x1a, 0x92, 0xea, 0xad, 0x2f, 0xbc, 0x18, 0xa9, 0x05, 0x90, 0x43, 0xdc, 0xa9,
	0x25, 0x52, 0xe1, 0x52, 0xda, 0x08, 0x51, 0xb7, 0x3d, 0xdc, 0xaa, 0xf0, 0x8d, 0x70, 0x29, 0x39,
	0xc2, 0xdb, 0xf0, 0x62, 0x94, 0x5d, 0x8b, 0x60, 0xb4, 0x4c, 0x69, 0x73, 0xf5, 0x6d, 0xa7, 0xe2,
	0x97, 0x25, 0x26, 0x11, 0xde, 0x98, 0xc1, 0x2c, 0x06, 0xeb, 0x68, 0x2b, 0x6c, 0x76, 0x37, 0x23,
	0xca, 0xaa, 0x6a, 0x65, 0x13, 0x76, 0x17, 0x3f, 0x82, 0xf4, 0x83, 0xfe, 0xf0, 0x47, 0x3c, 0x7f,
	0x5e, 0x16, 0x01, 0xfa, 0x28, 0x02, 0xf4, 0x59, 0x04, 0xe8, 0xed, 0x2b, 0xf8, 0xf7, 0x70, 0x91,
	0x28, 0x9f, 0xce, 0x63, 0x26, 0xf5, 0x8c, 0x67, 0xce, 0x48, 0xd9, 0x9b, 0x40, 0xce, 0x33, 0xd0,
	0x53, 0xd7, 0x13, 0x46, 0xf5, 0x12, 0xcd, 0xff, 0x9a, 0xf4, 0xb4, 0x4e, 0xde, 0x1b, 0x7b, 0xb7,
	0xa0, 0xaf, 0x46, 0xec, 0xec, 0x7e, 0x50, 0x1e, 0xaf, 0x8f, 0x15, 0xff, 0x5f, 0x8d, 0x78, 0xf2,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x95, 0x67, 0x64, 0x9a, 0x01, 0x00, 0x00,
}

func (m *StorageGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StorageGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ExpirationEpoch != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ExpirationEpoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValidationHash) > 0 {
		i -= len(m.ValidationHash)
		copy(dAtA[i:], m.ValidationHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidationHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.ValidationDataSize != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ValidationDataSize))
		i--
		dAtA[i] = 0x8
	}
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
func (m *StorageGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidationDataSize != 0 {
		n += 1 + sovTypes(uint64(m.ValidationDataSize))
	}
	l = len(m.ValidationHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ExpirationEpoch != 0 {
		n += 1 + sovTypes(uint64(m.ExpirationEpoch))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
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
func (m *StorageGroup) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StorageGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidationDataSize", wireType)
			}
			m.ValidationDataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidationDataSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidationHash", wireType)
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
			m.ValidationHash = append(m.ValidationHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidationHash == nil {
				m.ValidationHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationEpoch", wireType)
			}
			m.ExpirationEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
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
			m.Members = append(m.Members, &grpc.ObjectID{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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