package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleErrorResponse(err error) error {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.Internal:
			return ErrorServerError
		default:
			return ErrorBadGateway
		}
	}

	return ErrorBadGateway
}
