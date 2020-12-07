package grpc

import (
	"fmt"
)

type errorBadGateWay struct {
	err error
}

func (e *errorBadGateWay) Error() string {
	return fmt.Sprintf("Bad gateway error: %s", e.err.Error())
}

type errorServerError struct {
	err error
}

func (e *errorServerError) Error() string {
	return fmt.Sprintf("Internal server error: %s", e.err.Error())
}
