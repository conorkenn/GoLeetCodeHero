package leetcode

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		t := n % 10
		n /= 10
		product *= t
		sum += t
	}

	return product - sum
}
