package array

//两数之和返回下标
func twoSum(numbers []int, target int) []int {
	idxMap := make(map[int][]int)
	for idx, number := range numbers {
		if _, ok := idxMap[number]; !ok {
			idxMap[number] = []int{idx}
		}
		idxMap[number] = append(idxMap[number], idx)
	}
	for idx, number := range numbers {
		need := target - number
		if idxList, ok := idxMap[need]; ok {
			if idxList[len(idxList)-1] > idx {
				return []int{idx, idxList[len(idxList)-1]}
			}
		}
	}
	return nil
}
