package smore

import (
	"log"
	"strings"
)

func tokenSplit(token string, tokens, parts []string) []string {
	var out = make([]string, 0, cap(parts))
PartLoop:
	for _, part := range parts {
		for _, t := range tokens {
			if part == t {
				out = append(out, part)
				continue PartLoop
			}
		}
		split := strings.Split(part, token)
		for i, s := range split {
			if i != len(split)-1 {
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
	var tokens = []string{
		"```", // codeBlock
		"\n\n",
		"**", // boldA
		"__", // underline
		"\n",
		"*",  // italicA
		"_",  // italicU
		"`",  // code
		"\\", // escaped
	}
	for _, token := range tokens {
		parts = tokenSplit(token, tokens, parts)
	}

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
		case "\n\n":
			out += "\n"
		case "\n":
			out += " "
		default:
			out += part
		}
	}
	return out
}
