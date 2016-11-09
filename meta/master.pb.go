// Code generated by protoc-gen-gogo.
// source: master.proto
// DO NOT EDIT!

package meta

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NewIDRequest struct {
}

func (m *NewIDRequest) Reset()                    { *m = NewIDRequest{} }
func (m *NewIDRequest) String() string            { return proto.CompactTextString(m) }
func (*NewIDRequest) ProtoMessage()               {}
func (*NewIDRequest) Descriptor() ([]byte, []int) { return fileDescriptorMaster, []int{0} }

type NewIDResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ID     int64           `protobuf:"varint,2,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
}

func (m *NewIDResponse) Reset()                    { *m = NewIDResponse{} }
func (m *NewIDResponse) String() string            { return proto.CompactTextString(m) }
func (*NewIDResponse) ProtoMessage()               {}
func (*NewIDResponse) Descriptor() ([]byte, []int) { return fileDescriptorMaster, []int{1} }

func (m *NewIDResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type RegionGetRequest struct {
	Host string `protobuf:"bytes,1,opt,name=Host,json=host,proto3" json:"Host,omitempty"`
}

func (m *RegionGetRequest) Reset()                    { *m = RegionGetRequest{} }
func (m *RegionGetRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionGetRequest) ProtoMessage()               {}
func (*RegionGetRequest) Descriptor() ([]byte, []int) { return fileDescriptorMaster, []int{2} }

type RegionGetResponse struct {
	Header  *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Regions []Region        `protobuf:"bytes,2,rep,name=Regions,json=regions" json:"Regions"`
}

func (m *RegionGetResponse) Reset()                    { *m = RegionGetResponse{} }
func (m *RegionGetResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionGetResponse) ProtoMessage()               {}
func (*RegionGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorMaster, []int{3} }

func (m *RegionGetResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RegionGetResponse) GetRegions() []Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

func init() {
	proto.RegisterType((*NewIDRequest)(nil), "candy.meta.NewIDRequest")
	proto.RegisterType((*NewIDResponse)(nil), "candy.meta.NewIDResponse")
	proto.RegisterType((*RegionGetRequest)(nil), "candy.meta.RegionGetRequest")
	proto.RegisterType((*RegionGetResponse)(nil), "candy.meta.RegionGetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Master service

type MasterClient interface {
	NewID(ctx context.Context, in *NewIDRequest, opts ...grpc.CallOption) (*NewIDResponse, error)
	// RegionGet 获取region，如果host传空，返回所有region
	RegionGet(ctx context.Context, in *RegionGetRequest, opts ...grpc.CallOption) (*RegionGetResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) NewID(ctx context.Context, in *NewIDRequest, opts ...grpc.CallOption) (*NewIDResponse, error) {
	out := new(NewIDResponse)
	err := grpc.Invoke(ctx, "/candy.meta.Master/NewID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) RegionGet(ctx context.Context, in *RegionGetRequest, opts ...grpc.CallOption) (*RegionGetResponse, error) {
	out := new(RegionGetResponse)
	err := grpc.Invoke(ctx, "/candy.meta.Master/RegionGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	NewID(context.Context, *NewIDRequest) (*NewIDResponse, error)
	// RegionGet 获取region，如果host传空，返回所有region
	RegionGet(context.Context, *RegionGetRequest) (*RegionGetResponse, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_NewID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).NewID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candy.meta.Master/NewID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).NewID(ctx, req.(*NewIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_RegionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RegionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candy.meta.Master/RegionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RegionGet(ctx, req.(*RegionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "candy.meta.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewID",
			Handler:    _Master_NewID_Handler,
		},
		{
			MethodName: "RegionGet",
			Handler:    _Master_RegionGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorMaster,
}

func (m *NewIDRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NewIDRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *NewIDResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NewIDResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMaster(data, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ID != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintMaster(data, i, uint64(m.ID))
	}
	return i, nil
}

func (m *RegionGetRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RegionGetRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Host) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMaster(data, i, uint64(len(m.Host)))
		i += copy(data[i:], m.Host)
	}
	return i, nil
}

func (m *RegionGetResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RegionGetResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMaster(data, i, uint64(m.Header.Size()))
		n2, err := m.Header.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Regions) > 0 {
		for _, msg := range m.Regions {
			data[i] = 0x12
			i++
			i = encodeVarintMaster(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Master(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Master(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMaster(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *NewIDRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *NewIDResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovMaster(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovMaster(uint64(m.ID))
	}
	return n
}

func (m *RegionGetRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovMaster(uint64(l))
	}
	return n
}

func (m *RegionGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovMaster(uint64(l))
	}
	if len(m.Regions) > 0 {
		for _, e := range m.Regions {
			l = e.Size()
			n += 1 + l + sovMaster(uint64(l))
		}
	}
	return n
}

func sovMaster(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMaster(x uint64) (n int) {
	return sovMaster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NewIDRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NewIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMaster(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaster
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
func (m *NewIDResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NewIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaster
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMaster(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaster
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
func (m *RegionGetRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMaster
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaster(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaster
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
func (m *RegionGetResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaster
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaster
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Regions = append(m.Regions, Region{})
			if err := m.Regions[len(m.Regions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaster(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaster
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
func skipMaster(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMaster
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMaster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMaster
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMaster
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMaster(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMaster = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMaster   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("master.proto", fileDescriptorMaster) }

var fileDescriptorMaster = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2c, 0x2e,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4e, 0xcc, 0x4b, 0xa9, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xeb, 0x83, 0x58, 0x10,
	0x15, 0x52, 0x3c, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x9e, 0x12, 0x1f, 0x17, 0x8f, 0x5f,
	0x6a, 0xb9, 0xa7, 0x4b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x52, 0x30, 0x17, 0x2f, 0x94,
	0x5f, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc4, 0xc5, 0x96, 0x91, 0x9a, 0x98, 0x92, 0x5a,
	0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x87, 0xb0, 0x41, 0x0f, 0xa6, 0xca, 0x03,
	0xac, 0x22, 0x08, 0xaa, 0x52, 0x88, 0x8f, 0x8b, 0xc9, 0xd3, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x39, 0x88, 0x29, 0xd3, 0x45, 0x49, 0x8d, 0x4b, 0x20, 0x28, 0x35, 0x3d, 0x33, 0x3f, 0xcf, 0x3d,
	0xb5, 0x04, 0x6a, 0x91, 0x90, 0x10, 0x17, 0x8b, 0x47, 0x7e, 0x71, 0x09, 0xd8, 0x54, 0xce, 0x20,
	0x96, 0x8c, 0xfc, 0xe2, 0x12, 0xa5, 0x6a, 0x2e, 0x41, 0x24, 0x75, 0x14, 0x38, 0xc0, 0x88, 0x8b,
	0x1d, 0x62, 0x50, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x10, 0xaa, 0x26, 0x90, 0x94,
	0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0xec, 0x45, 0x10, 0x85, 0x46, 0x13, 0x18, 0xb9, 0xd8,
	0x7c, 0xc1, 0x41, 0x29, 0x64, 0xc3, 0xc5, 0x0a, 0x0e, 0x04, 0x21, 0x09, 0x64, 0x6d, 0xc8, 0xe1,
	0x24, 0x25, 0x89, 0x45, 0x06, 0xea, 0x60, 0x0f, 0x2e, 0x4e, 0xb8, 0x2f, 0x84, 0x64, 0x30, 0x2d,
	0x46, 0x04, 0x82, 0x94, 0x2c, 0x0e, 0x59, 0x88, 0x49, 0x4e, 0x62, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x51, 0x2c, 0x20, 0x95,
	0x11, 0x0c, 0x49, 0x6c, 0xe0, 0xd8, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0x30, 0xa8,
	0x2e, 0xfd, 0x01, 0x00, 0x00,
}
