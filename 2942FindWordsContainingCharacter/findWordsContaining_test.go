package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	words  []string
	x      byte
	result []int
}

func TestFindWordsContaining(t *testing.T) {

	test1 := TestCase{[]string{"leet", "code"}, byte('e'), []int{0, 1}}
	test2 := TestCase{[]string{"abc", "bcd", "aaaa", "cbc"}, byte('a'), []int{0, 2}}
	test3 := TestCase{[]string{"abc", "bcd", "aaaa", "cbc"}, byte('z'), []int{}}

	t.Run("Test case 1", func(t *testing.T) {
		got := findWordsContaining(test1.words, test1.x)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := findWordsContaining(test2.words, test2.x)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := findWordsContaining(test3.words, test3.x)
		want := test3.result
		checkResult(t, got, want)

	})

}

func checkResult(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
