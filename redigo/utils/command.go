package utils

import (
	"strings"
)

func GenCommand(s ...string) string {
	command := strings.Join(s, " ")
	return command + "\n"
}
