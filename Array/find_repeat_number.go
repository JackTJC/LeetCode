package array

func findRepeatNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		var ori int
		ori = nums[i]
		if nums[i] < 0 {
			ori = nums[i] + n
		}
		if nums[ori] < 0 {
			return ori
		}
		nums[ori] -= n
	}
	return 0
}
