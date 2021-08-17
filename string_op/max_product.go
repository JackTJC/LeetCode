package string_op

import "sort"

/*
剑指 Offer II 005. 单词长度的最大乘积
给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
*/
func maxProduct(words []string) int {
	//低26位作为bit mask
	wordBitMasks := make([]int, len(words))
	maskMap := make(map[byte]int)
	base := 1
	for i := byte('a'); i <= byte('z'); i++ {
		maskMap[i] = base
		base <<= 1
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for i, word := range words {
		for j := range word {
			wordBitMasks[i] |= maskMap[word[j]]
		}
	}
	n := len(words)
	ret := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if wordBitMasks[i]&wordBitMasks[j] == 0 {
				ret = max(ret, len(words[i])*len(words[j]))
			}
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
