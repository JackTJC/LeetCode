package slidewindow

/*
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	begin := 0
	product := 1
	cnt := 0
	for end, value := range nums {
		product *= value
		for begin <= end && product >= k {
			product /= nums[begin]
			begin++
		}
		if end >= begin {
			cnt += end - begin + 1
		}
	}
	return cnt
}
