package apperr

import "errors"

func OutmostInfo(err error, key string) (info any) {
	if err == nil {
		return nil
	}
	infoGetter, ok := err.(infoGettable)
	if !ok {
		return nil
	}
	info, ok = infoGetter.info(key)
	if !ok {
		return OutmostInfo(errors.Unwrap(err), key)
	}
	return info
}

type infoGettable interface {
	info(key string) (val any, ok bool)
}

func GroupInfos(infos ...InfoSetter) InfoSetter {
	return func(i infoSettable) {
		for _, setInfo := range infos {
			setInfo(i)
		}
	}
}

func Info(key string, val any) InfoSetter {
	return func(i infoSettable) {
		i.setInfo(key, val)
	}
}

type InfoSetter func(infoSettable)

type infoSettable interface {
	setInfo(key string, val any)
}

var _ infoSettable = (*infoSetGetter)(nil)
var _ infoGettable = (*infoSetGetter)(nil)

type infoSetGetter struct {
	infoMap map[string]any
}

func (i *infoSetGetter) setInfo(key string, val any) {
	if i.infoMap == nil {
		i.infoMap = make(map[string]any)
	}
	i.infoMap[key] = val
}

func (i *infoSetGetter) info(key string) (any, bool) {
	val, ok := i.infoMap[key]
	return val, ok
}
