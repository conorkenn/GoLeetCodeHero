package leetcode

func makeFancyString(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) >= 2 && stack[len(stack)-1] == s[i] && stack[len(stack)-2] == s[i] {
			continue
		}
		stack = append(stack, s[i])
	}

	return string(stack)
}
