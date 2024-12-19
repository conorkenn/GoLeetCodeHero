package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result string
}

func TestIsPalindrome(t *testing.T) {

	test1 := TestCase{"the sky is blue", "blue is sky the"}
	test2 := TestCase{"  hello world  ", "world hello"}
	test3 := TestCase{"a good   example", "example good a"}

	t.Run("Test case 1", func(t *testing.T) {
		got := reverseWords(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := reverseWords(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := reverseWords(test3.s)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
