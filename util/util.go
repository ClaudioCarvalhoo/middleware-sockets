package util

import "strings"

func TrimString(text string) string {
	text = strings.TrimRight(text, "\n")
	text = strings.TrimRight(text, "\r\n")
	return text
}
