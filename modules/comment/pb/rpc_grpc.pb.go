// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: modules/comment/pb/rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
	ListComment(ctx context.Context, in *ListCommentRequest, opts ...grpc.CallOption) (*ListCommentResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	DeleteCommentByVideoID(ctx context.Context, in *DeleteCommentByVideoIDRequest, opts ...grpc.CallOption) (*DeleteCommentByVideoIDResponse, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) ListComment(ctx context.Context, in *ListCommentRequest, opts ...grpc.CallOption) (*ListCommentResponse, error) {
	out := new(ListCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/ListComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DeleteCommentByVideoID(ctx context.Context, in *DeleteCommentByVideoIDRequest, opts ...grpc.CallOption) (*DeleteCommentByVideoIDResponse, error) {
	out := new(DeleteCommentByVideoIDResponse)
	err := c.cc.Invoke(ctx, "/comment.pb.Comment/DeleteCommentByVideoID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
	ListComment(context.Context, *ListCommentRequest) (*ListCommentResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	DeleteCommentByVideoID(context.Context, *DeleteCommentByVideoIDRequest) (*DeleteCommentByVideoIDResponse, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedCommentServer) ListComment(context.Context, *ListCommentRequest) (*ListCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedCommentServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServer) DeleteCommentByVideoID(context.Context, *DeleteCommentByVideoIDRequest) (*DeleteCommentByVideoIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommentByVideoID not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/ListComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).ListComment(ctx, req.(*ListCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DeleteCommentByVideoID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentByVideoIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DeleteCommentByVideoID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.pb.Comment/DeleteCommentByVideoID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DeleteCommentByVideoID(ctx, req.(*DeleteCommentByVideoIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.pb.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _Comment_Healthz_Handler,
		},
		{
			MethodName: "ListComment",
			Handler:    _Comment_ListComment_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Comment_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Comment_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Comment_DeleteComment_Handler,
		},
		{
			MethodName: "DeleteCommentByVideoID",
			Handler:    _Comment_DeleteCommentByVideoID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/comment/pb/rpc.proto",
}
