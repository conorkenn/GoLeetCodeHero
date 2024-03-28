package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	n      int
	m      int
	result int
}

func TestDifferenceOfSums(t *testing.T) {

	test1 := TestCase{10, 3, 19}
	test2 := TestCase{5, 6, 15}
	test3 := TestCase{5, 1, -15}

	t.Run("Test case 1", func(t *testing.T) {
		got := differenceOfSums(test1.n, test1.m)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := differenceOfSums(test2.n, test2.m)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := differenceOfSums(test3.n, test3.m)
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
