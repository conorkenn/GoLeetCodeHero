package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	num1   int
	num2   int
	result int
}

func TestSum(t *testing.T) {

	test1 := TestCase{12, 5, 17}
	test2 := TestCase{-10, 4, -6}

	t.Run("Test case 1", func(t *testing.T) {
		got := sum(test1.num1, test1.num2)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := sum(test2.num1, test2.num2)
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
