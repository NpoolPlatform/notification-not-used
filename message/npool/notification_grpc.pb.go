// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error)
	GetAnnouncementsByApp(ctx context.Context, in *GetAnnouncementsByAppRequest, opts ...grpc.CallOption) (*GetAnnouncementsByAppResponse, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*UpdateNotificationResponse, error)
	GetNotificationsByAppUser(ctx context.Context, in *GetNotificationsByAppUserRequest, opts ...grpc.CallOption) (*GetNotificationsByAppUserResponse, error)
	CreateReadUser(ctx context.Context, in *CreateReadUserRequest, opts ...grpc.CallOption) (*CreateReadUserResponse, error)
	CheckReadUser(ctx context.Context, in *CheckReadUserRequest, opts ...grpc.CallOption) (*CheckReadUserResponse, error)
	CreateMail(ctx context.Context, in *CreateMailRequest, opts ...grpc.CallOption) (*CreateMailResponse, error)
	UpdateMail(ctx context.Context, in *UpdateMailRequest, opts ...grpc.CallOption) (*UpdateMailResponse, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/CreateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error) {
	out := new(UpdateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/UpdateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) GetAnnouncementsByApp(ctx context.Context, in *GetAnnouncementsByAppRequest, opts ...grpc.CallOption) (*GetAnnouncementsByAppResponse, error) {
	out := new(GetAnnouncementsByAppResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/GetAnnouncementsByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/CreateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*UpdateNotificationResponse, error) {
	out := new(UpdateNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/UpdateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) GetNotificationsByAppUser(ctx context.Context, in *GetNotificationsByAppUserRequest, opts ...grpc.CallOption) (*GetNotificationsByAppUserResponse, error) {
	out := new(GetNotificationsByAppUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/GetNotificationsByAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CreateReadUser(ctx context.Context, in *CreateReadUserRequest, opts ...grpc.CallOption) (*CreateReadUserResponse, error) {
	out := new(CreateReadUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/CreateReadUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CheckReadUser(ctx context.Context, in *CheckReadUserRequest, opts ...grpc.CallOption) (*CheckReadUserResponse, error) {
	out := new(CheckReadUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/CheckReadUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CreateMail(ctx context.Context, in *CreateMailRequest, opts ...grpc.CallOption) (*CreateMailResponse, error) {
	out := new(CreateMailResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/CreateMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) UpdateMail(ctx context.Context, in *UpdateMailRequest, opts ...grpc.CallOption) (*UpdateMailResponse, error) {
	out := new(UpdateMailResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.Notification/UpdateMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
// All implementations must embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error)
	GetAnnouncementsByApp(context.Context, *GetAnnouncementsByAppRequest) (*GetAnnouncementsByAppResponse, error)
	CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	UpdateNotification(context.Context, *UpdateNotificationRequest) (*UpdateNotificationResponse, error)
	GetNotificationsByAppUser(context.Context, *GetNotificationsByAppUserRequest) (*GetNotificationsByAppUserResponse, error)
	CreateReadUser(context.Context, *CreateReadUserRequest) (*CreateReadUserResponse, error)
	CheckReadUser(context.Context, *CheckReadUserRequest) (*CheckReadUserResponse, error)
	CreateMail(context.Context, *CreateMailRequest) (*CreateMailResponse, error)
	UpdateMail(context.Context, *UpdateMailRequest) (*UpdateMailResponse, error)
	mustEmbedUnimplementedNotificationServer()
}

// UnimplementedNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedNotificationServer) CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedNotificationServer) UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnnouncement not implemented")
}
func (UnimplementedNotificationServer) GetAnnouncementsByApp(context.Context, *GetAnnouncementsByAppRequest) (*GetAnnouncementsByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncementsByApp not implemented")
}
func (UnimplementedNotificationServer) CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationServer) UpdateNotification(context.Context, *UpdateNotificationRequest) (*UpdateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotification not implemented")
}
func (UnimplementedNotificationServer) GetNotificationsByAppUser(context.Context, *GetNotificationsByAppUserRequest) (*GetNotificationsByAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationsByAppUser not implemented")
}
func (UnimplementedNotificationServer) CreateReadUser(context.Context, *CreateReadUserRequest) (*CreateReadUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadUser not implemented")
}
func (UnimplementedNotificationServer) CheckReadUser(context.Context, *CheckReadUserRequest) (*CheckReadUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReadUser not implemented")
}
func (UnimplementedNotificationServer) CreateMail(context.Context, *CreateMailRequest) (*CreateMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMail not implemented")
}
func (UnimplementedNotificationServer) UpdateMail(context.Context, *UpdateMailRequest) (*UpdateMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMail not implemented")
}
func (UnimplementedNotificationServer) mustEmbedUnimplementedNotificationServer() {}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CreateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CreateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/CreateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CreateAnnouncement(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_UpdateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).UpdateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/UpdateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).UpdateAnnouncement(ctx, req.(*UpdateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_GetAnnouncementsByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncementsByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).GetAnnouncementsByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/GetAnnouncementsByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).GetAnnouncementsByApp(ctx, req.(*GetAnnouncementsByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/CreateNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_UpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).UpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/UpdateNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).UpdateNotification(ctx, req.(*UpdateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_GetNotificationsByAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsByAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).GetNotificationsByAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/GetNotificationsByAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).GetNotificationsByAppUser(ctx, req.(*GetNotificationsByAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CreateReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CreateReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/CreateReadUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CreateReadUser(ctx, req.(*CreateReadUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CheckReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReadUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CheckReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/CheckReadUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CheckReadUser(ctx, req.(*CheckReadUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CreateMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CreateMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/CreateMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CreateMail(ctx, req.(*CreateMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_UpdateMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).UpdateMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.Notification/UpdateMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).UpdateMail(ctx, req.(*UpdateMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.v1.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Notification_Version_Handler,
		},
		{
			MethodName: "CreateAnnouncement",
			Handler:    _Notification_CreateAnnouncement_Handler,
		},
		{
			MethodName: "UpdateAnnouncement",
			Handler:    _Notification_UpdateAnnouncement_Handler,
		},
		{
			MethodName: "GetAnnouncementsByApp",
			Handler:    _Notification_GetAnnouncementsByApp_Handler,
		},
		{
			MethodName: "CreateNotification",
			Handler:    _Notification_CreateNotification_Handler,
		},
		{
			MethodName: "UpdateNotification",
			Handler:    _Notification_UpdateNotification_Handler,
		},
		{
			MethodName: "GetNotificationsByAppUser",
			Handler:    _Notification_GetNotificationsByAppUser_Handler,
		},
		{
			MethodName: "CreateReadUser",
			Handler:    _Notification_CreateReadUser_Handler,
		},
		{
			MethodName: "CheckReadUser",
			Handler:    _Notification_CheckReadUser_Handler,
		},
		{
			MethodName: "CreateMail",
			Handler:    _Notification_CreateMail_Handler,
		},
		{
			MethodName: "UpdateMail",
			Handler:    _Notification_UpdateMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notification.proto",
}
