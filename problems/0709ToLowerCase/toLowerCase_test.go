package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result string
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"Hello",
		"hello",
	}

	test2 := TestCase{
		"here",
		"here",
	}

	test3 := TestCase{
		"LOVELY",
		"lovely",
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := toLowerCase(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := toLowerCase(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := toLowerCase(test3.s)
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
