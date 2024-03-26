package leetcode

import "slices"

func lengthOfLongestSubstring(s string) int {
	substring := []rune{}
	ans := 0
	runes := []rune(s)
	for _, c := range runes {
		if !slices.Contains(substring, c) {
			substring = append(substring, c)
		} else {
			for slices.Contains(substring, c) {
				substring = substring[1:]
			}
			substring = append(substring, c)
		}
		if len(substring) > ans {
			ans = len(substring)
		}
	}

	return ans
}
