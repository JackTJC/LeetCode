package DynamicProgramming;

public class SellStockWithFee {
    static class Solution {
        public int maxProfit(int[] prices, int fee) {
            /*
            cash:the maximum if we have no stock
            hold:the maximum if we have a stock
            when started, if we want to have a stock we must buy a stock at the first day,
            and if we don't have a stock we should have 0 $
             */
            int cash = 0, hold = -prices[0];
            for (int i = 1; i < prices.length; i++) {
                cash = Math.max(cash, hold + prices[i] - fee);
                hold = Math.max(hold, cash - prices[i]);
            }
            //because at last we have no stock according to the problem, so we return the cash
            return cash;
        }
    }
}
