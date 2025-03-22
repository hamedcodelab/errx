package errx

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcErrorInterceptor() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			var errorModel ErrXModel
			if errors.As(err, &errorModel) {
				switch errorModel.GetType() {
				case ErrorValidation:
					return nil, status.Error(codes.InvalidArgument, errorModel.GetMessage())
				case ErrorForbidden:
					return nil, status.Error(codes.PermissionDenied, errorModel.GetMessage())
				case ErrorUnauthorized:
					return nil, status.Error(codes.Unauthenticated, errorModel.GetMessage())
				case ErrorNotFound:
					return nil, status.Error(codes.NotFound, errorModel.GetMessage())
				case ErrorTooManyRequests:
					return nil, status.Error(codes.ResourceExhausted, errorModel.GetMessage())
				case ErrorConflict:
					return nil, status.Error(codes.AlreadyExists, errorModel.GetMessage())
				default:
					//fmt.Println(errorModel.GetStack())
					return nil, status.Error(codes.Internal, InternalServerError(err).GetMessage())
				}
			} else {
				switch ConvertGRPCToErrorType(err) {
				case ErrorValidation:
					return nil, status.Error(codes.InvalidArgument, err.Error())
				case ErrorForbidden:
					return nil, status.Error(codes.PermissionDenied, err.Error())
				case ErrorUnauthorized:
					return nil, status.Error(codes.Unauthenticated, err.Error())
				case ErrorNotFound:
					return nil, status.Error(codes.NotFound, err.Error())
				case ErrorTooManyRequests:
					return nil, status.Error(codes.ResourceExhausted, err.Error())
				case ErrorConflict:
					return nil, status.Error(codes.AlreadyExists, err.Error())
				default:
					return nil, status.Error(codes.Internal, InternalServerError(err).GetMessage())
				}
			}
		}

		return resp, nil
	})
}
