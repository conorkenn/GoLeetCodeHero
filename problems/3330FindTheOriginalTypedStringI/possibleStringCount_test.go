package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	word   string
	result int
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"abbcccc",
		5,
	}

	test2 := TestCase{
		"abcd",
		1,
	}

	test3 := TestCase{
		"aaaa",
		4,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := possibleStringCount(test1.word)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := possibleStringCount(test2.word)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := possibleStringCount(test3.word)
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
