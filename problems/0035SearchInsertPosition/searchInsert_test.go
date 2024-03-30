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

func TestMostWordsFound(t *testing.T) {

	test1 := TestCase{[]int{1, 3, 5, 6}, 5, 2}
	test2 := TestCase{[]int{1, 3, 5, 6}, 2, 1}
	test3 := TestCase{[]int{1, 3, 5, 6}, 7, 4}

	t.Run("Test case 1", func(t *testing.T) {
		got := searchInsert(test1.nums, test1.target)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := searchInsert(test2.nums, test2.target)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := searchInsert(test3.nums, test3.target)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
