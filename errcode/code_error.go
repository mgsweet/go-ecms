package errcode

import "fmt"

// CodeError error with error code, and customize ErrorMessage
type CodeError struct {
	Code    int32
	Message string
}

// NewCodeError create a new CodeError
func NewCodeError(code int32) *CodeError {
	return &CodeError{
		Code:    code,
		Message: CodeDefaultDesc(code),
	}
}

// NewCodeErrorWithMessage create a new CodeError with message
func NewCodeErrorWithMessage(code int32, message string) *CodeError {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

// Error implement error interface
func (err *CodeError) Error() string {
	return fmt.Sprintf("error code: %v, message: %v", err.Code, err.Message)
}
