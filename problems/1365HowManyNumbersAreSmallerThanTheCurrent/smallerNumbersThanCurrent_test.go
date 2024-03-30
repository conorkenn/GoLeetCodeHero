package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result []int
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	test1 := TestCase{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}}
	test2 := TestCase{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}}
	test3 := TestCase{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}}

	t.Run("Test case 1", func(t *testing.T) {
		got := smallerNumbersThanCurrent(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := smallerNumbersThanCurrent(test2.nums)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := smallerNumbersThanCurrent(test3.nums)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
