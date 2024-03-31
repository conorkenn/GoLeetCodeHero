package leetcode

func isAnagram(s string, t string) bool {
	dict := make(map[rune]int)

	for _, c := range s {
		dict[c]++
	}

	for _, c := range t {
		dict[c]--
	}

	for _, d := range dict {
		if d != 0 {
			return false
		}
	}
	return true

}
