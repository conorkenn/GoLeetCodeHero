package leetcode

func reverse(x int) int {
	ans := 0
	isNegative := 1
	if x < 0 {
		isNegative = -1
		x *= isNegative
	}

	for x > 0 {
		ans = ans*10 + x%10
		if ans > 2147483647 || ans < -2147483648 {
			return 0
		}
		x /= 10
	}
	return ans * isNegative
}
