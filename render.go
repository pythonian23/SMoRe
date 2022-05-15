package smore

import (
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
	parts = tokenSplit("```", parts) // codeBlock
	parts = tokenSplit("\n\n", parts)
	parts = tokenSplit("**", parts) // boldA
	parts = tokenSplit("__", parts) // underline
	parts = tokenSplit("\n", parts)
	parts = tokenSplit("*", parts)  // italicA
	parts = tokenSplit("_", parts)  // italicU
	parts = tokenSplit("`", parts)  // code
	parts = tokenSplit("\\", parts) // escaped

	var out string
	var (
		codeBlock, boldA, underline, italicA, italicU, code, escaped bool
	)
	for _, part := range parts {
		if escaped || (codeBlock && part != "```") || (code && part != "`") {
			out += part
			continue
		}
		switch part {
		case "```":
			if codeBlock {
				// codeBlock terminate
			} else {
				// codeBlock start
			}
			codeBlock = !codeBlock
		case "**":
			if boldA {
				// boldA terminate
			} else {
				// boldA start
			}
			boldA = !boldA
		case "__":
			if underline {
				// underline terminate
			} else {
				// underline start
			}
			underline = !underline
		case "*":
			if italicA {
				// italicA terminate
			} else {
				// italicA start
			}
			italicA = !italicA
		case "_":
			if italicU {
				// italicU terminate
			} else {
				// italicU start
			}
			italicU = !italicU
		case "`":
			if code {
				// code terminate
			} else {
				// code start
			}
			code = !code
		case "\\":
			if escaped {
				// escaped terminate
			} else {
				// escaped start
			}
			escaped = !escaped
		default:
			out += part
		}
	}
	return out
}
