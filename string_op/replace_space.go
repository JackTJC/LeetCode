package string_op

import "bytes"

func replaceSpace(s string) string {
	buf := bytes.Buffer{}
	for _, c := range s {
		if c == rune(' ') {
			buf.WriteString("%20")
			continue
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
