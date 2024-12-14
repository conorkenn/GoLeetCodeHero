package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	allowed string
	words   []string
	result  int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		"ab",
		[]string{"ad", "bd", "aaab", "baa", "badab"},
		2,
	}

	test2 := TestCase{
		"abc",
		[]string{"a", "b", "c", "ab", "ac", "bc", "abc"},
		7,
	}

	test3 := TestCase{
		"cad",
		[]string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
		4,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := countConsistentStrings(test1.allowed, test1.words)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := countConsistentStrings(test2.allowed, test2.words)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := countConsistentStrings(test3.allowed, test3.words)
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
