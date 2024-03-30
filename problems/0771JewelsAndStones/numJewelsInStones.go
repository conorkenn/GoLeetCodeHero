package leetcode

func numJewelsInStones(jewels string, stones string) int {
	j := map[rune]bool{}

	for _, jewel := range jewels {
		j[jewel] = true
	}

	ans := 0

	for _, stone := range stones {
		if j[stone] {
			ans += 1
		}
	}

	return ans
}
