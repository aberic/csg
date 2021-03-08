// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/multi_sign.proto

package common

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// voting status
type VoteStatus int32

const (
	// disagree
	VoteStatus_DISAGREE VoteStatus = 0
	// agree
	VoteStatus_AGREE VoteStatus = 1
)

var VoteStatus_name = map[int32]string{
	0: "DISAGREE",
	1: "AGREE",
}

var VoteStatus_value = map[string]int32{
	"DISAGREE": 0,
	"AGREE":    1,
}

func (x VoteStatus) String() string {
	return proto.EnumName(VoteStatus_name, int32(x))
}

func (VoteStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c10139b8b6f992e, []int{0}
}

// Multi signature status
type MultiSignStatus int32

const (
	// processing
	MultiSignStatus_PROCESSING MultiSignStatus = 0
	// adopted
	MultiSignStatus_ADOPTED MultiSignStatus = 1
	// refused
	MultiSignStatus_REFUSED MultiSignStatus = 2
	// expired
	MultiSignStatus_EXPIRED MultiSignStatus = 3
)

var MultiSignStatus_name = map[int32]string{
	0: "PROCESSING",
	1: "ADOPTED",
	2: "REFUSED",
	3: "EXPIRED",
}

var MultiSignStatus_value = map[string]int32{
	"PROCESSING": 0,
	"ADOPTED":    1,
	"REFUSED":    2,
	"EXPIRED":    3,
}

func (x MultiSignStatus) String() string {
	return proto.EnumName(MultiSignStatus_name, int32(x))
}

func (MultiSignStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c10139b8b6f992e, []int{1}
}

type MultiSignInfo struct {
	// trancation identifier
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// transaction type in case of multiple signatures
	TxType string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	// bytes of payload
	PayloadBytes []byte `protobuf:"bytes,3,opt,name=payload_bytes,json=payloadBytes,proto3" json:"payload_bytes,omitempty"`
	// the height of the block to the deadline
	DeadlineBlock uint64 `protobuf:"varint,4,opt,name=deadline_block,json=deadlineBlock,proto3" json:"deadline_block,omitempty"`
	// multi signature information
	VoteInfos []*MultiSignVoteInfo `protobuf:"bytes,5,rep,name=voteInfos,proto3" json:"voteInfos,omitempty"`
	// status of mutil signature
	Status MultiSignStatus `protobuf:"varint,6,opt,name=status,proto3,enum=common.MultiSignStatus" json:"status,omitempty"`
}

func (m *MultiSignInfo) Reset()         { *m = MultiSignInfo{} }
func (m *MultiSignInfo) String() string { return proto.CompactTextString(m) }
func (*MultiSignInfo) ProtoMessage()    {}
func (*MultiSignInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c10139b8b6f992e, []int{0}
}
func (m *MultiSignInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSignInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSignInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSignInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSignInfo.Merge(m, src)
}
func (m *MultiSignInfo) XXX_Size() int {
	return m.Size()
}
func (m *MultiSignInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSignInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSignInfo proto.InternalMessageInfo

func (m *MultiSignInfo) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *MultiSignInfo) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *MultiSignInfo) GetPayloadBytes() []byte {
	if m != nil {
		return m.PayloadBytes
	}
	return nil
}

func (m *MultiSignInfo) GetDeadlineBlock() uint64 {
	if m != nil {
		return m.DeadlineBlock
	}
	return 0
}

func (m *MultiSignInfo) GetVoteInfos() []*MultiSignVoteInfo {
	if m != nil {
		return m.VoteInfos
	}
	return nil
}

func (m *MultiSignInfo) GetStatus() MultiSignStatus {
	if m != nil {
		return m.Status
	}
	return MultiSignStatus_PROCESSING
}

type MultiSignVoteInfo struct {
	// voting status
	Vote VoteStatus `protobuf:"varint,1,opt,name=vote,proto3,enum=common.VoteStatus" json:"vote,omitempty"`
	// signature information
	Endorsement *EndorsementEntry `protobuf:"bytes,2,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
}

func (m *MultiSignVoteInfo) Reset()         { *m = MultiSignVoteInfo{} }
func (m *MultiSignVoteInfo) String() string { return proto.CompactTextString(m) }
func (*MultiSignVoteInfo) ProtoMessage()    {}
func (*MultiSignVoteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c10139b8b6f992e, []int{1}
}
func (m *MultiSignVoteInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSignVoteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSignVoteInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSignVoteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSignVoteInfo.Merge(m, src)
}
func (m *MultiSignVoteInfo) XXX_Size() int {
	return m.Size()
}
func (m *MultiSignVoteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSignVoteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSignVoteInfo proto.InternalMessageInfo

func (m *MultiSignVoteInfo) GetVote() VoteStatus {
	if m != nil {
		return m.Vote
	}
	return VoteStatus_DISAGREE
}

func (m *MultiSignVoteInfo) GetEndorsement() *EndorsementEntry {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// mutil signature response message
type MultiSignResp struct {
	// transaction identifier
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// payload hash
	PayloadHash string `protobuf:"bytes,2,opt,name=payload_hash,json=payloadHash,proto3" json:"payload_hash,omitempty"`
}

func (m *MultiSignResp) Reset()         { *m = MultiSignResp{} }
func (m *MultiSignResp) String() string { return proto.CompactTextString(m) }
func (*MultiSignResp) ProtoMessage()    {}
func (*MultiSignResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c10139b8b6f992e, []int{2}
}
func (m *MultiSignResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSignResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSignResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSignResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSignResp.Merge(m, src)
}
func (m *MultiSignResp) XXX_Size() int {
	return m.Size()
}
func (m *MultiSignResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSignResp.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSignResp proto.InternalMessageInfo

func (m *MultiSignResp) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *MultiSignResp) GetPayloadHash() string {
	if m != nil {
		return m.PayloadHash
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.VoteStatus", VoteStatus_name, VoteStatus_value)
	proto.RegisterEnum("common.MultiSignStatus", MultiSignStatus_name, MultiSignStatus_value)
	proto.RegisterType((*MultiSignInfo)(nil), "common.MultiSignInfo")
	proto.RegisterType((*MultiSignVoteInfo)(nil), "common.MultiSignVoteInfo")
	proto.RegisterType((*MultiSignResp)(nil), "common.MultiSignResp")
}

func init() { proto.RegisterFile("common/multi_sign.proto", fileDescriptor_3c10139b8b6f992e) }

var fileDescriptor_3c10139b8b6f992e = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0xcd, 0x34, 0x97, 0xfe, 0xf9, 0x72, 0xf9, 0xc3, 0x80, 0x14, 0xc3, 0xc2, 0x0a, 0x41, 0x45,
	0x51, 0xa5, 0xc6, 0x92, 0x59, 0x20, 0xb1, 0x6b, 0xf0, 0x90, 0x7a, 0x01, 0x8d, 0xc6, 0x05, 0x21,
	0x36, 0x96, 0x1d, 0x4f, 0x1d, 0x2b, 0xb1, 0xc7, 0x78, 0x26, 0x10, 0xbf, 0x05, 0x8f, 0xc5, 0xb2,
	0x4b, 0x96, 0x28, 0x79, 0x04, 0x5e, 0x00, 0xf9, 0x12, 0x12, 0x51, 0x76, 0x3e, 0xb7, 0x4f, 0xc7,
	0x47, 0x03, 0xfd, 0x39, 0x0f, 0x43, 0x1e, 0x69, 0xe1, 0x7a, 0x25, 0x03, 0x5b, 0x04, 0x7e, 0x34,
	0x8e, 0x13, 0x2e, 0x39, 0x6e, 0x14, 0xc2, 0x93, 0x47, 0xa5, 0x21, 0x61, 0x9f, 0xd7, 0x4c, 0xc8,
	0x42, 0x1d, 0xfe, 0x42, 0xd0, 0x79, 0x9b, 0x45, 0xac, 0xc0, 0x8f, 0xcc, 0xe8, 0x96, 0xe3, 0x87,
	0x50, 0x97, 0x1b, 0x3b, 0xf0, 0x14, 0x34, 0x40, 0xa3, 0x26, 0xad, 0xc9, 0x8d, 0xe9, 0xe1, 0x3e,
	0x9c, 0xca, 0x8d, 0x2d, 0xd3, 0x98, 0x29, 0x27, 0x39, 0xdd, 0x90, 0x9b, 0x9b, 0x34, 0x66, 0xf8,
	0x19, 0x74, 0x62, 0x27, 0x5d, 0x71, 0xc7, 0xb3, 0xdd, 0x54, 0x32, 0xa1, 0x54, 0x07, 0x68, 0xd4,
	0xa6, 0xed, 0x92, 0x9c, 0x64, 0x1c, 0x3e, 0x83, 0xae, 0xc7, 0x1c, 0x6f, 0x15, 0x44, 0xcc, 0x76,
	0x57, 0x7c, 0xbe, 0x54, 0x6a, 0x03, 0x34, 0xaa, 0xd1, 0xce, 0x9e, 0x9d, 0x64, 0x24, 0x7e, 0x09,
	0xcd, 0x2f, 0x5c, 0xb2, 0xac, 0x85, 0x50, 0xea, 0x83, 0xea, 0xa8, 0xa5, 0x3f, 0x1e, 0x17, 0xad,
	0xc7, 0x7f, 0x3a, 0x7e, 0x28, 0x1d, 0xf4, 0xe0, 0xc5, 0x1a, 0x34, 0x84, 0x74, 0xe4, 0x5a, 0x28,
	0x8d, 0x01, 0x1a, 0x75, 0xf5, 0xfe, 0xbd, 0x94, 0x95, 0xcb, 0xb4, 0xb4, 0x0d, 0xbf, 0xc2, 0x83,
	0x7b, 0x07, 0xf1, 0x73, 0xa8, 0x65, 0x27, 0xf3, 0xff, 0xee, 0xea, 0x78, 0x7f, 0x23, 0xd3, 0xcb,
	0x78, 0xae, 0xe3, 0x57, 0xd0, 0x62, 0x91, 0xc7, 0x13, 0xc1, 0x42, 0x16, 0xc9, 0x7c, 0x8f, 0x96,
	0xae, 0xec, 0xed, 0xe4, 0x20, 0x91, 0x48, 0x26, 0x29, 0x3d, 0x36, 0x0f, 0xa7, 0x47, 0x6b, 0x53,
	0x26, 0xe2, 0x7f, 0xaf, 0xfd, 0x14, 0xf6, 0xfb, 0xd9, 0x0b, 0x47, 0x2c, 0xca, 0xc9, 0x5b, 0x25,
	0x77, 0xe5, 0x88, 0xc5, 0xf9, 0x19, 0xc0, 0xa1, 0x18, 0x6e, 0xc3, 0x7f, 0x86, 0x69, 0x5d, 0x4e,
	0x29, 0x21, 0xbd, 0x0a, 0x6e, 0x42, 0xbd, 0xf8, 0x44, 0xe7, 0x57, 0xf0, 0xff, 0x5f, 0x1b, 0xe0,
	0x2e, 0xc0, 0x8c, 0x5e, 0xbf, 0x26, 0x96, 0x65, 0xbe, 0x9b, 0xf6, 0x2a, 0xb8, 0x05, 0xa7, 0x97,
	0xc6, 0xf5, 0xec, 0x86, 0x18, 0x3d, 0x94, 0x01, 0x4a, 0xde, 0xbc, 0xb7, 0x88, 0xd1, 0x3b, 0xc9,
	0x00, 0xf9, 0x38, 0x33, 0x29, 0x31, 0x7a, 0xd5, 0xc9, 0xed, 0xf7, 0xad, 0x8a, 0xee, 0xb6, 0x2a,
	0xfa, 0xb9, 0x55, 0xd1, 0xb7, 0x9d, 0x5a, 0xb9, 0xdb, 0xa9, 0x95, 0x1f, 0x3b, 0xb5, 0x02, 0x0a,
	0x4f, 0xfc, 0xf1, 0x7c, 0xe1, 0x04, 0x51, 0xe8, 0x2c, 0x59, 0x32, 0x8e, 0xdd, 0x72, 0x8b, 0x4f,
	0xfa, 0x11, 0xcb, 0x13, 0x5f, 0x3b, 0xc0, 0x0b, 0xe1, 0x2d, 0x2f, 0x7c, 0xae, 0xc5, 0xae, 0x96,
	0xbf, 0x43, 0x9f, 0x6b, 0x45, 0xc6, 0x6d, 0xe4, 0xf8, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0x54, 0x73, 0x99, 0xd0, 0x02, 0x00, 0x00,
}

func (m *MultiSignInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSignInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSignInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMultiSign(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.VoteInfos) > 0 {
		for iNdEx := len(m.VoteInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VoteInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMultiSign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.DeadlineBlock != 0 {
		i = encodeVarintMultiSign(dAtA, i, uint64(m.DeadlineBlock))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PayloadBytes) > 0 {
		i -= len(m.PayloadBytes)
		copy(dAtA[i:], m.PayloadBytes)
		i = encodeVarintMultiSign(dAtA, i, uint64(len(m.PayloadBytes)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintMultiSign(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintMultiSign(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MultiSignVoteInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSignVoteInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSignVoteInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Endorsement != nil {
		{
			size, err := m.Endorsement.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMultiSign(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Vote != 0 {
		i = encodeVarintMultiSign(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MultiSignResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSignResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSignResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PayloadHash) > 0 {
		i -= len(m.PayloadHash)
		copy(dAtA[i:], m.PayloadHash)
		i = encodeVarintMultiSign(dAtA, i, uint64(len(m.PayloadHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintMultiSign(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMultiSign(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultiSign(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiSignInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovMultiSign(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovMultiSign(uint64(l))
	}
	l = len(m.PayloadBytes)
	if l > 0 {
		n += 1 + l + sovMultiSign(uint64(l))
	}
	if m.DeadlineBlock != 0 {
		n += 1 + sovMultiSign(uint64(m.DeadlineBlock))
	}
	if len(m.VoteInfos) > 0 {
		for _, e := range m.VoteInfos {
			l = e.Size()
			n += 1 + l + sovMultiSign(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovMultiSign(uint64(m.Status))
	}
	return n
}

func (m *MultiSignVoteInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vote != 0 {
		n += 1 + sovMultiSign(uint64(m.Vote))
	}
	if m.Endorsement != nil {
		l = m.Endorsement.Size()
		n += 1 + l + sovMultiSign(uint64(l))
	}
	return n
}

func (m *MultiSignResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovMultiSign(uint64(l))
	}
	l = len(m.PayloadHash)
	if l > 0 {
		n += 1 + l + sovMultiSign(uint64(l))
	}
	return n
}

func sovMultiSign(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultiSign(x uint64) (n int) {
	return sovMultiSign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiSignInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultiSign
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
			return fmt.Errorf("proto: MultiSignInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSignInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayloadBytes = append(m.PayloadBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.PayloadBytes == nil {
				m.PayloadBytes = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeadlineBlock", wireType)
			}
			m.DeadlineBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeadlineBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteInfos = append(m.VoteInfos, &MultiSignVoteInfo{})
			if err := m.VoteInfos[len(m.VoteInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MultiSignStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMultiSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultiSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiSignVoteInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultiSign
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
			return fmt.Errorf("proto: MultiSignVoteInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSignVoteInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= VoteStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endorsement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Endorsement == nil {
				m.Endorsement = &EndorsementEntry{}
			}
			if err := m.Endorsement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultiSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultiSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiSignResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultiSign
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
			return fmt.Errorf("proto: MultiSignResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSignResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultiSign
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
				return ErrInvalidLengthMultiSign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultiSign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayloadHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultiSign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultiSign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMultiSign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultiSign
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
					return 0, ErrIntOverflowMultiSign
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
					return 0, ErrIntOverflowMultiSign
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
				return 0, ErrInvalidLengthMultiSign
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultiSign
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultiSign
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultiSign        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultiSign          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultiSign = fmt.Errorf("proto: unexpected end of group")
)
