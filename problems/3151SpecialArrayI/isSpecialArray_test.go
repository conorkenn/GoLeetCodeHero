package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	result bool
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		[]int{1},
		true,
	}

	test2 := TestCase{
		[]int{2, 1, 4},
		true,
	}

	test3 := TestCase{
		[]int{4, 3, 1, 6},
		false,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := isArraySpecial(test1.nums)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isArraySpecial(test2.nums)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := isArraySpecial(test3.nums)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want bool) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
