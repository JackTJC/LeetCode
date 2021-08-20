package string_op

func reverseLeftWords(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	absMove := n % len(s)
	return s[absMove:] + s[:absMove]
}
