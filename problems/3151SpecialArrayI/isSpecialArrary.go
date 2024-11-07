package leetcode

func isArraySpecial(nums []int) bool {

	for i := 1; i < len(nums); i++ {
		t1 := nums[i-1] % 2
		t2 := nums[i] % 2
		if t1 == t2 {
			return false
		}
	}
	return true
}
