package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	t      string
	result bool
}

func TestIsAnagram(t *testing.T) {

	test1 := TestCase{"anagram", "nagaram", true}
	test2 := TestCase{"rat", "car", false}

	t.Run("Test case 1", func(t *testing.T) {
		got := isAnagram(test1.s, test1.t)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isAnagram(test2.s, test2.t)
		want := test2.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want bool) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
