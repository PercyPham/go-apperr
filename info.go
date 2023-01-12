package apperr

import "errors"

func OutmostInfo(err error, key string) (info any) {
	if err == nil {
		return nil
	}
	infoGettable, ok := err.(interface {
		info(key string) (val any, ok bool)
	})
	if !ok {
		return nil
	}
	info, ok = infoGettable.info(key)
	if !ok {
		return OutmostInfo(errors.Unwrap(err), key)
	}
	return info
}

func GroupInfos(infos ...InfoSetter) InfoSetter {
	return infoSetterGroup(infos)
}

type infoSetterGroup []InfoSetter

func (g infoSetterGroup) setInfo(e infoSettable) {
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

func (s *infoSetter) setInfo(e infoSettable) {
	e.setInfo(s.key, s.val)
}

type InfoSetter interface {
	setInfo(infoSettable)
}

type infoSettable interface {
	setInfo(key string, val any)
}
