// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: payment/v1/payment.proto

package v1

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

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	ClientToken(ctx context.Context, in *ClientTokenRequest, opts ...grpc.CallOption) (*ClientTokenResponse, error)
	CalculateTransactionFees(ctx context.Context, in *CalculateTransactionFeesRequest, opts ...grpc.CallOption) (*CalculateTransactionFeesResponse, error)
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
	RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error)
	CapturePayment(ctx context.Context, in *CapturePaymentRequest, opts ...grpc.CallOption) (*CapturePaymentResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ClientToken(ctx context.Context, in *ClientTokenRequest, opts ...grpc.CallOption) (*ClientTokenResponse, error) {
	out := new(ClientTokenResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/ClientToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CalculateTransactionFees(ctx context.Context, in *CalculateTransactionFeesRequest, opts ...grpc.CallOption) (*CalculateTransactionFeesResponse, error) {
	out := new(CalculateTransactionFeesResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/CalculateTransactionFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	out := new(ProcessPaymentResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/ProcessPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error) {
	out := new(RefundPaymentResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/RefundPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CapturePayment(ctx context.Context, in *CapturePaymentRequest, opts ...grpc.CallOption) (*CapturePaymentResponse, error) {
	out := new(CapturePaymentResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/CapturePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/payment.v1.PaymentService/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	ClientToken(context.Context, *ClientTokenRequest) (*ClientTokenResponse, error)
	CalculateTransactionFees(context.Context, *CalculateTransactionFeesRequest) (*CalculateTransactionFeesResponse, error)
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error)
	CapturePayment(context.Context, *CapturePaymentRequest) (*CapturePaymentResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) ClientToken(context.Context, *ClientTokenRequest) (*ClientTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientToken not implemented")
}
func (UnimplementedPaymentServiceServer) CalculateTransactionFees(context.Context, *CalculateTransactionFeesRequest) (*CalculateTransactionFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateTransactionFees not implemented")
}
func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServiceServer) RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundPayment not implemented")
}
func (UnimplementedPaymentServiceServer) CapturePayment(context.Context, *CapturePaymentRequest) (*CapturePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CapturePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_ClientToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ClientToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/ClientToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ClientToken(ctx, req.(*ClientTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CalculateTransactionFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateTransactionFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CalculateTransactionFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/CalculateTransactionFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CalculateTransactionFees(ctx, req.(*CalculateTransactionFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/ProcessPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_RefundPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).RefundPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/RefundPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).RefundPayment(ctx, req.(*RefundPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CapturePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapturePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CapturePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/CapturePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CapturePayment(ctx, req.(*CapturePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.v1.PaymentService/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.v1.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientToken",
			Handler:    _PaymentService_ClientToken_Handler,
		},
		{
			MethodName: "CalculateTransactionFees",
			Handler:    _PaymentService_CalculateTransactionFees_Handler,
		},
		{
			MethodName: "ProcessPayment",
			Handler:    _PaymentService_ProcessPayment_Handler,
		},
		{
			MethodName: "RefundPayment",
			Handler:    _PaymentService_RefundPayment_Handler,
		},
		{
			MethodName: "CapturePayment",
			Handler:    _PaymentService_CapturePayment_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _PaymentService_GetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/v1/payment.proto",
}
