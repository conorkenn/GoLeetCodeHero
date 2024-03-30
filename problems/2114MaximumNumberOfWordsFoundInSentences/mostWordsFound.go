package leetcode

import "strings"

func mostWordsFound(sentences []string) int {
	ans := 0
	count := 0
	for _, sentence := range sentences {
		count = strings.Count(sentence, " ") + 1
		if count > ans {
			ans = count
		}
	}
	return ans
}
