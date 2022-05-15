package smore

import (
	"strings"
)

func tokenSplit(token string, parts []string) []string {
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

var tokens = [...]string{
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

func escape(values ...string) string {
	return "\u001b[" + strings.Join(values, ";") + "m"
}

type state struct {
	codeBlock bool
	boldA     bool
	underline bool
	italicA   bool
	italicU   bool
	code      bool
	escaped   bool
}

func Render(md string) string {
	var parts = []string{md}

	for _, token := range tokens {
		parts = tokenSplit(token, parts)
	}

	var out string
	var current = state{}
	var previous = state{}
	for _, part := range parts {
		if current.escaped || (current.codeBlock && part != "```") || (current.code && part != "`") {
			out += part
			continue
		}
		switch part {
		case "```":
			current.codeBlock = !current.codeBlock
		case "**":
			current.boldA = !current.boldA
		case "__":
			current.underline = !current.underline
		case "*":
			current.italicA = !current.italicA
		case "_":
			current.italicU = !current.italicU
		case "`":
			current.code = !current.code
		case "\\":
			current.escaped = !current.escaped
		case "\n\n":
			out += "\n"
		case "\n":
			out += " "
		default:
			out += part
		}
		previous = current
		switch {
		case current.codeBlock != previous.codeBlock:
		case current.boldA != previous.boldA:
		case current.underline != previous.underline:
		case current.italicA != previous.italicA:
		case current.italicU != previous.italicU:
		case current.code != previous.code:
		case current.escaped != previous.escaped:
		}
	}
	return out
}
