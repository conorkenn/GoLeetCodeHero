package leetcode

func checkIfPangram(sentence string) bool {
	d := make(map[rune]int)

	for _, c := range sentence {
		if _, ok := d[c]; ok {
			d[c] += 1
		} else {
			d[c] = 1
		}
	}
	return len(d) == 26
}
