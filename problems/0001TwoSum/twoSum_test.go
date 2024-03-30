package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	result []int
}

func TestTwoSum(t *testing.T) {

	test1 := TestCase{
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 0},
	}

	test2 := TestCase{
		[]int{3, 2, 4},
		6,
		[]int{2, 1},
	}

	test3 := TestCase{
		[]int{3, 3},
		6,
		[]int{1, 0},
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := twoSum(test1.nums, test1.target)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := twoSum(test2.nums, test2.target)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := twoSum(test3.nums, test3.target)
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
