package leetcode

func minChanges(s string) int {
	c := 0

	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			c += 1
		}
	}
	return c
}
