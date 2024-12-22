package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result int
}

func TestTwoSum(t *testing.T) {

	test1 := TestCase{
		[]int{1, 1, 4, 2, 1, 3},
		3,
	}

	test2 := TestCase{
		[]int{5, 1, 2, 3, 4},
		5,
	}

	test3 := TestCase{
		[]int{1, 2, 3, 4, 5},
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := heightChecker(test1.nums)
		want := test1.result
		checkResult(t, got, want)
	})

	t.Run("Test case 2", func(t *testing.T) {
		got := heightChecker(test2.nums)
		want := test2.result
		checkResult(t, got, want)
	})

	t.Run("Test case 3", func(t *testing.T) {
		got := heightChecker(test3.nums)
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
