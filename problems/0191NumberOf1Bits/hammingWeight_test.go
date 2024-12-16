package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	result int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		11,
		3,
	}

	test2 := TestCase{
		128,
		1,
	}

	test3 := TestCase{
		2147483645,
		30,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := hammingWeight(test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := hammingWeight(test2.n)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := hammingWeight(test3.n)
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
