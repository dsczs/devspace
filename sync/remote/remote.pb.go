// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package remote

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

type ChangeType int32

const (
	ChangeType_CHANGE ChangeType = 0
	ChangeType_DELETE ChangeType = 1
)

var ChangeType_name = map[int32]string{
	0: "CHANGE",
	1: "DELETE",
}

var ChangeType_value = map[string]int32{
	"CHANGE": 0,
	"DELETE": 1,
}

func (x ChangeType) String() string {
	return proto.EnumName(ChangeType_name, int32(x))
}

func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{0}
}

type Watch struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Exclude              []string `protobuf:"bytes,2,rep,name=Exclude,proto3" json:"Exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Watch) Reset()         { *m = Watch{} }
func (m *Watch) String() string { return proto.CompactTextString(m) }
func (*Watch) ProtoMessage()    {}
func (*Watch) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{0}
}

func (m *Watch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Watch.Unmarshal(m, b)
}
func (m *Watch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Watch.Marshal(b, m, deterministic)
}
func (m *Watch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Watch.Merge(m, src)
}
func (m *Watch) XXX_Size() int {
	return xxx_messageInfo_Watch.Size(m)
}
func (m *Watch) XXX_DiscardUnknown() {
	xxx_messageInfo_Watch.DiscardUnknown(m)
}

var xxx_messageInfo_Watch proto.InternalMessageInfo

func (m *Watch) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Watch) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type Change struct {
	ChangeType           ChangeType `protobuf:"varint,1,opt,name=ChangeType,proto3,enum=remote.ChangeType" json:"ChangeType,omitempty"`
	Path                 string     `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	Mtime                int64      `protobuf:"varint,3,opt,name=Mtime,proto3" json:"Mtime,omitempty"`
	Size                 int64      `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	IsDir                bool       `protobuf:"varint,5,opt,name=IsDir,proto3" json:"IsDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{1}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetChangeType() ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return ChangeType_CHANGE
}

func (m *Change) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Change) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *Change) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Change) GetIsDir() bool {
	if m != nil {
		return m.IsDir
	}
	return false
}

type Excluded struct {
	Paths                []string `protobuf:"bytes,1,rep,name=Paths,proto3" json:"Paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Excluded) Reset()         { *m = Excluded{} }
func (m *Excluded) String() string { return proto.CompactTextString(m) }
func (*Excluded) ProtoMessage()    {}
func (*Excluded) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{2}
}

func (m *Excluded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Excluded.Unmarshal(m, b)
}
func (m *Excluded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Excluded.Marshal(b, m, deterministic)
}
func (m *Excluded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Excluded.Merge(m, src)
}
func (m *Excluded) XXX_Size() int {
	return xxx_messageInfo_Excluded.Size(m)
}
func (m *Excluded) XXX_DiscardUnknown() {
	xxx_messageInfo_Excluded.DiscardUnknown(m)
}

var xxx_messageInfo_Excluded proto.InternalMessageInfo

func (m *Excluded) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Path struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{3}
}

func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (m *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(m, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{4}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{5}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("remote.ChangeType", ChangeType_name, ChangeType_value)
	proto.RegisterType((*Watch)(nil), "remote.Watch")
	proto.RegisterType((*Change)(nil), "remote.Change")
	proto.RegisterType((*Excluded)(nil), "remote.Excluded")
	proto.RegisterType((*Path)(nil), "remote.Path")
	proto.RegisterType((*Chunk)(nil), "remote.Chunk")
	proto.RegisterType((*Empty)(nil), "remote.Empty")
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor_eefc82927d57d89b) }

var fileDescriptor_eefc82927d57d89b = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0xc6, 0x29, 0xb0, 0x31, 0x4e, 0x78, 0x09, 0x69, 0xde, 0x8b, 0x86, 0xab, 0xd9, 0x78, 0x31,
	0x49, 0x44, 0x32, 0xe3, 0x07, 0x30, 0x63, 0x51, 0x13, 0x35, 0xa6, 0x42, 0xbc, 0x9e, 0xd0, 0x38,
	0x02, 0x6b, 0x97, 0xad, 0xa8, 0xf8, 0x29, 0xfc, 0xc8, 0xa6, 0x2d, 0xe3, 0x4f, 0xa2, 0x77, 0xe7,
	0xd9, 0x79, 0x76, 0x7e, 0xe7, 0x39, 0x29, 0x74, 0x0a, 0x9e, 0x49, 0xc5, 0x87, 0x79, 0x21, 0x95,
	0xc4, 0xae, 0x55, 0xf4, 0x0a, 0x9c, 0x97, 0x44, 0xcd, 0x52, 0x8c, 0xa1, 0xf9, 0x94, 0xa8, 0x94,
	0x20, 0x1f, 0x05, 0x6d, 0x66, 0x6a, 0x4c, 0xa0, 0x15, 0x7f, 0xce, 0x56, 0xeb, 0x39, 0x27, 0x75,
	0xbf, 0x11, 0xb4, 0x59, 0x25, 0xe9, 0x37, 0x02, 0x37, 0x4a, 0x13, 0xf1, 0xc6, 0x71, 0x08, 0x60,
	0xab, 0xc9, 0x26, 0xe7, 0xe6, 0xf7, 0x6e, 0x88, 0x87, 0x5b, 0xd8, 0xbe, 0xc3, 0x0e, 0x5c, 0x3b,
	0x58, 0xfd, 0x00, 0xf6, 0x1f, 0x9c, 0x07, 0xb5, 0xc8, 0x38, 0x69, 0xf8, 0x28, 0x68, 0x30, 0x2b,
	0xb4, 0xf3, 0x79, 0xf1, 0xc5, 0x49, 0xd3, 0x7c, 0x34, 0xb5, 0x76, 0xde, 0x95, 0xe3, 0x45, 0x41,
	0x1c, 0x1f, 0x05, 0x1e, 0xb3, 0x82, 0xfa, 0xe0, 0x6d, 0xb7, 0x9b, 0x6b, 0x87, 0x9e, 0x59, 0x12,
	0x64, 0xd6, 0xb6, 0x82, 0xf6, 0x2d, 0xf5, 0xb7, 0xa8, 0xf4, 0x04, 0x9c, 0x28, 0x5d, 0x8b, 0xa5,
	0xce, 0x1c, 0x49, 0xa1, 0xb8, 0x50, 0xa6, 0xdf, 0x61, 0x95, 0xa4, 0x2d, 0x70, 0xe2, 0x2c, 0x57,
	0x9b, 0xc1, 0xe9, 0x61, 0x62, 0x0c, 0xe0, 0x46, 0xb7, 0xd7, 0x8f, 0x37, 0x71, 0xaf, 0xa6, 0xeb,
	0x71, 0x7c, 0x1f, 0x4f, 0xe2, 0x1e, 0x0a, 0x57, 0x00, 0x63, 0xf9, 0x21, 0x4a, 0x55, 0xf0, 0x24,
	0xc3, 0xe7, 0xe0, 0x69, 0xb5, 0x92, 0xc9, 0x1c, 0x77, 0xaa, 0xeb, 0x68, 0x72, 0xff, 0xdf, 0xfe,
	0x56, 0x6b, 0xb1, 0xa4, 0xb5, 0x00, 0x8d, 0x10, 0xbe, 0x80, 0x96, 0x45, 0x94, 0xb8, 0x57, 0xf5,
	0xab, 0x74, 0xfd, 0xee, 0xf1, 0x75, 0x69, 0x6d, 0x84, 0xc2, 0x04, 0xbc, 0x69, 0xbe, 0x65, 0x0d,
	0xc0, 0x9d, 0xe6, 0x86, 0x74, 0x3c, 0x7b, 0x8f, 0x32, 0x39, 0x34, 0x0a, 0x9f, 0x81, 0xcb, 0x78,
	0x26, 0xdf, 0xf9, 0x5f, 0x5b, 0xed, 0xac, 0xaf, 0xae, 0x79, 0x39, 0x97, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x5e, 0xdc, 0x55, 0x49, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownstreamClient is the client API for Downstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownstreamClient interface {
	Download(ctx context.Context, opts ...grpc.CallOption) (Downstream_DownloadClient, error)
	Changes(ctx context.Context, in *Excluded, opts ...grpc.CallOption) (Downstream_ChangesClient, error)
}

type downstreamClient struct {
	cc *grpc.ClientConn
}

func NewDownstreamClient(cc *grpc.ClientConn) DownstreamClient {
	return &downstreamClient{cc}
}

func (c *downstreamClient) Download(ctx context.Context, opts ...grpc.CallOption) (Downstream_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Downstream_serviceDesc.Streams[0], "/remote.Downstream/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &downstreamDownloadClient{stream}
	return x, nil
}

type Downstream_DownloadClient interface {
	Send(*Path) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type downstreamDownloadClient struct {
	grpc.ClientStream
}

func (x *downstreamDownloadClient) Send(m *Path) error {
	return x.ClientStream.SendMsg(m)
}

func (x *downstreamDownloadClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *downstreamClient) Changes(ctx context.Context, in *Excluded, opts ...grpc.CallOption) (Downstream_ChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Downstream_serviceDesc.Streams[1], "/remote.Downstream/Changes", opts...)
	if err != nil {
		return nil, err
	}
	x := &downstreamChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Downstream_ChangesClient interface {
	Recv() (*Change, error)
	grpc.ClientStream
}

type downstreamChangesClient struct {
	grpc.ClientStream
}

func (x *downstreamChangesClient) Recv() (*Change, error) {
	m := new(Change)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DownstreamServer is the server API for Downstream service.
type DownstreamServer interface {
	Download(Downstream_DownloadServer) error
	Changes(*Excluded, Downstream_ChangesServer) error
}

func RegisterDownstreamServer(s *grpc.Server, srv DownstreamServer) {
	s.RegisterService(&_Downstream_serviceDesc, srv)
}

func _Downstream_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DownstreamServer).Download(&downstreamDownloadServer{stream})
}

type Downstream_DownloadServer interface {
	Send(*Chunk) error
	Recv() (*Path, error)
	grpc.ServerStream
}

type downstreamDownloadServer struct {
	grpc.ServerStream
}

func (x *downstreamDownloadServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *downstreamDownloadServer) Recv() (*Path, error) {
	m := new(Path)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Downstream_Changes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Excluded)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownstreamServer).Changes(m, &downstreamChangesServer{stream})
}

type Downstream_ChangesServer interface {
	Send(*Change) error
	grpc.ServerStream
}

type downstreamChangesServer struct {
	grpc.ServerStream
}

func (x *downstreamChangesServer) Send(m *Change) error {
	return x.ServerStream.SendMsg(m)
}

var _Downstream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.Downstream",
	HandlerType: (*DownstreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Downstream_Download_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Changes",
			Handler:       _Downstream_Changes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote.proto",
}

// UpstreamClient is the client API for Upstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpstreamClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Upstream_UploadClient, error)
	Remove(ctx context.Context, opts ...grpc.CallOption) (Upstream_RemoveClient, error)
}

type upstreamClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamClient(cc *grpc.ClientConn) UpstreamClient {
	return &upstreamClient{cc}
}

func (c *upstreamClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Upstream_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upstream_serviceDesc.Streams[0], "/remote.Upstream/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamUploadClient{stream}
	return x, nil
}

type Upstream_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type upstreamUploadClient struct {
	grpc.ClientStream
}

func (x *upstreamUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *upstreamUploadClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *upstreamClient) Remove(ctx context.Context, opts ...grpc.CallOption) (Upstream_RemoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upstream_serviceDesc.Streams[1], "/remote.Upstream/Remove", opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamRemoveClient{stream}
	return x, nil
}

type Upstream_RemoveClient interface {
	Send(*Path) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type upstreamRemoveClient struct {
	grpc.ClientStream
}

func (x *upstreamRemoveClient) Send(m *Path) error {
	return x.ClientStream.SendMsg(m)
}

func (x *upstreamRemoveClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpstreamServer is the server API for Upstream service.
type UpstreamServer interface {
	Upload(Upstream_UploadServer) error
	Remove(Upstream_RemoveServer) error
}

func RegisterUpstreamServer(s *grpc.Server, srv UpstreamServer) {
	s.RegisterService(&_Upstream_serviceDesc, srv)
}

func _Upstream_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpstreamServer).Upload(&upstreamUploadServer{stream})
}

type Upstream_UploadServer interface {
	SendAndClose(*Empty) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type upstreamUploadServer struct {
	grpc.ServerStream
}

func (x *upstreamUploadServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *upstreamUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Upstream_Remove_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpstreamServer).Remove(&upstreamRemoveServer{stream})
}

type Upstream_RemoveServer interface {
	SendAndClose(*Empty) error
	Recv() (*Path, error)
	grpc.ServerStream
}

type upstreamRemoveServer struct {
	grpc.ServerStream
}

func (x *upstreamRemoveServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *upstreamRemoveServer) Recv() (*Path, error) {
	m := new(Path)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Upstream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.Upstream",
	HandlerType: (*UpstreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Upstream_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Remove",
			Handler:       _Upstream_Remove_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "remote.proto",
}
