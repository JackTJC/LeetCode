package slide_window

import (
	"strings"
	"testing"
)

func TestLgest(t *testing.T) {
	lengthOfLongestSubstring("abcabcbb")
}

func TestSlice(t *testing.T) {
	s := "123456789"
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			t.Log(s[i:j])
		}
	}
}

func TestIndex(t *testing.T) {
	t.Log(strings.Index("123456", "234"))
}
