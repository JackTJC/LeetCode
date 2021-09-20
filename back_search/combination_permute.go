package backtrack_search

/*
排列和组合算法汇总
*/

import (
	"sort"
)

/*
给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
*/

func combinationSum2(candidates []int, target int) [][]int {
	sum := func(arr []int) int {
		s := 0
		for _, v := range arr {
			s += v
		}
		return s
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	ret := make([][]int, 0)
	arr := make([]int, 0)
	var backtrack func(int)
	backtrack = func(start int) {
		cur := sum(arr)
		if cur == target {
			temp := make([]int, len(arr))
			copy(temp, arr)
			ret = append(ret, temp)
			return
		}
		if cur > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			//防止出现重复
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			arr = append(arr, candidates[i])
			backtrack(i + 1)
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return ret
}
func permute(nums []int) [][]int {
	sel := make(map[int]struct{})
	ret := make([][]int, 0)
	arr := make([]int, 0)
	var helper func()
	helper = func() {
		if len(arr) == len(nums) {
			t := make([]int, len(arr))
			copy(t, arr)
			ret = append(ret, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			if _, ok := sel[nums[i]]; ok {
				continue
			}
			arr = append(arr, nums[i])
			sel[nums[i]] = struct{}{}
			helper()
			delete(sel, nums[i])
			arr = arr[:len(arr)-1]
		}
	}
	helper()
	return ret
}

/*
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。

要解决重复问题，我们只要设定一个规则，保证在填第 idx 个数的时候重复数字只会被填入一次即可。
而在本题解中，我们选择对原数组排序，保证相同的数字都相邻，然后每次填入的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」
*/
func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	arr := make([]int, 0)
	visied := make(map[int]bool)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var helper func()
	helper = func() {
		if len(arr) == len(nums) {
			t := make([]int, len(arr))
			copy(t, arr)
			ret = append(ret, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			//核心去重代码
			if visied[i] || i > 0 && nums[i] == nums[i-1] && !visied[i-1] {
				continue
			}
			arr = append(arr, nums[i])
			visied[i] = true
			helper()
			visied[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	helper()
	return ret
}
