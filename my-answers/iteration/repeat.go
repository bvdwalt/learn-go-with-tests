package iteration

import (
	"strings"
)

func Repeat(s string, n int) string {
	var result strings.Builder
	for range n {
		result.WriteString(s)
	}
	return result.String()
}
