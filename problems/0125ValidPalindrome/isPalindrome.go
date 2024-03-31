package leetcode

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
