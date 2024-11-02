package leetcode

import "strings"

func isCircularSentence(sentence string) bool {

	words := strings.Split(sentence, " ")

	for i := 1; i < len(words); i++ {
		prev := words[i-1]
		curr := words[i]

		if prev[len(prev)-1] != curr[0] {
			return false
		}
	}
	firstWord := words[0]
	lastWord := words[len(words)-1]
	return lastWord[len(lastWord)-1] == firstWord[0]
}
