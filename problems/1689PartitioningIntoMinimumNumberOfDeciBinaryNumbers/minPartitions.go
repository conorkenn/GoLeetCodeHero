package leetcode

//int(v-'0') converts the character c representing a digit to its corresponding integer value. This integer value can then be used for numerical calculations or comparisons.

func minPartitions(n string) int {
	ans := 0
	for _, c := range n {
		if int(c-'0') > ans {
			ans = int(c - '0')
		}
	}
	return ans
}
