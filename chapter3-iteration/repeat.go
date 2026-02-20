package iteration

import "strings"

func Repeat(str string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(str)
	}
	return repeated.String()
}
