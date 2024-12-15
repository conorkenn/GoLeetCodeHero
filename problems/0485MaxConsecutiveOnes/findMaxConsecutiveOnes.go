package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count += 1
			ans = max(ans, count)
		} else {
			count = 0
		}
	}
	return ans
}
