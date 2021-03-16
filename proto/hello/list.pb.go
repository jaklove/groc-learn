// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello/list.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_526de67ae34f0998, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Pager struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalRows            int64    `protobuf:"varint,3,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_526de67ae34f0998, []int{1}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pager) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetTotalRows() int64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

type List struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_526de67ae34f0998, []int{2}
}

func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *List) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type GetListReply struct {
	List                 []*List  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListReply) Reset()         { *m = GetListReply{} }
func (m *GetListReply) String() string { return proto.CompactTextString(m) }
func (*GetListReply) ProtoMessage()    {}
func (*GetListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_526de67ae34f0998, []int{3}
}

func (m *GetListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListReply.Unmarshal(m, b)
}
func (m *GetListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListReply.Marshal(b, m, deterministic)
}
func (m *GetListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListReply.Merge(m, src)
}
func (m *GetListReply) XXX_Size() int {
	return xxx_messageInfo_GetListReply.Size(m)
}
func (m *GetListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetListReply proto.InternalMessageInfo

func (m *GetListReply) GetList() []*List {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetListReply) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "hello.ListRequest")
	proto.RegisterType((*Pager)(nil), "hello.Pager")
	proto.RegisterType((*List)(nil), "hello.List")
	proto.RegisterType((*GetListReply)(nil), "hello.GetListReply")
}

func init() { proto.RegisterFile("proto/hello/list.proto", fileDescriptor_526de67ae34f0998) }

var fileDescriptor_526de67ae34f0998 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x69, 0xfe, 0x00, 0xb9, 0x14, 0x86, 0x03, 0xa1, 0x08, 0x84, 0x28, 0x9e, 0x3a, 0xa5,
	0x52, 0xe0, 0x03, 0x20, 0x31, 0xb0, 0x30, 0x20, 0x67, 0x60, 0xac, 0x0c, 0x3d, 0x15, 0x4b, 0x06,
	0x07, 0xdb, 0x50, 0xd1, 0x4f, 0x8f, 0x7c, 0x8e, 0xa0, 0x53, 0x2e, 0xef, 0x67, 0xbd, 0x7b, 0xef,
	0xe0, 0x6c, 0x70, 0x36, 0xd8, 0xc5, 0x1b, 0x19, 0x63, 0x17, 0x46, 0xfb, 0xd0, 0xb2, 0x80, 0x25,
	0x2b, 0xe2, 0x1a, 0xea, 0x47, 0xed, 0x83, 0xa4, 0xcf, 0x2f, 0xf2, 0x01, 0x11, 0x8a, 0x0f, 0xf5,
	0x4e, 0xcd, 0x64, 0x36, 0x99, 0x57, 0x92, 0x67, 0xf1, 0x0c, 0xe5, 0x93, 0x5a, 0x93, 0x8b, 0x70,
	0x50, 0xeb, 0x04, 0x73, 0xc9, 0x33, 0x5e, 0x40, 0x15, 0xbf, 0x4b, 0xaf, 0xb7, 0xd4, 0x64, 0x0c,
	0x0e, 0xa3, 0xd0, 0xeb, 0x2d, 0xe1, 0x25, 0x40, 0xb0, 0x41, 0x99, 0xa5, 0xb3, 0x1b, 0xdf, 0xe4,
	0x4c, 0x2b, 0x56, 0xa4, 0xdd, 0x78, 0x71, 0x07, 0x45, 0xdc, 0x8d, 0xc7, 0x90, 0xe9, 0xd5, 0xe8,
	0x9a, 0xe9, 0xd5, 0x5f, 0x88, 0xec, 0x3f, 0x04, 0x9e, 0x42, 0xe9, 0x83, 0x0a, 0xc4, 0x2e, 0x47,
	0x32, 0xfd, 0x88, 0x1e, 0xa6, 0x0f, 0x14, 0x52, 0x81, 0xc1, 0xfc, 0xe0, 0x15, 0x14, 0xb1, 0x62,
	0x33, 0x99, 0xe5, 0xf3, 0xba, 0xab, 0x5b, 0xee, 0xd8, 0x32, 0x67, 0x80, 0x02, 0xca, 0x98, 0xce,
	0xb1, 0x77, 0xdd, 0x4d, 0xc7, 0x17, 0xdc, 0x4f, 0x26, 0xd4, 0xdd, 0xa7, 0x93, 0xf4, 0xe4, 0xbe,
	0xf5, 0x2b, 0xe1, 0x2d, 0x1c, 0x8c, 0x3b, 0x10, 0x77, 0x0d, 0xd3, 0xc5, 0xce, 0x4f, 0x46, 0x6d,
	0x37, 0x87, 0xd8, 0x7b, 0xd9, 0xe7, 0x2b, 0xdf, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x84, 0x15,
	0xd6, 0x45, 0x7f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListServiceClient is the client API for ListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListServiceClient interface {
	GetList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*GetListReply, error)
}

type listServiceClient struct {
	cc *grpc.ClientConn
}

func NewListServiceClient(cc *grpc.ClientConn) ListServiceClient {
	return &listServiceClient{cc}
}

func (c *listServiceClient) GetList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*GetListReply, error) {
	out := new(GetListReply)
	err := c.cc.Invoke(ctx, "/hello.ListService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServiceServer is the server API for ListService service.
type ListServiceServer interface {
	GetList(context.Context, *ListRequest) (*GetListReply, error)
}

// UnimplementedListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListServiceServer struct {
}

func (*UnimplementedListServiceServer) GetList(ctx context.Context, req *ListRequest) (*GetListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterListServiceServer(s *grpc.Server, srv ListServiceServer) {
	s.RegisterService(&_ListService_serviceDesc, srv)
}

func _ListService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.ListService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServiceServer).GetList(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.ListService",
	HandlerType: (*ListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _ListService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello/list.proto",
}
