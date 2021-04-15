// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: raft.proto

package raftpb

import (
	"fmt"
	"io"
)

type SnapshotHeader struct {
	SessionSize     uint64
	DataStoreSize   uint64
	UnreliableTime  uint64
	GitVersion      string
	HeaderChecksum  []byte
	PayloadChecksum []byte
	ChecksumType    ChecksumType
	Version         uint64
	CompressionType CompressionType
}

func (m *SnapshotHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.SessionSize))
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.DataStoreSize))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.UnreliableTime))
	dAtA[i] = 0x22
	i++
	i = encodeVarintRaft(dAtA, i, uint64(len(m.GitVersion)))
	i += copy(dAtA[i:], m.GitVersion)
	if m.HeaderChecksum != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.HeaderChecksum)))
		i += copy(dAtA[i:], m.HeaderChecksum)
	}
	if m.PayloadChecksum != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.PayloadChecksum)))
		i += copy(dAtA[i:], m.PayloadChecksum)
	}
	dAtA[i] = 0x38
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.ChecksumType))
	dAtA[i] = 0x40
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Version))
	dAtA[i] = 0x48
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.CompressionType))
	return i, nil
}

func (m *SnapshotHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.SessionSize))
	n += 1 + sovRaft(uint64(m.DataStoreSize))
	n += 1 + sovRaft(uint64(m.UnreliableTime))
	l = len(m.GitVersion)
	n += 1 + l + sovRaft(uint64(l))
	if m.HeaderChecksum != nil {
		l = len(m.HeaderChecksum)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.PayloadChecksum != nil {
		l = len(m.PayloadChecksum)
		n += 1 + l + sovRaft(uint64(l))
	}
	n += 1 + sovRaft(uint64(m.ChecksumType))
	n += 1 + sovRaft(uint64(m.Version))
	n += 1 + sovRaft(uint64(m.CompressionType))
	return n
}

func (m *SnapshotHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: SnapshotHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionSize", wireType)
			}
			m.SessionSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataStoreSize", wireType)
			}
			m.DataStoreSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataStoreSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnreliableTime", wireType)
			}
			m.UnreliableTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnreliableTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderChecksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderChecksum = append(m.HeaderChecksum[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderChecksum == nil {
				m.HeaderChecksum = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadChecksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRaft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayloadChecksum = append(m.PayloadChecksum[:0], dAtA[iNdEx:postIndex]...)
			if m.PayloadChecksum == nil {
				m.PayloadChecksum = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChecksumType", wireType)
			}
			m.ChecksumType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChecksumType |= ChecksumType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionType", wireType)
			}
			m.CompressionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionType |= CompressionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRaft
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