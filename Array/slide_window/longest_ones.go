package slidewindow

/*
1004. 最大连续1的个数 III
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。
*/
func longestOnes(nums []int, k int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	start := 0
	ret := 0
	n := len(nums)
	end := 0
	for {
		cnt := 0
		for cnt < k && end < n {
			if nums[end] == 0 {
				cnt++
			}
			end++
		}
		for end < n && nums[end] == 1 {
			end++
		}
		ret = max(ret, end-start)
		if end == n {
			break
		}
		start++
		end = start
	}
	return ret
}
