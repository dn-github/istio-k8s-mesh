// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reviews.proto

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
	return fileDescriptor_5b2a6bc9211d2c89, []int{0}
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

type Review struct {
	Review               string   `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	Rating               int64    `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2a6bc9211d2c89, []int{1}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *Review) GetRating() int64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*Book)(nil), "pb.Book")
	proto.RegisterType((*Review)(nil), "pb.Review")
}

func init() { proto.RegisterFile("reviews.proto", fileDescriptor_5b2a6bc9211d2c89) }

var fileDescriptor_5b2a6bc9211d2c89 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0xcb,
	0x4c, 0x2d, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe2,
	0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x2c, 0xb8, 0xd8, 0x82, 0xc0, 0x1a, 0x84, 0xc4, 0xb8,
	0xd8, 0x20, 0x5a, 0xa1, 0xf2, 0x50, 0x1e, 0x58, 0x3c, 0xb1, 0x24, 0x33, 0x2f, 0x5d, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x39, 0x08, 0xca, 0x33, 0x32, 0xe2, 0xe2, 0x85, 0xe8, 0x0c, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0x52, 0xe4, 0x62, 0x87, 0xda, 0x2d, 0xc4, 0xa1, 0x57, 0x90, 0xa4, 0x07,
	0xb2, 0x53, 0x8a, 0x0b, 0xc4, 0x82, 0xa8, 0x53, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xca, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x20, 0xa8, 0x23, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	Reviews(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Review, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Reviews(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/pb.ReviewService/reviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	Reviews(context.Context, *Book) (*Review, error)
}

// UnimplementedReviewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (*UnimplementedReviewServiceServer) Reviews(ctx context.Context, req *Book) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reviews not implemented")
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_Reviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Reviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReviewService/Reviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Reviews(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reviews",
			Handler:    _ReviewService_Reviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}