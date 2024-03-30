package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == num
}
