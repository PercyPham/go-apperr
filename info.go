package apperr

import "errors"

func OutmostInfo(err error, key string) (info any) {
	if err == nil {
		return nil
	}
	wrappedErr, ok := err.(*WrappedError)
	if !ok {
		return nil
	}
	info, ok = wrappedErr.infoMap[key]
	if !ok {
		return OutmostInfo(errors.Unwrap(err), key)
	}
	return info
}

func GroupInfos(infos ...InfoSetter) InfoSetter {
	return infoSetterGroup(infos)
}

type infoSetterGroup []InfoSetter

func (g infoSetterGroup) setInfo(e *WrappedError) {
	for _, setter := range g {
		setter.setInfo(e)
	}
}

func Info(key string, val any) InfoSetter {
	return &infoSetter{key: key, val: val}
}

type infoSetter struct {
	key string
	val any
}

func (s *infoSetter) setInfo(e *WrappedError) {
	if e.infoMap == nil {
		e.infoMap = make(map[string]any)
	}
	e.infoMap[s.key] = s.val
}

type InfoSetter interface {
	setInfo(*WrappedError)
}
