package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		[]int{1, 1, 0, 1, 1, 1},
		3,
	}

	test2 := TestCase{
		[]int{1, 0, 1, 1, 0, 1},
		2,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := findMaxConsecutiveOnes(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := findMaxConsecutiveOnes(test2.nums)
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
