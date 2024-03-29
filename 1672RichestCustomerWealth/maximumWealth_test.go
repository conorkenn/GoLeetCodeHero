package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	accounts [][]int
	result   int
}

func TestMaximumWealth(t *testing.T) {
	test1 := TestCase{[][]int{{1, 2, 3}, {3, 2, 1}}, 6}
	test2 := TestCase{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10}
	test3 := TestCase{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17}

	t.Run("Test case 1", func(t *testing.T) {
		got := maximumWealth(test1.accounts)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := maximumWealth(test2.accounts)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := maximumWealth(test3.accounts)
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
