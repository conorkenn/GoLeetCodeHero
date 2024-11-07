package leetcode

func isBalanced(num string) bool {
	evenSum := 0
	oddSum := 0

	for i := 0; i < len(num); i++ {
		if i%2 == 0 {
			evenSum += int(num[i] - '0')
		} else {
			oddSum += int(num[i] - '0')
		}
	}
	return evenSum == oddSum
}
