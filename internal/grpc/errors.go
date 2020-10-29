package grpc

import "errors"

var ErrorInvalidArgument = errors.New("invalid argument provided in request")
var ErrorBadGateway = errors.New("error response returned from external service")
var ErrorServerError = errors.New("internal server error")
var ErrorNotFound = errors.New("the resource requested does not exist")
