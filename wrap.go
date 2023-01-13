package apperr

import (
	"fmt"
)

func Wrap(err error, logMsg string) *wrappedError {
	if err == nil {
		return nil
	}
	return &wrappedError{
		err:        err,
		logMsg:     logMsg,
		stackTrace: getCallerStackTrace(),
	}
}

func Wrapf(err error, logMsgFormat string, args ...any) *wrappedError {
	if err == nil {
		return nil
	}
	return &wrappedError{
		err:        err,
		logMsg:     fmt.Sprintf(logMsgFormat, args...),
		stackTrace: getCallerStackTrace(),
	}
}

type wrappedError struct {
	err error

	logMsg     string
	stackTrace string

	infoSetGetter
}

func (e *wrappedError) With(infos ...InfoSetter) *wrappedError {
	for _, setInfo := range infos {
		setInfo(e)
	}
	return e
}

func (e *wrappedError) Error() string {
	return e.logMsg + ": " + e.err.Error()
}

func (e *wrappedError) StackTrace() string {
	if len(e.infoMap) > 0 {
		return fmt.Sprintf("%s | %s\n\t%s", e.logMsg, serializeInfoMap(e.infoMap), e.stackTrace)
	}
	return fmt.Sprintf("%s\n\t%s", e.logMsg, e.stackTrace)
}

func (e *wrappedError) Unwrap() error {
	return e.err
}
