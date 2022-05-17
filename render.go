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
	"```",  // codeBlock
	"\n* ", // uList
	"\n- ", // uList (same variable as above)
	"\n\n",
	"**", // boldA
	"__", // underline
	"\n",
	"*",  // italicA
	"_",  // italicU
	"`",  // code
	"\\", // escaped
}

var debug = false

func escape(values ...string) string {
	if len(values) == 0 {
		return ""
	}
	if debug {
		return "\u001b[" + strings.Join(values, ";") + "m[" + strings.Join(values, ",") + "]"
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
	var headerLevel int
	var escaped bool
	var current = state{}
	var previous = state{}
	for i, part := range parts {
		var escapes []string
		var styleReset bool
		header := headerLevel > 0

		if escaped || (current.codeBlock && part != "```") || (current.code && part != "`") {
			out += part
			continue
		}

		if i == 0 || (parts[i-1] == "\n" || parts[i-1] == "\n\n") {
			if strings.HasPrefix(part, "# ") {
				headerLevel = 1
				if i != 0 {
					out += "\n"
				}
				out += escape("1", "4", "42")
			} else if strings.HasPrefix(part, "## ") {
				if i != 0 {
					out += "\n"
				}
				headerLevel = 2
				out += escape("43")
			} else if strings.HasPrefix(part, "### ") {
				if i != 0 {
					out += "\n"
				}
				headerLevel = 3
				out += escape("44", "2")
			}
		}

		switch part {
		case "```":
			current.codeBlock = !current.codeBlock
		case "\n* ":
			fallthrough
		case "\n- ":
			if header {
				headerLevel = 0
				header = false
			}
			out += escape("0", "1") + "\n - " + escape("0")
			styleReset = true
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
			if header {
				headerLevel = 0
				out += escape("0")
			}
			out += "\n"
		case "\n":
			if header {
				headerLevel = 0
				out += escape("0") + "\n"
			} else {
				out += " "
			}
		default:
			out += part
		}

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
			escapes = append(escapes, "4")
		case (styleReset || (current.italicA != previous.italicA)) && current.italicA:
			fallthrough
		case (styleReset || (current.italicU != previous.italicU)) && current.italicU:
			escapes = append(escapes, "3")
		case (current.codeBlock != previous.codeBlock) && current.codeBlock:
		case (current.code != previous.code) && current.code:
		}
		if styleReset && header {
			if headerLevel == 1 {
				escapes = append(escapes, "1", "4", "42")
			} else if headerLevel == 2 {
				escapes = append(escapes, "43")
			} else if headerLevel == 3 {
				escapes = append(escapes, "44", "2")
			}
		}

		out += escape(escapes...)
		previous = current
	}
	return out
}
