package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result []int
}

func TestBuildArray(t *testing.T) {
	test1 := TestCase{
		[]int{0, 2, 1, 5, 3, 4},
		[]int{0, 1, 2, 4, 5, 3},
	}

	test2 := TestCase{
		[]int{5, 0, 1, 2, 3, 4},
		[]int{4, 5, 0, 1, 2, 3},
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := buildArray(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := buildArray(test2.nums)
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
