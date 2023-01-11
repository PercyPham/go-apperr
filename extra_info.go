package apperr

import "errors"

func OutmostExtraInfo(err error, key string) (extraInfo any) {
	if err == nil {
		return nil
	}
	wrappedErr, ok := err.(*WrappedError)
	if !ok {
		return nil
	}
	extraInfo, ok = wrappedErr.extraInfoMap[key]
	if !ok {
		return OutmostExtraInfo(errors.Unwrap(err), key)
	}
	return extraInfo
}

func GroupExtraInfos(extraInfos ...ExtraInfoSetter) ExtraInfoSetter {
	return extraInfoSetterGroup(extraInfos)
}

type extraInfoSetterGroup []ExtraInfoSetter

func (g extraInfoSetterGroup) setExtraInfo(e *WrappedError) {
	for _, setter := range g {
		setter.setExtraInfo(e)
	}
}

func ExtraInfo(key string, val any) ExtraInfoSetter {
	return &extraInfoSetter{key: key, val: val}
}

type extraInfoSetter struct {
	key string
	val any
}

func (s *extraInfoSetter) setExtraInfo(e *WrappedError) {
	if e.extraInfoMap == nil {
		e.extraInfoMap = make(map[string]any)
	}
	e.extraInfoMap[s.key] = s.val
}

type ExtraInfoSetter interface {
	setExtraInfo(*WrappedError)
}
