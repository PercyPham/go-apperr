package apperr

import "fmt"

func NewProgrammerError(text string) *programmerError {
	return &programmerError{
		text:       text,
		stackTrace: getCallerStackTrace(),
	}
}

type programmerError struct {
	text       string
	stackTrace string

	infoSetGetter
}

func (e *programmerError) With(infos ...InfoSetter) *programmerError {
	for _, setInfo := range infos {
		setInfo(e)
	}
	return e
}

func (e *programmerError) Error() string {
	return e.text
}

func (e *programmerError) StackTrace() string {
	if len(e.infoMap) > 0 {
		return fmt.Sprintf("%s | %s\n\t%s", e.text, serializeInfoMap(e.infoMap), e.stackTrace)
	}
	return fmt.Sprintf("%s\n\t%s", e.text, e.stackTrace)
}
