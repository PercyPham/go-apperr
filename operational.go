package apperr

import (
	"errors"
	"fmt"
)

func New(text string) *operationalError {
	return &operationalError{
		text:       text,
		stackTrace: getCallerStackTrace(),
	}
}

func IsOperational(err error) bool {
	return errors.Is(err, errOperational)
}

var errOperational = errors.New("operational")

type operationalError struct {
	text       string
	stackTrace string

	infoSetGetter
}

func (e *operationalError) With(infos ...InfoSetter) *operationalError {
	for _, setInfo := range infos {
		setInfo(e)
	}
	return e
}

func (e *operationalError) Unwrap() error {
	return errOperational
}

func (e *operationalError) Error() string {
	return e.text
}

func (e *operationalError) StackTrace() string {
	if len(e.infoMap) > 0 {
		return fmt.Sprintf("%s | %s\n\t%s", e.text, serializeInfoMap(e.infoMap), e.stackTrace)
	}
	return fmt.Sprintf("%s\n\t%s", e.text, e.stackTrace)
}
