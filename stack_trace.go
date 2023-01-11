package apperr

import "errors"

func StackTrace(err error) string {
	if err == nil {
		return ""
	}
	if err == errOperational {
		return ""
	}

	unwrappedErr := errors.Unwrap(err)

	stackTrace := StackTrace(unwrappedErr)
	if stackTrace != "" {
		stackTrace += "\n"
	}

	traceableErr, ok := err.(interface {
		StackTrace() string
	})
	if !ok {
		return stackTrace + err.Error()
	}

	return stackTrace + traceableErr.StackTrace()
}
