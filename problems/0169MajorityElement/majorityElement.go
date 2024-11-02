package leetcode

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	ans := 0
	max := 0
	for _, num := range nums {
		counts[num] += 1
		if counts[num] > max {
			max = counts[num]
			ans = num
		}
	}

	return ans

}
