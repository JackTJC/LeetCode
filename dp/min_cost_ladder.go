package dp

/*
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。
请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
*/
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0, 0)
	pp, p := cost[0], cost[1]
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 2; i < len(cost); i++ {
		t := min(pp, p) + cost[i]
		pp = p
		p = t
	}
	return p
}
