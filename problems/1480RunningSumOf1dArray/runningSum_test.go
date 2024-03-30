package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	hours  []int
	result []int
}

func TestRunningSum(t *testing.T) {
	test1 := TestCase{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}}
	test2 := TestCase{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}}
	test3 := TestCase{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}}

	t.Run("Test case 1", func(t *testing.T) {
		got := runningSum(test1.hours)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := runningSum(test2.hours)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := runningSum(test3.hours)
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
