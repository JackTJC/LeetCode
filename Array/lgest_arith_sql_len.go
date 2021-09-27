package array

/*
*1027. 最长等差数列
*给定一个整数数组 A，返回 A 中最长等差子序列的长度。
*回想一下，A 的子序列是列表 A[i_1], A[i_2], ..., A[i_k] 其中 0 <= i_1 < i_2 < ... < i_k <= A.length - 1。并且如果 B[i+1] - B[i]( 0 <= i < B.length - 1) 的值都相同，那么序列 B 是等差的。
**/
func longestArithSeqLength(nums []int) int {
	max := func(i, j int16) int16 {
		if i > j {
			return i
		}
		return j
	}
	n := len(nums)
	var ret int16
	dp := make([][]int16, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int16, 20000+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j] + 10000
			dp[i][diff] = max(dp[j][diff]+1, dp[i][diff])
			ret = max(ret, dp[i][diff])
		}
	}
	return int(ret + 1)
}

func longestArithSeqLength2(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	var ret int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] = max(dp[j][diff]+1, dp[i][diff])
			ret = max(ret, dp[i][diff])
		}
	}
	return ret
}
