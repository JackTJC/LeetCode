package backtrack_search

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for numbers[right] == numbers[left] && left < right {
		left++
	}
	for right-left > 1 {
		mid := (left + right) / 2
		if numbers[mid] <= numbers[right] {
			right = mid
		} else {
			left = mid
		}
	}
	return min(numbers[left], numbers[right])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
