package helpers

import (
	"strings"
)

func StringBuild(args []string) string {
	var s strings.Builder
	for _, arg := range args {
		s.WriteString(arg)
	}
	return s.String()
}
