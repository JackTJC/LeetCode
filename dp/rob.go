package dp

func rob(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	steal, notSteal := nums[0], 0
	for i := 1; i < len(nums); i++ {
		//如果今天偷窃
		t := steal
		steal = notSteal + nums[i]
		//今天不偷窃
		notSteal = max(t, notSteal)
	}
	return max(steal, notSteal)
}
