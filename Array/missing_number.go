package array

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
func missingNumber(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(nums) == 0 {
		return 1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > mid {
			right = max(left, mid-1)
		} else {
			left = min(right, mid+1)
		}
	}
	if nums[right] > right {
		return nums[right] - 1
	}
	return nums[right] + 1
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
