package apperr

import (
	"fmt"
)

func Wrap(err error, logMsg string) *WrappedError {
	if err == nil {
		return nil
	}
	return &WrappedError{
		err:        err,
		logMsg:     logMsg,
		stackTrace: getCallerStackTrace(),
	}
}

func Wrapf(err error, logMsgFormat string, args ...any) *WrappedError {
	if err == nil {
		return nil
	}
	return &WrappedError{
		err:        err,
		logMsg:     fmt.Sprintf(logMsgFormat, args...),
		stackTrace: getCallerStackTrace(),
	}
}

type WrappedError struct {
	err error

	logMsg     string
	stackTrace string

	extraInfoMap map[string]any
}

func (e *WrappedError) With(extraInfos ...ExtraInfoSetter) *WrappedError {
	for _, extraInfoSetter := range extraInfos {
		extraInfoSetter.setExtraInfo(e)
	}
	return e
}

func (e *WrappedError) Error() string {
	return e.logMsg + ": " + e.err.Error()
}

func (e *WrappedError) StackTrace() string {
	return e.logMsg + "\n\t" + e.stackTrace
}

func (e *WrappedError) Unwrap() error {
	return e.err
}
