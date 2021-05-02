package errors

import "fmt"

type errorCode string

const (
	IllegalArgument errorCode = "IllegalArgument"
	Unauthorized    errorCode = "Unauthorized"
	Forbidden       errorCode = "Forbidden"
	NotFound        errorCode = "NotFound"
	Conflict        errorCode = "Conflict"
	GeneralError    errorCode = "GeneralError"
)

// New returns an error that formats as the given text and aligns with builtin error
func New(status errorCode, message string) error {
	return &ServiceError{status, message}
}

type ServiceError struct {
	Status  errorCode
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%v: %v", e.Status, e.Message)
}
