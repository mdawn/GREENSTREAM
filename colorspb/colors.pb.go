// Code generated by protoc-gen-go. DO NOT EDIT.
// source: colorspb/colors.proto

package colorspb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type Coloring struct {
	Adjective            string   `protobuf:"bytes,1,opt,name=adjective,proto3" json:"adjective,omitempty"`
	BaseColor            string   `protobuf:"bytes,2,opt,name=base_color,json=baseColor,proto3" json:"base_color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coloring) Reset()         { *m = Coloring{} }
func (m *Coloring) String() string { return proto.CompactTextString(m) }
func (*Coloring) ProtoMessage()    {}
func (*Coloring) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16566239bd98cfd, []int{0}
}

func (m *Coloring) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coloring.Unmarshal(m, b)
}
func (m *Coloring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coloring.Marshal(b, m, deterministic)
}
func (m *Coloring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coloring.Merge(m, src)
}
func (m *Coloring) XXX_Size() int {
	return xxx_messageInfo_Coloring.Size(m)
}
func (m *Coloring) XXX_DiscardUnknown() {
	xxx_messageInfo_Coloring.DiscardUnknown(m)
}

var xxx_messageInfo_Coloring proto.InternalMessageInfo

func (m *Coloring) GetAdjective() string {
	if m != nil {
		return m.Adjective
	}
	return ""
}

func (m *Coloring) GetBaseColor() string {
	if m != nil {
		return m.BaseColor
	}
	return ""
}

type ColorRequest struct {
	Colors               *Coloring `protobuf:"bytes,1,opt,name=colors,proto3" json:"colors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ColorRequest) Reset()         { *m = ColorRequest{} }
func (m *ColorRequest) String() string { return proto.CompactTextString(m) }
func (*ColorRequest) ProtoMessage()    {}
func (*ColorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16566239bd98cfd, []int{1}
}

func (m *ColorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColorRequest.Unmarshal(m, b)
}
func (m *ColorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColorRequest.Marshal(b, m, deterministic)
}
func (m *ColorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColorRequest.Merge(m, src)
}
func (m *ColorRequest) XXX_Size() int {
	return xxx_messageInfo_ColorRequest.Size(m)
}
func (m *ColorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ColorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ColorRequest proto.InternalMessageInfo

func (m *ColorRequest) GetColors() *Coloring {
	if m != nil {
		return m.Colors
	}
	return nil
}

type ColorResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColorResponse) Reset()         { *m = ColorResponse{} }
func (m *ColorResponse) String() string { return proto.CompactTextString(m) }
func (*ColorResponse) ProtoMessage()    {}
func (*ColorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16566239bd98cfd, []int{2}
}

func (m *ColorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColorResponse.Unmarshal(m, b)
}
func (m *ColorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColorResponse.Marshal(b, m, deterministic)
}
func (m *ColorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColorResponse.Merge(m, src)
}
func (m *ColorResponse) XXX_Size() int {
	return xxx_messageInfo_ColorResponse.Size(m)
}
func (m *ColorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ColorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ColorResponse proto.InternalMessageInfo

func (m *ColorResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type ColorEverywhereRequest struct {
	Coloring             *Coloring `protobuf:"bytes,1,opt,name=coloring,proto3" json:"coloring,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ColorEverywhereRequest) Reset()         { *m = ColorEverywhereRequest{} }
func (m *ColorEverywhereRequest) String() string { return proto.CompactTextString(m) }
func (*ColorEverywhereRequest) ProtoMessage()    {}
func (*ColorEverywhereRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16566239bd98cfd, []int{3}
}

func (m *ColorEverywhereRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColorEverywhereRequest.Unmarshal(m, b)
}
func (m *ColorEverywhereRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColorEverywhereRequest.Marshal(b, m, deterministic)
}
func (m *ColorEverywhereRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColorEverywhereRequest.Merge(m, src)
}
func (m *ColorEverywhereRequest) XXX_Size() int {
	return xxx_messageInfo_ColorEverywhereRequest.Size(m)
}
func (m *ColorEverywhereRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ColorEverywhereRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ColorEverywhereRequest proto.InternalMessageInfo

func (m *ColorEverywhereRequest) GetColoring() *Coloring {
	if m != nil {
		return m.Coloring
	}
	return nil
}

type ColorEverywhereResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColorEverywhereResponse) Reset()         { *m = ColorEverywhereResponse{} }
func (m *ColorEverywhereResponse) String() string { return proto.CompactTextString(m) }
func (*ColorEverywhereResponse) ProtoMessage()    {}
func (*ColorEverywhereResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16566239bd98cfd, []int{4}
}

func (m *ColorEverywhereResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColorEverywhereResponse.Unmarshal(m, b)
}
func (m *ColorEverywhereResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColorEverywhereResponse.Marshal(b, m, deterministic)
}
func (m *ColorEverywhereResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColorEverywhereResponse.Merge(m, src)
}
func (m *ColorEverywhereResponse) XXX_Size() int {
	return xxx_messageInfo_ColorEverywhereResponse.Size(m)
}
func (m *ColorEverywhereResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ColorEverywhereResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ColorEverywhereResponse proto.InternalMessageInfo

func (m *ColorEverywhereResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Coloring)(nil), "colors.Coloring")
	proto.RegisterType((*ColorRequest)(nil), "colors.ColorRequest")
	proto.RegisterType((*ColorResponse)(nil), "colors.ColorResponse")
	proto.RegisterType((*ColorEverywhereRequest)(nil), "colors.ColorEverywhereRequest")
	proto.RegisterType((*ColorEverywhereResponse)(nil), "colors.ColorEverywhereResponse")
}

func init() { proto.RegisterFile("colorspb/colors.proto", fileDescriptor_a16566239bd98cfd) }

var fileDescriptor_a16566239bd98cfd = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xc9,
	0x2f, 0x2a, 0x2e, 0x48, 0xd2, 0x87, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20,
	0x3c, 0x25, 0x77, 0x2e, 0x0e, 0x67, 0x10, 0x2b, 0x33, 0x2f, 0x5d, 0x48, 0x86, 0x8b, 0x33, 0x31,
	0x25, 0x2b, 0x35, 0xb9, 0x24, 0xb3, 0x2c, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21,
	0x20, 0x24, 0xcb, 0xc5, 0x95, 0x94, 0x58, 0x9c, 0x1a, 0x0f, 0xd6, 0x28, 0xc1, 0x04, 0x91, 0x06,
	0x89, 0x80, 0xf5, 0x2b, 0x59, 0x70, 0xf1, 0x80, 0x19, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x1a, 0x5c, 0x50, 0x2b, 0xc0, 0x26, 0x71, 0x1b, 0x09, 0xe8, 0x41, 0xed, 0x87, 0x59, 0x17,
	0x04, 0x73, 0x82, 0x3a, 0x17, 0x2f, 0x54, 0x67, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18,
	0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0xd4, 0x11, 0x50, 0x9e, 0x92, 0x1b, 0x97, 0x18,
	0x58, 0xa1, 0x6b, 0x59, 0x6a, 0x51, 0x65, 0x79, 0x46, 0x6a, 0x51, 0x2a, 0xcc, 0x32, 0x1d, 0x2e,
	0x8e, 0x64, 0xa8, 0xb1, 0x38, 0xad, 0x83, 0xab, 0x50, 0x32, 0xe4, 0x12, 0xc7, 0x30, 0x07, 0xbf,
	0xd5, 0x46, 0x19, 0x50, 0xdf, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x45, 0x70, 0xf1,
	0xa3, 0x19, 0x21, 0x24, 0x87, 0x62, 0x23, 0x86, 0x1b, 0xa5, 0xe4, 0x71, 0xca, 0x43, 0xec, 0x56,
	0x62, 0xd0, 0x60, 0x34, 0x60, 0x74, 0xe2, 0x8a, 0xe2, 0x80, 0xc5, 0x58, 0x12, 0x1b, 0x38, 0xae,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x44, 0x40, 0x5f, 0xc4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ColorServiceClient is the client API for ColorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ColorServiceClient interface {
	// BiDi Streaming
	ColorEverywhere(ctx context.Context, opts ...grpc.CallOption) (ColorService_ColorEverywhereClient, error)
}

type colorServiceClient struct {
	cc *grpc.ClientConn
}

func NewColorServiceClient(cc *grpc.ClientConn) ColorServiceClient {
	return &colorServiceClient{cc}
}

func (c *colorServiceClient) ColorEverywhere(ctx context.Context, opts ...grpc.CallOption) (ColorService_ColorEverywhereClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ColorService_serviceDesc.Streams[0], "/colors.ColorService/ColorEverywhere", opts...)
	if err != nil {
		return nil, err
	}
	x := &colorServiceColorEverywhereClient{stream}
	return x, nil
}

type ColorService_ColorEverywhereClient interface {
	Send(*ColorEverywhereRequest) error
	Recv() (*ColorEverywhereResponse, error)
	grpc.ClientStream
}

type colorServiceColorEverywhereClient struct {
	grpc.ClientStream
}

func (x *colorServiceColorEverywhereClient) Send(m *ColorEverywhereRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *colorServiceColorEverywhereClient) Recv() (*ColorEverywhereResponse, error) {
	m := new(ColorEverywhereResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ColorServiceServer is the server API for ColorService service.
type ColorServiceServer interface {
	// BiDi Streaming
	ColorEverywhere(ColorService_ColorEverywhereServer) error
}

func RegisterColorServiceServer(s *grpc.Server, srv ColorServiceServer) {
	s.RegisterService(&_ColorService_serviceDesc, srv)
}

func _ColorService_ColorEverywhere_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ColorServiceServer).ColorEverywhere(&colorServiceColorEverywhereServer{stream})
}

type ColorService_ColorEverywhereServer interface {
	Send(*ColorEverywhereResponse) error
	Recv() (*ColorEverywhereRequest, error)
	grpc.ServerStream
}

type colorServiceColorEverywhereServer struct {
	grpc.ServerStream
}

func (x *colorServiceColorEverywhereServer) Send(m *ColorEverywhereResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *colorServiceColorEverywhereServer) Recv() (*ColorEverywhereRequest, error) {
	m := new(ColorEverywhereRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ColorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "colors.ColorService",
	HandlerType: (*ColorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ColorEverywhere",
			Handler:       _ColorService_ColorEverywhere_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "colorspb/colors.proto",
}
