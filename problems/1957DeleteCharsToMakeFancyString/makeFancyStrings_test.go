package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	sentence string
	result   string
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"leeetcode",
		"leetcode",
	}

	test2 := TestCase{
		"aaabaaaa",
		"aabaa",
	}

	test3 := TestCase{
		"aab",
		"aab",
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := makeFancyString(test1.sentence)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := makeFancyString(test2.sentence)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := makeFancyString(test3.sentence)
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
