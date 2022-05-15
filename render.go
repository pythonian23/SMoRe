package smore

import (
	"fmt"
	"strings"
)

func tokenSplit(token string, parts []string) []string {
	var out = make([]string, 0, cap(parts))
	for _, part := range parts {
		split := strings.Split(part, token)
		for i, s := range split {
			if i != 0 && i != len(split)-1 {
				out = append(out, s, token)
			} else {
				out = append(out, s)
			}
		}
	}
	return out
}

func Render(md string) string {
	var parts = []string{md}
	parts = tokenSplit("```", parts)
	parts = tokenSplit("**", parts)
	parts = tokenSplit("__", parts)
	parts = tokenSplit("*", parts)
	parts = tokenSplit("_", parts)
	parts = tokenSplit("`", parts)
	return fmt.Sprint(parts)
}
