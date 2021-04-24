package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	s1 := []byte(ransomNote)
	s2 := []byte(magazine)
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for _, c1 := range s1 {
		m1[c1]++
	}
	for _, c2 := range s2 {
		m2[c2]++
	}
	for k, v := range m1 {
		if cnt, ok := m2[k]; ok {
			if cnt < v {
				return false
			}
			continue
		}
		return false
	}
	return true
}
