package sort_alg

func arrMinSum(arr []int) int {
	exist := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for j := -100; j <= 100; j++ {
			_, ok := exist[j]
			if !ok {
				exist[j] = 0
			}
			if arr[i] >= j {

			}
		}
	}
}
