package dp

import "math"

/*
* 最长公共子序列
* */
func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(arr ...int) int {
		ret := math.MinInt32
		for _, v := range arr {
			if v > ret {
				ret = v
			}
		}
		return ret
	}
	threeOp := func(cond bool, a, b int) int {
		if cond {
			return a
		}
		return b
	}
	l1, l2 := len(text1), len(text2)
	dpArr := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dpArr[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 || j == 0 {
				dpArr[i][j] = 0
				continue
			}
			dpArr[i][j] = max(dpArr[i][j-1], dpArr[i-1][j], dpArr[i-1][j-1]+threeOp(text1[i-1] == text2[j-1], 1, 0))
		}
	}
	return dpArr[l1][l2]
}
