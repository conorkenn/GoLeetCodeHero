package leetcode

func maxPower(s string) int {
	ans := 1
	count := 1
	prev := rune(s[0])

	for _, char := range s[1:] {
		if char == prev {
			count += 1
			ans = max(ans, count)
		} else {
			count = 1
		}
		prev = char
	}
	return ans
}
