package leetcode

func maximumWealth(accounts [][]int) int {
	ans := 0

	for _, account := range accounts {
		accountTotal := 0
		for _, amount := range account {
			accountTotal += amount
		}
		if accountTotal > ans {
			ans = accountTotal
		}
	}
	return ans
}
