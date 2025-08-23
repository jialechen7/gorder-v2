package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DomainError interface for typed errors
type DomainError interface {
	error
	GRPCCode() codes.Code
}

// ErrorInterceptor converts domain errors to gRPC status errors
func ErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			// Check if it's a domain error
			var domainErr DomainError
			if errors.As(err, &domainErr) {
				return nil, status.Error(domainErr.GRPCCode(), domainErr.Error())
			}

			// Check if it's already a gRPC status error
			if _, ok := status.FromError(err); ok {
				return nil, err
			}

			// Default to internal error for unknown errors
			return nil, status.Error(codes.Internal, err.Error())
		}
		return resp, nil
	}
}
