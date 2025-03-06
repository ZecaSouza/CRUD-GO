package rest_err

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"` // "omitempty" remove a chave JSON se a lista estiver vazia
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code: http.StatusBadRequest,
	}
}


func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// 401 - Unauthorized
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

// 403 - Forbidden
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

// 404 - Not Found
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// 408 - Request Timeout
func NewRequestTimeoutError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "request_timeout",
		Code:    http.StatusRequestTimeout,
	}
}

// 409 - Conflict
func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "conflict",
		Code:    http.StatusConflict,
	}
}

// 422 - Unprocessable Entity
func NewUnprocessableEntityError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unprocessable_entity",
		Code:    http.StatusUnprocessableEntity,
		Causes:  causes,
	}
}

// 429 - Too Many Requests
func NewTooManyRequestsError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "too_many_requests",
		Code:    http.StatusTooManyRequests,
	}
}

// 500 - Internal Server Error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

// 501 - Not Implemented
func NewNotImplementedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_implemented",
		Code:    http.StatusNotImplemented,
	}
}

// 503 - Service Unavailable
func NewServiceUnavailableError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "service_unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

// 504 - Gateway Timeout
func NewGatewayTimeoutError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "gateway_timeout",
		Code:    http.StatusGatewayTimeout,
	}
}

func (r *RestErr) Error() string {
	return r.Message
}