package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	result int
}

func TestTwoSum(t *testing.T) {

	test1 := TestCase{
		7,
		21,
	}

	test2 := TestCase{
		10,
		40,
	}

	test3 := TestCase{
		9,
		30,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := sumOfMultiples(test1.n)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := sumOfMultiples(test2.n)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := sumOfMultiples(test3.n)
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
