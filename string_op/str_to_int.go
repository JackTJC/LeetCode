package string_op

import (
	"math"
	"unicode"
)

const (
	PLUS  = byte('+')
	MINUS = byte('-')
	BLANK = byte(' ')
)

func strToInt(str string) int {
	helper := func(i, j int, cond bool) int {
		if cond {
			return i
		}
		return j

	}
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == BLANK {
			continue
		}
		if str[i] == MINUS || str[i] == PLUS || unicode.IsDigit(rune(str[i])) {
			start = i
			break
		}
		return 0
	}
	ret := 0
	signed := PLUS
	for i := start; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			ret = 10*ret + int(str[i]-'0')
			if ret >= math.MaxInt32 {
				break
			}
		} else {
			if i == start {
				signed = str[i]
			} else {
				break
			}

		}
	}
	if signed == MINUS {
		return helper(-ret, math.MinInt32, -ret > math.MinInt32)
	}
	return helper(ret, math.MaxInt32, ret < math.MaxInt32)
}
