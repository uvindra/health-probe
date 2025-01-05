package common

type ServiceError struct {
	message    string
	httpStatus int
}

func NewServiceError(message string, httpStatus int) *ServiceError {
	return &ServiceError{message: message, httpStatus: httpStatus}
}

func (e *ServiceError) Error() string {
	return e.message
}

func (e *ServiceError) HttpStatus() int {
	return e.httpStatus
}
