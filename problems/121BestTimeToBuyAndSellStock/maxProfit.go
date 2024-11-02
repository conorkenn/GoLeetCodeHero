package leetcode

func maxProfit(prices []int) int {

	maxProfit := 0
	minPurchase := prices[0]

	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-minPurchase)
		minPurchase = min(minPurchase, prices[i])
	}

	return maxProfit
}
