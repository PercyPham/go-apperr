package apperr

import (
	"encoding/json"
	"runtime/debug"
	"strings"
)

func getCallerStackTrace() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	callerStackTraceRaw := stack[8]
	return strings.Split(strings.TrimSpace(callerStackTraceRaw), " ")[0]
}

func serializeInfoMap(infoMap map[string]any) string {
	b, err := json.Marshal(infoMap)
	if err != nil {
		return "unable to serialize info map"
	}
	return string(b)
}
