//122. 买卖股票的最佳时机 II
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
package SellStock

func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	hold := -prices[0]
	cash := 0
	for i := 1; i < n; i++ {
		hold = max(hold, cash-prices[i])
		cash = max(cash, hold+prices[i])
	}
	return cash
}
func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
