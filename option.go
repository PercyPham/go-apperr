package apperr

import "fmt"

type option interface {
	apply(*appError)
}

func PredefineOptions(opts ...option) option {
	return predefinedOptions(opts)
}

type predefinedOptions []option

func (p predefinedOptions) apply(ae *appError) {
	for _, opt := range p {
		opt.apply(ae)
	}
}

func HTTPStatusCode(code int) option {
	return httpStatusCodeOption(code)
}

type httpStatusCodeOption int

func (h httpStatusCodeOption) apply(ae *appError) {
	httpStatusCode := int(h)
	ae.httpStatusCode = &httpStatusCode
}

func Code(code int) option {
	return codeOption(code)
}

type codeOption int

func (o codeOption) apply(ae *appError) {
	code := int(o)
	ae.code = &code
}

func PublicMessage(format string, args ...interface{}) option {
	publicMsg := fmt.Sprintf(format, args...)
	return publicMsgOption(publicMsg)
}

type publicMsgOption string

func (p publicMsgOption) apply(ae *appError) {
	publicMsg := string(p)
	ae.publicMsg = &publicMsg
}
