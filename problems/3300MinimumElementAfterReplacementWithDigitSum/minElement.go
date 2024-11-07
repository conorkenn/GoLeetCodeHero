package leetcode

func minElement(nums []int) int {
	ans := int(^uint(0) >> 1) // MaxInt for int
	for _, num := range nums {
		t := num
		sum := 0
		for t > 0 {
			sum += (t % 10)
			t /= 10
		}
		if sum < ans {
			ans = sum
		}
	}
	return ans
}
