package apperr

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
)

const (
	CodeUndefined = 0

	defaultHttpStatusCode = http.StatusInternalServerError
	defaultCode           = CodeUndefined
	defaultPublicMsg      = "Internal Server Error"
)

func New(text, logMsg string, opts ...Option) *AppError {
	appErr := &AppError{
		err:        errors.New(text),
		logMsg:     logMsg,
		stackTrace: getCallerStackTraceWithLog(logMsg),
	}
	return appErr.With(opts...)
}

func Wrap(err error, logMsg string) *AppError {
	if err == nil {
		return nil
	}
	return &AppError{
		err:        err,
		logMsg:     logMsg,
		stackTrace: getCallerStackTraceWithLog(logMsg),
	}
}

func Wrapf(err error, logMsgFormat string, args ...any) *AppError {
	if err == nil {
		return nil
	}
	logMsg := fmt.Sprintf(logMsgFormat, args...)
	stackTrace := getCallerStackTraceWithLog(logMsg)
	return &AppError{
		err:        err,
		logMsg:     logMsg,
		stackTrace: stackTrace,
	}
}

func getCallerStackTraceWithLog(logMsg string) string {
	stack := strings.Split(string(debug.Stack()), "\n")
	callerMethod := stack[7]
	callerTraceRaw := stack[8]
	callerTrace := strings.Split(strings.TrimSpace(callerTraceRaw), " ")[0]
	return fmt.Sprintf("%s\n\t%s\n\t%s", callerMethod, logMsg, callerTrace)
}

func RootError(err error) error {
	appErr, ok := err.(*AppError)
	if !ok {
		return err
	}
	return RootError(appErr.err)
}

type AppError struct {
	err error

	logMsg     string
	stackTrace string

	httpStatusCode *int
	code           *int
	publicMsg      *string
}

func (ae *AppError) With(opts ...Option) *AppError {
	if ae == nil {
		return nil
	}
	for _, opt := range opts {
		opt.apply(ae)
	}
	return ae
}

func (ae *AppError) StackTrace() string {
	stack := ""
	if nestedAppErr, ok := ae.err.(*AppError); ok {
		stack = nestedAppErr.StackTrace() + "\n"
	}
	stack += ae.stackTrace
	return stack
}

func (ae *AppError) Error() string {
	return ae.logMsg + ": " + ae.err.Error()
}

// support go1.13 error functions: Is & As
func (ae *AppError) Unwrap() error {
	return ae.err
}

func (ae *AppError) HTTPStatusCode() int {
	if ae.httpStatusCode != nil {
		return *ae.httpStatusCode
	}
	nestedAppErr, ok := ae.err.(*AppError)
	if !ok {
		return defaultHttpStatusCode
	}
	return nestedAppErr.HTTPStatusCode()
}

func (ae *AppError) Code() int {
	if ae.code != nil {
		return *ae.code
	}
	nestedAppErr, ok := ae.err.(*AppError)
	if !ok {
		return defaultCode
	}
	return nestedAppErr.Code()
}

func (ae *AppError) PublicMessage() string {
	if ae.publicMsg != nil {
		return *ae.publicMsg
	}
	nestedAppErr, ok := ae.err.(*AppError)
	if !ok {
		return defaultPublicMsg
	}
	return nestedAppErr.PublicMessage()
}
