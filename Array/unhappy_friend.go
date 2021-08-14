package Array

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	ret := 0
	//hash 表保存
	pairMap := make(map[int]int)
	for _, pair := range pairs {
		pairMap[pair[0]] = pair[1]
		pairMap[pair[1]] = pair[0]
	}
	for i := 0; i < n; i++ {
		pair := pairMap[i]
		for j := 0; j < n; j++ {
			current := preferences[i][j]
			if current == pair {
				break
			}
			// 现在需要验证current与i的亲密度是否大于current的原配
			pairOfCur := pairMap[current]
			unhappy := false
			for k := 0; k < n; k++ {
				//先遇到原配
				if preferences[current][k] == pairOfCur {
					break
				}
				//先遇到i
				if preferences[current][k] == i {
					unhappy = true
					ret++
					break
				}
			}
			if unhappy {
				break
			}
		}
	}
	return ret
}
