package leetcode

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		value, ok := seen[target-num]
		if ok {
			return []int{i, value}
		}
		seen[num] = i
	}
	return nil
}
