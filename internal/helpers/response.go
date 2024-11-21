package helpers

import (
	"strings"
)

func IsTrue(string string) bool {
	return strings.TrimSpace(string) == "true"
}
