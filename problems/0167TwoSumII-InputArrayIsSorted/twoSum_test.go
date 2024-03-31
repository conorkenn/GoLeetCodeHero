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

func TestTwoSum2(t *testing.T) {

	test1 := TestCase{
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	}

	test2 := TestCase{
		[]int{2, 3, 4},
		6,
		[]int{1, 3},
	}

	test3 := TestCase{
		[]int{-1, 0},
		-1,
		[]int{1, 2},
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := twoSum2(test1.nums, test1.target)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := twoSum2(test2.nums, test2.target)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := twoSum2(test3.nums, test3.target)
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
