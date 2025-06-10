package web

import (
	"strconv"
	"strings"
)

func NewInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func NewString(s string) string {
	s = strings.TrimSpace(s)
	return s
}
