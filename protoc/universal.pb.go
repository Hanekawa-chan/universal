// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: universal.proto

package protoc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_universal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_universal_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_universal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_universal_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender   string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_universal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_universal_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

var File_universal_proto protoreflect.FileDescriptor

var file_universal_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x22, 0x36, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x26, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x32, 0xfb, 0x01, 0x0a,
	0x09, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_universal_proto_rawDescOnce sync.Once
	file_universal_proto_rawDescData = file_universal_proto_rawDesc
)

func file_universal_proto_rawDescGZIP() []byte {
	file_universal_proto_rawDescOnce.Do(func() {
		file_universal_proto_rawDescData = protoimpl.X.CompressGZIP(file_universal_proto_rawDescData)
	})
	return file_universal_proto_rawDescData
}

var file_universal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_universal_proto_goTypes = []interface{}{
	(*User)(nil),     // 0: protoc.User
	(*Response)(nil), // 1: protoc.Response
	(*Message)(nil),  // 2: protoc.Message
}
var file_universal_proto_depIdxs = []int32{
	0, // 0: protoc.Universal.signUp:input_type -> protoc.User
	0, // 1: protoc.Universal.signIn:input_type -> protoc.User
	0, // 2: protoc.Universal.signOut:input_type -> protoc.User
	2, // 3: protoc.Universal.sendMessage:input_type -> protoc.Message
	0, // 4: protoc.Universal.receiveMessage:input_type -> protoc.User
	1, // 5: protoc.Universal.signUp:output_type -> protoc.Response
	1, // 6: protoc.Universal.signIn:output_type -> protoc.Response
	1, // 7: protoc.Universal.signOut:output_type -> protoc.Response
	1, // 8: protoc.Universal.sendMessage:output_type -> protoc.Response
	2, // 9: protoc.Universal.receiveMessage:output_type -> protoc.Message
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_universal_proto_init() }
func file_universal_proto_init() {
	if File_universal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_universal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_universal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_universal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_universal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_universal_proto_goTypes,
		DependencyIndexes: file_universal_proto_depIdxs,
		MessageInfos:      file_universal_proto_msgTypes,
	}.Build()
	File_universal_proto = out.File
	file_universal_proto_rawDesc = nil
	file_universal_proto_goTypes = nil
	file_universal_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UniversalClient is the client API for Universal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UniversalClient interface {
	SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	SignIn(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	SignOut(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Universal_SendMessageClient, error)
	ReceiveMessage(ctx context.Context, in *User, opts ...grpc.CallOption) (Universal_ReceiveMessageClient, error)
}

type universalClient struct {
	cc grpc.ClientConnInterface
}

func NewUniversalClient(cc grpc.ClientConnInterface) UniversalClient {
	return &universalClient{cc}
}

func (c *universalClient) SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universalClient) SignIn(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universalClient) SignOut(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protoc.Universal/signOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universalClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Universal_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Universal_serviceDesc.Streams[0], "/protoc.Universal/sendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &universalSendMessageClient{stream}
	return x, nil
}

type Universal_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type universalSendMessageClient struct {
	grpc.ClientStream
}

func (x *universalSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *universalSendMessageClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *universalClient) ReceiveMessage(ctx context.Context, in *User, opts ...grpc.CallOption) (Universal_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Universal_serviceDesc.Streams[1], "/protoc.Universal/receiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &universalReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Universal_ReceiveMessageClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type universalReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *universalReceiveMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UniversalServer is the server API for Universal service.
type UniversalServer interface {
	SignUp(context.Context, *User) (*Response, error)
	SignIn(context.Context, *User) (*Response, error)
	SignOut(context.Context, *User) (*Response, error)
	SendMessage(Universal_SendMessageServer) error
	ReceiveMessage(*User, Universal_ReceiveMessageServer) error
}

// UnimplementedUniversalServer can be embedded to have forward compatible implementations.
type UnimplementedUniversalServer struct {
}

func (*UnimplementedUniversalServer) SignUp(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedUniversalServer) SignIn(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedUniversalServer) SignOut(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedUniversalServer) SendMessage(Universal_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedUniversalServer) ReceiveMessage(*User, Universal_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}

func RegisterUniversalServer(s *grpc.Server, srv UniversalServer) {
	s.RegisterService(&_Universal_serviceDesc, srv)
}

func _Universal_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversalServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoc.Universal/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversalServer).SignUp(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universal_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversalServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoc.Universal/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversalServer).SignIn(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universal_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversalServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoc.Universal/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversalServer).SignOut(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Universal_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UniversalServer).SendMessage(&universalSendMessageServer{stream})
}

type Universal_SendMessageServer interface {
	SendAndClose(*Response) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type universalSendMessageServer struct {
	grpc.ServerStream
}

func (x *universalSendMessageServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *universalSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Universal_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UniversalServer).ReceiveMessage(m, &universalReceiveMessageServer{stream})
}

type Universal_ReceiveMessageServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type universalReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *universalReceiveMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Universal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoc.Universal",
	HandlerType: (*UniversalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "signUp",
			Handler:    _Universal_SignUp_Handler,
		},
		{
			MethodName: "signIn",
			Handler:    _Universal_SignIn_Handler,
		},
		{
			MethodName: "signOut",
			Handler:    _Universal_SignOut_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendMessage",
			Handler:       _Universal_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "receiveMessage",
			Handler:       _Universal_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "universal.proto",
}
