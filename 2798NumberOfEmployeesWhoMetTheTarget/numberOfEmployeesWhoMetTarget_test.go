package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	hours  []int
	target int
	result int
}

func TestNumberOfEmployeesWhoMetTarget(t *testing.T) {
	test1 := TestCase{[]int{0, 1, 2, 3, 4}, 2, 3}
	test2 := TestCase{[]int{5, 1, 4, 2, 2}, 6, 0}

	t.Run("Test case 1", func(t *testing.T) {
		got := numberOfEmployeesWhoMetTarget(test1.hours, test1.target)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := numberOfEmployeesWhoMetTarget(test2.hours, test2.target)
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
