package leetcode

import "unicode"

func countKeyChanges(s string) int {
	ans := 0

	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i-1])) {
			ans++
		}
	}
	return ans
}
