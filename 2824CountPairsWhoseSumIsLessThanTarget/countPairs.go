package leetcode

func countPairs(nums []int, target int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if 0 <= i && i < j && j < len(nums) && nums[i]+nums[j] < target {
				ans++
			}
		}
	}
	return ans
}
