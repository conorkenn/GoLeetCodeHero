package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	goal   string
	result bool
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"abcde",
		"cdeab",
		true,
	}

	test2 := TestCase{
		"abcde",
		"abced",
		false,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := rotateString(test1.s, test1.goal)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := rotateString(test2.s, test2.goal)
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
