package leetcode

func runningSum(nums []int) []int {
	ans := []int{}
	runningSum := 0
	for _, num := range nums {
		runningSum += num
		ans = append(ans, runningSum)
	}
	return ans
}
