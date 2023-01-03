package apperr

import (
	"fmt"
	"net/http"
)

const (
	CodeUndefined = 0

	defaultHttpStatusCode = http.StatusOK
	defaultCode           = CodeUndefined
	defaultPublicMsg      = "internal server error"
)

type AppError interface {
	error

	// support go1.13 error functions: Is & As
	Unwrap() error

	HTTPStatusCode() int
	Code() int
	PublicMessage() string
}

var _ AppError = (*appError)(nil)

type appError struct {
	err    error
	logMsg string

	httpStatusCode *int
	code           *int
	publicMsg      *string
}

func Wrap(err error, logMsgFormat string, args ...interface{}) *appError {
	if err == nil {
		return nil
	}
	return &appError{
		err:    err,
		logMsg: fmt.Sprintf(logMsgFormat, args...),
	}
}

func (ae *appError) With(opts ...option) *appError {
	if ae == nil {
		return nil
	}
	for _, opt := range opts {
		opt.apply(ae)
	}
	return ae
}

func (ae *appError) Error() string {
	return ae.logMsg + ": " + ae.err.Error()
}

func (ae *appError) Unwrap() error {
	return ae.err
}

func (ae *appError) HTTPStatusCode() int {
	if ae.httpStatusCode != nil {
		return *ae.httpStatusCode
	}
	nestedAppErr, ok := ae.err.(*appError)
	if !ok {
		return defaultHttpStatusCode
	}
	return nestedAppErr.HTTPStatusCode()
}

func (ae *appError) Code() int {
	if ae.code != nil {
		return *ae.code
	}
	nestedAppErr, ok := ae.err.(*appError)
	if !ok {
		return defaultCode
	}
	return nestedAppErr.Code()
}

func (ae *appError) PublicMessage() string {
	if ae.publicMsg != nil {
		return *ae.publicMsg
	}
	nestedAppErr, ok := ae.err.(*appError)
	if !ok {
		return defaultPublicMsg
	}
	return nestedAppErr.PublicMessage()
}
