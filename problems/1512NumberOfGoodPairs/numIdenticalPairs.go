package leetcode

func numIdenticalPairs(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				ans += 1
			}
		}
	}
	return ans

}
