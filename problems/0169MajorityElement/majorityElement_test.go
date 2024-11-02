package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result int
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		[]int{3, 2, 3},
		3,
	}

	test2 := TestCase{
		[]int{2, 2, 1, 1, 1, 2, 2},
		2,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := majorityElement(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := majorityElement(test2.nums)
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
