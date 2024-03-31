package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result bool
}

func TestIsPalindrome(t *testing.T) {

	test1 := TestCase{"A man, a plan, a canal: Panama", true}
	test2 := TestCase{"race a car", false}
	test3 := TestCase{" ", true}

	t.Run("Test case 1", func(t *testing.T) {
		got := isPalindrome(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isPalindrome(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := isPalindrome(test3.s)
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
