package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	n      int
	result []int
}

func TestShuffleTheArray(t *testing.T) {

	test1 := TestCase{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}}
	test2 := TestCase{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}}
	test3 := TestCase{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}}

	t.Run("Test case 1", func(t *testing.T) {
		got := shuffle(test1.nums, test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := shuffle(test2.nums, test2.n)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := shuffle(test3.nums, test3.n)
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
