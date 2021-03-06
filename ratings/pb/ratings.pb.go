// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ratings.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Book struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf07e2fe2aafa921, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Rating struct {
	Rating               int64    `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rating) Reset()         { *m = Rating{} }
func (m *Rating) String() string { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()    {}
func (*Rating) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf07e2fe2aafa921, []int{1}
}

func (m *Rating) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rating.Unmarshal(m, b)
}
func (m *Rating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rating.Marshal(b, m, deterministic)
}
func (m *Rating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rating.Merge(m, src)
}
func (m *Rating) XXX_Size() int {
	return xxx_messageInfo_Rating.Size(m)
}
func (m *Rating) XXX_DiscardUnknown() {
	xxx_messageInfo_Rating.DiscardUnknown(m)
}

var xxx_messageInfo_Rating proto.InternalMessageInfo

func (m *Rating) GetRating() int64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*Book)(nil), "pb.Book")
	proto.RegisterType((*Rating)(nil), "pb.Rating")
}

func init() { proto.RegisterFile("ratings.proto", fileDescriptor_bf07e2fe2aafa921) }

var fileDescriptor_bf07e2fe2aafa921 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2c, 0xc9,
	0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe2,
	0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb8, 0xd8, 0x82, 0xc0, 0x1a, 0x84, 0xc4, 0xb8,
	0xd8, 0x20, 0x5a, 0xc1, 0xf2, 0xcc, 0x41, 0x50, 0x9e, 0x91, 0x11, 0x17, 0x2f, 0x44, 0x45, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x22, 0x17, 0x3b, 0xd4, 0x0e, 0x21, 0x0e, 0xbd, 0x82,
	0x24, 0x3d, 0x90, 0xd9, 0x52, 0x5c, 0x20, 0x16, 0x44, 0x9d, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x72,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xcd, 0x13, 0x63, 0x8d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RatingServiceClient interface {
	Ratings(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Rating, error)
}

type ratingServiceClient struct {
	cc *grpc.ClientConn
}

func NewRatingServiceClient(cc *grpc.ClientConn) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) Ratings(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Rating, error) {
	out := new(Rating)
	err := c.cc.Invoke(ctx, "/pb.RatingService/ratings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
type RatingServiceServer interface {
	Ratings(context.Context, *Book) (*Rating, error)
}

// UnimplementedRatingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (*UnimplementedRatingServiceServer) Ratings(ctx context.Context, req *Book) (*Rating, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ratings not implemented")
}

func RegisterRatingServiceServer(s *grpc.Server, srv RatingServiceServer) {
	s.RegisterService(&_RatingService_serviceDesc, srv)
}

func _RatingService_Ratings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Ratings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RatingService/Ratings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Ratings(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

var _RatingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ratings",
			Handler:    _RatingService_Ratings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratings.proto",
}
