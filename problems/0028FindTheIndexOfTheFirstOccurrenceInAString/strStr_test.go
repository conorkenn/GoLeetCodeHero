package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	haystack string
	needle   string
	result   int
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"sadbutsad",
		"sad",
		0,
	}

	test2 := TestCase{
		"leetcode",
		"leeto",
		-1,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := strStr(test1.haystack, test1.needle)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := strStr(test2.haystack, test2.needle)
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
