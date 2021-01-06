package SellStock

//123. 买卖股票的最佳时机 III
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
import "math"

func maxProfit3(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	/**
	  对于任意一天考虑四个变量:
	  fstBuy: 在该天第一次买入股票可获得的最大收益
	  fstSell: 在该天第一次卖出股票可获得的最大收益
	  secBuy: 在该天第二次买入股票可获得的最大收益
	  secSell: 在该天第二次卖出股票可获得的最大收益
	  分别对四个变量进行相应的更新, 最后secSell就是最大
	  收益值(secSell >= fstSell)
	  **/
	firstBuy := math.MinInt32
	firstSell := 0
	secondBuy := math.MinInt32
	secondSell := 0
	for _, i := range prices {
		firstBuy = Max(firstBuy, -i)
		firstSell = Max(firstSell, firstBuy+i)
		secondBuy = Max(secondBuy, firstSell-i)
		secondSell = Max(secondSell, secondBuy+i)
	}
	return secondSell
}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
