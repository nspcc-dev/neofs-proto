// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bootstrap/types.proto

package bootstrap

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type SpreadMap struct {
	Epoch                uint64     `protobuf:"varint,1,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	NetMap               []NodeInfo `protobuf:"bytes,2,rep,name=NetMap,proto3" json:"NetMap"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SpreadMap) Reset()      { *m = SpreadMap{} }
func (*SpreadMap) ProtoMessage() {}
func (*SpreadMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_423083266369adee, []int{0}
}
func (m *SpreadMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpreadMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpreadMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpreadMap.Merge(m, src)
}
func (m *SpreadMap) XXX_Size() int {
	return m.Size()
}
func (m *SpreadMap) XXX_DiscardUnknown() {
	xxx_messageInfo_SpreadMap.DiscardUnknown(m)
}

var xxx_messageInfo_SpreadMap proto.InternalMessageInfo

func (m *SpreadMap) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *SpreadMap) GetNetMap() []NodeInfo {
	if m != nil {
		return m.NetMap
	}
	return nil
}

type NodeInfo struct {
	Address              string     `protobuf:"bytes,1,opt,name=Address,proto3" json:"address"`
	PubKey               []byte     `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"pubkey,omitempty"`
	Options              []string   `protobuf:"bytes,3,rep,name=Options,proto3" json:"options,omitempty"`
	Status               NodeStatus `protobuf:"varint,4,opt,name=Status,proto3,customtype=NodeStatus" json:"status"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NodeInfo) Reset()      { *m = NodeInfo{} }
func (*NodeInfo) ProtoMessage() {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_423083266369adee, []int{1}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NodeInfo) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *NodeInfo) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*SpreadMap)(nil), "bootstrap.SpreadMap")
	proto.RegisterType((*NodeInfo)(nil), "bootstrap.NodeInfo")
}

func init() { proto.RegisterFile("bootstrap/types.proto", fileDescriptor_423083266369adee) }

var fileDescriptor_423083266369adee = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcb, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x33, 0x6d, 0x4d, 0xcd, 0xd4, 0x85, 0xc6, 0x16, 0x82, 0x48, 0x12, 0x0a, 0x42, 0x90,
	0x36, 0xc1, 0xcb, 0x0b, 0x18, 0x10, 0x14, 0x69, 0x95, 0xd4, 0x95, 0xbb, 0x5c, 0xa6, 0x17, 0xa4,
	0xf9, 0x87, 0xcc, 0x44, 0xc8, 0xce, 0xc7, 0xf0, 0x89, 0xa4, 0x4b, 0x97, 0xc5, 0x45, 0xd0, 0xb8,
	0xcb, 0x53, 0x48, 0x27, 0x6d, 0xe9, 0xee, 0x3f, 0xe7, 0x7c, 0x33, 0xf3, 0xcf, 0xc1, 0x9d, 0x00,
	0x80, 0x33, 0x9e, 0xf8, 0xd4, 0xe1, 0x19, 0x25, 0xcc, 0xa6, 0x09, 0x70, 0x50, 0x95, 0xad, 0x7d,
	0xd2, 0x9f, 0xcc, 0xf8, 0x34, 0x0d, 0xec, 0x10, 0xe6, 0xce, 0x04, 0x26, 0xe0, 0x08, 0x22, 0x48,
	0xc7, 0x42, 0x09, 0x21, 0xa6, 0xea, 0x64, 0xf7, 0x19, 0x2b, 0x23, 0x9a, 0x10, 0x3f, 0x1a, 0xf8,
	0x54, 0x6d, 0xe3, 0xbd, 0x5b, 0x0a, 0xe1, 0x54, 0x43, 0x26, 0xb2, 0x1a, 0x5e, 0x25, 0xd4, 0x0b,
	0x2c, 0x0f, 0x09, 0x1f, 0xf8, 0x54, 0xab, 0x99, 0x75, 0xab, 0x75, 0x79, 0x6c, 0x6f, 0x5f, 0xb3,
	0x87, 0x10, 0x91, 0xfb, 0x78, 0x0c, 0x6e, 0x63, 0x91, 0x1b, 0x92, 0xb7, 0x06, 0xbb, 0x9f, 0x08,
	0xef, 0x6f, 0x22, 0xf5, 0x0c, 0x37, 0x6f, 0xa2, 0x28, 0x21, 0x8c, 0x89, 0x7b, 0x15, 0xb7, 0x55,
	0xe6, 0x46, 0xd3, 0xaf, 0x2c, 0x6f, 0x93, 0xa9, 0x3d, 0x2c, 0x3f, 0xa5, 0xc1, 0x03, 0xc9, 0xb4,
	0x9a, 0x89, 0xac, 0x03, 0xb7, 0x5d, 0xe6, 0xc6, 0x21, 0x4d, 0x83, 0x57, 0x92, 0xf5, 0x60, 0x3e,
	0xe3, 0x64, 0x4e, 0x79, 0xe6, 0xad, 0x19, 0xd5, 0xc1, 0xcd, 0x47, 0xca, 0x67, 0x10, 0x33, 0xad,
	0x6e, 0xd6, 0x2d, 0xc5, 0xed, 0x94, 0xb9, 0x71, 0x04, 0x95, 0xb5, 0xc3, 0x6f, 0x28, 0xf5, 0x1a,
	0xcb, 0x23, 0xee, 0xf3, 0x94, 0x69, 0x8d, 0xd5, 0xe7, 0xdc, 0xd3, 0xd5, 0xc2, 0xdf, 0xb9, 0x81,
	0x57, 0x7b, 0x56, 0x49, 0x99, 0x1b, 0x32, 0x13, 0x93, 0xb7, 0x66, 0xdd, 0xbb, 0xe5, 0xaf, 0x2e,
	0xbd, 0x17, 0xba, 0xb4, 0x28, 0x74, 0xf4, 0x55, 0xe8, 0x68, 0x59, 0xe8, 0xe8, 0xa7, 0xd0, 0xd1,
	0xc7, 0x9f, 0x2e, 0xbd, 0x9c, 0xef, 0x74, 0x1d, 0x33, 0x1a, 0x86, 0xfd, 0x88, 0xbc, 0x39, 0x31,
	0x81, 0x31, 0xeb, 0x57, 0x4d, 0x6f, 0x9b, 0x0a, 0x64, 0x61, 0x5c, 0xfd, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x71, 0xeb, 0x37, 0x57, 0xc2, 0x01, 0x00, 0x00,
}

func (m *SpreadMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpreadMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpreadMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NetMap) > 0 {
		for iNdEx := len(m.NetMap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetMap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Options[iNdEx])
			copy(dAtA[i:], m.Options[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Options[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *SpreadMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovTypes(uint64(m.Epoch))
	}
	if len(m.NetMap) > 0 {
		for _, e := range m.NetMap {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
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
func (m *SpreadMap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SpreadMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpreadMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetMap", wireType)
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
			m.NetMap = append(m.NetMap, NodeInfo{})
			if err := m.NetMap[len(m.NetMap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *NodeInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= NodeStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
