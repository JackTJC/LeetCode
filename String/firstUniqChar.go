package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("aaabbc"))
}
func firstUniqChar(s string) int {
	m := make(map[byte]int)
	sByte := []byte(s)
	for _, c := range sByte {
		m[c]++
	}
	for k, v := range m {
		if v != 1 {
			delete(m, k)
		}
	}
	for index, c := range sByte {
		if m[c] == 1 {
			return index
		}
	}
	return -1
}
