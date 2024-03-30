package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	ans := []string{}

	for i := 1; i <= n; i++ {
		insert := ""
		if i%3 == 0 {
			insert += "Fizz"
		}
		if i%5 == 0 {
			insert += "Buzz"
		}
		if insert == "" {
			insert = strconv.Itoa(i)
		}

		ans = append(ans, insert)
	}
	return ans
}
