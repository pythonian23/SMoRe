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
	if len(values) == 0 {
		return ""
	}
	return "\u001b[" + strings.Join(values, ";") + "m"
}

type state struct {
	codeBlock bool
	boldA     bool
	underline bool
	italicA   bool
	italicU   bool
	code      bool
}

func Render(md string) string {
	var parts = []string{md}

	for _, token := range tokens {
		parts = tokenSplit(token, parts)
	}

	var out string
	var escaped bool
	var current = state{}
	var previous = state{}
	for _, part := range parts {
		if escaped || (current.codeBlock && part != "```") || (current.code && part != "`") {
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
			escaped = !escaped
		case "\n\n":
			out += "\n"
		case "\n":
			out += " "
		default:
			out += part
		}

		var escapes []string
		var styleReset bool
		switch {
		case (current.boldA != previous.boldA) && !current.boldA:
			fallthrough
		case (current.underline != previous.underline) && !current.underline:
			fallthrough
		case (current.italicA != previous.italicA) && !current.italicA:
			fallthrough
		case (current.italicU != previous.italicU) && !current.italicU:
			escapes = append(escapes, "0")
			styleReset = true
		case (current.codeBlock != previous.codeBlock) && !current.codeBlock:
		case (current.code != previous.code) && !current.code:
		case (styleReset || (current.boldA != previous.boldA)) && current.boldA:
			escapes = append(escapes, "1")
		case (styleReset || (current.underline != previous.underline)) && current.underline:
			escapes = append(escapes, "3")
		case (styleReset || (current.italicA != previous.italicA)) && current.italicA:
			fallthrough
		case (styleReset || (current.italicU != previous.italicU)) && current.italicU:
			escapes = append(escapes, "4")
		case (current.codeBlock != previous.codeBlock) && current.codeBlock:
		case (current.code != previous.code) && current.code:
		}
		out += escape(escapes...)
		previous = current
	}
	return out
}
