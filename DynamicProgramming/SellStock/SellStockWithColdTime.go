package SellStock

//309. 最佳买卖股票时机含冷冻期
//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
import "math"

func maxProfitWithColdTime(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	Max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	hold := math.MinInt32
	sell := 0
	cool := 0
	for _, price := range prices {
		hold = Max(hold, sell-price)
		sell = Max(cool, sell)
		cool = hold + price
	}
	return Max(sell, cool)
}
