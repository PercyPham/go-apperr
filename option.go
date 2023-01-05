package apperr

import "fmt"

type Option interface {
	apply(*AppError)
}

func GroupOptions(opts ...Option) Option {
	return optionGroup(opts)
}

type optionGroup []Option

func (g optionGroup) apply(ae *AppError) {
	for _, opt := range g {
		opt.apply(ae)
	}
}

func HTTPStatusCode(code int) Option {
	return httpStatusCodeOption(code)
}

type httpStatusCodeOption int

func (h httpStatusCodeOption) apply(ae *AppError) {
	httpStatusCode := int(h)
	ae.httpStatusCode = &httpStatusCode
}

func Code(code int) Option {
	return codeOption(code)
}

type codeOption int

func (o codeOption) apply(ae *AppError) {
	code := int(o)
	ae.code = &code
}

func PublicMessage(text string, args ...any) Option {
	return publicMsgOption(text)
}

func PublicMessagef(format string, args ...any) Option {
	publicMsg := fmt.Sprintf(format, args...)
	return publicMsgOption(publicMsg)
}

type publicMsgOption string

func (p publicMsgOption) apply(ae *AppError) {
	publicMsg := string(p)
	ae.publicMsg = &publicMsg
}
