package array

/*
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。
*/
func search(nums []int, target int) int {
	n := len(nums)
	idx := 0
	left, right := 0, n-1
	find := false
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			idx = mid
			find = true
			break
		}
	}
	if !find {
		return 0
	}
	forwards, backwards := idx+1, idx-1
	cnt := 0
	for forwards < n && nums[forwards] == nums[idx] {
		forwards++
		cnt++
	}
	for backwards >= 0 && nums[backwards] == nums[idx] {
		backwards--
		cnt++
	}
	return cnt + 1

}
