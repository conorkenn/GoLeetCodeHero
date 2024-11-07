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
		[]int{10, 12, 13, 14},
		1,
	}

	test2 := TestCase{
		[]int{1, 2, 3, 4},
		1,
	}

	test3 := TestCase{
		[]int{999, 19, 199},
		10,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := minElement(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := minElement(test2.nums)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := minElement(test3.nums)
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
