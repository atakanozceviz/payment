package server

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validatable interface {
	ValidateAll() error
	Validate() error
}

type validationError interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

type multiError interface {
	AllErrors() []error
}

func validate(req interface{}) error {
	switch v := req.(type) {
	case validatable:
		if err := v.ValidateAll(); err != nil {
			multiError, ok := err.(multiError)
			if !ok {
				return status.Error(codes.InvalidArgument, err.Error())
			}

			br := &errdetails.BadRequest{FieldViolations: make([]*errdetails.BadRequest_FieldViolation, 0)}
			for _, e := range multiError.AllErrors() {
				ve, ok := e.(validationError)
				if !ok {
					return status.Error(codes.InvalidArgument, err.Error())
				}
				br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       ve.Field(),
					Description: ve.Reason(),
				})
			}
			s, err := status.New(codes.InvalidArgument, "validation failed").WithDetails(br)
			if err != nil {
				return err
			}
			return s.Err()
		}
	}
	return nil
}

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validate(req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// UnaryClientInterceptor returns a new unary client interceptor that validates outgoing messages.
//
// Invalid messages will be rejected with `InvalidArgument` before sending the request to server.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if err := validate(req); err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that validates incoming messages.
//
// The stage at which invalid messages will be rejected with `InvalidArgument` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{stream}
		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}

	if err := validate(m); err != nil {
		return err
	}

	return nil
}
