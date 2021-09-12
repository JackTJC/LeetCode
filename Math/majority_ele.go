package math

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//常规解法hash map
//排序取中位数

//摩尔投票算法
//众数投+1，其他数投-1，可以确定的是最后的投票结果一定为正数
//众数与其他不同的数会一一抵消
//模拟这个一一抵消过程
func majorityElement(nums []int) int {
	cnt := 0
	var res int
	for _, num := range nums {
		if cnt == 0 {
			res = num
			cnt++
			continue
		}
		//遇到相同就投+1，不同就投-1，相同的居多，最后结果一定是众数
		if res == num {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
