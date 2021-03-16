package errcode

import (
	"go-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(&proto.Error{Code: int32(err.Code()), Message: err.Msg()})
	return s.Err()
}

func ToRpcStatus(code int, msg string) *Status {
	details, _ := status.New(ToRPCCode(code), msg).WithDetails(&proto.Error{Code: int32(code), Message: msg})
	return &Status{details}
}

func ToRPCCode(code int) codes.Code {
	var statuscode codes.Code
	switch code {
	case Fail.code:
		statuscode = codes.Internal
	case InvalidParams.code:
		statuscode = codes.InvalidArgument
	case Unauthorized.code:
		statuscode = codes.Unauthenticated
	case AccessDenied.code:
		statuscode = codes.PermissionDenied
	case DeadlineExceeded.code:
		statuscode = codes.DeadlineExceeded
	case NotFound.code:
		statuscode = codes.NotFound
	case LimitExceed.code:
		statuscode = codes.ResourceExhausted
	case MethodNotAllowed.code:
		statuscode = codes.Unimplemented
	default:
		statuscode = codes.Unknown
	}

	return statuscode
}

type Status struct {
	*status.Status
}

func FormError(err error) *Status {
	fromError, _ := status.FromError(err)
	return &Status{fromError}
}
