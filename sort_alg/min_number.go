package sort_alg

/*
解题思路：
此题求拼接起来的最小数字，本质上是一个排序问题。设数组 nums中任意两数字的字符串为 x 和 y ，则规定 排序判断规则 为：

若拼接字符串 x + y > y + x，则 x “大于” y ；
反之，若 x + y < y + x，则 x “小于” y ；
*/

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
	strNums := make([]string, 0)
	for _, num := range nums {
		strNums = append(strNums, fmt.Sprintf("%d", num))
	}
	sort.Slice(strNums, func(i, j int) bool {
		return compareStr(strNums[i], strNums[j])
	})
	return strings.Join(strNums, "")
}

func compareStr(s1, s2 string) bool {
	concat := func(s1, s2 string) string {
		return fmt.Sprintf("%s%s", s1, s2)
	}
	return concat(s1, s2) < concat(s2, s1)
}
