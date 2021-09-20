package array

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	begin, end := 0, 0
	var sum int
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	ret := int(1e5) + 1
	for end < n {
		sum += nums[end]
		for sum >= target {
			ret = min(ret, end-begin+1)
			sum -= nums[begin]
			begin++
		}
		end++
	}
	if ret == 100001 {
		return 0
	}
	return ret
}
