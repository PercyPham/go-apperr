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
