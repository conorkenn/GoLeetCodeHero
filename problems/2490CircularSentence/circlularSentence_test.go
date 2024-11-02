package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	sentence string
	result   bool
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"leetcode exercises sound delightful",
		true,
	}

	test2 := TestCase{
		"eetcode",
		true,
	}

	test3 := TestCase{
		"Leetcode is cool",
		false,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := isCircularSentence(test1.sentence)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := isCircularSentence(test2.sentence)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := isCircularSentence(test3.sentence)
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
