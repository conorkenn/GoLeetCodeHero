package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	result int
}

func TestMaximumWealth(t *testing.T) {
	test1 := TestCase{[]int{-1, 1, 2, 3, 1}, 2, 3}
	test2 := TestCase{[]int{-6, 2, 5, -2, -7, -1, 3}, -2, 10}

	t.Run("Test case 1", func(t *testing.T) {
		got := countPairs(test1.nums, test1.target)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := countPairs(test2.nums, test2.target)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
