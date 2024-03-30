package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	x      int
	result bool
}

func TestIsPalindromer(t *testing.T) {

	test1 := TestCase{121, true}
	test2 := TestCase{-123, false}
	test3 := TestCase{10, false}

	t.Run("Test case 1", func(t *testing.T) {
		got := isPalindrome(test1.x)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isPalindrome(test2.x)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := isPalindrome(test3.x)
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
