package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result int
}

func TestCircularSentence(t *testing.T) {

	test1 := TestCase{
		"1001",
		2,
	}

	test2 := TestCase{
		"10",
		1,
	}

	test3 := TestCase{
		"0000",
		0,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := minChanges(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := minChanges(test2.s)
		want := test2.result
		checkResult(t, got, want)

	})

	t.Run("Test case 3", func(t *testing.T) {
		got := minChanges(test3.s)
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
