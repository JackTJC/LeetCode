package dp

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	ret := make([]int, n+1)
	ret[0] = 0
	for i := 1; i < n+1; i++ {
		if i%2 == 0 {
			ret[i] = ret[i/2]
		} else {
			ret[i] = ret[i-1] + 1
		}
	}
	return ret
}
