package response

import "health-probe/probe"

type ServiceResponse struct {
	message    string
	httpStatus int
	isError    bool
}

func NewSuccessResponse(message string, httpStatus int, probe *probe.BaseProbe) ServiceResponse {
	probe.IncrementSuccessCount()
	return ServiceResponse{message: message, httpStatus: httpStatus, isError: false}
}

func NewErrorResponse(message string, httpStatus int, probe *probe.BaseProbe) ServiceResponse {
	probe.IncrementErrorCount()
	return ServiceResponse{message: message, httpStatus: httpStatus, isError: true}
}

func (e *ServiceResponse) Error() string {
	return e.message
}

func (e *ServiceResponse) HttpStatus() int {
	return e.httpStatus
}

func (e *ServiceResponse) IsError() bool {
	return e.isError
}
