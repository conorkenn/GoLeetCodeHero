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
		"leetcode",
		2,
	}

	test2 := TestCase{
		"abbcccddddeeeeedcba",
		5,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := maxPower(test1.s)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := maxPower(test2.s)
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
