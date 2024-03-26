package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result int
}

func TestConvertTemperature(t *testing.T) {
	test1 := TestCase{[]int{1, 2, 3, 1, 1, 3}, 4}
	test2 := TestCase{[]int{1, 1, 1, 1}, 6}
	test3 := TestCase{[]int{1, 2, 3}, 0}

	t.Run("Test case 1", func(t *testing.T) {
		got := numIdenticalPairs(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := numIdenticalPairs(test2.nums)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := numIdenticalPairs(test3.nums)
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
