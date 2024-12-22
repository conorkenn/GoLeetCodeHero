package leetcode

func countPoints(rings string) int {
	m := make(map[rune]map[rune]bool)

	for i := 0; i < len(rings); i += 2 {
		color := rune(rings[i])
		rod := rune(rings[i+1])

		if _, exists := m[rod]; !exists {
			m[rod] = make(map[rune]bool)
		}

		m[rod][color] = true
	}

	count := 0

	for _, colors := range m {
		if colors['R'] && colors['G'] && colors['B'] {
			count++
		}
	}
	return count
}
