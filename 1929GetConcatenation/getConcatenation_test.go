package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result []int
}

func TestGetConcatenation(t *testing.T) {
	test1 := TestCase{
		[]int{1, 2, 1},
		[]int{1, 2, 1, 1, 2, 1},
	}

	test2 := TestCase{
		[]int{1, 3, 2, 1},
		[]int{1, 3, 2, 1, 1, 3, 2, 1},
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := getConcatenation(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := getConcatenation(test2.nums)
		want := test2.result
		checkResult(t, got, want)

	})
}

func checkResult(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
