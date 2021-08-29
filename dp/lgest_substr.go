package dp

import "strings"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	l, r := 0, 0
	ret := 0
	for r < n {
		idx := strings.Index(s[l:r], string([]byte{s[r]}))
		for r < n && idx == -1 {
			r++
			idx = strings.Index(s[l:r], string([]byte{s[r]}))
		}
		ret = max(ret, r-l)
		l = idx + 1
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
