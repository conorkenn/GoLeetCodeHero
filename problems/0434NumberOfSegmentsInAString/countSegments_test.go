package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s      string
	result int
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		"Hello, my name is John",
		5,
	}

	test2 := TestCase{
		"Hello",
		1,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := countSegments(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := countSegments(test2.s)
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
