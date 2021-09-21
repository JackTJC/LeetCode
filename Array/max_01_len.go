package array

import "sort"

/*
剑指 Offer II 011. 0 和 1 个数相同的子数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
*/
func findMaxLength(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	nums = append([]int{0}, nums...)
	prefix := make(map[int][]int)
	sum := 0
	for idx, v := range nums {
		sum += v
		if _, ok := prefix[sum]; !ok {
			prefix[sum] = make([]int, 0)
		}
		prefix[sum] = append(prefix[sum], idx)
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ret := 0
	for _, idxs := range prefix {
		sort.Ints(idxs)
		ret = max(ret, idxs[len(idxs)-1]-idxs[0])
	}
	return ret
}
