package SellStock

//121. 买卖股票的最佳时机
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
//注意：你不能在买入股票前卖出股票。
import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 1; i < n; i++ {
		min = int(math.Min(float64(prices[i]), float64(min)))
		max = int(math.Max(float64(prices[i]-min), float64(max)))
	}
	return max
}
