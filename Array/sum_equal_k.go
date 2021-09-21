package array

/*
剑指 Offer II 010. 和为 k 的子数组
给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
*/
func subarraySum(nums []int, k int) int {
	nums = append([]int{0}, nums...)
	prefixSum := make(map[int][]int)
	sum := 0
	for idx, v := range nums {
		sum += v
		if _, ok := prefixSum[sum]; !ok {
			prefixSum[sum] = make([]int, 0)
		}
		prefixSum[sum] = append(prefixSum[sum], idx)
	}
	cnt := 0
	for prefix, idxList := range prefixSum {
		targetIdxList, ok := prefixSum[k+prefix]
		if !ok {
			continue
		}
		for _, idx := range idxList {
			for _, targetIdx := range targetIdxList {
				if idx < targetIdx {
					cnt++
				}
			}
		}
	}
	return cnt
}
