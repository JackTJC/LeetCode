package string_op

import "unsafe"

func reverseStr(s string, k int) string {
	// ret := *(*[]byte)(unsafe.Pointer(&s))
	ret := []byte(s)
	n := len(ret)
	for i := 0; i < n; i += 2 * k {
		left, right := i, min(i+k, n-1)
		for left < right {
			t := ret[left]
			ret[left] = ret[right]
			ret[right] = t
			left++
			right--
		}
	}
	return *(*string)(unsafe.Pointer(&ret))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
