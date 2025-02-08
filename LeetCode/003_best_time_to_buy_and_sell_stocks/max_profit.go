/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock
and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Note that you cannot sell a stock before you buy one.

Tip: find local min and search for local max, sliding window;
*/

package maxprofit

// Brute Force Solution
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func maxProfitBF(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > maxProfit {
				maxProfit = prices[i] - prices[j]
			}
		}
	}
	return maxProfit
}

// dynamic programming
// for every price, find the max profit, then record the current minimum price.
// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	var minValue = prices[0]
	var maxProfit int
	//{7, 1, 5, 3, 6, 4}, {7, 6, 4, 3, 1}
	for _, price := range prices {
		if price-minValue > maxProfit {
			maxProfit = price - minValue
		}
		if price < minValue {
			minValue = price
		}
	}
	return maxProfit
}
