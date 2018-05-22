package util

import (
	"strings"
)

func Substring(s string, start int64, end int64) string {
	separator := ""
	slice := strings.Split(s, separator)
	if start < 0 {
		start = 0
	}

	if length := int64(len(slice)); end > length {
		end = length
	}
	return strings.Join(slice[start:end], separator)
}
