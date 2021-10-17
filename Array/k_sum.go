package array

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	ret := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ret = append(ret, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ret
}

/*
* 自己写的一种解法
* 感觉更好理解一些
 */
func _threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	n := len(nums)
	//先排序
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		//跳过重复的
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -nums[first]
		//设定左右指针
		second, third := first+1, n-1
		for second < third {
			//跳过重复的
			if second > first+1 && second < n && nums[second] == nums[second-1] {
				second++
				continue
			}
			current := nums[second] + nums[third]
			if current > target {
				third--
			} else if current < target {
				second++
			} else {
				ret = append(ret, []int{nums[first], nums[second], nums[third]})
				second++ //不能直接break，因为后面可能还有候选值
			}
		}
	}
	return ret
}
