package slidewindow

import "strings"

func lengthOfLongestSubstring(s string) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	n := len(s)
	l, r := 0, 0
	ret := 0
	for r < n {
		var idx int
		for r < n {
			idx = strings.Index(s[l:r], string([]byte{s[r]}))
			if idx != -1 {
				break
			}
			r++
		}
		ret = max(ret, r-l)
		l = l + idx + 1
		r++
	}
	return ret
}
