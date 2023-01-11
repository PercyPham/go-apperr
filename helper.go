package apperr

import (
	"runtime/debug"
	"strings"
)

func getCallerStackTrace() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	callerStackTraceRaw := stack[8]
	return strings.Split(strings.TrimSpace(callerStackTraceRaw), " ")[0]
}
