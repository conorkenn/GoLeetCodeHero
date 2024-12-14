package leetcode

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	ans := 0

	for _, word := range words {
		consistent := true
		for i := 0; i < len(word); i++ {
			if !strings.Contains(allowed, string(word[i])) {
				consistent = false
				break
			}
		}
		if consistent {
			ans += 1
		}
	}
	return ans
}
