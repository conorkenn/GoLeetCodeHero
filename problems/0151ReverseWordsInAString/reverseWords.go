package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	l := 0
	r := len(words) - 1

	for l < r {
		words[l], words[r] = words[r], words[l]
		l += 1
		r -= 1
	}

	return strings.Join(words, " ")
}
