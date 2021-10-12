package math

import "math/rand"

func rand10() int {
	var first, second int
	second = 8
	first = 8
	for first > 6 {
		first = rand7()
	}
	for second > 5 {
		second = rand7()
	}
	if first&1 > 0 {
		return second
	}
	return second + 5
}
func rand7() int {
	return rand.Intn(7) + 1
}
