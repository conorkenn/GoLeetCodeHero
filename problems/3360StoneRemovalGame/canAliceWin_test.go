package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	num    int
	result bool
}

func TestCanAliceWin(t *testing.T) {

	test1 := TestCase{
		12,
		true,
	}

	test2 := TestCase{
		1,
		false,
	}

	t.Run("Test case 1", func(t *testing.T) {
		got := canAliceWin(test1.num)
		want := test1.result
		checkResult(t, got, want)

	})

	t.Run("Test case 2", func(t *testing.T) {
		got := canAliceWin(test2.num)
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
