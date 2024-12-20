// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: intravision/labour/tx.proto

package labour

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

const (
	Msg_UpdateParams_FullMethodName = "/intravision.labour.Msg/UpdateParams"
	Msg_Work_FullMethodName         = "/intravision.labour.Msg/Work"
	Msg_BeginTask_FullMethodName    = "/intravision.labour.Msg/BeginTask"
	Msg_FinishTask_FullMethodName   = "/intravision.labour.Msg/FinishTask"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	Work(ctx context.Context, in *MsgWork, opts ...grpc.CallOption) (*MsgWorkResponse, error)
	BeginTask(ctx context.Context, in *MsgBeginTask, opts ...grpc.CallOption) (*MsgBeginTaskResponse, error)
	FinishTask(ctx context.Context, in *MsgFinishTask, opts ...grpc.CallOption) (*MsgFinishTaskResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Work(ctx context.Context, in *MsgWork, opts ...grpc.CallOption) (*MsgWorkResponse, error) {
	out := new(MsgWorkResponse)
	err := c.cc.Invoke(ctx, Msg_Work_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BeginTask(ctx context.Context, in *MsgBeginTask, opts ...grpc.CallOption) (*MsgBeginTaskResponse, error) {
	out := new(MsgBeginTaskResponse)
	err := c.cc.Invoke(ctx, Msg_BeginTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FinishTask(ctx context.Context, in *MsgFinishTask, opts ...grpc.CallOption) (*MsgFinishTaskResponse, error) {
	out := new(MsgFinishTaskResponse)
	err := c.cc.Invoke(ctx, Msg_FinishTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	Work(context.Context, *MsgWork) (*MsgWorkResponse, error)
	BeginTask(context.Context, *MsgBeginTask) (*MsgBeginTaskResponse, error)
	FinishTask(context.Context, *MsgFinishTask) (*MsgFinishTaskResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) Work(context.Context, *MsgWork) (*MsgWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedMsgServer) BeginTask(context.Context, *MsgBeginTask) (*MsgBeginTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginTask not implemented")
}
func (UnimplementedMsgServer) FinishTask(context.Context, *MsgFinishTask) (*MsgFinishTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTask not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWork)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Work_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Work(ctx, req.(*MsgWork))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BeginTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBeginTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BeginTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BeginTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BeginTask(ctx, req.(*MsgBeginTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FinishTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFinishTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FinishTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FinishTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FinishTask(ctx, req.(*MsgFinishTask))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "intravision.labour.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "Work",
			Handler:    _Msg_Work_Handler,
		},
		{
			MethodName: "BeginTask",
			Handler:    _Msg_BeginTask_Handler,
		},
		{
			MethodName: "FinishTask",
			Handler:    _Msg_FinishTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intravision/labour/tx.proto",
}
