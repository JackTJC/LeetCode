package string_op

import "testing"

func TestRs(t *testing.T) {
	t.Logf("result = %s", replaceSpace("this is a test for replace space"))
}

func TestRLW(t *testing.T) {
	t.Logf("%s", reverseLeftWords("abcde", 1))
	t.Logf("%s", reverseLeftWords("abcde", 8))
	t.Logf("%s", reverseLeftWords("", 1))
}

func TestRV(t *testing.T) {
	// t.Logf("%s", reverseVowels("leetcode"))
	// t.Log(reverseVowels(""))
	// t.Log(reverseVowels("xxxaxxx"))
	// t.Log(reverseVowels(" "))
	t.Log(reverseVowels("hello"))
}

func TestRsk(t *testing.T) {
	t.Log(reverseStr("abcdefg", 2))
}
