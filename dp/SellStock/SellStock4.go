package SellStock

//188. 买卖股票的最佳时机 IV
//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
import "math"

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k == 0 {
		return 0
	}
	Max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	profit := make([]int, 2*k)
	for i := 0; i < k; i++ {
		profit[2*i] = math.MinInt32
		profit[2*i+1] = 0
	}
	for _, price := range prices {
		for i := 0; i < k; i++ {
			if i == 0 {
				profit[2*i] = Max(profit[2*i], -price)
				profit[2*i+1] = Max(profit[2*i+1], profit[2*i]+price)
			} else {
				profit[2*i] = Max(profit[2*i], profit[2*i-1]-price)
				profit[2*i+1] = Max(profit[2*i+1], profit[2*i]+price)
			}
		}
	}
	return profit[2*k-1]
}
