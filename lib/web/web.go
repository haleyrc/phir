package web

import "strings"

func NewString(s string) string {
	s = strings.TrimSpace(s)
	return s
}
