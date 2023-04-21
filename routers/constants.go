package routers

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var code2Http = map[codes.Code]int{
	codes.Aborted:            http.StatusExpectationFailed,
	codes.AlreadyExists:      http.StatusConflict,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.DataLoss:           http.StatusExpectationFailed,
	codes.Unknown:            http.StatusBadRequest,
	codes.InvalidArgument:    http.StatusFailedDependency,
	codes.NotFound:           http.StatusNotFound,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.Internal:           http.StatusInternalServerError,
	codes.ResourceExhausted:  http.StatusExpectationFailed,
	codes.OK:                 http.StatusOK,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.Unimplemented:      http.StatusNotImplemented,
}
