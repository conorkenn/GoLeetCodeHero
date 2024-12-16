package leetcode

import "strconv"

func hammingWeight(n int) int {
	bin := strconv.FormatInt(int64(n), 2)
	count := 0
	for _, bit := range bin {
		if bit == '1' {
			count += 1
		}
	}
	return count
}
