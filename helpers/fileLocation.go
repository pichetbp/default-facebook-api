package helpers

import (
	"runtime"
	"strconv"
)

func SetFileLocation() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return StringBuild([]string{file, ":", strconv.Itoa(line)})
}
