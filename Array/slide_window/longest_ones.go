package slidewindow

func longestOnes(nums []int, k int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	start := 0
	ret := 0
	n := len(nums)
	end := 0
	for {
		cnt := 0
		for cnt <= k && end < n {
			if nums[end] == 0 {
				cnt++
			}
			end++
		}
		ret = max(ret, end-start)
		if end == n {
			break
		}
		start++
		end = start
	}
	return ret
}
