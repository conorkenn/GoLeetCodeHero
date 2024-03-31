package leetcode

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	var t int
	for l < r {
		t = numbers[l] + numbers[r]
		if t == target {
			return []int{l + 1, r + 1}
		}

		if t > target {
			r -= 1
		} else {
			l += 1
		}
	}
	return nil
}
