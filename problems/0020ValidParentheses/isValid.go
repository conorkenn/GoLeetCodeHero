package leetcode

import "strings"

func isValid(s string) bool {

	lparen := "[{("
	rparen := "]})"
	stack := []rune{}
	for _, paren := range s {
		if strings.ContainsRune(lparen, paren) {
			stack = append(stack, paren)
		} else {
			if len(stack) == 0 {
				return false
			}
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if strings.IndexRune(lparen, t) != strings.IndexRune(rparen, paren) {
				return false
			}
		}
	}
	return len(stack) == 0
}
