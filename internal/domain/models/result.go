package models

type Result struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func newSuccessResult(code int, data any, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func newWarningResult(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func newErrorResult(code int, errMessage string) *Result {
	return &Result{
		Code:    code,
		Message: errMessage,
		Data:    nil,
	}
}

func NewOkResult(data any) *Result {
	return newSuccessResult(
		200,
		data,
		"The operation was performed successfully.",
	)
}

func NewCreatedResult(data any) *Result {
	return newSuccessResult(
		201,
		data,
		"The registration was completed successfully.",
	)
}

func NewNoCotentResult() *Result {
	return newSuccessResult(
		204,
		nil,
		"",
	)
}

func NewBadRequestResult() *Result {
	return newWarningResult(
		400,
		"The request does not meet the expected format. Please check the data and try again.",
	)
}

func NewNotFoundResult() *Result {
	return newWarningResult(
		404,
		"The target was not found.",
	)
}

func NewConflictResult() *Result {
	return newWarningResult(
		409,
		"The target already exists.",
	)
}

func NewInternalServerErrorResult() *Result {
	return newErrorResult(
		500,
		"Sorry, something went wrong in our system. Please try again later.",
	)
}

func NewGatewayTimeoutErrorResult() *Result {
	return newErrorResult(
		504,
		"The server took too long to respond. Please try again later.",
	)
}
